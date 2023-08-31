package router

import (
	"mongoapi/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/getalldata", controllers.Getalldata)
	router.POST("/getdatabydate", controllers.Getalldataformpost)
	router.POST("/getsumformpost", controllers.Getsumformpost)
	router.POST("/getdatabyid", controllers.Getdatabyid)
	router.POST("/createprofile", controllers.CreateProfile)
	router.PUT("/updateprofile/:id", controllers.UpdateProfile)
	router.DELETE("/deleteprofile/:id", controllers.Deleteprofile)
	router.DELETE("/deleteallprofile", controllers.Deleteallprofile)
	router.POST("/moneytransfer", controllers.Depositcontroll)

	return router
}