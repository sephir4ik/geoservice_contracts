package contracts

type JsonSearchAddressesApiRequest struct {
	Query string `json:"query"`
}

type JsonSearchAddressesApiResponse struct {
	Addresses []struct {
		City   string `json:"city"`
		Street string `json:"street"`
		House  string `json:"house"`
		Lat    string `json:"lat"`
		Lon    string `json:"lon"`
	}
}
