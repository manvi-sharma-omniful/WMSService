package requests

type CreateHubRequest struct {
	Name     string
	ID       uint
	TenantID uint
	Address1 string
	Address2 string
	City     string
	State    string
	Country  string
	Pincode  string
}
