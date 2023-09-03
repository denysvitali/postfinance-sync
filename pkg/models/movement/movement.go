package movement

type Address struct {
	Addition   string `json:"addition,omitempty"`
	CountryIso int    `json:"countryIso,omitempty"`
	Language   string `json:"language,omitempty"`
	Location   string `json:"location"`
	PostalCode string `json:"postalCode,omitempty"`
	ShowInMaps bool   `json:"showInMaps"`
	StreetName string `json:"streetName,omitempty"`
	StreetNr   string `json:"streetNr,omitempty"`
}

type BookingAmount struct {
	Amount   float64 `json:"amount"`
	Currency int     `json:"currency"`
}

type IdLiteral struct {
	ID      int    `json:"id"`
	Literal string `json:"literal"`
}

type Movement struct {
	Actions                      []string      `json:"actions"`
	Address                      *Address      `json:"address,omitempty"`
	BancsUniqueSequenceNumber    string        `json:"bancsUniqueSequenceNumber,omitempty"`
	BookingAdditionalInformation string        `json:"bookingAdditionalInformation,omitempty"`
	BookingAmount                BookingAmount `json:"bookingAmount"`
	BookingDate                  string        `json:"bookingDate"`
	BookingRequestID             string        `json:"bookingRequestId,omitempty"`
	BookingType                  int           `json:"bookingType"`
	BusinessEventID              string        `json:"businessEventId"`
	CardNumber                   string        `json:"cardNumber,omitempty"`
	CounterAccount               string        `json:"counterAccount,omitempty"`
	CounterAccountName           string        `json:"counterAccountName,omitempty"`
	DLevelInformation            int           `json:"dLevelInformation,omitempty"`
	Date                         string        `json:"date"`
	EodBalance                   float64       `json:"eodBalance,omitempty"`
	ForeignExchange              struct {
		BookingExchangeRate float64 `json:"bookingExchangeRate,omitempty"`
	} `json:"foreignExchange"`
	InstrumentID         string     `json:"instrumentId,omitempty"`
	Investigation        string     `json:"investigation,omitempty"`
	LongBookingText      string     `json:"longBookingText"`
	MainCategory         int        `json:"mainCategory,omitempty"`
	Messages             string     `json:"messages,omitempty"`
	MovementForm         *IdLiteral `json:"movementForm,omitempty"`
	MovementID           string     `json:"movementId"`
	MovementType         *IdLiteral `json:"movementType"`
	OrderID              string     `json:"orderId"`
	OrderingSourceSystem int        `json:"orderingSourceSystem"`
	OwnAccount           string     `json:"ownAccount"`
	PaymentType          int        `json:"paymentType"`
	Personalisation      *struct {
		Hashtags            []any `json:"hashtags"`
		PersonalSubCategory int   `json:"personalSubCategory"`
	} `json:"personalisation,omitempty"`
	ProductCharacteristic int    `json:"productCharacteristic,omitempty"`
	ShortBookingText      string `json:"shortBookingText"`
	ShortText             string `json:"shortText"`
	ShortTextID           int    `json:"shortTextId,omitempty"`
	Storno                bool   `json:"storno"`
	SubCategory           int    `json:"subCategory,omitempty"`
	TransactionDate       string `json:"transactionDate,omitempty"`
	TransactionID         string `json:"transactionId,omitempty"`
	URILogo               string `json:"uriLogo,omitempty"`
	URILogoCategory       string `json:"uriLogoCategory,omitempty"`
	ValutaDate            string `json:"valutaDate,omitempty"`
}

type Result struct {
	MoreRecords       bool       `json:"moreRecords"`
	Movements         []Movement `json:"movements"`
	NumberOfMovements int        `json:"numberOfMovements"`
	RestartKey        string     `json:"restartKey"`
	SearchParams      struct {
		AccountList      []string `json:"accountList"`
		MaximumDateRange struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"maximumDateRange"`
	} `json:"searchParams"`
	SelectedValues struct {
		AccountList   []string `json:"accountList"`
		QueryDateFrom string   `json:"queryDateFrom"`
		QueryDateTo   string   `json:"queryDateTo"`
	} `json:"selectedValues"`
}

type Response struct {
	Result Result `json:"result"`
}
