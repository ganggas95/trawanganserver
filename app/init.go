package app

import (
	"fmt"
	"github.com/ganggas95/trawanganserver/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

var GORM *gorm.DB

const (
	DB_HOST = "ec2-107-22-251-225.compute-1.amazonaws.com"
	DB_USER = "vcpyraxvhxkcxc"
	DB_PASS = "gHkAPj8weKd4a8_KGH_y1-bOYh"
	DB_NAME = "daalefbcj1ekfm"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

func InitDB() {
	var err error
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", DB_HOST, DB_USER, DB_PASS, DB_NAME)
	GORM, err = gorm.Open(revel.Config.StringDefault("db.driver", "postgres"), revel.Config.StringDefault("db.spec", dbinfo))
	if err != nil {
		revel.INFO.Println("Db Error ", err)
	}
	createTable(GORM)
	GORM.LogMode(true)
	revel.INFO.Println("Db COnnected")
}
func createTable(db *gorm.DB) {
	db.SingularTable(true)
	db.AutoMigrate(&models.AgentTravel{},
		&models.User{},
		&models.UserToken{},
		&models.Person{},
		&models.AgentService{},
		&models.UserFoto{},
		&models.FotoService{},
		&models.AddOnService{},
	)

}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
