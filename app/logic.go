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
func (a *authorService) CreateAuthor(author *Author) (*Author, error) {
	if err := validate.Validate(author); err != nil {
		return nil, errs.Wrap(ErrInvalid, "service.Author.Create")
	}
	author.Id = uuid.New().String()
	return a.authorRepo.CreateAuthor(author)
}

func (a *authorService) ReadAuthor(id string) (*Author, error) {
	return a.authorRepo.ReadAuthor(id)
}

func (a *authorService) ReadAuthorAll() ([]*Author, error) {
	return a.authorRepo.ReadAuthorAll()
}

func (a *authorService) UpdateAuthor(author *Author) (*Author, error) {
	return a.authorRepo.UpdateAuthor(author)
}

func (a *authorService) DeleteAuthor(id string) error {
	return a.authorRepo.DeleteAuthor(id)
}

// Article service methods
func (a *articleService) CreateArticle(article *Article) (*Article, error) {
	if err := validate.Validate(article); err != nil {
		return nil, errs.Wrap(ErrInvalid, "service.Article.Create")
	}
	article.Id = uuid.New().String()
	article.CreateAt = time.Now().UTC().Unix()
	return a.articleRepo.CreateArticle(article)
}

func (a *articleService) ReadArticle(id string) (*Article, error) {
	return a.articleRepo.ReadArticle(id)
}

func (a *articleService) ReadArticleAll() ([]*Article, error) {
	return a.articleRepo.ReadArticleAll()
}
func (a *articleService) UpdateArticle(article *Article) (*Article, error) {
	return a.articleRepo.UpdateArticle(article)
}

func (a *articleService) DeleteArticle(id string) error {
	return a.articleRepo.DeleteArticle(id)
}
