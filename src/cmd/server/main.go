package main

import (
	"cat_alog/src/infrastructure/cassandra"
	"fmt"
)

func main() {
	cassandraCheck, err := cassandra.CheckCassandraConnection()
	if cassandraCheck {
		fmt.Println("Cassandra connection successful")
	} else if err != nil {
		fmt.Println("Cassandra connection failed: " + err.Error())
	}
}
