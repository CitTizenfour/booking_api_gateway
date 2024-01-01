// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateBikeResponse struct {
	BikeStatus   string  `json:"bike_status"`
	BikeNumber   string  `json:"bike_number"`
	BikeModel    string  `json:"bike_model"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type CreateBusResponse struct {
	ID           string  `json:"id"`
	DriverID     string  `json:"driver_id"`
	MerchantID   string  `json:"merchant_id"`
	RegionID     string  `json:"region_id"`
	BusStatus    string  `json:"bus_status"`
	BusName      string  `json:"bus_name"`
	BusNumber    string  `json:"bus_number"`
	BusType      string  `json:"bus_type"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	BookingCount int     `json:"booking_count"`
	PlaceCount   int     `json:"place_count"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type CreateCarResponse struct {
	IsCovid      bool    `json:"is_covid"`
	IsAirCondt   bool    `json:"is_air_condt"`
	CarStatus    string  `json:"car_status"`
	CarNumber    string  `json:"car_number"`
	CarName      string  `json:"car_name"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type CreateDriverResponse struct {
	DriverName  string `json:"driver_name"`
	PhoneNumber string `json:"phone_number"`
	Experience  string `json:"experience"`
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateKekeResponse struct {
	KekeStatus   string  `json:"keke_status"`
	KekeNumber   string  `json:"keke_number"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type CreateMerchantResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type CreateOrderResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CreateReviewReplyResponse struct {
	ID        string `json:"id"`
	SenderID  string `json:"sender_id"`
	ReviewID  string `json:"review_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateReviewResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	MerchantID    string  `json:"merchant_id"`
	Rate          float64 `json:"rate"`
	ReviewContent string  `json:"review_content"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type CreateUserResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type GetListBikesReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListBikesResponse struct {
	GetBikes []*GetSingleBikeResponse `json:"getBikes"`
	Count    int                      `json:"count"`
}

type GetListBusesReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListBusesResponse struct {
	GetBuses []*GetSingleBusResponse `json:"getBuses"`
	Count    int                     `json:"count"`
}

type GetListCarsReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListCarsResponse struct {
	GetCars []*GetSingleCarResponse `json:"getCars"`
	Count   int                     `json:"count"`
}

type GetListDriversReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListDriversResponse struct {
	GetDrivers []*GetSingleDriverResponse `json:"getDrivers"`
	Count      int                        `json:"count"`
}

type GetListKekesReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListKekesResponse struct {
	GetKekes []*GetSingleKekeResponse `json:"getKekes"`
	Count    int                      `json:"count"`
}

type GetListMerchantReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListMerchantResponse struct {
	Getmerchants []*GetSingleMerchantResponse `json:"getmerchants"`
	Count        int                          `json:"count"`
}

type GetListOrdersReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListOrdersResponse struct {
	Getorders []*GetSingleOrderResponse `json:"getorders"`
	Count     int                       `json:"count"`
}

type GetListReviewsReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListReviewsResponse struct {
	GetReviews []*GetSingleReviewResponse `json:"getReviews"`
	Count      int                        `json:"count"`
}

type GetListUsersReq struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetListUsersResponse struct {
	GetUsers []*GetSingleUserResponse `json:"getUsers"`
	Count    int                      `json:"count"`
}

type GetReplyReviewsResponse struct {
	GetReviews []*ReviewComment `json:"getReviews"`
	Count      int              `json:"count"`
}

type GetSingleBikeResponse struct {
	BikeStatus   string  `json:"bike_status"`
	BikeNumber   string  `json:"bike_number"`
	BikeModel    string  `json:"bike_model"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GetSingleBusResponse struct {
	ID           string  `json:"id"`
	DriverID     string  `json:"driver_id"`
	MerchantID   string  `json:"merchant_id"`
	RegionID     string  `json:"region_id"`
	BusStatus    string  `json:"bus_status"`
	BusName      string  `json:"bus_name"`
	BusNumber    string  `json:"bus_number"`
	BusType      string  `json:"bus_type"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	BookingCount int     `json:"booking_count"`
	PlaceCount   int     `json:"place_count"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GetSingleCarResponse struct {
	IsCovid      bool    `json:"is_covid"`
	IsAirCondt   bool    `json:"is_air_condt"`
	CarStatus    string  `json:"car_status"`
	CarNumber    string  `json:"car_number"`
	CarName      string  `json:"car_name"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GetSingleDriverResponse struct {
	DriverName  string `json:"driver_name"`
	PhoneNumber string `json:"phone_number"`
	Experience  string `json:"experience"`
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetSingleKekeResponse struct {
	KekeStatus   string  `json:"keke_status"`
	KekeNumber   string  `json:"keke_number"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GetSingleMerchantResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type GetSingleOrderResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type GetSingleReviewResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	MerchantID    string  `json:"merchant_id"`
	Rate          float64 `json:"rate"`
	ReviewContent string  `json:"review_content"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type GetSingleUserResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type ReviewComment struct {
	ID        string `json:"id"`
	ReviewID  string `json:"review_id"`
	SenderID  string `json:"sender_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateBikeResponse struct {
	BikeStatus   string  `json:"bike_status"`
	BikeNumber   string  `json:"bike_number"`
	BikeModel    string  `json:"bike_model"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type UpdateBusResponse struct {
	ID           string  `json:"id"`
	BusStatus    string  `json:"bus_status"`
	BusName      string  `json:"bus_name"`
	BusNumber    string  `json:"bus_number"`
	BusType      string  `json:"bus_type"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	BookingCount int     `json:"booking_count"`
	PlaceCount   int     `json:"place_count"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type UpdateCarResponse struct {
	IsCovid      bool    `json:"is_covid"`
	IsAirCondt   bool    `json:"is_air_condt"`
	CarStatus    string  `json:"car_status"`
	CarNumber    string  `json:"car_number"`
	CarName      string  `json:"car_name"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type UpdateDriverResponse struct {
	DriverName  string `json:"driver_name"`
	PhoneNumber string `json:"phone_number"`
	Experience  string `json:"experience"`
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateKekeResponse struct {
	KekeStatus   string  `json:"keke_status"`
	KekeNumber   string  `json:"keke_number"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type UpdateMerchantResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type UpdateOrderResponse struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateReviewReplyResponse struct {
	ID        string `json:"id"`
	SenderID  string `json:"sender_id"`
	ReviewID  string `json:"review_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateReviewResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	MerchantID    string  `json:"merchant_id"`
	Rate          float64 `json:"rate"`
	ReviewContent string  `json:"review_content"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type UpdateUserResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type CreateBikeReq struct {
	BikeStatus   string  `json:"bike_status"`
	BikeNumber   string  `json:"bike_number"`
	BikeModel    string  `json:"bike_model"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	Colour       string  `json:"colour"`
}

type CreateBusReq struct {
	DriverID     string  `json:"driver_id"`
	MerchantID   string  `json:"merchant_id"`
	RegionID     string  `json:"region_id"`
	BusStatus    string  `json:"bus_status"`
	BusName      string  `json:"bus_name"`
	BusNumber    string  `json:"bus_number"`
	BusType      string  `json:"bus_type"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	BookingCount int     `json:"booking_count"`
	PlaceCount   int     `json:"place_count"`
}

type CreateCarReq struct {
	IsCovid      bool    `json:"is_covid"`
	IsAirCondt   bool    `json:"is_air_condt"`
	CarStatus    string  `json:"car_status"`
	CarNumber    string  `json:"car_number"`
	CarName      string  `json:"car_name"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
}

type CreateDriverReq struct {
	DriverName  string `json:"driver_name"`
	PhoneNumber string `json:"phone_number"`
	Experience  string `json:"experience"`
}

type CreateKekeReq struct {
	KekeStatus   string  `json:"keke_status"`
	KekeNumber   string  `json:"keke_number"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
}

type CreateMerchantReq struct {
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
}

type CreateOrderReq struct {
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
}

type CreateReviewReplyReq struct {
	SenderID string `json:"sender_id"`
	ReviewID string `json:"review_id"`
	Content  string `json:"content"`
}

type CreateReviewReq struct {
	UserID        string  `json:"user_id"`
	MerchantID    string  `json:"merchant_id"`
	Rate          float64 `json:"rate"`
	ReviewContent string  `json:"review_content"`
}

type CreateUserReq struct {
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
}

type UpdateBikeReq struct {
	BikeStatus   string  `json:"bike_status"`
	BikeNumber   string  `json:"bike_number"`
	BikeModel    string  `json:"bike_model"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
}

type UpdateBusReq struct {
	BusStatus    string  `json:"bus_status"`
	BusName      string  `json:"bus_name"`
	BusNumber    string  `json:"bus_number"`
	BusType      string  `json:"bus_type"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	BookingCount int     `json:"booking_count"`
	PlaceCount   int     `json:"place_count"`
	ID           string  `json:"id"`
}

type UpdateCarReq struct {
	IsCovid      bool    `json:"is_covid"`
	IsAirCondt   bool    `json:"is_air_condt"`
	CarStatus    string  `json:"car_status"`
	CarNumber    string  `json:"car_number"`
	CarName      string  `json:"car_name"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
}

type UpdateDriverReq struct {
	DriverName  string `json:"driver_name"`
	PhoneNumber string `json:"phone_number"`
	Experience  string `json:"experience"`
	ID          string `json:"id"`
}

type UpdateKekeReq struct {
	KekeStatus   string  `json:"keke_status"`
	KekeNumber   string  `json:"keke_number"`
	StartedPrice float64 `json:"started_price"`
	EndPrice     float64 `json:"end_price"`
	PlaceCount   int     `json:"place_count"`
	Colour       string  `json:"colour"`
	ID           string  `json:"id"`
}

type UpdateMerchantReq struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
}

type UpdateOrderReq struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
}

type UpdateReviewReplyReq struct {
	ID       string `json:"id"`
	SenderID string `json:"sender_id"`
	Content  string `json:"content"`
}

type UpdateReviewReq struct {
	ID            string  `json:"id"`
	Rate          float64 `json:"rate"`
	ReviewContent string  `json:"review_content"`
}

type UpdateUserReq struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PaymentHistory string `json:"payment_history"`
	PhoneNumber    string `json:"phone_number"`
}
