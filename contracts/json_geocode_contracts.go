package contracts

type JsonGeocodeAddressesApiRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type JsonGeocodeAddressesApiResponse struct {
	Addresses []struct {
		City   string `json:"city"`
		Street string `json:"street"`
		House  string `json:"house"`
		Lat    string `json:"lat"`
		Lon    string `json:"lon"`
	}
}
