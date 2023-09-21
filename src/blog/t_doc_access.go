package blog

import (
	"time"

	"gorm.io/gorm"
)

// 数据库 - 单条文章
type T_doc_access struct {
	// gorm.Model
	Id          int       `gorm:"primary_key"`
	Create_Time time.Time `gorm:"autoCreateTime;"`
	Url         string    `gorm:"not null;"`
	Num         int       `gorm:"not null;"`
	Status      int       `gorm:"not null;"`
}

func (T_doc_access) TableName() string {
	return "t_doc_access"
}

func (t *T_doc_access) getPv(db *gorm.DB, url string) int {
	var num int = 0
	result := db.Debug().First(&t, "url = ?", url)

	if result.RowsAffected > 0 {
		num = t.Num
		t.Num += 1
		db.Debug().Model(&t).Update("num", t.Num)

		return num
	} else {
		t.Num = 1
		t.Url = url
		db.Debug().Create(&t)
		return 1
	}
}
