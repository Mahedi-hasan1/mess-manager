package service

import (
	"errors"
	"log"
	"mess-manager/internal/dto"
	"mess-manager/internal/model"
	"mess-manager/internal/repository"
	"time"
)

func CreateMeal(mealCreateReq *dto.CreateMealRequest, userID string) error {
	meal := &model.Meal{
		UserID:    userID,
		// Content:   postCreateReq.Content,
		// ImageURLs: postCreateReq.ImageURLs,
	}
	return repository.CreateMeal(meal)
}
func CreateMealType(mealType *model.MealType,) error {
	return repository.CreateMealType(mealType)
}

func GetUserMeals(username string, limit int) ([]dto.UserPostResponse, error) {
	return repository.GetMealsByUsername(username, limit)
}
func BulkCreateMeal(mealsCreateReq *[]dto.CreateMealRequest, username string) error {
	now := time.Now()
	user, err := repository.GetUser("", "", username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("No User found of this username")
	}

	for i, mealCreateReq := range *mealsCreateReq {
		createdAt := now.Add(time.Duration(-i*24) * time.Hour)
		meal := &model.Meal{
			UserID:    user.ID,
			// Content:   postCreateReq.Content,
			// ImageURLs: postCreateReq.ImageURLs,
			CreatedAt: createdAt,
		}
		if err := repository.CreateMeal(meal); err != nil {
			log.Panicln("post not created. details: ", meal)
		} else {
			log.Println("post created: details", mealCreateReq)
		}
	}
	return nil
}
