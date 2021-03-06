package repositories

import (
	"database/sql"

	"github.com/jacobtie/learning-resource-tracker/pkg/models"
)

// TagRepository ...
type TagRepository struct {
	DB *sql.DB
}

// GetCategoryTags ...
func (tr *TagRepository) GetCategoryTags(categoryID int) (models.Tags, error) {
	rows, err := tr.DB.Query(`
							SELECT
								ta.tag_id,
								ta.title
							FROM
								Tag ta
								INNER JOIN CategoryTag ct USING (tag_id)
							WHERE ct.category_id = ?
							;
							`, categoryID)
	if err != nil {
		return nil, err
	}

	tags := make(models.Tags, 0)
	for rows.Next() {
		var tag models.Tag
		err = rows.Scan(&tag.TagID, &tag.Title)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

// GetTopicTags ...
func (tr *TagRepository) GetTopicTags(topicID int) (models.Tags, error) {
	rows, err := tr.DB.Query(`
							SELECT
								ta.tag_id,
								ta.title
							FROM
								Tag ta
								INNER JOIN TopicTag tt USING (tag_id)
							WHERE tt.topic_id = ?
							;
							`, topicID)
	if err != nil {
		return nil, err
	}

	tags := make(models.Tags, 0)
	for rows.Next() {
		var tag models.Tag
		err = rows.Scan(&tag.TagID, &tag.Title)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

// GetResourceTags ...
func (tr *TagRepository) GetResourceTags(resourceID int) (models.Tags, error) {
	rows, err := tr.DB.Query(`
							SELECT
								ta.tag_id,
								ta.title
							FROM
								Tag ta
								INNER JOIN ResourceTag rt USING (tag_id)
							WHERE rt.resource_item_id = ?
							;
							`, resourceID)
	if err != nil {
		return nil, err
	}

	tags := make(models.Tags, 0)
	for rows.Next() {
		var tag models.Tag
		err = rows.Scan(&tag.TagID, &tag.Title)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
