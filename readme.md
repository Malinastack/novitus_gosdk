# Novitus client SDK for Go
This is a client SDK for the Novitus API, written in Go. It provides a simple and easy-to-use interface for interacting with the Novitus API.
# Installation
```go get github.com/Hkozacz/novitus_gosdk```

# Usage

## Client creation
To start using the Novitus client, you need to create a new client instance. You can do this by providing the base URL of the Novitus API and optionaly Bearer token.
If you provide empty token (`""`"), the client will generate new one on init.
```go
client, err := novitus_gosdk.NewNovitusClient(baseUrl, token)
```
(Base URL should be in the format `https://example.com`)

## API calls
API calls that require authentication will automatically try to refresh the token before making the request. But you can also manually refresh the token if needed.

### ObtainToken
To obtain a new token, you can use the `ObtainToken` method. This method will return a `TokenResponse` struct containing the token and its expiration time.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
tokenResponse, err := client.ObtainToken()
```
### RefreshToken
RefreshToken will automatically refresh the token inside client instance. It does not return response.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
err := client.RefreshToken()
```

### GetQueueStatus
To get the status of the queue, you can use the `GetQueueStatus` method. This method will return a `QueueStatusResponse` struct containing the status of the queue.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
queueStatus, err := client.GetQueueStatus()
```

### DeleteQueue
To delete the queue, you can use the `DeleteQueue` method. This method will return a `DeleteQueueResponse` struct containing the status of the deletion.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
deleteQueueResponse, err := client.DeleteQueue()
```

### SendDocument
SendDocument method allows you to send a document to the Novitus API. It requires a `documentType string` and `Document` struct as an argument and returns a `SendDocumentResponse` struct containing the status of the document.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
SendDocumentResponse, err := client.SendDocument("invoice", invoiceToSend)
```

### Confirm 
Confirm method allows you to confirm a request to the Novitus API. It requires a `objectType String` and `requestId` as an argument and returns a `SendDocumentResponse` struct containing the status of the confirmation.
requestID can be found in the response of the `SendDocument` method.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
confirmResponse, err := client.Confirm("invoice", requestId)
```

### CheckDocumentStatus
CheckDocumentStatus method allows you to check the status of a document in the Novitus API. It requires a `objectType String` and `requestId` as an argument and returns a `CheckDocumentStatusResponse` struct containing the status of the document.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
checkDocumentStatusResponse, err := client.CheckDocumentStatus("invoice", requestId)
```

### DeleteDocument
DeleteDocument method allows you to delete a document in the Novitus API. It requires a `objectType String` and `requestId` as an argument and returns a `DeleteDocumentResponse` struct containing the status of the deletion.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
deleteDocumentResponse, err := client.DeleteDocument("invoice", requestId)
```

### SendReceipt
SendReceipt is a wrapper for the `SendDocument` method, specifically for sending receipts. It requires a `Receipt` struct as an argument and `confirm bool` and returns a `SendDocumentResponse` struct containing the status of the receipt.
If `confirm` is set to true, the method will automatically confirm the receipt after sending it.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
receipt := novitus_gosdk.Receipt{
    // fill in the receipt details
}
sendReceiptResponse, err := client.SendReceipt(receipt, true)
```

### SendInvoice
SendInvoice is a wrapper for the `SendDocument` method, specifically for sending invoices. It requires an `Invoice` struct as an argument and `confirm bool` and returns a `SendDocumentResponse` struct containing the status of the invoice.
If `confirm` is set to true, the method will automatically confirm the invoice after sending it.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
invoice := novitus_gosdk.Invoice{
    // fill in the invoice details
}
sendInvoiceResponse, err := client.SendInvoice(invoice, true)
```

### SendNFPrintout
SendNFPrintout is a wrapper for the `SendDocument` method, specifically for sending nonfiscal printouts. It requires an `NFPrintout` struct as an argument and `confirm bool` and returns a `SendDocumentResponse` struct containing the status of the NF printout.
If `confirm` is set to true, the method will automatically confirm the NF printout after sending it.
```go
client, err := novitus_gosdk.NewNovitusClient("example.com", "")
nfPrintout := novitus_gosdk.NFPrintout{
    // fill in the NF printout details
}
sendNFPrintoutResponse, err := client.SendNFPrintout(nfPrintout, true)
```


## Structs

```go
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
	Items         []interface{}    `json:"items"` // Required: true
	Payments      []interface{}    `json:"payments"`
	Summary       `json:"summary"` // Required: true
	PrintoutLines []interface{}    `json:"printout_lines"`
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
	Info           `json:"info"`   // Required: true
	Buyer          `json:"buyer"`  // Required: true
	Recipient      TransactionSide `json:"recipient"`
	Seller         TransactionSide `json:"seller"`
	Options        `json:"options"`
	Items          []interface{}    `json:"items"` // Required: true
	Payments       []interface{}    `json:"payments"`
	Summary        `json:"summary"` // Required: true
	PrintoutLines  []interface{}    `json:"printout_lines"`
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
	Lines         []string        `json:"lines"` // Required: true
	EDocument     `json:"e_document"`
	SystemInfo    `json:"system_info"`
	DeviceControl `json:"device_control"`
}

// Items

type Article struct {
	Name           string `json:"name"`            // Required: true
	PTU            string `json:"ptu"`             // Enum: "A" - "G" Required: true
	Quantity       string `json:"quantity"`        // Quantity in units, e.g. "1.00" Required: true
	Price          string `json:"price"`           // Price in currency, e.g. "1.00" Required: true
	Value          string `json:"value"`           // Total value for the item, e.g. "1.00" Required: true
	Unit           string `json:"unit"`            // Enum: "szt" - "kg", etc.
	DiscountMarkup string `json:"discount_markup"` // Optional, e.g. "0.00"
	Code           string `json:"code"`            // Optional, e.g. "1234567890123" Can be set only if Description is not Set
	Description    string `json:"description"`     // Optional, e.g. "Sample Item"
}

type Advance struct {
	Description string `json:"description"` // Required: true, e.g. "Advance Payment"
	PTU         string `json:"ptu"`         // Enum: "A" - "G" Required: true
	Value       string `json:"value"`       // Value in currency, e.g. "100.00" Required: true
}

type AdvanceReturn struct {
	Description string `json:"description"` // Required: true, e.g. "Advance Return"
	PTU         string `json:"ptu"`         // Enum: "A" - "G" Required: true
	Value       string `json:"value"`       // Value in currency, e.g. "50.00" Required: true
}

type Container struct {
	Name     string `json:"name"`     // e.g. "Container Name"
	Number   string `json:"number"`   // e.g. "12345"
	Quantity string `json:"quantity"` // Quantity in units, e.g. "10.00"
	Value    string `json:"value"`    // Total value for the container, e.g. "100.00" Required: true
}

type ContainerReturn struct {
	Name     string `json:"name"`     // e.g. "Container Name"
	Number   string `json:"number"`   // e.g. "12345"
	Quantity string `json:"quantity"` // Quantity in units, e.g. "10.00"
	Value    string `json:"value"`    // Total value for the container, e.g. "100.00" Required: true
}

// Payments

type Cash struct {
	Value string `json:"value"` // Value in currency, e.g. "100.00" Required: true
}

type TypicalPaymentMethod struct {
	Name  string `json:"name"`  // enum "card", cheque, coupon, other, credit, account, transfer, mobile, voucher
	Value string `json:"value"` // Value in currency, e.g. "100.00" Required: true
}

type Currency struct {
	Course        string `json:"course"`         // e.g. "1.00" Required: true
	CurrencyValue string `json:"currency_value"` // e.g. "USD" Required: true
	LocalValue    string `json:"local_value"`    // e.g. "100.00" Required: true
	IsChange      bool   `json:"is_change"`      // true if this is a change, false otherwise Required: true
	Name          string `json:"name"`           // e.g. "USD" Required: true
}

// Printout Lines

type PrintoutLine struct {
	Text 	  string `json:"text"`        // The text to be printed, e.g. "Sample Text" Required: true
	Masked    bool   `json:"masked"`      // true if the text should be masked, false otherwise Required: true
}

type TextLine struct {
	Bold       bool   `json:"bold"`        // true if the text should be bold, false otherwise
	Invers     bool   `json:"invers"`      // true if the text should be inverted, false otherwise
	Center     bool   `json:"center"`      // true if the text should be centered, false otherwise
	FontNumber int    `json:"font_number"` // Font number, e.g. 1, 2, 3
	Big        bool   `json:"big"`         // true if the text should be big, false otherwise
	Height     int    `json:"height"`      // Height of the text in points, e.g. 12
	Width      int    `json:"width"`       // Width of the text in points, e.g. 100
	Text       string `json:"text"`        // The text to be printed, e.g. "Sample Text" Required: true
	Masked     bool   `json:"masked"`      // true if the text should be masked, false otherwise Required: true
}

```