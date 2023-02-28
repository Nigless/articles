package article

import (
	"rest-api/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Article struct {
	*sqlx.DB
}

func New(db *sqlx.DB) *Article {
	return &Article{db}
}

func (self *Article) Get(id int64) (entity.Article, error) {
	var article entity.Article
	return article, self.QueryRowx("SELECT * FROM article WHERE id=$1", id).StructScan(&article)
}

func (self *Article) Create(article entity.Article) (int64, error) {
	res, err := self.NamedQuery(`
		INSERT INTO article (title, content)
		VALUES (:title, :content)
    	RETURNING id;
	`, article)
	if err != nil {
		return 0, err
	}
	res.Next()
	var id int64
	res.Scan(&id)
	return id, nil
}
