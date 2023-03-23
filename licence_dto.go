package kervan

type CheckLicencePayload struct {
	Token string `json:"token"`
}

type CheckLicenceResponse struct {
	Token string `json:"token"`
}
