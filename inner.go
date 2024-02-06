package main

import (
	"fmt"
	"log"
)

func join() {
	query := `
	SELECT orders.orderid, customer.customername
	FROM orders
	FULL OUTER JOIN customer ON orders.customerid = customer.customerid;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			customername, orderid any
		)

		err := rows.Scan(&orderid, &customername)
		if err != nil {
			log.Fatal(err)
		}
		if customername == nil {
			customername = ""
		}
		if orderid == nil {
			orderid = 0
		}
		fmt.Printf("orderid: %d, customername: %s\n", orderid, customername)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
