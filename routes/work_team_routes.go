package routes

import (
	"services-work-team-delete/controllers"

	"github.com/gin-gonic/gin"
)

func WorkTeamRoutes(r *gin.RouterGroup) {
	r.DELETE("/work-team/delete/:teamCode", controllers.DeleteWorkTeam)
}
