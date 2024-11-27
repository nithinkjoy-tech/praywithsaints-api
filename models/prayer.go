package models

type PrayerStruct struct {
	PrayerID    uint     `json:"prayer_id"`
	PrayerName  string   `json:"prayer_name" binding:"required,min=3,max=255"`
	Prayer      string   `json:"prayer" binding:"required,min=3,max=100000"`
	HowtoPray   string   `json:"how_to_pray" binding:"omitempty,min=3,max=100000"`
	Impact      string   `json:"impact" binding:"min=3,max=100000"`
	TalkLinks   []string `json:"talk_links" binding:"omitempty,dive,url,min=3,max=2048"`
	AuthorName  string   `json:"author_name" binding:"required,min=2,max=255"`
	AuthorLink  string   `json:"author_link" binding:"omitempty,url,min=3,max=2048"`
	ApprovedBy  string   `json:"approved_by" binding:"required,min=2,max=255"`
	Testimonies []string `json:"testimonies" binding:"required,dive,min=10,max=1000"`
	References  []string `json:"references" binding:"required,dive,url,min=3,max=2048"`
	IsApproved  bool     `json:"is_approved" binding:"omitempty"`
}
