// https://tutorialedge.net/golang/golang-mysql-tutorial/

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type Tag struct {
	ID   int    `json:"id",db:"id"`
	Name string `json:"name",db:"name"`
}

func main() {

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/peatio_development")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO `tags`(`name`) VALUES ( 'SVIP' )")

	// if there is an error inserting, handle it
	if err != nil {
		log.Printf("ignore this err: %+v", err)
	} else {
		defer insert.Close()
	}
	// be careful deferring Queries if you are using transactions


	results, err := db.Query("SELECT id, name FROM tags")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}


	var tag1 Tag
	// Execute the query
	err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag1.ID, &tag1.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(tag1.ID)
	log.Println(tag1.Name)

}
