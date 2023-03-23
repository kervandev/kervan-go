package kervan

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// KervanAPI is the main struct for the Tapsilat API
type API struct {
	EndPoint string `json:"end_point"`
	Token    string `json:"token"`
}

// NewAPI creates a new KervanAPI struct
func NewAPI(token string) *API {
	return &API{
		EndPoint: "https://api.kervan.dev/api/v1",
		Token:    token,
	}
}

// NewAPIWithEndpoint creates a new KervanAPI struct with a custom endpoint
func NewCustomAPI(endpoint, token string) *API {
	return &API{
		EndPoint: endpoint,
		Token:    token,
	}
}

func (t *API) post(path string, payload interface{}, response interface{}) error {
	url := t.EndPoint + path
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	return t.do(req, response)
}

func (t *API) get(path string, payload interface{}, response interface{}) error {
	url := t.EndPoint + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	return t.do(req, response)
}

func (t *API) do(req *http.Request, response interface{}) error {
	req.Header.Set("Authorization", "Bearer "+t.Token)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	decode := json.NewDecoder(bytes.NewReader(body))
	decode.DisallowUnknownFields()
	decode.UseNumber()
	if err := decode.Decode(response); err != nil {
		return err
	}
	return nil
}