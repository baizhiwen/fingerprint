package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"time"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {

	//使用默认中间件(logger 和 recovery 中间件)创建 gin 路由
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	router.GET("/users", func(c *gin.Context) {})
	router.GET("/user/:id", func(c *gin.Context) {})
	router.GET("/user/:id/mmm/*action", func(c *gin.Context) {})

	router.GET("/bookable", getBookable)
	router.GET("/testing", startPage)

	router.GET("/someGet/:name/*act", getting)
	router.POST("/somePost", posting)

	router.GET("/json", func(c *gin.Context) {
		var AA struct {
			Html string `json:"html"`
		}
		a := AA
		a.Html = "<b>Hello, world!</b>"
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// Serves literal characters
	router.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)
	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。.
	router.Run()
	//router.Run(":3000") hardcode 端口号

}

func startPage(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}

	c.String(200, "Success")
}

func getting(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(c.Request.URL)
	user := ""
	fmt.Println(c.DefaultQuery("user", user))

	c.String(http.StatusOK, "Hello %s", name)
}
func posting(c *gin.Context) {
	user := ""
	fmt.Println(c.DefaultPostForm("user", user))
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	fmt.Printf("ids: %v; names: %v", ids, names)
	c.String(http.StatusOK, "Hello %s", user)
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
