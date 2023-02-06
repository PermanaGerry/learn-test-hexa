package main

// C:\Users\Majoo Indonesia\workspace\go\src\learn-test-hexa\internal\handlers\notifhdl\notif.go
import (
	"fmt"
	"learn-test-hexa/internal/core/services/notifsrv"
	"learn-test-hexa/internal/handlers/notifhdl"
	"learn-test-hexa/internal/repositories/notifrepo"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	notifRepository := notifrepo.NewMemNotif()

	fmt.Println(notifRepository)
	os.Exit(0)

	notifService := notifsrv.New(notifRepository)
	notifHandler := notifhdl.NewHTTPHadler(notifService)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/send", notifHandler.Send)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	router.Run(":8000") // listen and serve on 0.0.0.0:8000

	// uuid := uuid.New().String()

	// r := gin.New()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": uuid,
	// 	})
	// })
	// r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
