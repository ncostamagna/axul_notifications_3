package notification

import (
	"time"

	"gorm.io/gorm"
)

// Notification entity
type Notification struct {
	ID         string         `json:"id" gorm:"not null;primary_key;unique_index;AUTO_INCREMENT"`
	UserID     string         `json:"user_id" gorm:"type:varchar(50);not null;unique"`
	NotifyType string         `json:"alert_type" gorm:"type:varchar(20);not null;unique"`
	Message    string         `json:"message" gorm:"type:varchar(200);not null;unique"`
	Checked    bool           `json:"checked"`
	DeletedAt  gorm.DeletedAt `json:"-"`
	CreatedAt  time.Time      `json:"-" gorm:"type:timestamp;column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `json:"-" gorm:"type:timestamp;column:updated_at;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (d Notification) TableName() string {
	return "notification"
}
