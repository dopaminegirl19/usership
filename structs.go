package main

// User struct info contains data fro 1 row of users table
type User struct {
	ID       int64
	Name     string
	Contact  string
	Contact2 string
}

// Item struct info contains data fro 1 row of items table
type Item struct {
	ID            int64
	Name          string
	Description   string
	CurrentUserID int64
}
