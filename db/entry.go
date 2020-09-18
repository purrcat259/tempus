package db

import (
	"errors"
	"math"
	"sort"
	"time"

	"github.com/jinzhu/gorm"
)

func GetOngoingEntry(projectID uint) (bool, *ProjectEntry, error) {
	hasEntries := ProjectHasEntries(projectID)
	if !hasEntries {
		return false, nil, nil
	}
	noOngoingEntry := DB.Where("project_id = ? AND close_time IS NULL", projectID).Find(&ProjectEntry{}).RecordNotFound()
	if noOngoingEntry {
		return false, nil, nil
	}
	var ongoingEntry ProjectEntry
	err := DB.Order("open_time DESC").Where("project_id = ? AND close_time IS NULL", projectID).Find(&ongoingEntry).Error
	if err != nil {
		return false, nil, err
	}
	return true, &ongoingEntry, nil
}

func CreateEntry(db *gorm.DB, projectID uint, entryType string, startedWithContextSwitch bool) error {
	newEntry := &ProjectEntry{
		EntryType:                entryType,
		ProjectID:                projectID,
		OpenTime:                 time.Now(),
		StartedWithContextSwitch: startedWithContextSwitch,
	}
	err := db.Create(&newEntry).Error
	return err
}

func EntryExists(entryID uint) bool {
	doesNotExist := DB.Where("id = ?", entryID).Find(&ProjectEntry{}).RecordNotFound()
	return !doesNotExist
}

func CloseEntry(db *gorm.DB, entryID uint, endedWithContextSwitch bool) error {
	now := time.Now()
	err := db.Model(&ProjectEntry{}).Where("id = ?", entryID).Updates(&ProjectEntry{CloseTime: &now, EndedWithContextSwitch: endedWithContextSwitch}).Error
	return err
}

func GetEntriesBetweenDatetimes(projectID uint, startTime time.Time, endTime time.Time) ([]ProjectEntry, error) {
	entries := []ProjectEntry{}
	err := DB.Where("project_id = ? AND close_time IS NOT NULL AND open_time BETWEEN ? AND ?", projectID, startTime.Format("2006-01-02"), endTime.Format("2006-01-02")).Find(&entries).Error
	return entries, err
}

func SwitchEntry(projectID uint, targetEntryType string, contextSwitchHappening bool) error {
	hasOngoingEntry, ongoingEntry, err := GetOngoingEntry(projectID)
	if err != nil {
		return err
	}
	if !hasOngoingEntry {
		return errors.New("Project does not have an ongoing entry")
	}
	tx := DB.Begin()
	err = CloseEntry(tx, ongoingEntry.ID, contextSwitchHappening)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = CreateEntry(tx, projectID, targetEntryType, contextSwitchHappening)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type ProjectWithOngoingEntry struct {
	Project      Project
	OngoingEntry ProjectEntry
}

func GetAllOngoingEntriesForUser(userID uint) ([]ProjectWithOngoingEntry, error) {
	projectsWithOngoingEntry := []ProjectWithOngoingEntry{}
	if !UserExistsByID(userID) {
		return projectsWithOngoingEntry, errors.New("User does not exist")
	}
	user, err := GetUserByID(userID)
	if err != nil {
		return projectsWithOngoingEntry, err
	}
	if len(user.Projects) == 0 {
		return projectsWithOngoingEntry, nil
	}
	for _, project := range user.Projects {
		hasOngoingEntry, ongoingEntry, err := GetOngoingEntry(project.ID)
		if err != nil {
			return projectsWithOngoingEntry, err
		}
		if hasOngoingEntry {
			projectsWithOngoingEntry = append(projectsWithOngoingEntry, ProjectWithOngoingEntry{Project: project, OngoingEntry: *ongoingEntry})
		}
	}
	return projectsWithOngoingEntry, nil
}

type TimeBreakdown struct {
	Hours   float64
	Minutes float64
	Seconds float64
}

type TimeProportion struct {
	EntryType         string
	TimeBreakdown     TimeBreakdown
	PercentageOfTotal float64
}

type EntryStatistics struct {
	EntryProportions []TimeProportion
	TotalTime        TimeBreakdown
	Count            int
}

func CalculateEntryStatisticsToday(entries []ProjectEntry) EntryStatistics {
	var todayEntryStatistics []ProjectEntry
	for _, entry := range entries {
		if entry.OpenedToday() {
			todayEntryStatistics = append(todayEntryStatistics, entry)
		}
	}
	return CalculateEntriesStatistics(todayEntryStatistics)
}

func SecondsToHoursMinutesSeconds(seconds float64) (float64, float64, float64) {
	asMinutes := seconds / 60.0
	wholeMinutes, fracMinutes := math.Modf(asMinutes)
	actualSeconds := fracMinutes * 60.0
	asHours := wholeMinutes / 60.0
	wholeHours, fracHours := math.Modf(asHours)
	actualMinutes := fracHours * 60.0
	return wholeHours, actualMinutes, actualSeconds
}

func CalculateEntriesStatistics(entries []ProjectEntry) EntryStatistics {
	var totalSeconds float64
	for _, entry := range entries {
		var entryLengthSeconds float64 = 0.0
		if !entry.IsOngoing() {
			entryLengthSeconds = entry.CloseTime.Sub(entry.OpenTime).Seconds()
		} else {
			entryLengthSeconds = time.Now().Sub(entry.OpenTime).Seconds()
		}
		totalSeconds += entryLengthSeconds
	}

	wholeHours, actualMinutes, actualSeconds := SecondsToHoursMinutesSeconds(totalSeconds)

	// Calculate proprotions
	secondsByEntryType := make(map[string]float64)
	for _, entry := range entries {
		var entryLengthSeconds float64 = 0.0
		if !entry.IsOngoing() {
			entryLengthSeconds = entry.CloseTime.Sub(entry.OpenTime).Seconds()
		} else {
			entryLengthSeconds = time.Now().Sub(entry.OpenTime).Seconds()
		}

		if _, ok := secondsByEntryType[entry.EntryType]; ok {
			secondsByEntryType[entry.EntryType] += entryLengthSeconds
		} else {
			secondsByEntryType[entry.EntryType] = entryLengthSeconds
		}
	}
	var proportionsByEntryType []TimeProportion
	for entryType, totalSecondsForEntryType := range secondsByEntryType {
		hours, minutes, seconds := SecondsToHoursMinutesSeconds(totalSecondsForEntryType)
		timeProportion := TimeProportion{
			EntryType: entryType,
			TimeBreakdown: TimeBreakdown{
				Hours:   hours,
				Minutes: minutes,
				Seconds: seconds,
			},
			PercentageOfTotal: totalSecondsForEntryType / totalSeconds * 100.0,
		}
		proportionsByEntryType = append(proportionsByEntryType, timeProportion)
	}

	// Ensure they are in descending order
	sort.SliceStable(proportionsByEntryType, func(i, j int) bool {
		return proportionsByEntryType[i].PercentageOfTotal > proportionsByEntryType[j].PercentageOfTotal
	})

	totalTime := TimeBreakdown{Hours: wholeHours, Minutes: actualMinutes, Seconds: actualSeconds}
	stats := EntryStatistics{TotalTime: totalTime, Count: len(entries), EntryProportions: proportionsByEntryType}
	return stats
}
