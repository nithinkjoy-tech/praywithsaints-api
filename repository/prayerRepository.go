package repository

import (
	"prayer-book/config"
	"prayer-book/models"

	"github.com/lib/pq"
)

func GetAllPrayers() ([]models.PrayerStruct, error) {
	var prayers []models.PrayerStruct
	query := "select * from prayers where is_approved=true"
	rows, err := config.DB.Query(query)
	if err != nil {
		return []models.PrayerStruct{}, nil
	}
	defer rows.Close()

	for rows.Next() {
		var prayer models.PrayerStruct
		err = rows.Scan(
			&prayer.PrayerID,
			&prayer.PrayerName,
			&prayer.Prayer,
			&prayer.HowtoPray,
			&prayer.Impact,
			pq.Array(&prayer.TalkLinks),
			&prayer.AuthorName,
			&prayer.AuthorLink,
			&prayer.ApprovedBy,
			pq.Array(&prayer.Testimonies),
			pq.Array(&prayer.References),
			&prayer.IsApproved,
		)

		if err != nil {
			return []models.PrayerStruct{}, nil
		}

		prayers = append(prayers, prayer)
	}

	return prayers, nil

}

func GetPrayersByID(id int) (models.PrayerStruct, error) {
	query := "select * from prayers where prayer_id=$1"

	var prayer models.PrayerStruct
	err := config.DB.QueryRow(query, id).Scan(
		&prayer.PrayerID,
		&prayer.PrayerName,
		&prayer.Prayer,
		&prayer.HowtoPray,
		&prayer.Impact,
		pq.Array(&prayer.TalkLinks),
		&prayer.AuthorName,
		&prayer.AuthorLink,
		&prayer.ApprovedBy,
		pq.Array(&prayer.Testimonies),
		pq.Array(&prayer.References),
		&prayer.IsApproved,
	)

	if err != nil {
		return models.PrayerStruct{}, err
	}

	return prayer, nil
}

func InsertPrayer(prayer models.PrayerStruct) (int, error) {
	query := `
        INSERT INTO prayers (
            prayer_name,
            prayer,
            how_to_pray,
            impact,
            talk_links,
            author_name,
            author_link,
			approved_by,
            testimonies,
            "references"
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING prayer_id
    `

	var id int
	err := config.DB.QueryRow(query,
		prayer.PrayerName,
		prayer.Prayer,
		prayer.HowtoPray,
		prayer.Impact,
		pq.Array(prayer.TalkLinks),
		prayer.AuthorName,
		prayer.AuthorLink,
		prayer.ApprovedBy,
		pq.Array(prayer.Testimonies),
		pq.Array(prayer.References),
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}
