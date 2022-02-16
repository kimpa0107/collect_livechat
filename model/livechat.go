package model

import "time"

type LiveChat struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Platform  string    `gorm:"type:varchar(20)" json:"platform"`
	Room      string    `gorm:"type:varchar(50)" json:"room"`
	Date      string    `gorm:"type:date" json:"date"`
	Title     string    `gorm:"type:varchar(100)" json:"title"`
	Nickname  string    `gorm:"type:varchar(50)" json:"nick"`
	Message   string    `gorm:"type:text" json:"text"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

type DouyinLiveChat struct {
	LiveChat
}

func (DouyinLiveChat) TableName() string {
	return "livechat_douyin"
}

type KuaishouLiveChat struct {
	LiveChat
}

func (KuaishouLiveChat) TableName() string {
	return "livechat_kuaishou"
}

type DouyuLiveChat struct {
	LiveChat
}

func (DouyuLiveChat) TableName() string {
	return "livechat_douyu"
}

type MiguLiveChat struct {
	LiveChat
}

func (MiguLiveChat) TableName() string {
	return "livechat_migu"
}

type AfreecatvLiveChat struct {
	LiveChat
}

func (AfreecatvLiveChat) TableName() string {
	return "livechat_afreecatv"
}

type PandatvLiveChat struct {
	LiveChat
}

func (PandatvLiveChat) TableName() string {
	return "livechat_pandatv"
}

type FlextvLiveChat struct {
	LiveChat
}

func (FlextvLiveChat) TableName() string {
	return "livechat_flextv"
}
