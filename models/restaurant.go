package models

type (
	Pincode        string
	OrderCrieteria int
)

const (
	RatingsOrder OrderCrieteria = iota
	PriceOrder
)

type Restaurant struct {
	Name                string
	ServiceablePinCodes []Pincode
	FoodItem            string
	Quantity            uint64
	Price               float32

	// Misc:
	Owner   string
	Ratings Ratings
}

func (r *Restaurant) ToServiceable() ServiceableRestaurant {
	rating := float32(0)
	for _, rate := range r.Ratings {
		rating += float32(rate.Rating)
	}
	rating /= float32(len(r.Ratings))
	return ServiceableRestaurant{
		Name:     r.Name,
		FoodItem: r.FoodItem,
		Quantity: r.Quantity,
		Price:    r.Price,
		Rating:   rating,
	}
}

func (r *Restaurant) CanDeliver(pincode Pincode) bool {
	for _, p := range r.ServiceablePinCodes {
		if p == pincode {
			return true
		}
	}

	return false
}

type Ratings []Rating
type Rating struct {
	User    string
	Comment *string
	Rating  int
}

type ServiceableRestaurant struct {
	Name     string
	FoodItem string
	Quantity uint64

	Rating float32
	Price  float32
}
