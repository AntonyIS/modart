package http

import (
	"log"
	"net/http"

	"example.com/server/app"
	"example.com/server/repository"
	"github.com/gin-contrib/cors"
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

type ginArticleHandler struct {
	articleService app.ArticleService
}

func NewGinAuthorHandler(authorService app.AuthorService) GinRoutehandler {
	return &ginAuthorHandler{
		authorService,
	}
}

func NewGinArticleHandler(articleService app.ArticleService) GinRoutehandler {
	return &ginArticleHandler{
		articleService,
	}
}

// Author handler
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

// Article handler
func (a ginArticleHandler) GetAll(c *gin.Context) {
	articles, err := a.articleService.ReadArticleAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func (a ginArticleHandler) Get(c *gin.Context) {
	id := c.Param("id")
	article, err := a.articleService.ReadArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	if article == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "author not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": article,
	})
	return
}

func (a ginArticleHandler) Post(c *gin.Context) {
	var article app.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res, err := a.articleService.CreateArticle(&article)
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

func (a ginArticleHandler) Put(c *gin.Context) {
	var article app.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := a.articleService.UpdateArticle(&article)
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

func (a ginArticleHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := a.articleService.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"error": "article not deleted",
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "article deleted successfully",
	})

}
func SetupGinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())
	authorRepo, err := repository.NewAuthorRepository()

	if err != nil {
		log.Fatal(err)
	}

	articleRepo, err := repository.NewArticleRepository()
	if err != nil {
		log.Fatal(err)
	}

	authorSrv := app.NewAuthorService(authorRepo)
	articleSrv := app.NewArticleService(articleRepo)

	authorHandler := NewGinAuthorHandler(authorSrv)
	articleHandler := NewGinArticleHandler(articleSrv)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Mode",
		})
	})

	router.GET("/api/v1/authors", authorHandler.GetAll)
	router.GET("/api/v1/authors/:id", authorHandler.Get)
	router.POST("/api/v1/authors", authorHandler.Post)
	router.PUT("/api/v1/authors/:id", authorHandler.Put)
	router.DELETE("/api/v1/authors/:id", authorHandler.Delete)

	router.GET("/api/v1/articles", articleHandler.GetAll)
	router.GET("/api/v1/articles/:id", articleHandler.Get)
	router.POST("/api/v1/articles", articleHandler.Post)
	router.PUT("/api/v1/articles/:id", articleHandler.Put)
	router.DELETE("/api/v1/articles/:id", articleHandler.Delete)

	return router
}
