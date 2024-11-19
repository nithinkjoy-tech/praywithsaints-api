package services

import (
	"prayer-book/models"
	"prayer-book/repository"
)

func GetPrayers() ([]models.PrayerStruct, error) {
	prayers, err := repository.GetAllPrayers()
	if err != nil {
		return []models.PrayerStruct{}, err
	}

	return prayers, nil
}

func GetPrayersByID(id int) (models.PrayerStruct, error) {
	prayers, err := repository.GetPrayersByID(id)
	if err != nil {
		return models.PrayerStruct{}, err
	}

	return prayers, nil
}

func InsertPrayer(prayer models.PrayerStruct) (int, error) {
	id, err := repository.InsertPrayer(prayer)
	if err != nil {
		return 0, err
	}
	return id, nil
}
