package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/AntonyIS/modart/app"
	config "github.com/AntonyIS/modart/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postgresRepository struct {
	db *gorm.DB
}

func newPostgresDB() (*gorm.DB, error) {
	var (
		host     = config.GetEnvVariable("DB_HOST")
		port     = config.GetEnvVariable("DB_PORT")
		user     = config.GetEnvVariable("DB_USER")
		dbname   = config.GetEnvVariable("DB_NAME")
		password = config.GetEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open("postgres", conn)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.DB().SetMaxIdleConns(30)
	if err != nil {
		return nil, err
	}
	return db, nil

}

func NewAuthorRepository() (app.AuthorRepository, error) {
	repo := postgresRepository{}
	db, err := newPostgresDB()
	if err != nil {
		return nil, err
	}

	repo.db = db
	repo.db.AutoMigrate(app.Author{})
	return repo, nil
}

// func NewArticleRepository() (app.ArticleRepository, error) {
// 	repo := postgresRepository{}
// 	db, err := newPostgresDB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	repo.db = db
// 	repo.db.AutoMigrate(app.Article{})
// 	return repo, nil
// }

func (repo postgresRepository) CreateAuthor(author *app.Author) (*app.Author, error) {
	password, err := author.GenerateHashPassord()
	if err != nil {
		return nil, errors.New("error harshing password")
	}
	author.Password = password
	res := repo.db.Create(&author)
	if res.RowsAffected == 0 {
		return nil, errors.New("attendee not created")
	}
	return author, nil
}

func (repo postgresRepository) ReadAuthor(id string) (*app.Author, error) {
	var author app.Author
	res := repo.db.First(&author, id)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &author, nil
}

func (repo postgresRepository) ReadAuthorAll() ([]*app.Author, error) {

	var authors []*app.Author
	res := repo.db.Find(&authors)

	if res.Error != nil {
		return nil, errors.New("authors not found")
	}

	return authors, nil
}

func (repo postgresRepository) UpdateAuthor(author *app.Author) (*app.Author, error) {
	var updateAuthor app.Author
	result := repo.db.Model(&updateAuthor).Where("id = ?", author.ID).Updates(author)
	if result.RowsAffected == 0 {
		return &app.Author{}, errors.New("author not updated")
	}
	return &updateAuthor, nil
}

func (repo postgresRepository) DeleteAuthor(id string) error {
	var deletedauthor app.Author
	result := repo.db.Where("id = ?", id).Delete(&deletedauthor)
	if result.RowsAffected == 0 {
		return errors.New("author data not deleted")
	}
	return nil
}

func (repo postgresRepository) CreateArticle(article *app.Article) (*app.Article, error) {

	res := repo.db.Create(&article)
	if res.RowsAffected == 0 {
		return nil, errors.New("attendee not created")
	}

	return article, nil
}

func (repo postgresRepository) ReadArticle(id string) (*app.Article, error) {

	var article *app.Article

	res := repo.db.First(article, "id = ?", id)

	if res.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}
	return article, nil
}

func (repo postgresRepository) ReadArticleAll() ([]*app.Article, error) {

	var articles []*app.Article
	res := repo.db.Find(&articles)

	if res.Error != nil {
		return nil, errors.New("authors not found")
	}
	return articles, nil
}

func (repo postgresRepository) UpdateArticle(article *app.Article) (*app.Article, error) {

	var _author app.Article
	res := repo.db.First(&_author, "id = ?", article.ID)

	if res.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}

	res = repo.db.Model(_author).Where("id = ?", article.ID).Updates(_author)

	if res.RowsAffected == 0 {
		return nil, errors.New("error updating article")
	}
	return article, nil
}

func (repo postgresRepository) DeleteArticle(id string) error {

	var article app.Article
	res := repo.db.Where("id = ?", id).Delete(&article)

	if res.RowsAffected == 0 {
		return errors.New("article not deleted")
	}
	return nil
}
