package routes

import (
	controller "github.com/ChanchalS7/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateInvoice())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())

}
