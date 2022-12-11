package models

type Charity struct {
	ID          int           `json:"id" gorm:"primary_key:auto_increment"`
	Title       string        `json:"title" form:"name" gorm:"type:varchar(255)"`
	Image       string        `json:"image" form:"image" gorm:"type:varchar(255)"`
	Goal        int           `json:"goal" form:"goal" gorm:"type:int"`
	Donation    int           `json:"donation" form:"donation" gorm:"type:int"`
	Status      string        `json:"status" gorm:"type: varchar(255)"`
	Description string        `json:"description" gorm:"type: varchar(255)"`
	UserID      int           `json:"user_id" gorm:"type: int"`
	User        UserResponse  `json:"fundraiser"`
	Funder      []Transaction `json:"funder"`
}
