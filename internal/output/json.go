package output

import (
	"encoding/json"
	"os"

	"github.com/WesleyTran0/newsfeed/pkg/models"
)

// WriteJSON converts all articles into JSON and writes to the file located at path.
// If there was an error at any point in writing this JSON, or writing to the file,
// this function will terminate and write return the error
func WriteJSON(articles []models.Article, path string) error {
	finalJSON, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, finalJSON, 0o644); err != nil {
		return err
	}
	return nil
}
