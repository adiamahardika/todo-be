package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {

	router := gin.Default()
	router.Run(":3030")
}
