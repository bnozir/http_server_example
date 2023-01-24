package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bnozir/todoes/internal/models"
)

type DBMock struct{}

func (dbm *DBMock) Connect() (*sql.Conn, error) {
	return nil, nil
}

func (dbm *DBMock) Disconnect() error {
	return nil
}

func (dbm *DBMock) FindTodoByID(db *sql.Conn, id int, todo *models.Todoes) (bool, error) {
	var (
		notFound = true
		err      error
	)
	if id == 1 {
		notFound = false
		startTime := time.Now().Add(time.Hour * 12)
		expTime := startTime.Add(time.Hour * 48)
		*todo = models.Todoes{
			ID:             1,
			Name:           "Bala todo",
			StartTime:      &startTime,
			ExpirationTime: &expTime,
		}
	} else if id == -1 {
		notFound = true
		err = fmt.Errorf("wrong id value")
	}

	return notFound, err
}

func (dbm *DBMock) FindTodoByFields(db *sql.Conn, fields map[string]interface{}, todo *models.Todoes) (bool, error) {
	var (
		notFound = true
		err      error
	)
	if fields["id"].(int) == 1 && fields["name"].(string) == "Bala todo" {
		notFound = false
		startTime := time.Now().Add(time.Hour * 12)
		expTime := startTime.Add(time.Hour * 48)
		*todo = models.Todoes{
			ID:             1,
			Name:           "Bala todo",
			StartTime:      &startTime,
			ExpirationTime: &expTime,
		}
	} else if fields["id"].(int) == -1 {
		notFound = true
		err = fmt.Errorf("wrong id value")
	}

	return notFound, err
}
