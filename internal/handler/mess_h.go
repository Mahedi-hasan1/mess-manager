package handler

import (
	"net/http"
	"mess-manager/internal/dto"
	"mess-manager/internal/service"
	"mess-manager/internal/validators"
	"github.com/gin-gonic/gin"
	"fmt"
)
func AddMessMemer(c *gin.Context) {
	var memberReq dto.AddMessMemerRequest
	if err := c.ShouldBindJSON(&memberReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := validators.ValidateAddMessMember(&memberReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddMessMemer(&memberReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, memberReq)
}

func CreateMess(c *gin.Context){
	userId := c.GetString("user_id")
	fmt.Println("user id : ", userId);
	var messReq dto.CreateMessRequest
	if err := c.ShouldBindJSON(&messReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := validators.ValidateCreateMess(&messReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mess, err := service.CreateMess(messReq, userId);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create mess: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, mess)
}