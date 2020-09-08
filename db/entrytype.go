package db

func EntryTypeExistsInProject(entryType string, projectID uint) bool {
	doesNotExist := DB.Where("project_id = ? AND title = ?", projectID, entryType).Find(&ProjectEntryType{}).RecordNotFound()
	return !doesNotExist
}

func CreateEntryType(entryType string, projectID uint) error {
	newEntryType := &ProjectEntryType{Title: entryType, ProjectID: projectID, Colour: ""}
	err := DB.Create(&newEntryType).Error
	return err
}
