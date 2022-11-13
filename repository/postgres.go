package repository

import (
	"errors"
	"fmt"

	config "github.com/AntonyIS/modart"
	"github.com/AntonyIS/modart/app"
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

func NewItemRepository() (app.ItemRepository, error) {
	repo := postgresRepository{}
	db, err := newPostgresDB()
	if err != nil {
		return nil, err
	}
	repo.db = db
	repo.db.AutoMigrate(app.Article{}, app.Author{})
	return repo, nil
}

func (repo postgresRepository) Create(item interface{}) (*interface{}, error) {
	res := repo.db.Create(&item)
	if res.RowsAffected == 0 {
		return nil, errors.New("attendee not created")
	}
	return &item, nil
}

func (repo postgresRepository) Read(id string) (interface{}, error) {
	itemType := id[:4]
	if itemType == "auth-" {
		var author app.Author

		res := repo.db.First(&author, "id = ?", id)

		if res.RowsAffected == 0 {
			return nil, errors.New("author not found")
		}
		return author, nil

	}
	if itemType == "arti-" {
		var article app.Article

		res := repo.db.First(&article, "id = ?", id)

		if res.RowsAffected == 0 {
			return nil, errors.New("author not found")
		}
		return article, nil
	}

	return nil, nil

}

func (repo postgresRepository) Update(item interface{}) (interface{}, error) {
	articleValue, ok := item.(app.Article)
	if ok {
		var article app.Article
		res := repo.db.First(&article, "id = ?", articleValue.Id)
		if res.RowsAffected == 0 {
			return nil, errors.New("article not found")
		}

		res = repo.db.Model(article).Where("id = ?", articleValue.Id).Updates(article)

		if res.RowsAffected == 0 {
			return nil, errors.New("error updating article")
		}
		return &article, nil
	}

	authorValue, ok := item.(app.Author)
	if ok {
		var author app.Author
		res := repo.db.First(&author, "id = ?", authorValue.Id)
		if res.RowsAffected == 0 {
			return nil, errors.New("article not found")
		}

		res = repo.db.Model(author).Where("id = ?", authorValue.Id).Updates(author)

		if res.RowsAffected == 0 {
			return nil, errors.New("error updating article")
		}
		return &author, nil
	}

	return nil, nil
}

func (repo postgresRepository) Delete(id string) error {
	itemType := id[:4]
	if itemType == "auth-" {
		var author app.Author
		res := repo.db.Where("id = ?", id).Delete(&author)
		if res.RowsAffected == 0 {
			return errors.New("attendee not deleted")
		}
		return nil

	}
	if itemType == "arti-" {
		var article app.Article
		res := repo.db.Where("id = ?", id).Delete(&article)
		if res.RowsAffected == 0 {
			return errors.New("attendee not deleted")
		}
		return nil
	}

	return nil
}
