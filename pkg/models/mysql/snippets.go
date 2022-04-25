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
	sqlStatement := `SELECT id, title, content, created, expires FROM snippets
					 WHERE id = ? AND expires > UTC_TIMESTAMP()`

	// Query the Database with `id` as the placeholder
	row := snippetModel.DB.QueryRow(sqlStatement, id)

	// This is where we will store the result
	result := &models.Snippet{}

	// Scan the result and store it into s
	err := row.Scan(&result.ID, &result.Title, &result.Content, &result.Created, &result.Expires)

	// If we did not find that entry in the DB
	if err == sql.ErrNoRows {
		return nil, models.ErrorNoRecord // We have the appropriate error
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

// Latest Return the 10 most recently created snippets.
func (snippetModel *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
