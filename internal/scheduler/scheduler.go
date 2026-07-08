package scheduler

import (
	"context"
	"fmt"
	"kulkasku/internal/model"
	foodRepository "kulkasku/internal/repository/food"
	notificationRepository "kulkasku/internal/repository/notification"
	"log"
	"time"
)

type Scheduler struct {
	foodRepo  foodRepository.FoodRepository
	notifRepo notificationRepository.NotificationRepository
}

func New(foodRepo foodRepository.FoodRepository, notifRepo notificationRepository.NotificationRepository) *Scheduler {
	return &Scheduler{
		foodRepo:  foodRepo,
		notifRepo: notifRepo,
	}
}

func (s *Scheduler) Start() {
	go s.runDailyExpiredCheck()
	go s.runWeeklyRecipeRecommendation()
	log.Println("Scheduler started")
}

func (s *Scheduler) runDailyExpiredCheck() {
	now := time.Now()
	firstRun := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())
	if now.After(firstRun) {
		firstRun = firstRun.Add(24 * time.Hour)
	}

	delay := time.Until(firstRun)
	if delay < 0 {
		delay = 0
	}
	log.Printf("Scheduler: expired check starts in %v", delay)

	time.Sleep(delay)
	s.checkExpiredFoods()

	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		s.checkExpiredFoods()
	}
}

func (s *Scheduler) runWeeklyRecipeRecommendation() {
	now := time.Now()
	daysUntilMonday := (7 - int(now.Weekday()) + 1) % 7
	if daysUntilMonday == 0 {
		daysUntilMonday = 7
	}
	firstRun := time.Date(now.Year(), now.Month(), now.Day()+daysUntilMonday, 9, 0, 0, 0, now.Location())

	delay := time.Until(firstRun)
	if delay < 0 {
		delay = 0
	}
	log.Printf("Scheduler: recipe recommendation starts in %v", delay)

	time.Sleep(delay)
	s.checkRecipeRecommendation()

	ticker := time.NewTicker(7 * 24 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		s.checkRecipeRecommendation()
	}
}

func (s *Scheduler) checkExpiredFoods() {
	ctx := context.Background()
	now := time.Now()
	threeDaysLater := now.AddDate(0, 0, 3)

	log.Println("Scheduler: checking expired foods...")

	foods, err := s.foodRepo.GetExpiringFoods(ctx, now, threeDaysLater)
	if err != nil {
		log.Printf("Scheduler error: %v", err)
		return
	}

	created := 0
	for _, food := range foods {
		title := "Makanan Akan Kadaluarsa"
		message := fmt.Sprintf("%s akan kadaluarsa pada %s", food.Name, food.ExpiredAt.Format("02 Jan 2006"))

		_, err := s.notifRepo.Save(ctx, &model.Notification{
			UserID:   food.UserID,
			FoodID:   food.ID,
			Title:    title,
			Message:  message,
			Type:     "expired",
			NotifyAt: food.ExpiredAt,
		})
		if err != nil {
			log.Printf("Scheduler: failed to save notification for food %d: %v", food.ID, err)
			continue
		}
		created++
	}

	log.Printf("Scheduler: created %d expired notifications", created)
}

func (s *Scheduler) checkRecipeRecommendation() {
	ctx := context.Background()

	log.Println("Scheduler: checking recipe recommendations...")

	userIDs, err := s.foodRepo.GetUserIDsWithMinFoods(ctx, 5)
	if err != nil {
		log.Printf("Scheduler error: %v", err)
		return
	}

	created := 0
	for _, userID := range userIDs {
		foodID, err := s.foodRepo.GetFirstFoodIdByUserId(ctx, userID)
		if err != nil {
			log.Printf("Scheduler: failed to get food for user %d: %v", userID, err)
			continue
		}

		title := "Rekomendasi Resep"
		message := "Anda memiliki cukup bahan makanan! Coba generate resep baru sekarang."

		_, err = s.notifRepo.Save(ctx, &model.Notification{
			UserID:   userID,
			FoodID:   foodID,
			Title:    title,
			Message:  message,
			Type:     "recipe",
			NotifyAt: time.Now(),
		})
		if err != nil {
			log.Printf("Scheduler: failed to save recipe notification for user %d: %v", userID, err)
			continue
		}
		created++
	}

	log.Printf("Scheduler: created %d recipe recommendation notifications", created)
}
