package main

var currentId int

var menuitems MenuItems

// Give us some seed data
func init() {
	RepoCreateMenuItem(MenuItem{Name: "Veggie Taco", Price: 250})
	RepoCreateMenuItem(MenuItem{Name: "Chicken Taco", Price: 300})
	RepoCreateMenuItem(MenuItem{Name: "Beef Taco", Price: 300})
	RepoCreateMenuItem(MenuItem{Name: "Chorizo Taco", Price: 350})
}

func RepoCreateMenuItem(t MenuItem) MenuItem {
	currentId += 1
	t.Id = currentId
	menuitems = append(menuitems, t)
	return t
}

func RepoCalculateCost(items OrderItems) OrderCost {
	var totalCost = 0.0
	var totalQuantity = 0
	for _, item := range items {
		for _, menuitem := range menuitems {
			if menuitem.Name == item.Name {
				totalQuantity+=item.Quantity
				totalCost+=float64(item.Quantity)*float64(menuitem.Price)
			}
		}
	}

	if totalQuantity<4 {
		totalCost*=.8
	}

	return OrderCost{TotalCost: totalCost/100, TotalQuantity: totalQuantity}
}
