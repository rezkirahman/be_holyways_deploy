package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetTransactionByCharity(ID int) ([]models.Transaction, error)
	GetTransactionPendingByCharity(ID int) ([]models.Transaction, error)
	GetTransactionByUser(ID int) ([]models.Transaction, error)
	DeleteTransaction(ID int, order models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) (error)
	GetOneTransaction(ID string) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.db.Preload("Funder").Preload("Fundraiser").Preload("Charity.User").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) GetTransactionByCharity(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Where("charity_id=?", ID).Where("status='success'").Preload("Fundraiser").Preload("Funder").Preload("Charity.User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionPendingByCharity(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Where("charity_id=?", ID).Where("status!='success'").Preload("Fundraiser").Preload("Funder").Preload("Charity.User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByUser(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("funder_id=?", ID).Preload("Charity").Find(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(ID int, transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Exec("SET FOREIGN_KEY_CHECKS=0;").Raw("DELETE FROM transaction WHERE id=?", ID).Scan(&transaction).Error
	return transaction, err
}

// Create UpdateTransaction method here ...
func (r *repository) UpdateTransaction(status string, ID string) (error) {
	var transaction models.Transaction
	r.db.Preload("Charity").First(&transaction, ID)


	//update total donation on 
	if status != transaction.Status && status == "success" {
		var charity models.Charity
		r.db.First(&charity, transaction.Charity.ID)
		charity.Donation += transaction.Donation
		r.db.Save(&charity)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

// GetOneTransaction
func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Charity").Preload("Charity.User").Preload("Funder").Preload("Fundraiser").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
