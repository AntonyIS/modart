package app

import (
	"errors"
	"time"

	"github.com/google/uuid"
	errs "github.com/pkg/errors"
	"gopkg.in/dealancer/validate.v2"
)

var (
	ErrNotFound = errors.New("item not found")
	ErrInvalid  = errors.New("item invalid")
)

type authorService struct {
	authorRepo AuthorRepository
}

type articleService struct {
	articleRepo ArticleRepository
}

func NewAuthorService(authorRepo AuthorRepository) AuthorService {
	return &authorService{
		authorRepo,
	}
}

func NewArticleService(articleRepo ArticleRepository) ArticleService {
	return &articleService{
		articleRepo,
	}
}

// Author service methods
func (a *authorService) Create(author *Author) (*Author, error) {
	if err := validate.Validate(author); err != nil {
		return nil, errs.Wrap(ErrInvalid, "service.Author.Create")
	}
	author.Id = uuid.New().String()
	return a.authorRepo.Create(author)
}

func (a *authorService) Read(id string) (*Author, error) {
	return a.authorRepo.Read(id)
}

func (a *authorService) Update(author *Author) (*Author, error) {
	return a.authorRepo.Update(author)
}

func (a *authorService) Delete(id string) error {
	return a.authorRepo.Delete(id)
}

// Article service methods
func (a *articleService) Create(article *Article) (*Article, error) {
	if err := validate.Validate(article); err != nil {
		return nil, errs.Wrap(ErrInvalid, "service.Article.Create")
	}
	article.Id = uuid.New().String()
	article.CreateAt = time.Now().UTC().Unix()
	return a.articleRepo.Create(article)
}

func (a *articleService) Read(id string) (*Article, error) {
	return a.articleRepo.Read(id)
}

func (a *articleService) Update(article *Article) (*Article, error) {
	return a.articleRepo.Update(article)
}

func (a *articleService) Delete(id string) error {
	return a.articleRepo.Delete(id)
}
