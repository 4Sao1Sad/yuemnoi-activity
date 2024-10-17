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
	ViewActivityHistoryByUserId(user_id string) (*[]model.ActivityLog, error)
}

func NewActivityLogRepository(db *gorm.DB) ActivityLogRepository {
	return &ActivityLogRepositoryImpl{db}
}

func formatTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.00000")
}

func (r ActivityLogRepositoryImpl) CreateActivityLog(logdetail string, user_id string) (*model.ActivityLog, error) {
	activityLog := model.ActivityLog{
		LogDetail: logdetail,
		UserId:      user_id,
		Timestamp: formatTimestamp(time.Now()),
	}
	err := r.db.Create(&activityLog).Error
	if err != nil {
		return nil, err
	}

	return &activityLog, nil
}

func (r ActivityLogRepositoryImpl) ViewActivityHistoryByUserId(user_id string) (*[]model.ActivityLog, error) {
	var activityHistory []model.ActivityLog
	err := r.db.Where("user_id = ?", user_id).Find(&activityHistory).Error
	if err != nil {
		return nil, err
	}

	return &activityHistory, nil
}