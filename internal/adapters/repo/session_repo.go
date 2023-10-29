package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/imtiaz246/codera_oj/internal/adapters/repo/db"
	"github.com/imtiaz246/codera_oj/internal/core/domain/models"
	"github.com/imtiaz246/codera_oj/internal/core/ports"
	"github.com/imtiaz246/codera_oj/modules/cronera"
	"log"
	"time"
)

type sessionRepo struct {
	ports.GenericInterface[*models.Session]
	*db.Database
}

var _ ports.SessionRepoInterface = (*sessionRepo)(nil)

func NewSessionRepo(d *db.Database) (ports.SessionRepoInterface, error) {
	sr := &sessionRepo{
		Database:         d,
		GenericInterface: NewGenericRepo[*models.Session](d),
	}
	_, err := cronera.New().Every(1).Day().At("03:00").Do(context.Background(), sr.expiredSessionRemover)
	if err != nil {
		return nil, err
	}
	return sr, nil
}

func (sr *sessionRepo) GetSessionListOfUser(userID int64) ([]models.Session, error) {
	sessions := make([]models.Session, 0)
	err := sr.DB.Find(&sessions).Where("userID = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (sr *sessionRepo) GetSessionByTokenUUID(id uuid.UUID) (*models.Session, error) {
	session := new(models.Session)
	err := sr.DB.First(&session).Where("id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (sr *sessionRepo) expiredSessionRemover() {
	log.Println("removing expired session started")
	sessionRecords, err := sr.GetAllRecords()
	if err != nil {
		log.Println("session record getting error: ", err)
		return
	}

	for _, sessionRecord := range sessionRecords {
		if sessionRecord.ExpiresAt.Before(time.Now()) || sessionRecord.IsBlocked {
			if err = sr.DeleteRecord(sessionRecord); err != nil {
				log.Println("expired session deletion error: ", err)
			}
		}
	}
}
