package db

import (
	"database/sql"

	"github.com/bnozir/todoes/internal/models"
)

type Connector interface {
	Connect() (*sql.Conn, error)
	Disconnect() error
}

type Selector interface {
	FindTodoByID(db *sql.Conn, id int, todo *models.Todoes) (bool, error)
	FindTodoByFields(db *sql.Conn, fields map[string]interface{}, todo *models.Todoes) (bool, error)
}
