package main

import (
	"database/sql"
	"fmt"
	"github.com/atsuof/sampleapi/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbName := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)

	db, dbOpenError := sql.Open("mysql", dbConn)
	if dbOpenError != nil {
		fmt.Println("error occurred", dbOpenError)
	}
	defer db.Close()

	const sqlStr = `select article_id,title,contents,username,nice,created_at from articles;`

	rows, err := db.Query(sqlStr)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleSlice := []models.Article{}
	for rows.Next() {
		var article models.Article
		var createdAt sql.NullTime
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)

		if createdAt.Valid {
			article.CreatedAt = createdAt.Time
		}

		if err != nil {
			fmt.Println(err)
		} else {
			articleSlice = append(articleSlice, article)
		}
	}

	fmt.Printf("%+v\n", articleSlice)
}
