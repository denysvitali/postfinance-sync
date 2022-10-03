package models

type Response[T any] struct {
	Result T `json:"result"`
}

type SearchParams struct {
	Accounts []struct {
		ProductUniqueKey string `json:"productUniqueKey"`
	} `json:"accounts"`
	MinDate string `json:"minDate"`
	MaxDate string `json:"maxDate"`
}

type SelectedParams struct {
	DateFrom string `json:"dateFrom"`
	DateTo   string `json:"dateTo"`
}

type Amount struct {
	Amount   float64 `json:"amount"`
	Currency int     `json:"currency"`
}

type Categorization struct {
	MainCategoryId int `json:"mainCategoryId,omitempty"`
	IndustryId     int `json:"industryId,omitempty"`
}

type Address struct {
	City       string `json:"city"`
	Street     string `json:"street,omitempty"`
	Building   string `json:"building,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}

type Location struct {
	Address Address `json:"address"`
	Country int     `json:"country"`
}

type Partner struct {
	Name            string `json:"name"`
	MerchantLogoUrl string `json:"merchantLogoUrl,omitempty"`
	Account         string `json:"account,omitempty"`
}

type Booking struct {
	Actions               []string       `json:"actions"`
	Type                  int            `json:"type"`
	BancsUniqueSeqNr      string         `json:"bancsUniqueSeqNr"`
	BookingRequestId      string         `json:"bookingRequestId"`
	Date                  string         `json:"date"`
	ValueDate             string         `json:"valueDate"`
	BookingType           int            `json:"bookingType"`
	Storno                bool           `json:"storno"`
	Debit                 Amount         `json:"debit,omitempty"`
	Balance               Amount         `json:"balance,omitempty"`
	ShortText             string         `json:"shorttext"`
	Text                  string         `json:"text"`
	CompiledShortText     string         `json:"compiledShortText"`
	Categorisation        Categorization `json:"categorisation"`
	Location              Location       `json:"location,omitempty"`
	ProductCharacteristic int            `json:"productCharacteristic,omitempty"`
	CardNumber            string         `json:"cardNumber,omitempty"`
	TransactionPartner    Partner        `json:"transactionPartner,omitempty"`
	Credit                Amount         `json:"credit,omitempty"`
	Investigation         string         `json:"investigation,omitempty"`
}

type BookingResult struct {
	SearchParams   SearchParams   `json:"searchParams"`
	SelectedParams SelectedParams `json:"selectedParams"`
	ForwardId      string         `json:"forwardId"`
	Bookings       []Booking      `json:"bookings"`
}
