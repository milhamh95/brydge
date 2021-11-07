package handler

import "context"

type request struct {
	TrxId         string `json:"trxId"`
	ProductCode   string `json:"productCode"`
	PaymentAmount int    `json:"paymentAmount"`
}

type response struct {
	ID            string `json:"id"`
	PartnerID     string `json:"partnerId"`
	ProductCode   string `json:"productCode"`
	PaymentAmount int    `json:"paymentAmount"`
}

type useCase interface {
	Inquiry(ctx context.Context)
}
