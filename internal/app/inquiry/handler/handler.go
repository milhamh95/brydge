package handler

import "context"

type request struct {
	TrxId          string `json:"trxId"`
	Origin         string `json:"origin"`
	CustomerNumber string `json:"customerNumber"`
	Msisdn         string `json:"msisdn"`
	OrderCategory  string `json:"orderCategory"`
	Destination    string `json:"destination"`
	PaymentAmount  int    `json:"paymentAmount"`
	TrxDate        string `json:"trxDate"`
	Currency       string `json:"currency"`
	Merchant       string `json:"merchant"`
	Terminal       string `json:"terminal"`
	MerchantUsr    string `json:"merchantUsr"`
	MerchantPwd    string `json:"merchantPwd"`
	Topic          string `json:"topic"` // general, specific, mfs
	Product        string `json:"product"`
	Shortcode      string `json:"shortcode"`
	DenomId        string `json:"denomId"`
	DenomName      string `json:"denomName,omitempty"`
	DenomCode      string `json:"denomCode,omitempty"`
	DenomAmount    int    `json:"denomAmount,omitempty"`

	InquiryUrl string      `json:"inquiryUrl"`
	ExtValue   []extraData `json:"extraData"`
}

type response struct {
	TrxId                string      `json:"trxId"`
	Origin               string      `json:"origin"`
	Destination          string      `json:"destination"`
	DestinationAddress   string      `json:"destinationAddress,omitempty"`
	DestinationData      string      `json:"destinationData,omitempty"`
	DestinationReference string      `json:"destinationReference"`
	PaymentAmount        int         `json:"paymentAmount"`
	DenomAmount          interface{} `json:"denomAmount,omitempty"`

	ExtraData []extraData `json:"extraData,omitempty"`
}

// ExtraData is an extra data for smartbiller
type extraData struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	IsCopy bool   `json:"isCopy,omitempty"`
}

type useCase interface {
	Inquiry(ctx context.Context)
}
