package http

import (
	"log"
	"net/http"

	"github.com/AntonyIS/modart/app"
	"github.com/AntonyIS/modart/repository"
	"github.com/gin-gonic/gin"
)

type GinRoutehandler interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Post(*gin.Context)
	Put(*gin.Context)
	Delete(*gin.Context)
}

type ginAuthorHandler struct {
	authorService app.AuthorService
}

func NewGinAuthorHandler(authorService app.AuthorService) GinRoutehandler {
	return &ginAuthorHandler{
		authorService,
	}
}

func (a ginAuthorHandler) GetAll(c *gin.Context) {
	authors, err := a.authorService.ReadAuthorAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"authors": authors,
	})
}

func (a ginAuthorHandler) Get(c *gin.Context) {
	id := c.Param("id")
	author, err := a.authorService.ReadAuthor(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	if author == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "author not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"author": author,
	})
	return
}

func (a ginAuthorHandler) Post(c *gin.Context) {
	var author app.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res, err := a.authorService.CreateAuthor(&author)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"author": res,
	})
	return
}

func (a ginAuthorHandler) Put(c *gin.Context) {
	var author app.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res, err := a.authorService.UpdateAuthor(&author)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"author": res,
	})
	return
}

func (a ginAuthorHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	res := a.authorService.DeleteAuthor(id)
	c.JSON(http.StatusCreated, gin.H{
		"author": res,
	})

}
func SetupGinRouter() *gin.Engine {

	router := gin.Default()
	repo, err := repository.NewAuthorRepository()
	if err != nil {
		log.Fatal(err)
	}
	authorSrv := app.NewAuthorService(repo)
	handler := NewGinAuthorHandler(authorSrv)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Mode",
		})
	})

	router.GET("/api/v1/authors", handler.GetAll)
	router.GET("/api/v1/authors/:id", handler.Get)
	router.POST("/api/v1/authors", handler.Post)
	router.PUT("/api/v1/authors/:id", handler.Put)
	router.DELETE("/api/v1/authors/:id", handler.Delete)

	return router
}
