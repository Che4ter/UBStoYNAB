package ubsApi

type ChallengeResponse struct {
	Gui struct {
		Label    string `json:"label"`
		Name     string `json:"name"`
		Target   string `json:"target"`
		Language string `json:"language"`
		GuiElem  []struct {
			Name   string `json:"name"`
			Type   string `json:"type"`
			Label  string `json:"label"`
			Value  string `json:"value"`
			Length string `json:"length"`
		} `json:"GuiElem"`
	} `json:"Gui"`
	OperatingInfos          string `json:"OperatingInfos"`
	OperatingInfosEbanking  string `json:"OperatingInfosEbanking"`
	OperatingInfosQuotes    string `json:"OperatingInfosQuotes"`
	OperatingInfosNonClient string `json:"OperatingInfosNonClient"`
	IncidentInfosEbanking   string `json:"IncidentInfosEbanking"`
	IncidentInfosQuotes     string `json:"IncidentInfosQuotes"`
	IncidentInfosNonClient  string `json:"IncidentInfosNonClient"`
	IsMobileKeyDeployed     string `json:"isMobileKeyDeployed"`
}

type PingResponse struct {
	IsMobileKeyAvailable string `json:"isMobileKeyAvailable"`
}

type ChallengeCompleteResponse struct {
	NavigationItems     []string `json:"navigationItems"`
	UserDisplayName     string   `json:"userDisplayName"`
	Privileges          []string `json:"privileges"`
	PushElagOk          bool     `json:"pushElagOk"`
	PushAppleOffering   bool     `json:"pushAppleOffering"`
	PushAndroidOffering bool     `json:"pushAndroidOffering"`
	BookingCenterCd     string   `json:"bookingCenterCd"`
	ClientSegment       string   `json:"clientSegment"`
	LegalEntity         string   `json:"legalEntity"`
	ClientType          string   `json:"clientType"`
	Status              string   `json:"status"`
	Timestamp           string   `json:"timestamp"`
	RequestKey          string   `json:"requestKey"`
	EBanking            string   `json:"eBanking"`
}

type GetAccountsRequest struct {
	StartIdx   int    `json:"startIdx"`
	NumRecords int    `json:"numRecords"`
	SortBy     string `json:"sortBy"`
	Language   string `json:"language"`
	RequestKey string `json:"requestKey"`
}

type GetAccountsResponse struct {
	ValuationDate  string         `json:"valuationDate"`
	HasMoreRecords string         `json:"hasMoreRecords"`
	CashAccounts   []CashAccounts `json:"cashAccounts"`
	Status         string         `json:"status"`
	Timestamp      string         `json:"timestamp"`
}

type CashAccounts struct {
	ID          string   `json:"id"`
	Alias       string   `json:"alias"`
	CurrencyCd  string   `json:"currencyCd"`
	Balance     string   `json:"balance"`
	TrxList     []string `json:"trxList"`
	LastTrxTime string   `json:"lastTrxTime"`
	Intraday    string   `json:"intraday"`
	Fisca       string   `json:"fisca"`
}

type AccountBlockedResponse struct {
	Gui struct {
		Label    string `json:"label"`
		Name     string `json:"name"`
		Target   string `json:"target"`
		Language string `json:"language"`
		GuiElem  []struct {
			Name   string `json:"name"`
			Type   string `json:"type"`
			Label  string `json:"label"`
			Value  string `json:"value"`
			Length string `json:"length"`
		} `json:"GuiElem"`
	} `json:"Gui"`
	OperatingInfos          string `json:"OperatingInfos"`
	OperatingInfosEbanking  string `json:"OperatingInfosEbanking"`
	OperatingInfosQuotes    string `json:"OperatingInfosQuotes"`
	OperatingInfosNonClient string `json:"OperatingInfosNonClient"`
	IncidentInfosEbanking   string `json:"IncidentInfosEbanking"`
	IncidentInfosQuotes     string `json:"IncidentInfosQuotes"`
	IncidentInfosNonClient  string `json:"IncidentInfosNonClient"`
	IsMobileKeyDeployed     string `json:"isMobileKeyDeployed"`
}

type AccountBalanceResponse struct {
	HasMoreRecords string `json:"hasMoreRecords"`
	CashAccount    struct {
		ID          string   `json:"id"`
		Alias       string   `json:"alias"`
		AccountNo   string   `json:"accountNo"`
		Iban        string   `json:"iban"`
		Holder      string   `json:"holder"`
		Bic         string   `json:"bic"`
		CurrencyCd  string   `json:"currencyCd"`
		Balance     string   `json:"balance"`
		TrxList     []string `json:"trxList"`
		LastTrxTime string   `json:"lastTrxTime"`
		Intraday    string   `json:"intraday"`
		Fisca       string   `json:"fisca"`
	} `json:"cashAccount"`
	CashAccountTrxes []CashAccountTrxes `json:"cashAccountTrxes"`
	MinDate          string             `json:"minDate"`
	MaxDate          string             `json:"maxDate"`
	Status           string             `json:"status"`
	Timestamp        string             `json:"timestamp"`
}

type CashAccountTrxes struct {
	Description            string   `json:"description"`
	TrxAmount              string   `json:"trxAmount"`
	Balance                string   `json:"balance"`
	ValueDate              string   `json:"valueDate"`
	TransactionTextList    []string `json:"transactionTextList,omitempty"`
	PaymentInformationList []struct {
		Amount          string   `json:"amount"`
		Currency        string   `json:"currency"`
		DescriptionList []string `json:"descriptionList"`
	} `json:"paymentInformationList,omitempty"`
	PaymentInstructionList []string `json:"paymentInstructionList,omitempty"`
}

type AccountBalanceRequest struct {
	ID          string `json:"id"`
	StartIdx    int    `json:"startIdx"`
	NumRecords  int    `json:"numRecords"`
	ShowDetails bool   `json:"showDetails"`
	Refresh     bool   `json:"refresh"`
	Language    string `json:"language"`
	StartDate   string `json:"startDate"`
	RequestKey  string `json:"requestKey"`
}

type CreditCardOverviewRequest struct {
	Refresh    bool   `json:"refresh"`
	StartIdx   int    `json:"startIdx"`
	NumRecords int    `json:"numRecords"`
	Language   string `json:"language"`
	RequestKey string `json:"requestKey"`
}

type CreditCardOverviewResponse struct {
	HasMoreRecords     string               `json:"hasMoreRecords"`
	HasDebitCard       bool                 `json:"hasDebitCard"`
	NewUnseenInvoices  int                  `json:"newUnseenInvoices"`
	UnpaidInvoices     int                  `json:"unpaidInvoices"`
	CreditCardAccounts []CreditCardAccounts `json:"creditCardAccounts"`
	Status             string               `json:"status"`
	Timestamp          string               `json:"timestamp"`
}

type CreditCardAccounts struct {
	ID                     string `json:"id"`
	Alias                  string `json:"alias"`
	CurrencyCd             string `json:"currencyCd"`
	Balance                string `json:"balance"`
	AvailableBalance       string `json:"availableBalance"`
	Reserved               string `json:"reserved"`
	CurrentEntries         string `json:"currentEntries"`
	OpenStatementAmount    string `json:"openStatementAmount"`
	Liable                 string `json:"liable"`
	PrepayedCard           string `json:"prepayedCard"`
	AccountReferenceNumber string `json:"accountReferenceNumber"`
	SelfRegistered         bool   `json:"selfRegistered"`
	ShuffledIconElements   []struct {
		ProductType string `json:"productType"`
		Status      string `json:"status"`
	} `json:"shuffledIconElements"`
}

type CreditCardDetailsRequest struct {
	ID         string `json:"id"`
	StartIdx   int    `json:"startIdx"`
	NumRecords int    `json:"numRecords"`
	Refresh    bool   `json:"refresh"`
	Language   string `json:"language"`
	RequestKey string `json:"requestKey"`
}

type CreditCardDetailsResponse struct {
	HasMoreRecords    string `json:"hasMoreRecords"`
	HasDebitCard      bool   `json:"hasDebitCard"`
	CreditCardAccount struct {
		NewUnseenInvoice           bool   `json:"newUnseenInvoice"`
		HasInvoice                 bool   `json:"hasInvoice"`
		NewUnseenInvoiceAggregated bool   `json:"newUnseenInvoiceAggregated"`
		HasInvoiceAggregated       bool   `json:"hasInvoiceAggregated"`
		HasHistoricInvoice         bool   `json:"hasHistoricInvoice"`
		HasCurrentInvoice          bool   `json:"hasCurrentInvoice"`
		PeriodStartDate            string `json:"periodStartDate"`
		PeriodEndDate              string `json:"periodEndDate"`
		InvoiceAmount              string `json:"invoiceAmount"`
		ClosingBalance             string `json:"closingBalance"`
		ProductText                string `json:"productText"`
		StartDate                  string `json:"startDate"`
		SpendingLimit              string `json:"spendingLimit"`
		ID                         string `json:"id"`
		Alias                      string `json:"alias"`
		CurrencyCd                 string `json:"currencyCd"`
		Balance                    string `json:"balance"`
		AvailableBalance           string `json:"availableBalance"`
		Reserved                   string `json:"reserved"`
		CurrentEntries             string `json:"currentEntries"`
		OpenStatementAmount        string `json:"openStatementAmount"`
		Liable                     string `json:"liable"`
		PrepayedCard               string `json:"prepayedCard"`
		AccountReferenceNumber     string `json:"accountReferenceNumber"`
		SelfRegistered             bool   `json:"selfRegistered"`
		ShuffledIconElements       []struct {
			ProductType string `json:"productType"`
			Status      string `json:"status"`
		} `json:"shuffledIconElements"`
	} `json:"creditCardAccount"`
	CreditCards []CreditCards `json:"creditCards"`
	Status      string        `json:"status"`
	Timestamp   string        `json:"timestamp"`
}

type CreditCards struct {
	NewUnseenInvoice       bool   `json:"newUnseenInvoice"`
	HasInvoice             bool   `json:"hasInvoice"`
	HasHistoricInvoice     bool   `json:"hasHistoricInvoice"`
	HasCurrentInvoice      bool   `json:"hasCurrentInvoice"`
	CurrentEntries         string `json:"currentEntries"`
	Available              string `json:"available"`
	ProductText            string `json:"productText"`
	AccountReferenceNumber string `json:"accountReferenceNumber"`
	ID                     string `json:"id"`
	Alias                  string `json:"alias"`
	CurrencyCd             string `json:"currencyCd"`
	CardNumber             string `json:"cardNumber"`
	CardHolderName         string `json:"cardHolderName"`
	Liable                 string `json:"liable"`
	PrepayedCard           string `json:"prepayedCard"`
	Balance                string `json:"balance"`
	SpendingLimit          string `json:"spendingLimit"`
	ProductType            string `json:"productType"`
	ProductCode            string `json:"productCode"`
	Status                 string `json:"status"`
	ValidTo                string `json:"validTo"`
	AvailableBalance       string `json:"availableBalance"`
	SelfRegistered         bool   `json:"selfRegistered"`
	StartDate              string `json:"startDate"`
}

type CreditCardTransactionsRequest struct {
	AccountID                     string `json:"accountId"`
	CurrencyCd                    string `json:"currencyCd"`
	PeriodStartDate               string `json:"periodStartDate"`
	PeriodEndDate                 string `json:"periodEndDate"`
	StartIdxAccountTransactions   int    `json:"startIdxAccountTransactions"`
	NumRecordsAccountTransactions int    `json:"numRecordsAccountTransactions"`
	StartIdxCardTransactions      int    `json:"startIdxCardTransactions"`
	NumRecordsCardTransactions    int    `json:"numRecordsCardTransactions"`
	Refresh                       bool   `json:"refresh"`
	Language                      string `json:"language"`
	RequestKey                    string `json:"requestKey"`
}

type CreditCardTransactionsResponse struct {
	NewUnseenInvoice           bool                            `json:"newUnseenInvoice"`
	PeriodStartDate            string                          `json:"periodStartDate"`
	PeriodEndDate              string                          `json:"periodEndDate"`
	ClosingBalance             string                          `json:"closingBalance"`
	OpeningBalance             string                          `json:"openingBalance"`
	CurrencyCd                 string                          `json:"currencyCd"`
	TotalCredits               string                          `json:"totalCredits"`
	TotalDebits                string                          `json:"totalDebits"`
	InvoiceAmount              string                          `json:"invoiceAmount"`
	DocumentReferenceID        string                          `json:"documentReferenceId"`
	PaymentMethod              string                          `json:"paymentMethod"`
	Ebill                      bool                            `json:"ebill"`
	HasPdf                     bool                            `json:"hasPdf"`
	PaymentAllowed             bool                            `json:"paymentAllowed"`
	AccountID                  string                          `json:"accountId"`
	InvoiceID                  string                          `json:"invoiceId"`
	AccountTransactions        []CreditCardAccountTransactions `json:"accountTransactions"`
	CardTransactions           []CardTransactions              `json:"cardTransactions"`
	HasMoreAccountTransactions bool                            `json:"hasMoreAccountTransactions"`
	HasMoreCardTransactions    bool                            `json:"hasMoreCardTransactions"`
	Status                     string                          `json:"status"`
	Timestamp                  string                          `json:"timestamp"`
}

type CardTransactions struct {
	TransactionType      string `json:"transactionType"`
	MerchantCategoryText string `json:"merchantCategoryText"`
	PostingAmount        string `json:"postingAmount"`
	TransactionDate      string `json:"transactionDate"`
	BookingDate          string `json:"bookingDate"`
	TransactionText      string `json:"transactionText"`
	MoreDetailsAvailable bool   `json:"moreDetailsAvailable"`
	Alias                string `json:"alias"`
	ProductText          string `json:"productText"`
}

type CreditCardAccountTransactions struct {
	PostingAmount        string `json:"postingAmount"`
	TransactionDate      string `json:"transactionDate"`
	BookingDate          string `json:"bookingDate"`
	TransactionText      string `json:"transactionText"`
	MoreDetailsAvailable bool   `json:"moreDetailsAvailable"`
}
