package mysql

import (
	"database/sql"
	"snippetbox/pkg/models"
)

// Model which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (snippetModel *SnippetModel) Insert(title string, content string, expires string) (int, error) {
	return 0, nil
}

// Return a snippet based on ID.
func (snippetModel *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Return the 10 most recently created snippets.
func (snippetModel *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
