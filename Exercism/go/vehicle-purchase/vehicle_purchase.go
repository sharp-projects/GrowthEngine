package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	const message = "is clearly the better choice."
	if option1 < option2 {
		return fmt.Sprintf("%s %s", option1, message)
	}
	return fmt.Sprintf("%s %s", option2, message)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .8
	} else if age >= 10 {
		return originalPrice * .5
	} else {
		return originalPrice * .7
	}
}
