package repository

import (
	"time"

	"github.com/4Sao1Sad/yuemnoi-activity/internal/model"
	"gorm.io/gorm"
)

type ActivityLogRepositoryImpl struct {
	db *gorm.DB
}

type ActivityLogRepository interface {
	CreateActivityLog(logdetail string, user_id string) (*model.ActivityLog, error)
	ViewActivityHistoryByUserId(user_id string) (*model.ActivityLog, error)
}

func NewActivityLogRepository(db *gorm.DB) ActivityLogRepository {
	return &ActivityLogRepositoryImpl{db}
}

func (r ActivityLogRepositoryImpl) CreateActivityLog(logdetail string, user_id string) (*model.ActivityLog, error) {
	activityLog := model.ActivityLog{
		LogDetail: logdetail,
		User:      user_id,
		Timestamp: time.Now(),
	}
	err := r.db.Create(&activityLog).Error
	if err != nil {
		return nil, err
	}

	return &activityLog, nil
}

func (r ActivityLogRepositoryImpl) ViewActivityHistoryByUserId(user_id string) (*model.ActivityLog, error) {
	var activityHistory []model.ActivityLog
	err := r.db.Where("user = ?", user_id).Find(&activityHistory).Error
	if err != nil {
		return nil, err
	}

	return &activityHistory, nil
}