package models

type Prayer struct {
	PrayerID    uint     `json:"prayer_id"`
	PrayerName  string   `json:"prayer_name"`
	HowtoPray   string   `json:"how_to_pray"`
	Impact      string   `json:"impact"`
	TalkLink    string   `json:"talk_link"`
	Author      string   `json:"author"`
	Testimonies []string `json:"testimonies"`
	References  []string `json:"references"`
}
