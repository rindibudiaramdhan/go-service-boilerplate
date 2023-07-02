package configs

type Response struct {
	HTTPStatusCode int         // http response status code
	Status         bool        `json:"status"`
	Message        string      `json:"message"`
	Error          interface{} `json:"error,omitempty"`
	Data           interface{} `json:"data,omitempty"`
}
