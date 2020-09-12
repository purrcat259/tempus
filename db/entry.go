package db

import (
	"errors"
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

func CreateEntry(db *gorm.DB, projectID uint, entryType string) error {
	newEntry := &ProjectEntry{
		EntryType: entryType,
		ProjectID: projectID,
		OpenTime:  time.Now(),
	}
	err := db.Create(&newEntry).Error
	return err
}

func EntryExists(entryID uint) bool {
	doesNotExist := DB.Where("id = ?", entryID).Find(&ProjectEntry{}).RecordNotFound()
	return !doesNotExist
}

func CloseEntry(db *gorm.DB, entryID uint) error {
	err := db.Model(&ProjectEntry{}).Where("id = ?", entryID).Update("close_time", time.Now()).Error
	return err
}

func GetEntriesBetweenDatetimes(projectID uint, startTime time.Time, endTime time.Time) ([]ProjectEntry, error) {
	entries := []ProjectEntry{}
	err := DB.Where("project_id = ? AND close_time IS NOT NULL AND open_time BETWEEN ? AND ?", projectID, startTime.Format("2006-01-02"), endTime.Format("2006-01-02")).Find(&entries).Error
	return entries, err
}

func SwitchEntry(projectID uint, targetEntryType string) error {
	hasOngoingEntry, ongoingEntry, err := GetOngoingEntry(projectID)
	if err != nil {
		return err
	}
	if !hasOngoingEntry {
		return errors.New("Project does not have an ongoing entry")
	}
	tx := DB.Begin()
	err = CloseEntry(tx, ongoingEntry.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = CreateEntry(tx, projectID, targetEntryType)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
