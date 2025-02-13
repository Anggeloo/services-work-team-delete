package services

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"services-work-team-delete/database"
	"services-work-team-delete/models"
)

// DeleteWorkTeamService delete a work team based on teamCode
func DeleteWorkTeamService(teamCode string, updatedTeam *models.WorkTeam) (*models.WorkTeam, error) {
	var existingTeam models.WorkTeam

	if err := database.DB.Where("team_code = ?", teamCode).First(&existingTeam).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("team with teamCode %s not found", teamCode)
		}
		return nil, fmt.Errorf("error finding team: %w", err)
	}

	existingTeam.Status = false
	existingTeam.UpdatedAt = time.Now()

	if err := database.DB.Save(&existingTeam).Error; err != nil {
		return nil, fmt.Errorf("error updating team: %w", err)
	}

	return &existingTeam, nil
}
