package responses

type CreateHubResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	TenantID uint   `json:"tenant_id"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2,omitempty"` // Omits empty values
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Pincode  string `json:"pincode"`
}
