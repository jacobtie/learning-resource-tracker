package repositories

import (
	"database/sql"

	"github.com/jacobtie/learning-resource-tracker/pkg/models"
)

// TopicRepository ...
type TopicRepository struct {
	DB *sql.DB
}

// GetTopTopics ...
func (tr *TopicRepository) GetTopTopics() (models.Topics, error) {
	rows, err := tr.DB.Query(`
							SELECT
								t.topic_id,
								t.title,
								t.summary,
								s.label,
								p.label,
								p.worth
							FROM
								Topic t
								INNER JOIN State s USING (state_id)
								INNER JOIN Priority p USING (priority_id)
							WHERE t.category_id IS NULL AND s.label <> 'Deleted'
							;
							`)
	if err != nil {
		return nil, err
	}

	topics := make(models.Topics, 0)
	for rows.Next() {
		var topic models.Topic
		err = rows.Scan(
			&topic.TopicID,
			&topic.Title,
			&topic.Summary,
			&topic.State,
			&topic.Priority.Label,
			&topic.Priority.Worth,
		)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}

	return topics, nil
}

// GetCategoryTopics ...
func (tr *TopicRepository) GetCategoryTopics(categoryID int) (models.Topics, error) {
	rows, err := tr.DB.Query(`
							SELECT
								t.topic_id,
								t.title,
								t.summary,
								s.label,
								p.label,
								p.worth
							FROM
								Topic t
								INNER JOIN State s USING (state_id)
								INNER JOIN Priority p USING (priority_id)
							WHERE t.category_id = ? AND s.label <> 'Deleted'
							;
							`, categoryID)
	if err != nil {
		return nil, err
	}

	topics := make(models.Topics, 0)
	for rows.Next() {
		var topic models.Topic
		err = rows.Scan(
			&topic.TopicID,
			&topic.Title,
			&topic.Summary,
			&topic.State,
			&topic.Priority.Label,
			&topic.Priority.Worth,
		)
		if err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}

	return topics, nil
}
