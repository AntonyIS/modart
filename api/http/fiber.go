package http

import (
	"fmt"
	"log"

	"github.com/AntonyIS/modart/app"
	"github.com/AntonyIS/modart/repository"
	"github.com/gofiber/fiber"
)

type FiberRoutehandler interface {
	Get(*fiber.Ctx)
	GetAll(*fiber.Ctx)
	Post(*fiber.Ctx)
	Put(*fiber.Ctx)
	Delete(*fiber.Ctx)
}

type authorHandler struct {
	authorService app.AuthorService
}

type articleHandler struct {
	articleService app.ArticleService
}

func NewAuthorHandler(authorService app.AuthorService) FiberRoutehandler {
	return authorHandler{
		authorService,
	}
}
func NewArticleHandler(articleService app.ArticleService) FiberRoutehandler {
	return articleHandler{
		articleService,
	}
}

func (a authorHandler) Post(c *fiber.Ctx) {
	var author app.Author
	if err := c.BodyParser(author); err != nil {
		c.Status(503).Send(err)
	}
	a.authorService.CreateAuthor(&author)

	c.JSON(author)
}

func (a authorHandler) Get(c *fiber.Ctx) {
	id := c.Params("id")
	author, err := a.authorService.ReadAuthor(id)
	if err != nil {
		c.Status(404).Send(err)
	}

	c.JSON(author)
}

func (a authorHandler) GetAll(c *fiber.Ctx) {
	fmt.Println("Hello world")
	authors, err := a.authorService.ReadAuthorAll()

	if err != nil {
		c.Status(404).Send(err)
	}

	c.JSON(authors)
}

func (a authorHandler) Put(c *fiber.Ctx) {
	var author *app.Author
	if err := c.BodyParser(author); err != nil {
		c.Status(503).Send(err)
	}
	author, err := a.authorService.UpdateAuthor(author)
	if err != nil {
		c.Status(404).Send(err)
	}
}

func (a authorHandler) Delete(c *fiber.Ctx) {
	id := c.Params("id")
	err := a.authorService.DeleteAuthor(id)
	if err != nil {
		c.Status(404).Send(err)
	}
	c.Send("Author deleted successfully")

}

func (a articleHandler) Post(c *fiber.Ctx) {
	var article app.Article
	if err := c.BodyParser(article); err != nil {
		c.Status(503).Send(err)
	}
	a.articleService.CreateArticle(&article)

	c.JSON(article)
}

func (a articleHandler) Get(c *fiber.Ctx) {
	c.Send("Get article")
}

func (a articleHandler) GetAll(c *fiber.Ctx) {
	c.Send("Get articles")
}

func (a articleHandler) Put(c *fiber.Ctx) {
	c.Send("Update article")
}

func (a articleHandler) Delete(c *fiber.Ctx) {
	c.Send("Delete article")
}

func home(c *fiber.Ctx) {
	c.Send("Welcome to ModArt API ...")
}
func SetupRoutes(fiberApp *fiber.App) {
	authorRepo, err := repository.NewAuthorRepository()
	if err != nil {
		log.Fatal(err)
	}
	articleRepo, err := repository.NewArticleRepository()
	if err != nil {
		log.Fatal(err)
	}

	// Get services
	authorSrv := app.NewAuthorService(authorRepo)
	articleSrv := app.NewArticleService(articleRepo)
	// Get handlers
	appAuthorHandler := NewAuthorHandler(authorSrv)
	appArticleHandler := NewArticleHandler(articleSrv)

	fiberApp.Get("/", home)
	fiberApp.Get("/api/v1/author/:id", appAuthorHandler.Get)
	fiberApp.Get("/api/v1/author/all", appAuthorHandler.GetAll)
	fiberApp.Post("/api/v1/author", appAuthorHandler.Post)
	fiberApp.Delete("/api/v1/author/:id", appAuthorHandler.Delete)
	fiberApp.Get("/api/v1/article/:id", appArticleHandler.Get)
	fiberApp.Get("/api/v1/article/all", appArticleHandler.GetAll)
	fiberApp.Post("/api/v1/article", appArticleHandler.Post)
	fiberApp.Delete("/api/v1/author:id", appArticleHandler.Delete)

}
