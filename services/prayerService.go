package services

import (
	"prayer-book/models"
	"prayer-book/repository"
)

func GetPrayersService() ([]models.Prayer, error) {

	prayers, err := repository.GetAllPrayers()
	if err != nil {
		return []models.Prayer{}, err
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
