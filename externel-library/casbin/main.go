package main

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	//istioGrpcCheck()
	httpCheck()
	//dbCallback()
}

func httpCheck() {
	//enforcer, err := getEnforcerByFile()
	enforcer, err := getEnforcerByDB("host=192.168.1.205 port=5432 user=hsck password=hsck@2301 database=test_tenant sslmode=disable")

	if err != nil {
		fmt.Println(err)
		return
	}

	policies := enforcer.GetPolicy()
	for _, policy := range policies {
		fmt.Println("--> ", policy)
	}

	r := gin.Default()
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

func getEnforcerByFile() (*casbin.Enforcer, error) {
	return casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
}

func getEnforcerByDB(conn string) (*casbin.Enforcer, error) {
	// Increase the column size to 512.
	type CasbinRule struct {
		ID    uint   `gorm:"primaryKey;autoIncrement"`
		Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
		V0    string `gorm:"size:512;uniqueIndex:unique_index"`
		V1    string `gorm:"size:512;uniqueIndex:unique_index"`
		V2    string `gorm:"size:512;uniqueIndex:unique_index"`
		V3    string `gorm:"size:512;uniqueIndex:unique_index"`
		V4    string `gorm:"size:512;uniqueIndex:unique_index"`
		V5    string `gorm:"size:512;uniqueIndex:unique_index"`
	}

	// @ = %40
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	a, err := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{}, "saas.casbin_rule")
	if err != nil {
		return nil, err
	}

	return casbin.NewEnforcer("rbac_model.conf", a)

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
