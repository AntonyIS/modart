package http

import (
	"net/http"
	"os"
	"time"

	"github.com/AntonyIS/modart/app"
	"github.com/AntonyIS/modart/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type GinRoutehandler interface {
	LoginUser(*gin.Context)
	GetUser(*gin.Context)
	GetUsers(*gin.Context)
	PostUser(*gin.Context)
	PutUser(*gin.Context)
	DeleteUser(*gin.Context)
	GetArticle(*gin.Context)
	GetArticles(*gin.Context)
	PostArticle(*gin.Context)
	PutArticle(*gin.Context)
	DeleteArticle(*gin.Context)
}

type ginHandler struct {
	appService app.AppService
}

func NewHandler(appSrv app.AppService) GinRoutehandler {
	return &ginHandler{
		appSrv,
	}
}

// Author handler
func (a ginHandler) GetUsers(c *gin.Context) {
	authors, err := a.appService.ReadAuthors()
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

func (a ginHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	author, err := a.appService.ReadAuthor(id)
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

func (a ginHandler) LoginUser(c *gin.Context) {
	email := c.Param("email")
	author, err := a.appService.LoginAuthor(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": author.Email,
		"exp": time.Now().Add(time.Minute * 60).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}

func (a ginHandler) PostUser(c *gin.Context) {
	var author app.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := a.appService.CreateAuthor(&author)
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

func (a ginHandler) PutUser(c *gin.Context) {
	var author app.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := a.appService.UpdateAuthor(&author)
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

func (a ginHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	res := a.appService.DeleteAuthor(id)
	c.JSON(http.StatusCreated, gin.H{
		"author": res,
	})
}

// Article handler
func (a ginHandler) GetArticles(c *gin.Context) {
	articles, err := a.appService.ReadArticles()
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

func (a ginHandler) GetArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := a.appService.ReadArticle(id)
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

func (a ginHandler) PostArticle(c *gin.Context) {
	var article app.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res, err := a.appService.CreateArticle(&article)
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

func (a ginHandler) PutArticle(c *gin.Context) {
	var article app.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := a.appService.UpdateArticle(&article)
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

func (a ginHandler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	err := a.appService.DeleteArticle(id)
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
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.Use(cors.Default())

	dbClient := repository.NewPostgresqlDB()
	srv := app.NewItemService(dbClient)

	handler := NewHandler(srv)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Mode",
		})
	})

	r.GET("/api/v1/users", handler.GetUsers)
	r.GET("/api/v1/users/:id", handler.GetUser)
	r.POST("/api/v1/users", handler.PostUser)
	r.PUT("/api/v1/users/:id", handler.PutUser)
	r.DELETE("/api/v1/users/:id", handler.DeleteUser)
	r.GET("/api/v1/articles", handler.GetArticles)
	r.GET("/api/v1/articles/:id", handler.GetArticle)
	r.POST("/api/v1/articles", handler.PostArticle)
	r.PUT("/api/v1/articles/:id", handler.PutArticle)
	r.DELETE("/api/v1/articles/:id", handler.DeleteArticle)

	return r
}
