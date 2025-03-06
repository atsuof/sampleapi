package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// main tx sample
//func main() {
//	dbUser := "docker"
//	dbPassword := "docker"
//	dbName := "sampledb"
//	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)
//	db, dbOpenError := sql.Open("mysql", dbConn)
//	if dbOpenError != nil {
//		fmt.Println("error occurred", dbOpenError)
//		return
//	}
//	defer db.Close()
//
//	tx, beginErr := db.Begin()
//	if beginErr != nil {
//		fmt.Println(beginErr)
//		return
//	}
//	articleId := 1
//	const sqlStr = `
//	SELECT nice FROM articles WHERE article_id = ?;
//`
//	row := db.QueryRow(sqlStr, articleId)
//	if row.Err() != nil {
//		fmt.Println(row.Err())
//		tx.Rollback()
//		return
//	}
//
//	var niceNum int
//	scanError := row.Scan(&niceNum)
//	if scanError != nil {
//		fmt.Println(scanError)
//		tx.Rollback()
//		return
//	}
//
//	const sqlUpdateNiceNum = `
//	UPDATE articles set nice = ? WHERE article_id = ?;
//`
//	_, updateErr := db.Exec(sqlUpdateNiceNum, niceNum+1, articleId)
//
//	if updateErr != nil {
//		fmt.Println(updateErr)
//		tx.Rollback()
//		return
//	}
//	tx.Commit()
//}

// main select sample
//func main() {
//	dbUser := "docker"
//	dbPassword := "docker"
//	dbName := "sampledb"
//	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)
//
//	db, dbOpenError := sql.Open("mysql", dbConn)
//	if dbOpenError != nil {
//		fmt.Println("error occurred", dbOpenError)
//	}
//	defer db.Close()
//
//	articleId := 10
//
//	const sqlStr = `
//					SELECT * FROM articles
//					WHERE article_id = ?
//					;
//`
//	row := db.QueryRow(sqlStr, articleId)
//
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//defer rows.Close()
//
//	//articleSlice := []models.Article{}
//	var article models.Article
//	var createdAt sql.NullTime
//
//	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)
//
//	if createdAt.Valid {
//		article.CreatedAt = createdAt.Time
//	}
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Printf("%+v\n", article)
//
//	//for rows.Next() {
//	//	var article models.Article
//	//	var createdAt sql.NullTime
//	//	err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)
//	//
//	//	if createdAt.Valid {
//	//		article.CreatedAt = createdAt.Time
//	//	}
//	//
//	//	if err != nil {
//	//		fmt.Println(err)
//	//	} else {
//	//		articleSlice = append(articleSlice, article)
//	//	}
//	//}
//	//
//	//fmt.Printf("%+v\n", articleSlice)
//}

// main insert sample
//func main() {
//	dbUser := "docker"
//	dbPassword := "docker"
//	dbName := "sampledb"
//	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)
//
//	db, dbOpenError := sql.Open("mysql", dbConn)
//	if dbOpenError != nil {
//		fmt.Println("error occurred", dbOpenError)
//	}
//	defer db.Close()
//
//	//articleId := 10
//
//	const sqlStr = `
//					SELECT * FROM articles
//					WHERE article_id = ?
//					;
//`
//	const insertStr = `
//		insert into articles (title, contents, username, nice, created_at)
//		values (?, ?, ?, 0, now())
//		;
//`
//	article := models.Article{
//		Title:    "insert test",
//		Contents: "exec insert",
//		UserName: "test",
//	}
//	result, err := db.Exec(insertStr, article.Title, article.Contents, article.UserName)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(result.LastInsertId())
//	fmt.Println(result.RowsAffected())
//}
