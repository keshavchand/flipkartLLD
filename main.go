package main

import (
	"fmt"
	"log"

	"github.com/keshavchand/flipkarLLD/models"
	"github.com/keshavchand/flipkarLLD/service"
)

func main() {

	var err error
	rest := service.NewRestaurantService()
	rest.RegisterUser(models.User{
		Name:        "Pralove",
		Gender:      models.Male,
		PhoneNumber: "phoneNumber-1",
		Pincode:     "HSR",
	})
	rest.RegisterUser(models.User{
		Name:        "Nitesh",
		Gender:      models.Male,
		PhoneNumber: "phoneNumber-2",
		Pincode:     "BTM",
	})
	rest.RegisterUser(models.User{
		Name:        "Vatsal",
		Gender:      models.Male,
		PhoneNumber: "phoneNumber-3",
		Pincode:     "BTM",
	})

	err = rest.LoginUser("phoneNumber-1")
	if err != nil {
		log.Println("Login Error:", err)
	}
	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Food Court-1",
		ServiceablePinCodes: []models.Pincode{"BTM", "HSR"},
		FoodItem:            "NI Thali",
		Quantity:            5,
		Price:               100,
	})

	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Food Court-2",
		ServiceablePinCodes: []models.Pincode{"BTM"},
		FoodItem:            "Burger",
		Quantity:            3,
		Price:               120,
	})

	rest.LoginUser("phoneNumber-2")
	rest.RegisterRestaurant(models.Restaurant{
		Name:                "Food Court-3",
		ServiceablePinCodes: []models.Pincode{"HSR"},
		FoodItem:            "SI Thali",
		Quantity:            1,
		Price:               150,
	})

	err = rest.LoginUser("phoneNumber-3")
	if err != nil {
		log.Println("Login Error:", err)
	}

	serviceableRestaurants := rest.ShowRestaurants(models.PriceOrder)
	for _, rests := range serviceableRestaurants {
		fmt.Println(rests.Name, rests.FoodItem)
	}

	err = rest.CreateNewOrder("Food Court-1", 2)
	if err != nil {
		log.Println("Cannot place order: ", err.Error())
	} else {
		log.Println("Order placed successfully")
	}

	err = rest.CreateNewOrder("Food Court-2", 7)
	if err != nil {
		log.Println("Cannot place order: ", err.Error())
	} else {
		log.Println("Order placed successfully")
	}

	rest.RateRestaurant("Food Court-2", 3, "Good Food")
	rest.RateRestaurant("Food Court-1", 5, "Good Food")

	serviceableRestaurants = rest.ShowRestaurants(models.RatingsOrder)
	for _, rests := range serviceableRestaurants {
		fmt.Println(rests.Name, rests.FoodItem)
	}

	err = rest.LoginUser("phoneNumber-1")
	if err != nil {
		log.Println("Login Error:", err)
	}

	rest.UpdateQuantity("Food Court-2", 5)
}
