package main

import (
	"fmt"
	"log"
	"time"
    "GRPC/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)
func main(){
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewCRUDClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	fmt.Println("welcome to simple unary rpc implementation")
	fmt.Println("press 1 to create")
	var choice int
	fmt.Scanln(&choice) 
	if choice == 1 { 
		fmt.Println("enter the name")
		var name string
		fmt.Scanln(&name)
		fmt.Println("enter the ID")
		var id string
		fmt.Scanln(&id)
		item, err := c.CreateItem(ctx, &pb.Employee{Name: name, ID: id})
		if err != nil {
			log.Fatalf("Could not create a new item: %v", err)
		}
		fmt.Println("\nInserted", name, "with the ID", item.ID)
	}
	
}