package app

type AuthorRepository interface {
	Create(author *Author) (*Author, error)
	Read(id string) (*Author, error)
	Update(author *Author) (*Author, error)
	Delete(id string) error
}
type ArticleRepository interface {
	Create(Article *Article) (*Article, error)
	Read(id string) (*Article, error)
	Update(Article *Article) (*Article, error)
	Delete(id string) error
}

type ItemRepository interface {
	Create(item interface{}) (*interface{}, error)
	Read(id string) (interface{}, error)
	Update(Article interface{}) (interface{}, error)
	Delete(id string) error
}
