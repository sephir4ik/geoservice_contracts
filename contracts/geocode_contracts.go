package contracts

type GeocodeAddressesApiRequest struct {
	Lat string
	Lng string
}

type GeocodeAddressesApiResponse struct {
	Addresses []struct {
		City   string
		Street string
		House  string
		Lat    string
		Lon    string
	}
}

//
