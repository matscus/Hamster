package handlers

type request struct {
	Query string `json:"query"`
}

type fioResponce struct {
	Suggestion []suggestionFIO `json:"suggestion"`
}
type addressResponce struct {
	Suggestion []suggestionAddress `json:"suggestion"`
}
type orgResponce struct {
	Status          string `json:"status,omitempty"`
	ActualTimestamp int64  `json:"actualTimestamp,omitempty"`
	Data            struct {
		Suggestion []suggestionOrg `json:"suggestion"`
	} `json:"data"`
}
type suggestionFIO struct {
	Value             string `json:"value"`
	UnrestrictedValue string `json:"unrestricted_value"`
	Data              struct {
		Surname    string `json:"surname"`
		Name       string `json:"name"`
		Patronymic string `json:"patronymic"`
		Gender     string `json:"gender"`
		Source     string `json:"source"`
		Gc         string `json:"gc"`
	} `json:"data"`
}
type suggestionAddress struct {
	Value             string `json:"value"`
	UnrestrictedValue string `json:"unrestricted_value"`
	Data              struct {
		PostalCode           string `json:"postal_code,omitempty"`
		Country              string `json:"country,omitempty"`
		RegionFiasID         string `json:"region_fias_id,omitempty"`
		RegionKladrID        string `json:"region_kladr_id,omitempty"`
		RegionWithType       string `json:"region_with_type,omitempty"`
		RegionType           string `json:"region_type,omitempty"`
		RegionTypeFull       string `json:"region_type_full,omitempty"`
		Region               string `json:"region,omitempty"`
		AreaFiasID           string `json:"area_fias_id,omitempty"`
		AreaKladrID          string `json:"area_kladr_id,omitempty"`
		AreaWithType         string `json:"area_with_type,omitempty"`
		AreaType             string `json:"area_type,omitempty"`
		AreaTypeFull         string `json:"area_type_full,omitempty"`
		Area                 string `json:"area,omitempty"`
		CityFiasID           string `json:"city_fias_id,omitempty"`
		CityKladrID          string `json:"city_kladr_id,omitempty"`
		CityWithType         string `json:"city_with_type,omitempty"`
		CityType             string `json:"city_type,omitempty"`
		CityTypeFull         string `json:"city_type_full,omitempty"`
		City                 string `json:"city,omitempty"`
		CityArea             string `json:"city_area,omitempty"`
		CityDistrictFiasID   string `json:"city_district_fias_id,omitempty"`
		CityDistrictKladrID  string `json:"city_district_kladr_id,omitempty"`
		CityDistrictWithType string `json:"city_district_with_type,omitempty"`
		CityDistrictType     string `json:"city_district_type,omitempty"`
		CityDistrictTypeFull string `json:"city_district_type_full,omitempty"`
		CityDistrict         string `json:"city_district,omitempty"`
		SettlementFiasID     string `json:"settlement_fias_id,omitempty"`
		SettlementKladrID    string `json:"settlement_kladr_id,omitempty"`
		SettlementWithType   string `json:"settlement_with_type,omitempty"`
		SettlementType       string `json:"settlement_type,omitempty"`
		SettlementTypeFull   string `json:"settlement_type_full,omitempty"`
		Settlement           string `json:"settlement,omitempty"`
		StreetFiasID         string `json:"street_fias_id,omitempty"`
		StreetKladrID        string `json:"street_kladr_id,omitempty"`
		StreetWithType       string `json:"street_with_type,omitempty"`
		StreetType           string `json:"street_type,omitempty"`
		StreetTypeFull       string `json:"street_type_full,omitempty"`
		Street               string `json:"street,omitempty"`
		HouseFiasID          string `json:"house_fias_id,omitempty"`
		HouseKladrID         string `json:"house_kladr_id,omitempty"`
		HouseType            string `json:"house_type,omitempty"`
		HouseTypeFull        string `json:"house_type_full,omitempty"`
		House                string `json:"house,omitempty"`
		BlockType            string `json:"block_type,omitempty"`
		BlockTypeFull        string `json:"block_type_full,omitempty"`
		Block                string `json:"block,omitempty"`
		FlatType             string `json:"flat_type,omitempty"`
		FlatTypeFull         string `json:"flat_type_full,omitempty"`
		Flat                 string `json:"flat,omitempty"`
		FlatArea             string `json:"flat_area,omitempty"`
		SquareMeterPrice     string `json:"square_meter_price,omitempty"`
		FlatPrice            string `json:"flat_price,omitempty"`
		PostalBox            string `json:"postal_box,omitempty"`
		FiasID               string `json:"fias_id,omitempty"`
		FiasLevel            string `json:"fias_level,omitempty"`
		KladrID              string `json:"kladr_id,omitempty"`
		CapitalMarker        string `json:"capital_marker,omitempty"`
		Okato                string `json:"okato,omitempty"`
		Oktmo                string `json:"oktmo,omitempty"`
		TaxOffice            string `json:"tax_office,omitempty"`
		TaxOfficeLegal       string `json:"timezone,omitempty"`
		GeoLat               string `json:"geo_lat,omitempty"`
		GeoLon               string `json:"geo_lon,omitempty"`
		BeltwayHit           string `json:"beltway_hit,omitempty"`
		BeltwayDistance      string `json:"beltway_distance,omitempty"`
		QcGeo                string `json:"qc_geo,omitempty"`
		QcComplete           string `json:"qc_complete,omitempty"`
		QcHouse              string `json:"qc_house,omitempty"`
		UnparsedParts        string `json:"unparsed_parts,omitempty"`
		Qc                   string `json:"qc,omitempty"`
	} `json:"data"`
}
type suggestionOrg struct {
	Value             string `json:"value,omitempty"`
	UnrestrictedValue string `json:"unrestricted_value,omitempty"`
	Data              struct {
		Kpp        string `json:"kpp,omitempty"`
		Management struct {
			Name string `json:"name,omitempty"`
			Post string `json:"post,omitempty"`
		} `json:"management,omitempty"`
		BranchType  string `json:"branch_type,omitempty"`
		BranchCount int    `json:"branch_count,omitempty"`
		Type        string `json:"type,omitempty"`
		Opf         struct {
			Code  string `json:"code,omitempty"`
			Full  string `json:"full,omitempty"`
			Short string `json:"short,omitempty"`
		} `json:"opf,omitempty"`
		Name struct {
			FullWithOpf  string      `json:"full_with_opf,omitempty"`
			ShortWithOpf string      `json:"short_with_opf,omitempty"`
			Latin        interface{} `json:"latin,omitempty"`
			Full         string      `json:"full,omitempty"`
			Short        string      `json:"short,omitempty"`
		} `json:"name,omitempty"`
		Inn   string      `json:"inn,omitempty"`
		Ogrn  string      `json:"ogrn,omitempty"`
		Okpo  interface{} `json:"okpo,omitempty"`
		Okved interface{} `json:"okved,omitempty"`
		State struct {
			Status           string      `json:"status,omitempty"`
			ActualityDate    int64       `json:"actuality_date,omitempty"`
			RegistrationDate int64       `json:"registration_date,omitempty"`
			LiquidationDate  interface{} `json:"liquidation_date,omitempty"`
		} `json:"state,omitempty"`
		Address struct {
			Value             string `json:"value,omitempty"`
			UnrestrictedValue string `json:"unrestricted_value,omitempty"`
			Data              struct {
				Data struct {
					PostalCode           string `json:"postal_code,omitempty"`
					Country              string `json:"country,omitempty"`
					RegionFiasID         string `json:"region_fias_id,omitempty"`
					RegionKladrID        string `json:"region_kladr_id,omitempty"`
					RegionWithType       string `json:"region_with_type,omitempty"`
					RegionType           string `json:"region_type,omitempty"`
					RegionTypeFull       string `json:"region_type_full,omitempty"`
					Region               string `json:"region,omitempty"`
					AreaFiasID           string `json:"area_fias_id,omitempty"`
					AreaKladrID          string `json:"area_kladr_id,omitempty"`
					AreaWithType         string `json:"area_with_type,omitempty"`
					AreaType             string `json:"area_type,omitempty"`
					AreaTypeFull         string `json:"area_type_full,omitempty"`
					Area                 string `json:"area,omitempty"`
					CityFiasID           string `json:"city_fias_id,omitempty"`
					CityKladrID          string `json:"city_kladr_id,omitempty"`
					CityWithType         string `json:"city_with_type,omitempty"`
					CityType             string `json:"city_type,omitempty"`
					CityTypeFull         string `json:"city_type_full,omitempty"`
					City                 string `json:"city,omitempty"`
					CityArea             string `json:"city_area,omitempty"`
					CityDistrictFiasID   string `json:"city_district_fias_id,omitempty"`
					CityDistrictKladrID  string `json:"city_district_kladr_id,omitempty"`
					CityDistrictWithType string `json:"city_district_with_type,omitempty"`
					CityDistrictType     string `json:"city_district_type,omitempty"`
					CityDistrictTypeFull string `json:"city_district_type_full,omitempty"`
					CityDistrict         string `json:"city_district,omitempty"`
					SettlementFiasID     string `json:"settlement_fias_id,omitempty"`
					SettlementKladrID    string `json:"settlement_kladr_id,omitempty"`
					SettlementWithType   string `json:"settlement_with_type,omitempty"`
					SettlementType       string `json:"settlement_type,omitempty"`
					SettlementTypeFull   string `json:"settlement_type_full,omitempty"`
					Settlement           string `json:"settlement,omitempty"`
					StreetFiasID         string `json:"street_fias_id,omitempty"`
					StreetKladrID        string `json:"street_kladr_id,omitempty"`
					StreetWithType       string `json:"street_with_type,omitempty"`
					StreetType           string `json:"street_type,omitempty"`
					StreetTypeFull       string `json:"street_type_full,omitempty"`
					Street               string `json:"street,omitempty"`
					HouseFiasID          string `json:"house_fias_id,omitempty"`
					HouseKladrID         string `json:"house_kladr_id,omitempty"`
					HouseType            string `json:"house_type,omitempty"`
					HouseTypeFull        string `json:"house_type_full,omitempty"`
					House                string `json:"house,omitempty"`
					BlockType            string `json:"block_type,omitempty"`
					BlockTypeFull        string `json:"block_type_full,omitempty"`
					Block                string `json:"block,omitempty"`
					FlatType             string `json:"flat_type,omitempty"`
					FlatTypeFull         string `json:"flat_type_full,omitempty"`
					Flat                 string `json:"flat,omitempty"`
					FlatArea             string `json:"flat_area,omitempty"`
					SquareMeterPrice     string `json:"square_meter_price,omitempty"`
					FlatPrice            string `json:"flat_price,omitempty"`
					PostalBox            string `json:"postal_box,omitempty"`
					FiasID               string `json:"fias_id,omitempty"`
					FiasLevel            string `json:"fias_level,omitempty"`
					KladrID              string `json:"kladr_id,omitempty"`
					CapitalMarker        string `json:"capital_marker,omitempty"`
					Okato                string `json:"okato,omitempty"`
					Oktmo                string `json:"oktmo,omitempty"`
					TaxOffice            string `json:"tax_office,omitempty"`
					TaxOfficeLegal       string `json:"tax_office_legal,omitempty"`
					Timezone             string `json:"timezone,omitempty"`
					GeoLat               string `json:"geo_lat,omitempty"`
					GeoLon               string `json:"geo_lon,omitempty"`
					BeltwayHit           string `json:"beltway_hit,omitempty"`
					BeltwayDistance      string `json:"beltway_distance,omitempty"`
					QcGeo                string `json:"qc_geo,omitempty"`
					QcComplete           string `json:"qc_complete,omitempty"`
					QcHouse              string `json:"qc_house,omitempty"`
					UnparsedParts        string `json:"unparsed_parts,omitempty"`
					Qc                   string `json:"qc,omitempty"`
				} `json:"data,omitempty"`
			} `json:"data,omitempty"`
		} `json:"address,omitempty"`
	} `json:"data,omitempty"`
}

type errResponce struct {
	Status          string `json:"status"`
	ActualTimestamp int64  `json:"actualTimestamp"`
	Error           struct {
		ID             string `json:"id"`
		Code           string `json:"code"`
		DisplayMessage string `json:"displayMessage"`
		SystemID       string `json:"system-id"`
	} `json:"error"`
}
