package app

type AppRepository interface {
	CreateAuthor(author *Author) (*Author, error)
	ReadAuthor(id string) (*Author, error)
	ReadAuthors() ([]*Author, error)
	UpdateAuthor(author *Author) (*Author, error)
	DeleteAuthor(id string) error
	CreateArticle(Article *Article) (*Article, error)
	ReadArticle(id string) (*Article, error)
	ReadArticles() ([]*Article, error)
	UpdateArticle(Article *Article) (*Article, error)
	DeleteArticle(id string) error
}
