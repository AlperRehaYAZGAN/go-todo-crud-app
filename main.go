package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"gotodoapp/common"
	"gotodoapp/models"
)

var db *gorm.DB
var cache *redis.Client

func main() {
	// init db connection
	db = common.NewSqliteDB("myapp.db")
	common.SetupDBTables(db)

	// init cache connection
	cache = common.NewCacheConnection("localhost:6379", "1020304050")

	// http router start.
	r := gin.Default()

	r.GET("/get/:key", func(c *gin.Context) {
		key := c.Param("key")

		val, err := cache.Get(c, key).Result()
		if err != nil {
			c.String(http.StatusNotFound, "not found")
		} else {
			c.String(http.StatusOK, val)
		}

	})

	r.GET("/set/:key/:value", func(c *gin.Context) {
		key := c.Param("key")
		value := c.Param("value")

		err := cache.Set(c, key, value, time.Minute).Err()
		if err != nil {
			c.String(http.StatusInternalServerError, "set error")
		} else {
			c.String(http.StatusOK, "set ok")
		}

	})

	r.GET("/", func(c *gin.Context) {
		// read req.header.abc
		val := c.Request.Header.Get("abc")

		c.String(http.StatusOK, "hello world"+val)
	})

	r.GET("/todos", func(c *gin.Context) {
		var todos []models.Todo
		db.Find(&todos)

		c.JSON(http.StatusOK, todos)
	})

	/*
		// "GET /todos"  : get all todos
		// "POST /todos" : create a new todo
		// "PUT /todos/:id" : update a todo
		// "DELETE /todos/:id" : delete a todo
	*/

	// port
	r.Run(":5555")
}
