package model

type Episode struct {
	Id        string `xorm:"pk varchar(64)" form:"id" json:"id"`
	NovelId   string `xorm:"varchar(64)" form:"novelId" json:"novelId"`
	Title     string `xorm:"varchar(255)" form:"title" json:"title"`
	Content   string `xorm:"text" form:"content" json:"content"`
	CreatedAt string `xorm:"varchar(64)" form:"createdAt" json:"createdAt"`
	UpdatedAt string `xorm:"varchar(64)" form:"updatedAt" json:"updatedAt"`
}

type Chapter struct {
	Id       string    `xorm:"pk varchar(64)" form:"id" json:"id"`
	NovelId  string    `xorm:"varchar(64)" form:"novelId" json:"novelId"`
	Title    string    `xorm:"varchar(255)" form:"title" json:"title"`
	Episodes []Episode `xorm:"json" form:"episodes" json:"episodes"`
}

type Novel struct {
	Id        string    `xorm:"pk varchar(64)" form:"id" json:"id"`
	Title     string    `xorm:"varchar(255)" form:"title" json:"title"`
	Author    string    `xorm:"varchar(255)" form:"author" json:"author"`
	Category  string    `xorm:"varchar(255)" form:"category" json:"category"`
	Synopsis  string    `xorm:"text" form:"synopsis" json:"synopsis"`
	Image     string    `xorm:"varchar(255)" form:"image" json:"image"`
	CreatedAt string    `xorm:"varchar(64)" form:"createdAt" json:"createdAt"`
	UpdatedAt string    `xorm:"varchar(64)" form:"updatedAt" json:"updatedAt"`
	Chapters  []Chapter `xorm:"json" form:"chapters" json:"chapters"`
}
