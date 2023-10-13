package service

import (
	"errors"
	"testing"

	"github.com/keshavchand/flipkarLLD/models"
)

func TestRegisterUser(t *testing.T) {
	restaurants := NewRestaurantService()
	err := restaurants.RegisterUser(models.User{
		Name:        "User1",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-99",
		Pincode:     models.Pincode("BTM"),
	})

	if err != nil {
		t.Fatal(err)
	}

	err = restaurants.RegisterUser(models.User{
		Name:        "User1",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-99",
		Pincode:     models.Pincode("BTM"),
	})

	if !errors.Is(err, models.ErrPhoneNumberAlreadyRegistered) {
		t.Fatalf("expected %v got %v", models.ErrPhoneNumberAlreadyRegistered, err)
	}
}

func TestUserLogin(t *testing.T) {
	rest := NewRestaurantService()
	err := rest.RegisterUser(models.User{
		Name:        "User1",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-99",
		Pincode:     models.Pincode("BTM"),
	})

	if err != nil {
		t.Fatal(err)
	}

	rest.LoginUser("9999-9999-99")
}

func TestRegisterRestaurant(t *testing.T) {
	rest := NewRestaurantService()
	err := rest.RegisterUser(models.User{
		Name:        "User1",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-99",
		Pincode:     models.Pincode("BTM"),
	})

	if err != nil {
		t.Fatal(err)
	}

	rest.LoginUser("9999-9999-99")
	err = rest.RegisterRestaurant(models.Restaurant{
		Name:                "Restaurant1",
		ServiceablePinCodes: []models.Pincode{"HSR", "BTM"},
		FoodItem:            "NI Thali",
		Quantity:            5,
	})

	if err != nil {
		t.Fatal(err)
	}

	err = rest.RegisterRestaurant(models.Restaurant{
		Name:                "Restaurant1",
		ServiceablePinCodes: []models.Pincode{"HSR", "BTM"},
		FoodItem:            "NI Thali",
		Quantity:            5,
	})

	if !errors.Is(err, models.ErrRestaurantAlreadyRegistered) {
		t.Fatalf("Expected %v got %v", models.ErrRestaurantAlreadyRegistered, err)
	}
}

func TestRestuarantIncreaseQuantity(t *testing.T) {
	rest := NewRestaurantService()
	rest.RegisterUser(models.User{
		Name:        "User1",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-99",
		Pincode:     models.Pincode("BTM"),
	})

	rest.LoginUser("9999-9999-99")
	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Restaurant1",
		ServiceablePinCodes: []models.Pincode{"HSR", "BTM"},
		FoodItem:            "NI Thali",
		Quantity:            5,
	})

	err := rest.UpdateQuantity("Restaurant1", 5)
	if err != nil {
		t.Fatal(err)
	}

	err = rest.UpdateQuantity("Restaurant1", -100)
	if !errors.Is(err, models.ErrQuantityLessThanZero) {
		t.Fatalf("expected %v got %v", models.ErrQuantityLessThanZero, err)
	}

	rest.RegisterUser(models.User{
		Name:        "User2",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-98",
		Pincode:     models.Pincode("BTM"),
	})

	rest.LoginUser("9999-9999-98")
	err = rest.UpdateQuantity("Restaurant1", -100)
	if !errors.Is(err, models.ErrUnauthorizedUser) {
		t.Fatalf("expected %v got %v", models.ErrUnauthorizedUser, err)
	}

	err = rest.UpdateQuantity("Restaurant2", -100)
	if !errors.Is(err, models.ErrRestaurantNotRegistered) {
		t.Fatalf("expected %v got %v", models.ErrRestaurantNotRegistered, err)
	}
}

func TestUserRating(t *testing.T) {
	rest := NewRestaurantService()
	rest.RegisterUser(models.User{
		Name:        "User1",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-99",
		Pincode:     models.Pincode("BTM"),
	})

	rest.LoginUser("9999-9999-99")
	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Restaurant1",
		ServiceablePinCodes: []models.Pincode{"HSR", "BTM"},
		FoodItem:            "NI Thali",
		Quantity:            5,
	})

	err := rest.RateRestaurant("Restaurant1", 1, "bad restaurant")
	if err != nil {
		t.Fatal(err)
	}

	err = rest.RateRestaurant("Restaurant2", 1, "bad restaurant")
	if !errors.Is(err, models.ErrRestaurantNotRegistered) {
		t.Fatalf("Expected %v got %v", models.ErrRestaurantNotRegistered, err)
	}
}

func TestListRestaurants(t *testing.T) {
	rest := NewRestaurantService()
	rest.RegisterUser(models.User{
		Name:        "User1",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-99",
		Pincode:     models.Pincode("BTM"),
	})

	rest.LoginUser("9999-9999-99")
	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Restaurant1",
		ServiceablePinCodes: []models.Pincode{"HSR", "BTM"},
		FoodItem:            "NI Thali",
		Quantity:            5,
		Price:               10,
	})

	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Restaurant2",
		ServiceablePinCodes: []models.Pincode{"HSR", "BTM"},
		FoodItem:            "NI Thali",
		Quantity:            5,
		Price:               50,
	})

	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Restaurant3",
		ServiceablePinCodes: []models.Pincode{"HSR"},
		FoodItem:            "NI Thali",
		Quantity:            5,
		Price:               50,
	})

	rest.RateRestaurant("Restaurant1", 5, "")
	rest.RateRestaurant("Restaurant2", 1, "")

	serviceablerestaurants := rest.ShowRestaurants(models.RatingsOrder)
	expected := []string{"Restaurant1", "Restaurant2"}
	if len(serviceablerestaurants) != len(expected) {
		t.Fatalf("Wrong serviceable restaurant exprected %d got %d", len(expected), len(serviceablerestaurants))
	}

	for i := 0; i < len(expected); i++ {
		if serviceablerestaurants[i].Name != expected[i] {
			t.Errorf("Expected pos %d be %s got %s", i, expected[i], serviceablerestaurants[i].Name)
		}
	}

	serviceablerestaurants = rest.ShowRestaurants(models.PriceOrder)

	expected = []string{"Restaurant2", "Restaurant1"}
	if len(serviceablerestaurants) != len(expected) {
		t.Fatalf("Wrong serviceable restaurant exprected %d got %d", len(expected), len(serviceablerestaurants))
	}

	for i := 0; i < len(expected); i++ {
		if serviceablerestaurants[i].Name != expected[i] {
			t.Fatalf("Expected %d be %s got %s", i, expected[i], serviceablerestaurants[i].Name)
		}
	}
}

func TestCreateNewOrder(t *testing.T) {
	rest := NewRestaurantService()
	rest.RegisterUser(models.User{
		Name:        "User1",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-99",
		Pincode:     models.Pincode("BTM"),
	})

	rest.RegisterUser(models.User{
		Name:        "User2",
		Gender:      models.Male,
		PhoneNumber: "9999-9999-89",
		Pincode:     models.Pincode("HSR"),
	})

	rest.LoginUser("9999-9999-99")
	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Restaurant1",
		ServiceablePinCodes: []models.Pincode{"BTM"},
		FoodItem:            "NI Thali",
		Quantity:            5,
		Price:               10,
	})

	err := rest.CreateNewOrder("Restaurant1", 5)
	if err != nil {
		t.Fatalf("Expected nil got %v", err)
	}

	err = rest.CreateNewOrder("Restaurant1", 5)
	if !errors.Is(models.ErrNotEnoughQuantity, err) {
		t.Fatalf("Expected %v got %v", models.ErrNotEnoughQuantity, err)
	}

	err = rest.CreateNewOrder("Restaurant0", 5)
	if !errors.Is(models.ErrRestaurantNotRegistered, err) {
		t.Fatalf("Expected %v got %v", models.ErrRestaurantNotRegistered, err)
	}

	rest.LoginUser("9999-9999-89")
	err = rest.CreateNewOrder("Restaurant1", 5)
	if !errors.Is(models.ErrDoesNotDeliver, err) {
		t.Fatalf("Expected %v got %v", models.ErrDoesNotDeliver, err)
	}
}
