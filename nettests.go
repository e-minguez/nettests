package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type NetTest struct {
	Id               int
	InsertedDatetime time.Time
	Ping             float64
	Download         float64
	Upload           float64
}

func speedtest() (results string) {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "speedtest-cli"
	cmdArgs := []string{"--simple"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running speedtest-cli command: ", err)
		os.Exit(1)
	}
	results = string(cmdOut)
	return
}

func convert(results string) (ping, download, upload float64) {
	output := strings.Split(results, "\n")
	ping, _ = strconv.ParseFloat(strings.Split(output[0], " ")[1], 64)
	download, _ = strconv.ParseFloat(strings.Split(output[1], " ")[1], 64)
	upload, _ = strconv.ParseFloat(strings.Split(output[2], " ")[1], 64)
	return
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS nettests(
		Id integer PRIMARY KEY,
		InsertedDatetime DATETIME,
		Ping TEXT,
		Download TEXT,
		Upload TEXT
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func StoreItem(db *sql.DB, items []NetTest) {
	sql_additem := `
	INSERT OR REPLACE INTO nettests(
		Id,
		InsertedDatetime,
		Ping,
		Download,
		Upload
	) values(?, CURRENT_TIMESTAMP, ?, ?, ?)
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for _, item := range items {
		_, err2 := stmt.Exec(nil, item.Ping, item.Download, item.Upload)
		if err2 != nil {
			panic(err2)
		}
	}
}

func ReadItem(db *sql.DB) []NetTest {
	sql_readall := `
	SELECT * FROM nettests
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []NetTest
	for rows.Next() {
		item := NetTest{}
		err2 := rows.Scan(&item.Id, &item.InsertedDatetime, &item.Ping, &item.Download, &item.Upload)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}

func main() {
	ping, download, upload := convert(speedtest())

	const dbpath = "/data/nettests.db"

	db := InitDB(dbpath)
	defer db.Close()
	CreateTable(db)

	items := []NetTest{
		NetTest{0, time.Now(), ping, download, upload},
	}
	StoreItem(db, items)

	//readItems := ReadItem(db)
}
