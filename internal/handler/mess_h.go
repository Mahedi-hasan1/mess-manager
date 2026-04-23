package handler

import (
	"net/http"
	"mess-manager/internal/dto"
	"mess-manager/internal/service"
	"mess-manager/internal/validators"
	"github.com/gin-gonic/gin"
)
func AddMessMemer(c *gin.Context) {
	var memberReq dto.AddMessMemerRequest
	if err := c.ShouldBindJSON(&memberReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := validators.ValidateAddMessMemer(&memberReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.AddMessMemer(&memberReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, memberReq)
}