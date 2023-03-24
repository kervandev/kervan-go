package kervan

type GetVersionResponse struct {
	Version          string `json:"version"`
	Description      string `json:"description"`
	IsLatest         bool   `json:"is_latest"`
	IsSecurityUpdate bool   `json:"is_security_update"`
	UnattendedUpdate bool   `json:"unattended_update"`
	CreatedAt        string `json:"created_at"`
}
