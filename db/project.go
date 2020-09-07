package db

func ProjectExists(projectID uint) bool {
	doesNotExist := DB.Where("id = ?", projectID).Find(&Project{}).RecordNotFound()
	return !doesNotExist
}

func ProjectIsOwnedByUser(projectID uint, userID uint) (bool, error) {
	var project Project
	err := DB.Where("id = ?", projectID).Find(&project).Error
	if err != nil {
		return false, err
	}
	return project.UserID == userID, nil
}

func GetProjectByID(projectID uint) (Project, error) {
	var project Project
	err := DB.Where("id = ?", projectID).Preload("Entries").Preload("EntryTypes").Find(&project).Error
	if err != nil {
		return project, err
	}
	return project, nil
}
