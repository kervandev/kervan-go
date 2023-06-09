package kervan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// API is the main struct for the Kervan API
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

func (t *API) post(path string, payload interface{}, response interface{}, headers ...map[string]string) error {
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
	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Set(k, v)
		}
	}

	return t.do(req, response)
}

func (t *API) get(path string, payload interface{}, response interface{}, headers ...map[string]string) error {
	url := t.EndPoint + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Set(k, v)
		}
	}
	return t.do(req, response)
}

type ErrorResponse struct {
	Error string `json:"error"`
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
	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err := decode.Decode(&errorResponse); err != nil {
			return err
		}
		return fmt.Errorf("error while calling api: %s", errorResponse.Error)
	}

	if err := decode.Decode(response); err != nil {
		return err
	}
	return nil
}

func (t *API) CheckLicence(payload *CheckLicencePayload) (*CheckLicenceResponse, error) {
	var response CheckLicenceResponse
	headers := map[string]string{
		"x-ktp": t.Token,
	}
	err := t.post("/vlcs/licence/check", payload, &response, headers)
	if err != nil {
		return nil, fmt.Errorf("error while checking licence: %w", err)
	}
	if response.Token == "" {
		return nil, fmt.Errorf("error while checking licence: token is empty")
	}
	return &response, nil
}

func (t *API) GetVersion() (*GetVersionResponse, error) {
	var response GetVersionResponse
	headers := map[string]string{
		"x-ktp": t.Token,
	}
	if err := t.get("/vlcs/version", nil, &response, headers); err != nil {
		return nil, fmt.Errorf("error while getting version: %w", err)
	}
	return &response, nil
}
