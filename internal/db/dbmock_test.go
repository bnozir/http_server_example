package db_test

import (
	"testing"

	"github.com/bnozir/todoes/internal/db"
	"github.com/bnozir/todoes/internal/models"
)

func TestDBMockFindTodoByID(t *testing.T) {
	var dbMock *db.DBMock

	_, err := dbMock.Connect()
	if err != nil {
		t.Fatalf("failed on connect: %s", err)
	}
	defer dbMock.Disconnect()

	todo := &models.Todoes{}
	todoID := 1
	todoName := "Bala todo"
	notFound, err := dbMock.FindTodoByID(nil, todoID, todo)
	if err != nil {
		t.Fatalf("failed on FindTodoByID: %s", err)
	}
	if notFound {
		t.Logf("todo not found by id: %d", todoID)
	}
	if todo.ID != todoID && todo.Name != todoName {
		t.Fatalf("wrong todo data: expected: id: %d, name: %s; got: id: %d, name: %s", todoID, todoName, todo.ID, todo.Name)
	}

	_, err = dbMock.FindTodoByID(nil, -1, todo)
	if err == nil {
		t.Fatalf("not failed on FindTodoByID: expected: error, got: %s", err)
	}
	if err != nil {
		t.Logf("error ecurred as expected: %s", err)
	}
}
