package transactiondto

type CreateTransactionRequest struct {
	CharityID    int    `json:"charity_id"`
	FunderID     int    `json:"funder_id"`
	FundraiserID int    `json:"fundraiser_id"`
	Status       string `json:"status"`
	Donation     int `json:"donation"`
}
