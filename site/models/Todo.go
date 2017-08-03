package models

import (
	"time"
  "database/sql"
)

type Todo struct {
  ID uint `gorm:"primary_key"`
  Title string `gorm:"not null"`
  Description sql.NullString `sql:"type:text;" gorm:"default:null"`
  Priority int `sql:"DEFAULT:3" gorm:"not null;index;default:3"`
  EstimatedDuration sql.NullInt64  `gorm:"default:null"`
  ActualDuration sql.NullInt64 `gorm:"default:null"`
  DoOn time.Time `gorm:"default:null"`
  DoBefore time.Time `gorm:"default:null"`
  CompletedOn time.Time `gorm:"default:null"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}