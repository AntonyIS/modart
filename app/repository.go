package app

type AuthorRepository interface {
	Create(author *Author) error
	Read(id string) (*Author, error)
	Update(author *Author) (*Author, error)
	Delete(id string) (string, error)
}
type ArticleRepository interface {
	Create(Article *Article) error
	Read(id string) (*Article, error)
	Update(Article *Article) (*Article, error)
	Delete(id string) (string, error)
}
