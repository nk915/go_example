package main

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	enforcer, err := casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	r.Use(func(c *gin.Context) {
		sub := "alice"            // the user that wants to access a resource.
		obj := c.Request.URL.Path // the resource that is going to be accessed.
		act := c.Request.Method   // the operation that the user performs on the resource.

		ok, err := enforcer.Enforce(sub, obj, act)

		if err != nil || !ok {
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "You are not authorized"})
			c.Abort()
			return
		}

		c.Next()
	})

	r.GET("/admin/data1", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"data": "Admin Data"})
	})

	r.POST("/user/data2", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"data": "User Data"})
	})

	r.Run(":8080")

}

func casbin_authz() {

	// load the casbin model and policy from files, database is also supported.
	e, err := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	if err != nil {
		fmt.Printf("%v \n", err)
		return
	}

	// define your router, and use the Casbin authz middleware.
	// the access that is denied by authz will return HTTP 403 error.
	r := gin.New()
	r.Use(authz.NewAuthorizer(e))

	r.GET("/dataset2/resource2", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"data": "GET Data"})
	})

	r.Run(":8080")
}
