package contracts

type SearchAddressesApiRequest struct {
	Query string
}

type SearchAddressesApiResponse struct {
	Addresses []struct {
		City   string
		Street string
		House  string
		Lat    string
		Lon    string
	}
}
