package models

type PrayerStruct struct {
	PrayerID    uint     `json:"prayer_id"`
	PrayerName  string   `json:"prayer_name" binding:"required,max=255"`
	Prayer      string   `json:"prayer" binding:"required,max=1000000"`
	HowtoPray   string   `json:"how_to_pray"`
	Impact      string   `json:"impact"`
	TalkLinks   []string `json:"talk_links" binding:"omitempty,dive,url"`
	Author      string   `json:"author" binding:"required,max=255"`
	AuthorLink  string   `json:"author_link" binding:"omitempty,url"`
	ApprovedBy  string   `json:"approved_by" binding:"required"`
	Testimonies []string `json:"testimonies" binding:"required,dive,required"`
	References  []string `json:"references" binding:"required,dive,required"`
}
