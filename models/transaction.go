package models

type Transaction struct {
	ID           int          `json:"id" gorm:"primary_key:auto_increment"`
	CharityID    int          `json:"charity_id" gorm:"type:int"`
	Charity      Charity      `json:"charity"`
	FundraiserID int          `json:"fundraiser_id" gorm:"type:int"`
	Fundraiser   UserResponse `json:"fundraiser" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FunderID     int          `json:"funder_id" gorm:"type:int"`
	Funder       UserResponse `json:"funder" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status       string       `json:"status" gorm:"type:varchar(255)"`
	Donation     int          `json:"donation" gorm:"type:int"`
}
