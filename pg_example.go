package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strings"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tester"
	password = ""
	dbname   = "somedb"
)

type dbShardConfig struct {
	ID            int            `db:"id"`
	CompanyCode   string         `db:"company_code"`
	Environment   string         `db:"environment"`
	Host          string         `db:"host"`
	Adapter       string         `db:"adapter"`
	Database      string         `db:"database"`
	Username      string         `db:"username"`
	Password      sql.NullString `db:"password"`
	Port          string         `db:"port"`
	CreatedAt     *time.Time     `db:"created_at"`
	UpdatedAt     *time.Time     `db:"updated_at"`
	Pool          *int           `db:"pool"`
	Status        int            `db:"status"`
	RestartStatus int            `db:"restart_status"`
}

func main() {

	psqlInfo := buildConInfo()

	fmt.Print(psqlInfo)
	// db, err := sql.Open("postgres", psqlInfo)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	rows, err := db.Queryx("select * from db_shard_configs ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows.Columns())

	dsc := dbShardConfig{}
	dscList := [100]dbShardConfig{}
	results := make([]map[string]interface{}, 0)

	for rows.Next() {
		//err = rows.StructScan(&dsc)
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Printf("%#v\n", dsc)
		res := make(map[string]interface{})
		err = rows.MapScan(res)
		err = rows.StructScan(&dsc)

		fmt.Printf("%v\n", dsc)

		results = append(results, res)
		//
		//for k, v := range res {
		//	fmt.Printf("%s => %v\n", k, v)
		//}

	}

	for _, res := range results {
		fmt.Printf("%v\n", res["company_code"])
	}

	err = db.Get(&dsc, "select * from db_shared_configs")
	fmt.Printf("??????????%v\n", dsc)

	err = db.Select(&dscList, "select * from db_shared_configs where id > 10")
	fmt.Printf("+++++++++%v\n", dscList)

}

func buildConInfo() string {

	var buf bytes.Buffer

	if len(strings.TrimSpace(host)) > 0 {
		buf.WriteString(fmt.Sprintf("host=%s ", host))
	}

	buf.WriteString(fmt.Sprintf("port=%d ", port))

	if len(strings.TrimSpace(host)) > 0 {
		buf.WriteString(fmt.Sprintf("host=%s ", host))
	}

	if len(strings.TrimSpace(user)) > 0 {
		buf.WriteString(fmt.Sprintf("user=%s ", user))
	}

	if len(strings.TrimSpace(password)) > 0 {
		buf.WriteString(fmt.Sprintf("password=%s ", password))
	}

	if len(strings.TrimSpace(dbname)) > 0 {
		buf.WriteString(fmt.Sprintf("dbname=%s ", dbname))
	}

	buf.WriteString("sslmode=disable")

	return buf.String()
}
