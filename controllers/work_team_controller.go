package controllers

import (
	"fmt"
	"net/http"
	"services-work-team-delete/models"
	"services-work-team-delete/services"

	"github.com/gin-gonic/gin"
)

// UpdateWorkTeam godoc
// @Summary      Update work team
// @Description  Update a work team based on teamCode
// @Tags         work-team
// @Accept       json
// @Produce      json
// @Param        teamCode  path      string  true  "Team code"
// @Param        team      body      models.WorkTeam true  "Updated work team"
// @Success      200       {object}  models.ApiResponse[models.WorkTeam]
// @Failure      400       {object}  models.ApiResponse[models.WorkTeam]
// @Failure      404       {object}  models.ApiResponse[models.WorkTeam]
// @Router       /work-team/delete/{teamCode} [delete]
func DeleteWorkTeam(c *gin.Context) {
	teamCode := c.Param("teamCode")

	var updatedTeam models.WorkTeam
	if err := c.ShouldBindJSON(&updatedTeam); err != nil {
		c.JSON(http.StatusBadRequest, models.ApiResponse[models.WorkTeam]{
			Status:  "error",
			Data:    nil,
			Message: fmt.Sprintf("Invalid input: %s", err.Error()),
		})
		return
	}

	updatedWorkTeam, err := services.DeleteWorkTeamService(teamCode, &updatedTeam)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ApiResponse[models.WorkTeam]{
			Status:  "empty",
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ApiResponse[models.WorkTeam]{
		Status:  "success",
		Data:    updatedWorkTeam,
		Message: "Work team updated successfully",
	})
}
