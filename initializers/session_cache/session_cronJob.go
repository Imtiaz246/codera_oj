package session_cache

import (
	"github.com/go-co-op/gocron"
	"github.com/imtiaz246/codera_oj/initializers/db"
	"github.com/imtiaz246/codera_oj/models"
	"log"
	"time"
)

// sessionCleanUpTask deletes the outdated session data from database and session cache
var sessionCleanUpTask = func() {
	curTime := time.Now()
	log.Printf("\n\nsession clean up started, Date: %v, Month: %v, Year: %v, Time:%v\n", curTime.Day(), curTime.Month(), curTime.Year(), curTime)

	database := db.GetDB()
	var sessionRecords []models.Sessions
	if err := database.Find(&sessionRecords).Error; err != nil {
		log.Printf("database initialization failed while session clean up")
		return
	}

	for _, sessionRecord := range sessionRecords {
		if time.Now().After(sessionRecord.ExpiresAt) {
			if err := database.Delete(&sessionRecord).Error; err != nil {
				log.Printf("outdated session deletion failed: %v", sessionRecord.ID.String())
			}

			// delete from session_cache
			DeleteFromSessionCache(sessionRecord.ID.String())
		}
	}
}

// InitializeSessionCleanUp_CronJob sets up a cron job to clean up expired session data
func InitializeSessionCleanUp_CronJob() error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Day().At("03:00").Do(sessionCleanUpTask)
	if err != nil {
		return err
	}
	s.StartAsync()

	return nil
}
