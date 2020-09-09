package db

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // needed for sqlite
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Projects []Project
}

func (u *User) BeforeCreate() (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	u.Password = string(bytes)
	return
}

type Project struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Entries    []ProjectEntry
	EntryTypes []ProjectEntryType
	UserID     uint
}

func (p *Project) HasEntries() bool {
	return len(p.Entries) != 0
}

type ProjectEntryType struct {
	gorm.Model
	Title     string `gorm:"not null"`
	ProjectID uint
	Colour    string
}

type ProjectEntry struct {
	gorm.Model
	EntryType string
	ProjectID uint
	OpenTime  time.Time
	CloseTime *time.Time
}

func (pe *ProjectEntry) IsOngoing() bool {
	return pe.CloseTime == nil
}

func (pe *ProjectEntry) TimeTaken() (float64, float64, float64) {
	// https://stackoverflow.com/a/40262557
	if pe.IsOngoing() {
		return 0, 0, 0
	}
	diff := pe.CloseTime.Sub(pe.OpenTime)
	hs := diff.Hours()
	hs, mf := math.Modf(hs)
	ms := mf * 60
	ms, sf := math.Modf(ms)
	ss := sf * 60
	return hs, ms, ss
}

func (pe *ProjectEntry) TimeTakenHuman() string {
	if pe.IsOngoing() {
		return "Ongoing"
	}
	hours, minutes, seconds := pe.TimeTaken()
	sentence := ""
	delimiter := ""
	hasHours := math.Floor(hours) > 0
	hasMinutes := math.Floor(minutes) > 0
	hasSeconds := math.Floor(seconds) > 0
	if hasHours {
		sentence = fmt.Sprintf("%.0f Hours", hours)
	}
	if hasMinutes {
		if hasHours {
			delimiter = ", "
		}
		sentence = fmt.Sprintf("%s%s%.0f Minutes", sentence, delimiter, minutes)
	}
	if hasSeconds {
		if hasMinutes {
			delimiter = ", "
		}
		sentence = fmt.Sprintf("%s%s%.0f Seconds", sentence, delimiter, seconds)
	}
	return sentence
}

func Open() {
	var err error
	DB, err = gorm.Open("sqlite3", "data/tempus.db")
	if err != nil {
		log.Println("Unable to open DB!")
		panic(err)
	}
	DB.LogMode(true)
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Project{})
	DB.AutoMigrate(&ProjectEntryType{})
	DB.AutoMigrate(&ProjectEntry{})
}

func Clear() {
	info, err := os.Stat("./tempus.db")
	if err != nil {
		log.Println(err.Error())
		return
	}
	if !info.IsDir() {
		err := os.Remove("./tempus.db")
		if err != nil {
			panic(err)
		}
	}
}

const isoParseFormat = "2006-01-02 15:04:05-07:00"

func Seed() {
	// TODO: set default admin password in env
	baseUser := User{Name: "Admin", Email: "simon@agius-muscat.net", Password: "admin123"}

	notFound := DB.Where("email = ?", baseUser.Email).Find(&User{}).RecordNotFound()
	if notFound {
		DB.Create(&baseUser)
	}
	count := 0
	err := DB.Model(&Project{}).Count(&count).Error
	if err != nil {
		panic(err)
	}
	if count == 0 {
		project := Project{Title: "Test Project", UserID: 1}
		err := DB.Create(&project).Error
		if err != nil {
			panic(err)
		}
	}

}
