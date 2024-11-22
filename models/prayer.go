package models

type PrayerStruct struct {
	PrayerID    uint     `json:"prayer_id"`
	PrayerName  string   `json:"prayer_name" binding:"required,max=50"`
	Prayer      string   `json:"prayer" binding:"required,max=100000"`
	HowtoPray   string   `json:"how_to_pray" binding:"max=100000"`
	Impact      string   `json:"impact" binding:"max=100000"`
	TalkLinks   []string `json:"talk_links" binding:"omitempty,dive,url,max=2048"`
	AuthorName  string   `json:"author_name" binding:"required,max=50"`
	AuthorLink  string   `json:"author_link" binding:"omitempty,url,max=2048"`
	ApprovedBy  string   `json:"approved_by" binding:"required,max=50"`
	Testimonies []string `json:"testimonies" binding:"required,dive,max=1000"`
	References  []string `json:"references" binding:"required,dive,url,max=2048"`
	IsApproved  bool     `json:"is_approved" binding:"omitempty"`
}
