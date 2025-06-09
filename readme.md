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
