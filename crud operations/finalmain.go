
package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"
type Employee struct {
	Id int 
	Name string 
}
func main() {


	insert(Employee{101, "SHARAN"})
	insert(Employee{102, "VENGY"})

	updateById(Employee{101, "SHARAN AARUMUGAM"})
	results := selectAll()
	for results.Next() {
		var employee Employee
		results.Scan(&employee.Id, &employee.Name)
		fmt.Println(employee.Id, employee.Name)
	}
	result := selectById(101)
	var employee Employee
	result.Scan(&employee.Id, &employee.Name)
	fmt.Println(employee.Id, employee.Name)
	delete(102)
}
func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1505@/go_db")
	if err != nil {
		fmt.Println("Error! Getting connection...")
	}
	return db;
}
func insert(employee Employee) {
	db := connect()
	insert, err := db.Query("INSERT INTO employee VALUES (?, ?)", employee.Id, employee.Name)
	if err != nil {
		fmt.Println("Error! Inserting records...")
	}
	defer insert.Close()
	defer db.Close()
}
func selectAll() *sql.Rows {
	db := connect()
	results, err := db.Query("SELECT * FROM employee")
	if err != nil {
		fmt.Println("Error! Getting records...")
	}
	defer db.Close()
	return results
}
func selectById(id int) *sql.Row {
	db := connect()
	result := db.QueryRow("SELECT * FROM employee WHERE id=?", id)
	defer db.Close()
	return result
}
func updateById(employee Employee) {
	db := connect()
	db.QueryRow("UPDATE employee SET name=? WHERE id=?", employee.Name, employee.Id)
}
func delete(id int) {
	db := connect()
	db.QueryRow("DELETE FROM employee WHERE id=?", id)
}