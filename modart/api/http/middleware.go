package http

// func Authorize(c *gin.Context) {
// 	db, err := repository.NewPostgresqlDB()
// 	if err != nil {
// 		log.Fatal("error connection to the database")
// 	}
// 	tokenString, err := c.Cookie("Authorization")
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["sub"])
// 		}
// 		return []byte(os.Getenv("SECRET")), nil
// 	})
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		var author *app.Author
// 		db.Where("email= ?", claims["sub"]).First(author)
// 		repository.DB.Where("email =?", claims["sub"]).First(&author)

// 		if author.Email == "" {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		c.Set("user", author)
// 		c.Next()
// 	} else {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// }
