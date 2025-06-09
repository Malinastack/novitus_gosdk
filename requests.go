package novitus_gosdk

type Summary struct {
	DiscountMarkup string `json:"discount_markup"`
	Total          string `json:"total"`
	PayIn          string `json:"pay_in"`
	Change         string `json:"change"`
}

type EDocument struct {
	TransactionId string `json:"transaction_id"`
	Protocol      string `json:"protocol"`
	PrintSendMode string `json:"print_send_mode"`
}

type Buyer struct {
	Name      string   `json:"name"`
	IdType    string   `json:"id_type"`
	Id        string   `json:"id"`
	LabelType string   `json:"label_type"`
	Address   []string `json:"address"`
	Nip       string   `json:"nip"`
	EDocument `json:"e_document"`
}

type SystemInfo struct {
	CashierName  string `json:"cashier_name"`
	CashNumer    string `json:"cash_number"`
	SystemNumber string `json:"system_number"`
}

type DeviceControl struct {
	OpenDrawer        bool   `json:"open_drawer"`
	FeedAfterPrintout bool   `json:"feed_after_printout"`
	PaperCut          string `json:"paper_cut"`
}

type Receipt struct {
	Items         []struct{} `json:"items"`
	Payments      []struct{} `json:"payments"`
	Summary       `json:"summary"`
	PrintoutLines []struct{} `json:"printout_lines"`
	Buyer         `json:"buyer"`
	SystemInfo    `json:"system_info"`
	DeviceControl `json:"device_control"`
}

type Info struct {
	Number        string `json:"number"`
	CopyCount     int    `json:"copy_count"`
	DateOfSell    string `json:"date_of_sell"`
	DateOfPayment string `json:"date_of_payment"`
	PaymentForm   string `json:"payment_form"`
	Paid          string `json:"paid"`
}

type TransactionSide struct {
	Name      string `json:"name"`
	PrintInfo string `json:"print_info"` // Enum: "place_for_signature" "name_and_place_for_signature" "none"
}

type Options struct {
	SkipDescriptionValueToPay                  bool `json:"skip_description_value_to_pay"`
	SkipBlockGrossValueInAccountingTax         bool `json:"skip_block_gross_value_in_accounting_tax"`
	BuyerBold                                  bool `json:"buyer_bold"`
	SellerBold                                 bool `json:"seller_bold"`
	BuyerNipBold                               bool `json:"buyer_nip_bold"`
	SellerNipBold                              bool `json:"seller_nip_bold"`
	PrintLabelDescriptionSymbolInInvoiceHeader bool `json:"print_label_description_symbol_in_invoice_header"`
	PrintPositionNumberInInvoiceHeader         bool `json:"print_position_number_in_invoice_header"`
	PrintPositionNumberInvoice                 bool `json:"print_position_number_invoice"`
	ToPayLabelBeforeAcountingTaxBlock          bool `json:"to_pay_label_before_acounting_tax_block"`
	PrintCentsInWords                          bool `json:"print_cents_in_words"`
	DontPrintSellDateIfEqualCreateDate         bool `json:"dont_print_sell_date_if_equal_create_date"`
	DontPrintSellerDataInHeader                bool `json:"dont_print_seller_data_in_header"`
	DontPrintSellItemsDescription              bool `json:"dont_print_sell_items_description"`
	EnablePaymentForm                          bool `json:"enable_payment_form"`
	DontPrintCustomerData                      bool `json:"dont_print_customer_data"`
	PrintPaydInCash                            bool `json:"print_payd_in_cash"`
	SkipSellerLabel                            bool `json:"skip_seller_label"`
	PrintInvoiceTaxLabel                       bool `json:"print_invoice_tax_label"`
}

type AdditionalInfo struct {
	Text          string `json:"text"`
	Bold          bool   `json:"bold"`
	Justification string `json:"justification"` // Enum: "left" "center" "right"
}

type Invoice struct {
	Info           `json:"info"`
	Buyer          `json:"buyer"`
	Recipient      TransactionSide `json:"recipient"`
	Seller         TransactionSide `json:"seller"`
	Options        `json:"options"`
	Items          []struct{} `json:"items"`
	Payments       []struct{} `json:"payments"`
	Summary        `json:"summary"`
	PrintoutLines  []struct{}       `json:"printout_lines"`
	AdditionalInfo []AdditionalInfo `json:"additional_info"`
	DeviceControl  `json:"device_control"`
	SystemInfo     `json:"system_info"`
}

type PrintoutOptions struct {
	WithoutHeader    bool `json:"without_header"`
	LeftMargin       bool `json:"left_margin"`
	CopyOnly         bool `json:"copy_only"`
	FiscalMarginsOff bool `json:"fiscal_margins_off"`
}

type Printout struct {
	Options       PrintoutOptions `json:"options"`
	Lines         []string        `json:"lines"`
	EDocument     `json:"e_document"`
	SystemInfo    `json:"system_info"`
	DeviceControl `json:"device_control"`
}
