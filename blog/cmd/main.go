package ma

import (
   "github.com/perfectogo/gtm/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func main() {
 //  projectEnv := config.ProjectEnv()
      
   router := echo.New()
   {
      options := api.Options{}
      api.RegisterRoutes(router, options)
   }

   router.Start(":8080")
}