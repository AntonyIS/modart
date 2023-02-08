package app

type Item struct{}

func (i *Item) CreateItem(d *Item) (*Item, error) {
	return i.CreateItem(d)
}

func (i *Item) ReadItem(id string) (*Item, error) {
	return i.ReadItem(id)
}

type BuilderProcess interface {
	CreateItem(i *Item) (*Item, error)
	ReadItem(id string) (*Item, error)
	ReadItems() (*[]Item, error)
	UpdateItem(i *Item) (*Item, error)
	DeleteItem(id string) error
}
