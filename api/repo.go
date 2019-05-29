package api

import (
	"strings"

	"github.com/aelaster/go-taco/model"
)

var currentId int
var menuItems model.MenuItems

func init() {
	RepoCreateMenuItem(model.MenuItem{Name: "Veggie Taco", Price: 250})
	RepoCreateMenuItem(model.MenuItem{Name: "Chicken Taco", Price: 300})
	RepoCreateMenuItem(model.MenuItem{Name: "Beef Taco", Price: 300})
	RepoCreateMenuItem(model.MenuItem{Name: "Chorizo Taco", Price: 350})
}

func RepoCreateMenuItem(t model.MenuItem) model.MenuItem {
	currentId += 1
	t.Id = currentId
	menuItems = append(menuItems, t)
	return t
}

func RepoCalculateCost(items model.OrderItems) model.OrderCost {
	totalCost := 0
	totalQuantity := 0
	for _, item := range items {
		for _, menuItem := range menuItems {
			if strings.ToLower(menuItem.Name) == strings.ToLower(item.Name) {
				totalQuantity += item.Quantity
				totalCost += item.Quantity * menuItem.Price
			}
		}
	}

	if totalQuantity >= 4 {
		totalCost = (totalCost * 4) / 5
	}

	return model.OrderCost{TotalCost: totalCost, TotalQuantity: totalQuantity}
}
