package repository

import (
	"context"
	"time"

	"github.com/4Sao1Sad/yuemnoi-activity/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ActivityLogRepositoryImpl struct {
	db *mongo.Collection
}

type ActivityLogRepository interface {
	CreateActivityLog(logdetail string, user_id uint) (*model.ActivityLog, error)
	ViewActivityHistoryByUserId(user_id uint) (*[]model.ActivityLog, error)
}

func NewActivityLogRepository(db *mongo.Database) ActivityLogRepository {
	collection := db.Collection("activity_logs")
	return &ActivityLogRepositoryImpl{collection}
}

func formatTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.00000")
}

func (r ActivityLogRepositoryImpl) CreateActivityLog(logdetail string, user_id uint) (*model.ActivityLog, error) {
	activityLog := model.ActivityLog{
		LogDetail: logdetail,
		UserId:    user_id,
		Timestamp: formatTimestamp(time.Now()),
	}

	// Insert into MongoDB
	_, err := r.db.InsertOne(context.TODO(), activityLog)
	if err != nil {
		return nil, err
	}

	return &activityLog, nil
}

func (r ActivityLogRepositoryImpl) ViewActivityHistoryByUserId(user_id uint) (*[]model.ActivityLog, error) {
	var activityHistory []model.ActivityLog

	// Query for the user activity history
	filter := bson.M{"user_id": user_id}
	cursor, err := r.db.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	// Parse the results into activityHistory
	if err = cursor.All(context.TODO(), &activityHistory); err != nil {
		return nil, err
	}

	return &activityHistory, nil
}
