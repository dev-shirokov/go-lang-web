package repositories

import (
	"database/sql"
	"fmt"
	"ximlr/go-lang-web/database"
	"ximlr/go-lang-web/domain"

	_ "github.com/lib/pq"
)

type ArticleRepository struct {
}

func (r *ArticleRepository) GetById(id int) (*domain.ArticleModel, error) {

	dbContext := database.CreateContext()
	defer dbContext.Db.Close()

	sqlScript := fmt.Sprintf(`
		SELECT 
			id,
			title,
			content
		FROM articles
		WHERE id=%d`, id)

	var model = &domain.ArticleModel{}

	row := dbContext.Db.QueryRow(sqlScript)
	err := row.Scan(&model.Id, &model.Title, &model.Content)

	switch err {
	case sql.ErrNoRows:
		break
	case nil:
		break
	default:
		panic(err)
	}

	return model, err
}

func (r *ArticleRepository) Get() ([]*domain.ArticleModel, error) {

	dbContext := database.CreateContext()
	defer dbContext.Db.Close()

	sqlScript := `
		SELECT 
			id,
			title,
			content
		FROM articles
		ORDER BY id`

	rows, err := dbContext.Db.Query(sqlScript)
	switch err {
	case sql.ErrNoRows:
	case nil:
		break
	default:
		panic(err)
	}
	defer rows.Close()

	articles := make([]*domain.ArticleModel, 0)
	for rows.Next() {
		article := &domain.ArticleModel{}
		if err := rows.Scan(&article.Id, &article.Title, &article.Content); err != nil {
			panic(err)
		}

		articles = append(articles, article)
	}

	return articles, err
}

func (r *ArticleRepository) Create(model *domain.ArticleModel) error {
	dbContext := database.CreateContext()

	sqlScript := fmt.Sprintf("INSERT INTO articles (title, content) VALUES ('%s','%s')", model.Title, model.Content)

	_, err := dbContext.Db.Exec(sqlScript)
	if err != nil {
		panic(err)
	}

	defer dbContext.Db.Close()
	return err
}

func (r *ArticleRepository) Update(model *domain.ArticleModel) error {
	dbContext := database.CreateContext()

	sqlScript := fmt.Sprintf("UPDATE articles SET title='%s', content='%s' WHERE id=%d", model.Title, model.Content, &model.Id)
	_, err := dbContext.Db.Exec(sqlScript)
	if err != nil {
		panic(err)
	}

	defer dbContext.Db.Close()
	return err
}
