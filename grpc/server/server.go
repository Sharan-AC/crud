package main

import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"context"
	"GRPC/proto"
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type server struct{}
var db *gorm.DB
var err error
func main(){
	db, _ = gorm.Open("mysql", "root:1505@tcp(127.0.0.1:3306)/grpc_crud?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Hosting server on: %s", listen.Addr().String())
}
func (s *server)CreateItem(ctx context.Context, Employee *pb.Employee) (*pb.ID,error){
	if Employee.ID == ""{
		return nil, status.Error(codes.InvalidArgument, "ID is empty, please try again")
	}
	db.Create(&Employee)
	return &pb.ID{ID: Employee.ID},nil
}