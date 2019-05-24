__CODE CHALLENGE__

Taco Loco, a quick service fleet of taco trucks, is building a new mobile app to enable their customers to place orders for their delicious tacos. They’ve hired Detroit Labs to build the backend services to power this app.

As a backend developer, you’ve been tasked with creating a web service to calculate order totals. The service will take as input the items and quantities ordered, and respond with the order total.

Taco Loco’s menu consists of the the following items:

Veggie Taco, $2.50 ea.

Chicken or Beef Taco, $3.00 ea.

Chorizo Taco, $3.50 ea.

If a customer orders 4 or more tacos, then a 20% discount should be applied to the entire order.

Taco Loco wants you to define the data format for the request and response, and prefers using JSON.

---
Run `go run *.go` to launch the server.

__ENDPOINT:__
```POST http://localhost:8080/calculate```

__REQUEST:__
```json
[
    {
        "name": "Veggie Taco",
        "quantity": 1
    },
    {
        "name": "Chicken Taco",
        "quantity": 2
    },
    {
        "name": "Beef Taco",
        "quantity": 3
    },
    {
        "name": "Chorizo Taco",
        "quantity": 4
    }
]
```

__RESPONSE:__
```json
{
    "totalcost": 2520,
    "totalquantity": 10
}
```
Total cost is returned in cents.