package entity

import (
	"github.com/Hello-Storage/hello-storage-proxy/internal/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Miner struct {
	ID             uint   `gorm:"primarykey"                            json:"id"`
	Balance        string `gorm:"type:varchar(255);not null;" json:"balance"`
	OfferedStorage string `gorm:"type:varchar(255);" json:"offered_storage"`
	LastChallenge  string `gorm:"type:varchar(16);not null"             json:"last_challenge"`
	UserId         uint   `gorm:"column:user_id;not null;uniqueIndex" json:"user_id"` // Renamed field
}

// TableName returns the entity table name.
func (Miner) TableName() string {
	return "miners"
}

func (m *Miner) Create() error {
	db.Db().Logger = db.Db().Logger.LogMode(logger.Info)
	return db.Db().Create(m).Error
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *Miner) BeforeCreate(db *gorm.DB) error {
	//m.LastChallenge = rnd.GenerateRandomString(16)
	//db.Statement.SetColumn("nonce", m.LastChallenge)
	return nil
}

func (m *Miner) Save() error {
	return db.Db().Save(m).Error
}
