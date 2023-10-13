package service

import (
	"slices"

	"github.com/keshavchand/flipkarLLD/models"
)

type RestaurantService struct {
	users         map[string]models.User // Phone number to user
	currentUserID string                 // This is the registerd user's phone number

	restaurants map[string]models.Restaurant // Restaurant name to Resturant Details
	orders      map[string][]models.Order    // UserId to their order history
}

func NewRestaurantService() *RestaurantService {
	return &RestaurantService{
		users:       make(map[string]models.User),
		restaurants: make(map[string]models.Restaurant),
		orders:      make(map[string][]models.Order),
	}
}

func (r *RestaurantService) RegisterUser(user models.User) error {
	if _, registered := r.users[user.PhoneNumber]; registered {
		return models.ErrPhoneNumberAlreadyRegistered
	}
	r.users[user.PhoneNumber] = user
	return nil
}

func (r *RestaurantService) LoginUser(phoneNumber string) error {
	if _, registered := r.users[phoneNumber]; !registered {
		return models.ErrPhoneNumberNotRegistered
	}
	r.currentUserID = phoneNumber
	return nil
}

func (r *RestaurantService) RegisterRestaurant(rest models.Restaurant) error {
	if _, registered := r.restaurants[rest.Name]; registered {
		return models.ErrRestaurantAlreadyRegistered
	}

	rest.Owner = r.currentUserID
	r.restaurants[rest.Name] = rest
	return nil
}

func (r *RestaurantService) UpdateQuantity(name string, quantity int64) error {
	rest, registered := r.restaurants[name]
	if !registered {
		return models.ErrRestaurantNotRegistered
	}

	if rest.Owner != r.currentUserID {
		return models.ErrUnauthorizedUser
	}

	newQuantity := int64(rest.Quantity) + quantity
	if newQuantity < 0 {
		return models.ErrQuantityLessThanZero
	}

	rest.Quantity = uint64(quantity) // Ugh go :(
	r.restaurants[name] = rest
	return nil
}

func (r *RestaurantService) RateRestaurant(name string, rating int, comment string) error {
	rest, registered := r.restaurants[name]
	if !registered {
		return models.ErrRestaurantNotRegistered
	}

	if rating < 0 {
		rating = 1
	} else if rating > 5 {
		rating = 5
	}

	var comm *string
	if comment != "" {
		comm = &comment
	}

	rest.Ratings = append(rest.Ratings, models.Rating{
		User:    r.currentUserID,
		Comment: comm,
		Rating:  rating,
	})

	return nil
}

func (r *RestaurantService) ShowRestaurants(orderBy models.OrderCrieteria) []models.ServiceableRestaurant {
	var rests []models.ServiceableRestaurant
	pincode := r.users[r.currentUserID].Pincode
	for _, restaurant := range r.restaurants {
		if restaurant.CanDeliver(pincode) && restaurant.Quantity > 0 {
			rests = append(rests, restaurant.ToServiceable())
		}
	}

	switch orderBy {
	case models.RatingsOrder:
		slices.SortFunc(rests, func(a, b models.ServiceableRestaurant) int {
			if a.Rating > b.Rating {
				return -1
			}

			if a.Rating < b.Rating {
				return 1
			}
			return 0
		})
	case models.PriceOrder:
		slices.SortFunc(rests, func(a, b models.ServiceableRestaurant) int {
			if a.Price > b.Price {
				return -1
			}

			if a.Price < b.Price {
				return 1
			}
			return 0
		})
	}

	return rests
}

func (r *RestaurantService) CreateNewOrder(name string, quantity uint64) error {
	rest, registered := r.restaurants[name]
	if !registered {
		return models.ErrRestaurantNotRegistered
	}

	userPincode := r.users[r.currentUserID].Pincode

	if !rest.CanDeliver(userPincode) {
		return models.ErrDoesNotDeliver
	}

	if rest.Quantity < quantity {
		return models.ErrNotEnoughQuantity
	}

	order := models.Order{
		RestaurantName: name,
		ItemName:       rest.FoodItem,
		Quantity:       quantity,
		Price:          float32(quantity) * rest.Price,
	}

	rest.Quantity -= quantity
	prevOrders := r.orders[r.currentUserID]
	prevOrders = append(prevOrders, order)

	r.restaurants[name] = rest
	r.orders[r.currentUserID] = prevOrders
	return nil
}

func (r *RestaurantService) ListOrders() []models.Order {
	return r.orders[r.currentUserID]
}
