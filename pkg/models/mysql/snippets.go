package mysql

import (
	"database/sql"
	"snippetbox/pkg/models"
)

// SnippetModel Model which wraps a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (snippetModel *SnippetModel) Insert(title string, content string, expires string) (int, error) {
	sqlStatement := `INSERT INTO snippets (title, content, created, expires) 
			         VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := snippetModel.DB.Exec(sqlStatement, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// ID return is int64.
	return int(id), nil
}

// Get Return a snippet based on ID.
func (snippetModel *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest Return the 10 most recently created snippets.
func (snippetModel *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
