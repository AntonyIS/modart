package repository

import (
	"errors"
	"fmt"

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

func NewArticleRepository() (app.ArticleRepository, error) {
	repo := postgresRepository{}
	db, err := newPostgresDB()
	if err != nil {
		return nil, err
	}

	repo.db = db
	repo.db.AutoMigrate(app.Article{})
	return repo, nil
}

func (repo postgresRepository) CreateAuthor(author *app.Author) (*app.Author, error) {
	res := repo.db.Create(&author)
	if res.RowsAffected == 0 {
		return nil, errors.New("attendee not created")
	}
	defer repo.db.Close()
	return author, nil
}

func (repo postgresRepository) ReadAuthor(id string) (*app.Author, error) {

	var author *app.Author

	res := repo.db.First(author, "id = ?", id)
	defer repo.db.Close()
	if res.RowsAffected == 0 {
		return nil, errors.New("author not found")
	}
	return author, nil
}

func (repo postgresRepository) ReadAuthorAll() ([]*app.Author, error) {
	fmt.Println("Read all authors")
	var authors []*app.Author
	res := repo.db.Find(&authors)
	
	if res.Error != nil {
		return nil, errors.New("authors not found")
	}
	return authors, nil
}

func (repo postgresRepository) UpdateAuthor(author *app.Author) (*app.Author, error) {
	var article app.Article
	res := repo.db.First(&article, "id = ?", author.Id)
	defer repo.db.Close()
	if res.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}

	res = repo.db.Model(article).Where("id = ?", author.Id).Updates(article)

	if res.RowsAffected == 0 {
		return nil, errors.New("error updating article")
	}
	return author, nil
}

func (repo postgresRepository) DeleteAuthor(id string) error {
	var author app.Author
	res := repo.db.Where("id = ?", id).Delete(&author)
	defer repo.db.Close()
	if res.RowsAffected == 0 {
		return errors.New("author not deleted")
	}
	return nil
}

func (repo postgresRepository) CreateArticle(article *app.Article) (*app.Article, error) {
	res := repo.db.Create(&article)
	if res.RowsAffected == 0 {
		return nil, errors.New("attendee not created")
	}
	defer repo.db.Close()
	return article, nil
}

func (repo postgresRepository) ReadArticle(id string) (*app.Article, error) {

	var article *app.Article

	res := repo.db.First(article, "id = ?", id)
	defer repo.db.Close()
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
	res := repo.db.First(&_author, "id = ?", article.Id)
	defer repo.db.Close()
	if res.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}

	res = repo.db.Model(_author).Where("id = ?", article.Id).Updates(_author)

	if res.RowsAffected == 0 {
		return nil, errors.New("error updating article")
	}
	return article, nil
}

func (repo postgresRepository) DeleteArticle(id string) error {
	var article app.Article
	res := repo.db.Where("id = ?", id).Delete(&article)
	defer repo.db.Close()
	if res.RowsAffected == 0 {
		return errors.New("article not deleted")
	}
	return nil
}
