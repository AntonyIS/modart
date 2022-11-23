package app

type AuthorRepository interface {
	CreateAuthor(author *Author) (*Author, error)
	ReadAuthor(id string) (*Author, error)
	ReadAuthorAll() ([]*Author, error)
	UpdateAuthor(author *Author) (*Author, error)
	DeleteAuthor(id string) error
}
type ArticleRepository interface {
	CreateArticle(Article *Article) (*Article, error)
	ReadArticle(id string) (*Article, error)
	ReadArticleAll() ([]*Article, error)
	UpdateArticle(Article *Article) (*Article, error)
	DeleteArticle(id string) error
}
