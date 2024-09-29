package repository

import (
	"github.com/thiccpan/library_information_system/internal/entity"
	"gorm.io/gorm"
)

type TopicRepository struct {
	Db *gorm.DB
}

func NewTopicRepoImpl(db *gorm.DB) *TopicRepository {
	return &TopicRepository{
		Db: db,
	}
}

func (t *TopicRepository) Add(db *gorm.DB, topic *entity.Topic) error {
	tx := db.Create(topic)
	return tx.Error
}

func (t *TopicRepository) GetAll(db *gorm.DB) ([]entity.Topic, error) {
	var topics []entity.Topic
	tx := db.Find(&topics)
	if err := tx.Error; err != nil {
		return topics, err
	}
	return topics, nil
}

func (t *TopicRepository) GetById(db *gorm.DB, topic *entity.Topic) error {
	tx := db.First(&topic)
	return tx.Error
}

func (t *TopicRepository) UpdateById(db *gorm.DB, topic *entity.Topic) error {
	tx := db.Model(topic).Where("id = ?", topic.Id).Updates(topic)
	return tx.Error
}

func (t *TopicRepository) Delete(db *gorm.DB, topic *entity.Topic) error {
	tx := db.Delete(topic)
	return tx.Error
}
