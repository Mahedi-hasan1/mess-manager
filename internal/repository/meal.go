package repository

import (
	"mess-manager/internal/db"
	"mess-manager/internal/dto"
	"mess-manager/internal/model"
)

func CreateMeal(meal *model.Meal) error {
	// Logic to save post to the database
	if err := db.PgDb.Create(meal).Error; err != nil {
		return err
	}
	return nil
}
func CreateMealType(mealType *model.MealType) error {
	// Logic to save post to the database
	if err := db.PgDb.Create(mealType).Error; err != nil {
		return err
	}
	return nil
}

func GetMealsByUsername(username string, limit int) ([]dto.UserPostResponse, error) {
	var posts []dto.UserPostResponse

	err := db.PgDb.Model(&model.Meal{}).
		Select("posts.id", "user_id", "content", "image_urls", "like_count", "dislike_count", "comment_count", "posts.created_at", "posts.updated_at").
		Joins("JOIN users ON posts.user_id = users.id").
		Where("users.username = ?", username).
		Order("posts.created_at DESC").
		Limit(limit).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}

// func GetFeedPosts(followingIDs []string, seenPostIDs []string, limit int) ([]model.Post, error) {
// 	var posts []model.Post
// 	//offset := (pageNum-1)*limit
// 	// I think dont need offset as we getting always unseen post . frontend request limit amount of next posts
// 	// then this api will return those amount of posts.
// 	//log.Println("seen post id ", seenPostIDs)
// 	query := db.PgDb.Model(&model.Post{}).
// 		Where("user_id IN ?", followingIDs).
// 		Order("hot_score DESC").
// 		Order("created_at DESC").
// 		Limit(limit)

// 	if len(seenPostIDs) > 0 {
// 		query = query.Where("id NOT IN ?", seenPostIDs)
// 	}

// 	if err := query.Preload("User").Find(&posts).Error; err != nil {
// 		return nil, err
// 	}
// 	return posts, nil
// }

func CountMeal(userID string) int64 {
	var count int64
	db.PgDb.Model(model.Meal{}).Where("user_id = ?", userID).Count(&count)
	return count
}
