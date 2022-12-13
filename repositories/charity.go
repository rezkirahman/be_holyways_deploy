package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type CharityRepository interface {
	FindCharities() ([]models.Charity, error)
	CreateCharity(charity models.Charity) (models.Charity, error)
	GetCharity(ID int) (models.Charity, error)
	GetCharityByUser(ID int) ([]models.Charity, error)
	UpdateCharity(Charity models.Charity, ID int) (models.Charity, error)
	DeleteCharity(ID int, charity models.Charity) (models.Charity, error)
}

func RepositoryCharity(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCharities() ([]models.Charity, error) {
	var charity []models.Charity
	err := r.db.Where("status = 'Running'").Preload("User").Preload("Funder").Find(&charity).Error
	return charity, err
}

func (r *repository) CreateCharity(charity models.Charity) (models.Charity, error) {
	err := r.db.Debug().Create(&charity).Error
	return charity, err
}

func (r *repository) GetCharity(ID int) (models.Charity, error) {
	var charity models.Charity
	err := r.db.Preload("User").Preload("Funder").First(&charity, ID).Error
	return charity, err
}

func (r *repository) GetCharityByUser(ID int) ([]models.Charity, error) {
	var userCharity []models.Charity
	err := r.db.Debug().Where("user_id=?", ID).Preload("User").Find(&userCharity).Error
	return userCharity, err
}

func (r *repository) UpdateCharity(charity models.Charity, ID int) (models.Charity, error) {
	err := r.db.Save(&charity).Error
	return charity, err
}

func (r *repository) DeleteCharity(ID int, charity models.Charity) (models.Charity, error) {
	err := r.db.Exec("SET FOREIGN_KEY_CHECKS=0;").Raw("DELETE FROM charities WHERE id=?", ID).Scan(&charity).Error

	return charity, err
}
