package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Exzrgs/myapi/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, username, nice, created_at) values
	(?,?,?,0,now());
	`
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
	select article_id, title, contents, username, nice
	from articles
	limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		fmt.Println("error at Query in SelectArticleList")
		return nil, err
	}

	defer rows.Close()

	articleArray := make([]models.Article, 0)

	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
	select *
	from articles
	where article_id = ?;
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println("error at QueryRow in SelectArticleDetail")
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println("error at row.Scan in SelectArticleDetail")
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
	select nice
	from articles
	where article_id = ?;
	`

	const sqlUpdateNice = `
	update articles
	set nice = ?
	where article_id = ?;
	`

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error at bigen tx in UpdateNiceNum")
		fmt.Println(err)
		return err
	}

	var nice int
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		fmt.Println("error at getting query in UpdateNiceNum")
		return err
	}

	err = row.Scan(&nice)
	if err != nil {
		tx.Rollback()
		fmt.Println("error at scanning row in UpdateNiceNum")
		/* デバッグ用
		fmt.Printf("nice is %d\n", nice)
		fmt.Printf("articleID is %d\n", articleID)
		*/
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, nice+1, articleID)
	if err != nil {
		tx.Rollback()
		fmt.Println("error at exec niceNum in UpdateNiceNum")
		return err
	}

	if err := tx.Commit(); err != nil {
		fmt.Println("error at commit tx in UpdateNiceNum")
		return err
	}

	return nil
}
