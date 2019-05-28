package model

type MenuItem struct {
	Id        int      	`json:"id"`
	Name      string    	`json:"name"`
	Price	  int   	`json:"price"` //in cents
}

type MenuItems []MenuItem
