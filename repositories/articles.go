package repositories

import (
	"database/sql"
	"github.com/atsuof/sampleapi/models"
	_ "github.com/go-sql-driver/mysql"
)

// InsertArticle inserts a new article into the database and returns the inserted article's details.
// Use consistent indentation and formatting for back quotes in SQL queries to improve visibility.
// It takes a database connection (`db`) and a `models.Article` object as inputs.
// Returns the complete article model upon successful insertion or an error if any occurs.
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const insertArticleQuery = `
	insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
`
	result, err := db.Exec(insertArticleQuery, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}
	articleId, idErr := result.LastInsertId()
	if idErr != nil {
		return models.Article{}, idErr
	}
	newArticle, selectErr := SelectArticleDetail(db, int(articleId))
	if selectErr != nil {
		return models.Article{}, selectErr
	}

	return newArticle, nil
}

func SelectArticleDetail(db *sql.DB, articleId int) (models.Article, error) {
	const articleDetailSelectQuery = `
	SELECT * FROM articles WHERE article_id = ?;
`

	row := db.QueryRow(articleDetailSelectQuery, articleId)
	if row.Err() != nil {
		return models.Article{}, row.Err()
	}

	var createdAt sql.NullTime

	article := models.Article{}

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const articlesNumPerPage = 5
	const selectArticleQuery = `
	SELECT * FROM articles limit ? offset ?
`
	rows, err := db.Query(selectArticleQuery, articlesNumPerPage, (page-1)*articlesNumPerPage)
	if err != nil {
		return nil, err
	}
	articles := []models.Article{}

	for rows.Next() {
		article := models.Article{}
		var createdAt sql.NullTime
		err = rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)
		if err != nil {
			return nil, err
		}
		if createdAt.Valid {
			article.CreatedAt = createdAt.Time
		}
		articles = append(articles, article)

	}

	return articles, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = ` select nice
        from articles
        where article_id = ?;	
`

	row := db.QueryRow(sqlGetNice, articleID)
	var niceNum int
	err := row.Scan(&niceNum)
	if err != nil {
		return err
	}
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`

	_, err = db.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		return err
	}
	return nil
}
