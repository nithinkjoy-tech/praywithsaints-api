package repository

import (
	"prayer-book/config"
	"prayer-book/models"

	"github.com/lib/pq"
)

func GetAllPrayers() ([]models.Prayer, error) {
	var prayers []models.Prayer
	query := "select * from prayers"
	rows, err := config.DB.Query(query)
	if err != nil {
		return []models.Prayer{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		var prayer models.Prayer
		err = rows.Scan(
			&prayer.PrayerID,
			&prayer.PrayerName,
			&prayer.HowtoPray,
			&prayer.Impact,
			&prayer.TalkLink,
			&prayer.Author,
			&prayer.Testimonies,
			&prayer.References,
		)

		if err != nil {
			return []models.Prayer{}, nil
		}

		prayers = append(prayers, prayer)
	}

	return prayers, nil

}

func InsertPrayer(prayer models.Prayer) (int, error) {
	query := `
        INSERT INTO prayers (
            prayer_name,
            how_to_pray,
            impact,
            talk_link,
            author,
            testimonies,
            "references"
        ) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING prayer_id
    `

	var id int
	err := config.DB.QueryRow(query,
		prayer.PrayerName,
		prayer.HowtoPray,
		prayer.Impact,
		prayer.TalkLink,
		prayer.Author,
		pq.Array(prayer.Testimonies),
		pq.Array(prayer.References),
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
