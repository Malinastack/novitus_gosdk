package novitus_gosdk

type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Request struct {
	Status    string `json:"status"`
	Id        string `json:"id"`
	EDocument string `json:"e_document"`
	JPKID     string `json:"jpkid"`
	Error     `json:"error"`
}

type Device struct {
	Status string `json:"status"`
	Error  `json:"error"`
}

type TokenResponse struct {
	Token          string `json:"token"`
	ExpirationDate string `json:"expiration_date"`
}

type QueueResponse struct {
	RequestsInQueue int `json:"requests_in_queue"`
}

type DeleteQueueResponse struct {
	Status string `json:"status"`
}

type SendDocumentResponse struct {
	Request `json:"request"`
}

type ConfirmDocumentResponse struct {
	Request `json:"request"`
}

type CheckDocumentStatusResponse struct {
	DeviceObj Device `json:"device"`
	Request   `json:"request"`
}

type DeleteDocumentResponse struct {
	Request `json:"request"`
}

type ErrorResponse struct {
	Exception Error `json:"exception"`
}
