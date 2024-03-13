package piscine

type food struct {
	name     string
	preptime int
}

func FoodDeliveryTime(order string) int {
	foodPreparation := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}
	item, available := foodPreparation[order]
	if !available {
		return 404
	}
	return item.preptime
}
