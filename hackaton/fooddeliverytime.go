package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burger := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}

	if order == "burger" {
		return burger.preptime
	}
	if order == "chips" {
		return chips.preptime
	}
	if order == "nuggets" {
		return nuggets.preptime
	}
	return 404
}
