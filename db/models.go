package db

import (
	"time"
)


// модель для ссылок
type Link struct {
	ID			  uint 		 `gorm:"primarykey"					  	   json:"id"`
	ExternalLink  string	 `gorm:"type:text; not null; uniqueIndex"  json:"externalLink"`
	RedirectLink  string	 `gorm:"type:text; not null; uniqueIndex"  json:"redirectLink"`
	CreatedAt	  time.Time	 `gorm:"not null; autoCreateTime"	  	   json:"createdAt"`
}
