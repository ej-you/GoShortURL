package db

import (
	"time"
)


// модель для ссылок
type Link struct {
	ID			  uint 		 `gorm:"primarykey"					  			  json:"id"`
	ExternalLink  string	 `gorm:"type:text; not null; uniqueIndex"		  json:"externalLink"`
	LinkAlias	  string	 `gorm:"type:varchar(50); not null; uniqueIndex"  json:"linkAlias"`
	CreatedAt	  time.Time	 `gorm:"not null; autoCreateTime"	  	    	  json:"createdAt"`
}
