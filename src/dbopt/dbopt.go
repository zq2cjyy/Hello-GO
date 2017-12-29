package dbopt

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"time"
)

func Query() {
	db, err := sql.Open("mysql", "root:1985@tcp(localhost:3306)/babycare")
	defer db.Close()
	checkErr(err)
	rows, err := db.Query("select id,begintime,endtime from sleep where id = ?", 1)
	checkErr(err)
	for rows.Next() {
		var id int
		var begin string
		var end string
		err = rows.Scan(&id, &begin, &end)
		begindate, err := time.Parse("2006-1-02 15:04:05", begin)
		if err != nil {
			fmt.Println(err)
		}
		checkErr(err)
		fmt.Printf("id=%d begin:%s", id, begindate.Format("2006-1-2 15:04:05"))
	}
}

func checkErr(err error) {
	if (err != nil) {
		panic(err)
	}
}
