package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(Logger())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/hello", func(c *gin.Context) {
		// e := c.MustGet("example").(string)
		// c.String(200, e+"\n")
		// c.String(200, "hello world")
		name := c.Query("name")
		name2 := c.DefaultQuery("name", "default name")
		c.String(http.StatusOK, "get "+name)
		c.String(http.StatusOK, "get "+name2)
		// fmt.Println("hello world")
	})

	router.POST("/world", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(http.StatusOK, "post "+name)
	})
	groupStudy(router)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	/*
		srv := &http.Server{
			Addr:    ":8080",
			Handler: router,
		}

		go func() {
			// service connections
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		// kill (no param) default send syscanll.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			log.Println("timeout of 5 seconds.")
		}
		log.Println("Server exiting")
	*/
}

//middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func groupStudy(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("hello", func(c *gin.Context) {
			c.String(http.StatusOK, " v1 hello ")
		})
	}
	v2 := router.Group("v2")
	{
		v2.GET("hello", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 hello")
		})
	}
}
