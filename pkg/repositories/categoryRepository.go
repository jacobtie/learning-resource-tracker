package repositories

import (
	"database/sql"

	"github.com/jacobtie/learning-resource-tracker/pkg/models"
)

// CategoryRepository ...
type CategoryRepository struct {
	DB *sql.DB
}

// GetTopCategories ...
func (cr *CategoryRepository) GetTopCategories() (models.Categories, error) {
	rows, err := cr.DB.Query(`
							SELECT
								c.category_id,
								c.label
							FROM
								Category c
								INNER JOIN State s USING (state_id)
							WHERE c.parent_id IS NULL AND s.label <> 'Deleted'
							;
							`)
	if err != nil {
		return nil, err
	}

	categories := make(models.Categories, 0)
	for rows.Next() {
		var category models.Category
		err = rows.Scan(
			&category.CategoryID,
			&category.Label,
		)
		if err != nil {
			return nil, err
		}
		category.Depth = 1
		categories = append(categories, category)
	}

	return categories, nil
}

// GetSubCategories ...
func (cr *CategoryRepository) GetSubCategories(parentID int, depth int) (models.Categories, error) {
	rows, err := cr.DB.Query(`
							SELECT
								c.category_id,
								c.label,
								c.parent_id
							FROM
								Category c
								INNER JOIN State s USING (state_id)
							WHERE s.label <> 'Deleted' AND c.parent_id = ?
							;
							`, parentID)
	if err != nil {
		return nil, err
	}

	subCategories := make(models.Categories, 0)
	for rows.Next() {
		var category models.Category
		err = rows.Scan(
			&category.CategoryID,
			&category.Label,
			&category.ParentID,
		)
		if err != nil {
			return nil, err
		}
		category.Depth = depth
		subCategories = append(subCategories, category)
	}

	return subCategories, nil
}
