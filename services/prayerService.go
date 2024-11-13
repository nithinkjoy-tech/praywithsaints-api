package services

import "prayer-book/repository"

func GetPrayersService() {

	_, err := repository.GetAllPrayers()
	if err != nil {
		return err, nil
	}
}
