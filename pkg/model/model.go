package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Name string  `gorm:"not null"`
	Email string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Hospital struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Name string
	Address string
	ImageUrl string
	Rating float64
	Lantitude float64
	Longitude float64
	Doctor []Doctor `gorm:"foreignKey:DoctorId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Specialists struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Name string
	IconUrl string
	NoOfDoc int64
	Doctor []Doctor `gorm:"foreignKey:DoctorId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Doctor struct{
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Name string
	Experience float64
	Qualification string
	Doc_Details string
	Rating float64
	CreatedAt time.Time
	UpdatedAt time.Time
}