package novitus_gosdk

import (
	"fmt"
	"resty.dev/v3"
	"time"
)

type NovitusClient struct {
	host                string
	token               string
	tokenExpirationDate int64
}

func NewNovitusClient(host, token string) (*NovitusClient, error) {
	client := &NovitusClient{
		host: host,
	}
	if token != "" {
		client.token = token
		return client, nil
	}
	tokenResp, err := client.ObtainToken()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain token: %w", err)
	}
	client.token = tokenResp.Token
	return client, nil
}

func (n *NovitusClient) ObtainToken() (TokenResponse, error) {
	client := resty.New()
	defer client.Close()
	var tokenResponse TokenResponse
	var errorResponse ErrorResponse
	res, err := client.R().SetResult(&tokenResponse).SetError(&errorResponse).Get(n.host + "/api/v1/token")
	if err != nil {
		return TokenResponse{}, fmt.Errorf("failed to obtain token: %w", err)
	}
	if res.IsError() {
		return TokenResponse{}, fmt.Errorf("error obtaining token: %s", errorResponse.Exception.Description)
	}
	t, err := time.Parse(time.RFC3339, tokenResponse.ExpirationDate)
	if err != nil {
		return TokenResponse{}, fmt.Errorf("Failed to parse expiration date", err)
	}
	n.tokenExpirationDate = t.Unix()
	return tokenResponse, nil
}

func (n *NovitusClient) RefreshToken() error {
	client := resty.New()
	defer client.Close()
	var tokenResponse TokenResponse
	var errorResponse ErrorResponse
	res, err := client.R().SetResult(&tokenResponse).SetError(&errorResponse).
		SetHeader("Authorization", "Bearer "+n.token).
		Patch(n.host + "/api/v1/token")
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error refreshing token: %s", errorResponse.Exception.Description)
	}
	n.token = tokenResponse.Token
	t, err := time.Parse(time.RFC3339, tokenResponse.ExpirationDate)
	if err != nil {
		return fmt.Errorf("failed to parse expiration date: %w", err)
	}
	n.tokenExpirationDate = t.Unix()
	return nil

}

func (n *NovitusClient) RefreshIfNeeded() error {
	if n.token == "" {
		_, err := n.ObtainToken()
		return err
	}
	currentTime := time.Now().Unix()
	if currentTime >= n.tokenExpirationDate {
		err := n.RefreshToken()
		if err != nil {
			_, err := n.ObtainToken()
			return err
		}
	}
	if n.tokenExpirationDate-currentTime < 300 { // Refresh if less than 5 minutes left
		return n.RefreshToken()
	}
	return nil
}

func (n *NovitusClient) GetQueueStatus() (QueueResponse, error) {
	err := n.RefreshIfNeeded()
	if err != nil {
		return QueueResponse{}, fmt.Errorf("failed to refresh token before getting queue status: %w", err)
	}
	client := resty.New()
	defer client.Close()
	var queueResponse QueueResponse
	var errorResponse ErrorResponse
	res, err := client.R().SetResult(&queueResponse).SetError(&errorResponse).
		SetHeader("Authorization", "Bearer "+n.token).
		Get(n.host + "/api/v1/queue")
	if err != nil {
		return QueueResponse{}, fmt.Errorf("failed to get queue status: %w", err)
	}
	if res.IsError() {
		return QueueResponse{}, fmt.Errorf("error getting queue status: %s", errorResponse.Exception.Description)
	}
	return queueResponse, nil
}

func (n *NovitusClient) DeleteQueue() (DeleteQueueResponse, error) {
	err := n.RefreshIfNeeded()
	if err != nil {
		return DeleteQueueResponse{}, fmt.Errorf("failed to refresh token before getting queue status: %w", err)
	}
	client := resty.New()
	defer client.Close()
	var deleteQueueResponse DeleteQueueResponse
	var errorResponse ErrorResponse
	res, err := client.R().SetResult(&deleteQueueResponse).SetError(&errorResponse).
		SetHeader("Authorization", "Bearer "+n.token).
		Delete(n.host + "/api/v1/queue")
	if err != nil {
		return DeleteQueueResponse{}, fmt.Errorf("failed to delete queue: %w", err)
	}
	if res.IsError() {
		return DeleteQueueResponse{}, fmt.Errorf("error deleting queue: %s", errorResponse.Exception.Description)
	}
	return deleteQueueResponse, nil
}

func (n *NovitusClient) Confirm(objectType, requestId string) (SendDocumentResponse, error) {
	err := n.RefreshIfNeeded()
	if err != nil {
		return SendDocumentResponse{}, fmt.Errorf("failed to refresh token before confirming document: %w", err)
	}
	client := resty.New()
	defer client.Close()
	var confirmResponse SendDocumentResponse
	var errorResponse ErrorResponse
	res, err := client.R().
		SetResult(&confirmResponse).
		SetError(&errorResponse).
		SetHeader("Authorization", "Bearer "+n.token).
		Put(n.host + "/api/v1/" + objectType + "/" + requestId)
	if err != nil {
		return SendDocumentResponse{}, fmt.Errorf("failed to confirm document: %w", err)
	}
	if res.IsError() {
		return SendDocumentResponse{}, fmt.Errorf("error confirming document: %s", errorResponse.Exception.Description)
	}
	return confirmResponse, nil
}

func (n *NovitusClient) SendDocument(documentType string, document Document) (SendDocumentResponse, error) {
	err := document.Validate()
	if err != nil {
		return SendDocumentResponse{}, fmt.Errorf("Validation Error: %w", err)
	}
	err = n.RefreshIfNeeded()
	if err != nil {
		return SendDocumentResponse{}, fmt.Errorf("failed to refresh token before sending document: %w", err)
	}
	client := resty.New()
	defer client.Close()
	var sendDocumentResponse SendDocumentResponse
	var errorResponse ErrorResponse
	body := make(map[string]interface{})
	body[documentType] = document
	fmt.Println(body)
	res, err := client.R().
		SetResult(&sendDocumentResponse).
		SetError(&errorResponse).
		SetHeader("Authorization", "Bearer "+n.token).
		SetBody(body).
		Post(n.host + "/api/v1/" + documentType)
	if err != nil {
		return SendDocumentResponse{}, fmt.Errorf("failed to send document: %w", err)
	}
	if res.IsError() {
		return SendDocumentResponse{}, fmt.Errorf("error sending document: %s %s", string(rune(errorResponse.Exception.Code)), errorResponse.Exception.Description)
	}
	return sendDocumentResponse, nil
}

func (n *NovitusClient) CheckDocumentStatus(objectType, requestId string) (CheckDocumentStatusResponse, error) {
	err := n.RefreshIfNeeded()
	if err != nil {
		return CheckDocumentStatusResponse{}, fmt.Errorf("failed to refresh token before checking document status: %w", err)
	}
	client := resty.New()
	defer client.Close()
	var checkDocumentStatusResponse CheckDocumentStatusResponse
	var errorResponse ErrorResponse
	res, err := client.R().
		SetResult(&checkDocumentStatusResponse).
		SetError(&errorResponse).
		SetHeader("Authorization", "Bearer "+n.token).
		Get(n.host + "/api/v1/" + objectType + "/" + requestId)
	if err != nil {
		return CheckDocumentStatusResponse{}, fmt.Errorf("failed to check document status: %w", err)
	}
	if res.IsError() {
		return CheckDocumentStatusResponse{}, fmt.Errorf("error checking document status: %s", errorResponse.Exception.Description)
	}
	return checkDocumentStatusResponse, nil
}

func (n *NovitusClient) DeleteDocument(objectType, requestId string) (DeleteDocumentResponse, error) {
	err := n.RefreshIfNeeded()
	if err != nil {
		return DeleteDocumentResponse{}, fmt.Errorf("failed to refresh token before deleting document: %w", err)
	}
	client := resty.New()
	defer client.Close()
	var deleteDocumentResponse DeleteDocumentResponse
	var errorResponse ErrorResponse
	res, err := client.R().
		SetResult(&deleteDocumentResponse).
		SetError(&errorResponse).
		SetHeader("Authorization", "Bearer "+n.token).
		Delete(n.host + "/api/v1/" + objectType + "/" + requestId)
	if err != nil {
		return DeleteDocumentResponse{}, fmt.Errorf("failed to delete document: %w", err)
	}
	if res.IsError() {
		return DeleteDocumentResponse{}, fmt.Errorf("error deleting document: %s", errorResponse.Exception.Description)
	}
	return deleteDocumentResponse, nil
}

func (n *NovitusClient) SendReceipt(receipt *Receipt, confirm bool) (CheckDocumentStatusResponse, error) {
	sendDocumentResponse, err := n.SendDocument("receipt", receipt)
	if err != nil {
		return CheckDocumentStatusResponse{}, fmt.Errorf("failed to send receipt: %w", err)
	}
	if confirm {
		_, err := n.Confirm("receipt", sendDocumentResponse.Request.Id)
		if err != nil {
			return CheckDocumentStatusResponse{}, fmt.Errorf("failed to confirm document: %w", err)
		}
	}
	return n.CheckDocumentStatus("receipt", sendDocumentResponse.Request.Id)
}

func (n *NovitusClient) SendInvoice(invoice *Invoice, confirm bool) (CheckDocumentStatusResponse, error) {
	sendDocumentResponse, err := n.SendDocument("invoice", invoice)
	if err != nil {
		return CheckDocumentStatusResponse{}, fmt.Errorf("failed to send invoice: %w", err)
	}
	if confirm {
		_, err := n.Confirm("invoice", sendDocumentResponse.Request.Id)
		if err != nil {
			return CheckDocumentStatusResponse{}, fmt.Errorf("failed to confirm document: %w", err)
		}
	}
	return n.CheckDocumentStatus("invoice", sendDocumentResponse.Request.Id)
}

func (n *NovitusClient) SendNFPrintout(printout *Printout, confirm bool) (CheckDocumentStatusResponse, error) {
	sendDocumentResponse, err := n.SendDocument("nf_printout", printout)
	if err != nil {
		return CheckDocumentStatusResponse{}, fmt.Errorf("failed to send printout: %w", err)
	}
	if confirm {
		_, err := n.Confirm("nf_printout", sendDocumentResponse.Request.Id)
		if err != nil {
			return CheckDocumentStatusResponse{}, fmt.Errorf("failed to confirm document: %w", err)
		}
	}
	return n.CheckDocumentStatus("nf_printout", sendDocumentResponse.Request.Id)
}
