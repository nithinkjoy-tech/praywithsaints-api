package services

import (
	"prayer-book/models"
	"prayer-book/repository"
)

func GetPrayers() ([]models.Prayer, error) {
	prayers, err := repository.GetAllPrayers()
	if err != nil {
		return []models.Prayer{}, err
	}

	return prayers, nil
}

func GetPrayersByID(id int) (models.Prayer, error) {
	prayers, err := repository.GetPrayersByID(id)
	if err != nil {
		return models.Prayer{}, err
	}

	return prayers, nil
}

func InsertPrayer(prayer models.Prayer) (int, error) {
	id, err := repository.InsertPrayer(prayer)
	if err != nil {
		return 0, err
	}
	return id, nil
}
