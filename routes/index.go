package routes

import "github.com/gin-gonic/gin"

func RouterCombinerFunction() *gin.Engine {
	router := gin.New();
	router.Use(gin.Logger());

	UserRouters(router);

	return router;
}