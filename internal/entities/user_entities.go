package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/segmentio/ksuid"
)

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type KSUIDArray []ksuid.KSUID

// Implement GORM's Valuer interface for KSUIDArray
func (a KSUIDArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Implement GORM's Scanner interface for KSUIDArray
func (a *KSUIDArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, a)
}

type User struct {
	ID                  ksuid.KSUID `gorm:"primary_key;not null"`
	Username            string      `gorm:"not null;index"`
	Email               string      `gorm:"unique;not null;index"`
	Password            string      `gorm:"not null"`
	Role                string      `gorm:"type:text;not null;index;check:role IN ('admin', 'user')"` // Use CHECK constraint
	ProfilePicture      string      `gorm:"default:''"`
	ProfilePictureUrl   string      `gorm:"default:''"`
	IsVerified          bool        `gorm:"default:false"`
	VerificationToken   string      `gorm:"default:'';index"`
	ForgetPasswordToken string      `gorm:"default:'';index"`
	SubscriptionStatus  bool        `gorm:"default:false;index"`
	SubscriptionToken   string      `gorm:"default:'';index"`
	CreatedAt           time.Time   `gorm:"autoCreateTime;index"`
	UpdatedAt           time.Time   `gorm:"autoCreateTime;index"`
}

func NewUser(username, email, password string, role Role) (*User, error) {
	if role != AdminRole && role != UserRole {
		return nil, errors.New("invalid role")
	}

	user := &User{
		ID:        ksuid.New(),
		Username:  username,
		Email:     email,
		Password:  password,
		Role:      string(role),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}
