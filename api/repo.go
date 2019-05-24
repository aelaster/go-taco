package api

import (
	"github.com/aelaster/go-taco/model"
	"strings"
)

var currentId int

var menuitems model.MenuItems

func init() {
	RepoCreateMenuItem(model.MenuItem{Name: "Veggie Taco", Price: 250})
	RepoCreateMenuItem(model.MenuItem{Name: "Chicken Taco", Price: 300})
	RepoCreateMenuItem(model.MenuItem{Name: "Beef Taco", Price: 300})
	RepoCreateMenuItem(model.MenuItem{Name: "Chorizo Taco", Price: 350})
}

func RepoCreateMenuItem(t model.MenuItem) model.MenuItem {
	currentId += 1
	t.Id = currentId
	menuitems = append(menuitems, t)
	return t
}

func RepoCalculateCost(items model.OrderItems) model.OrderCost {
	var totalCost = 0
	var totalQuantity = 0
	for _, item := range items {
		for _, menuitem := range menuitems {
			if strings.ToLower(menuitem.Name) == strings.ToLower(item.Name) {
				totalQuantity+=item.Quantity
				totalCost+=item.Quantity*menuitem.Price
			}
		}
	}

	if totalQuantity >= 4 {
		totalCost = (totalCost * 4) / 5
	}

	return model.OrderCost{TotalCost: totalCost, TotalQuantity: totalQuantity}
}
