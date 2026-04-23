package handler

import (
	"mess-manager/internal/dto"
	"mess-manager/internal/model"
	"mess-manager/internal/service"
	"mess-manager/internal/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddMeal(c *gin.Context) {
	userID := c.GetString("user_id")
	var mealCreateReq dto.AddMealRequest
	if err := c.ShouldBindJSON(&mealCreateReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	if err := validators.ValidateCreateMeal(&mealCreateReq, userID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateMeal(&mealCreateReq, userID); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create post: " + err.Error()})
		return
	}
	c.JSON(201, mealCreateReq)
}
func AddMealType(c *gin.Context) {
	
	var mealTypeReq model.MealType
	if err := c.ShouldBindJSON(&mealTypeReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	// if err := validators.ValidateCreateMeal(&mealCreateReq, userID); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := service.CreateMealType(&mealTypeReq); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create Meal type: " + err.Error()})
		return
	}
	c.JSON(201, mealTypeReq)
}

func GetMealsByUsername(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	username := c.Query("username")
	if err != nil {
		c.JSON(400, gin.H{"error": "page and limit should be Integer Value " + err.Error()})
	}

	if err := validators.ValidateGetUserMeals(username, limit); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	posts, err := service.GetUserMeals(username, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get posts: " + err.Error()})
	}
	c.JSON(200, posts)
}

// func BulkCreatePosts(c *gin.Context) {
// 	username := c.Query("user_name")
// 	var postsCreateReq []dto.CreatePostRequest
// 	if err := c.ShouldBindJSON(&postsCreateReq); err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid request"})
// 		return
// 	}
// 	if err := service.BulkCreatePost(&postsCreateReq, username); err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to create post: " + err.Error()})
// 		return
// 	}
// 	c.JSON(201, postsCreateReq)
// }
