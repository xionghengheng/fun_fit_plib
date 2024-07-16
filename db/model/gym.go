package model

type GymInfoModel struct {
	GymID               int     `json:"gym_id,omitempty"`
	LocName             string  `json:"loc_name"`
	LocDetail           string  `json:"loc_detail"`
	Introduction        string  `json:"introduction"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	HeaderImage1        string  `json:"header_image1,omitempty"`
	HeaderImage2        string  `json:"header_image2,omitempty"`
	HeaderImage3        string  `json:"header_image3,omitempty"`
	HeaderImage4        string  `json:"header_image4,omitempty"`
	LocationGuideImage  string  `json:"location_guide_image,omitempty"`
	NearbySubwayStation string  `json:"nearby_subway_station,omitempty"`
	GymDiscount         string  `json:"gym_discount"` //场馆优惠信息
}
