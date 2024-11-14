package models

type Prayer struct {
	PrayerID    uint     `json:"prayer_id"`
	PrayerName  string   `json:"prayer_name" binding:"required,max=255"`
	HowtoPray   string   `json:"how_to_pray" binding:"required"`
	Impact      string   `json:"impact" binding:"required"`
	TalkLink    string   `json:"talk_link" binding:"omitempty,url"`
	Author      string   `json:"author" binding:"required,max=255"`
	Testimonies []string `json:"testimonies" binding:"required,dive,required"`
	References  []string `json:"references" binding:"required,dive,required"`
}
