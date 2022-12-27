package app

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrNotFound = errors.New("item not found")
	ErrInvalid  = errors.New("item invalid")
)

type appService struct {
	appRepo AppRepository
}

func NewItemService(appRepo AppRepository) AppService {
	return &appService{
		appRepo,
	}
}

// Author service methods
func (a *appService) CreateAuthor(author *Author) (*Author, error) {
	if err := validate.Validate(author); err != nil {
		return nil, errs.Wrap(ErrInvalid, "service.Author.Create")
	}
	// author.ID = uuid.New().String()
	return a.appRepo.CreateAuthor(author)
}

func (a *appService) LoginAuthor(email string) (*Author, error) {
	// author.ID = uuid.New().String()
	return a.appRepo.LoginAuthor(email)
}

func (a *appService) ReadAuthor(id string) (*Author, error) {
	return a.appRepo.ReadAuthor(id)
}

func (a *appService) ReadAuthors() ([]*Author, error) {
	return a.appRepo.ReadAuthors()
}

func (a *appService) UpdateAuthor(author *Author) (*Author, error) {
	return a.appRepo.UpdateAuthor(author)
}

func (a *appService) DeleteAuthor(id string) error {
	return a.appRepo.DeleteAuthor(id)
}

// Article service methods
func (a *appService) CreateArticle(article *Article) (*Article, error) {
	if err := validate.Validate(article); err != nil {
		return nil, errs.Wrap(ErrInvalid, "service.Article.Create")
	}
	// article.ID = uuid.New().String()
	article.CreateAt = time.Now().UTC().Unix()
	return a.appRepo.CreateArticle(article)
}

func (a *appService) ReadArticle(id string) (*Article, error) {
	return a.appRepo.ReadArticle(id)
}

func (a *appService) ReadArticles() ([]*Article, error) {
	return a.appRepo.ReadArticles()
}

func (a *appService) UpdateArticle(article *Article) (*Article, error) {
	return a.appRepo.UpdateArticle(article)
}

func (a *appService) DeleteArticle(id string) error {
	return a.appRepo.DeleteArticle(id)
}
