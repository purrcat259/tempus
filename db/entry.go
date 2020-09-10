package db

import (
	"time"
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

func CreateEntry(projectID uint, entryType string) error {
	newEntry := &ProjectEntry{
		EntryType: entryType,
		ProjectID: projectID,
		OpenTime:  time.Now(),
	}
	err := DB.Create(&newEntry).Error
	return err
}

func EntryExists(entryID uint) bool {
	doesNotExist := DB.Where("id = ?", entryID).Find(&ProjectEntry{}).RecordNotFound()
	return !doesNotExist
}

func CloseEntry(entryID uint) error {
	err := DB.Model(&ProjectEntry{}).Where("id = ?", entryID).Update("close_time", time.Now()).Error
	return err
}

func GetEntriesBetweenDatetimes(projectID uint, startTime time.Time, endTime time.Time) ([]ProjectEntry, error) {
	entries := []ProjectEntry{}
	err := DB.Where("project_id = ? AND close_time IS NOT NULL AND open_time BETWEEN ? AND ?", projectID, startTime.Format("2006-01-02"), endTime.Format("2006-01-02")).Find(&entries).Error
	return entries, err
}
