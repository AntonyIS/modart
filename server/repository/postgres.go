package repository

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	app "example.com/server/app"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

type postgresRepository struct {
	db *gorm.DB
}

func newPostgresDB() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var (
		host     = os.Getenv("DATABASE_HOST")
		port     = os.Getenv("DATABASE_PORT")
		user     = os.Getenv("POSTGRES_USER")
		dbname   = os.Getenv("POSTGRES_DB")
		password = os.Getenv("POSTGRES_PASSWORD")
	)
	fmt.Println("##############################")
	fmt.Println(host, port, user, dbname, password)

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
	db.AutoMigrate(app.Author{})
	db.AutoMigrate(app.Article{})

	if err != nil {
		return nil, err
	}
	return db, nil

}

func NewPostgresqlDB() app.AppRepository {
	db, err := newPostgresDB()
	if err != nil {
		log.Fatal("DB ERROR: ", err)
	}
	repo := postgresRepository{
		db: db,
	}
	return repo
}

func (r postgresRepository) CreateAuthor(author *app.Author) (*app.Author, error) {
	password, err := author.GenerateHashPassord()
	if err != nil {
		return nil, errors.New("error harshing password")
	}
	author.Password = password
	author.Id = uuid.New().String()
	author.Articles = []app.Article{}
	res := r.db.Create(&author)
	if res.RowsAffected == 0 {
		return nil, errors.New("attendee not created")
	}
	return author, nil
}

func (r postgresRepository) ReadAuthor(id string) (*app.Author, error) {
	var author app.Author
	res := r.db.First(&author, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("author with ID :%s not found", id))
	}
	return &author, nil
}

func (r postgresRepository) ReadAuthors() ([]*app.Author, error) {
	var authors []*app.Author
	res := r.db.Find(&authors)
	if res.Error != nil {
		return nil, errors.New("authors not found")
	}
	return authors, nil
}

func (r postgresRepository) UpdateAuthor(author *app.Author) (*app.Author, error) {
	var updateAuthor app.Author
	result := r.db.Model(&updateAuthor).Where("id = ?", author.Id).Updates(author)
	if result.RowsAffected == 0 {
		return &app.Author{}, errors.New("author not updated")
	}
	return &updateAuthor, nil
}

func (r postgresRepository) DeleteAuthor(id string) error {
	var deletedAuthor app.Author
	result := r.db.Where("id = ?", id).Delete(&deletedAuthor)
	if result.RowsAffected == 0 {
		return errors.New("author data not deleted")
	}
	return nil
}

func (r postgresRepository) LoginAuthor(email string) (*app.Author, error) {
	var author *app.Author
	DB.Where("email= ?", email).First(author)
	if author.Email == " " {
		return nil, errors.New("user not found")
	}
	return author, nil
}

func (r postgresRepository) CreateArticle(article *app.Article) (*app.Article, error) {
	article.Id = uuid.New().String()
	res := r.db.Create(&article)
	if res.RowsAffected == 0 {
		return nil, errors.New("article not created")
	}

	return article, nil
}

func (r postgresRepository) ReadArticle(id string) (*app.Article, error) {
	var article app.Article

	res := r.db.First(&article, "id = ?", id)

	if res.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}
	return &article, nil
}

func (r postgresRepository) ReadArticles() ([]*app.Article, error) {
	var articles []*app.Article
	res := r.db.Find(&articles)
	if res.Error != nil {
		return nil, errors.New("articles not found")
	}
	return articles, nil
}

func (r postgresRepository) UpdateArticle(article *app.Article) (*app.Article, error) {
	var updateArticle app.Article
	result := r.db.Model(&updateArticle).Where(article.AuthorID).Updates(article)
	if result.RowsAffected == 0 {
		return &app.Article{}, errors.New(fmt.Sprintf("article with ID: %v not updated", article.AuthorID))
	}
	return &updateArticle, nil
}

func (r postgresRepository) DeleteArticle(id string) error {
	var deletedArticle app.Article
	result := r.db.Where("id = ?", id).Delete(&deletedArticle)
	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("article with ID: %s not found", id))
	}
	return nil
}
