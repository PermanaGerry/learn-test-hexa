package main

// C:\Users\Majoo Indonesia\workspace\go\src\learn-test-hexa\internal\handlers\notifhdl\notif.go
import (
	"context"
	"learn-test-hexa/infrastructure/repositories/redis"
	"learn-test-hexa/internal/core/services/notifsrv"
	"learn-test-hexa/internal/handlers/notifhdl"
	"learn-test-hexa/internal/repositories/notifrepo"
	"learn-test-hexa/utils/config"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	config.LoadConfig()
}

func main() {

	notifRepository := notifrepo.NewMemNotif()
	notifService := notifsrv.New(notifRepository)
	notifHandler := notifhdl.NewHTTPHadler(notifService)

	notifRepositoryRedis := notifrepords.NewRepositoryNotif(config.DBRedis)
	notifServiceRedis := notifsrv.New(notifRepositoryRedis)
	notifHandlerRedis := notifhdl.NewHTTPHadler(notifServiceRedis)

	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.GET("/send", notifHandler.Send)
	router.GET("/show", notifHandler.List)
	groupRedis := router.Group("/redis")
	{
		groupRedis.GET("/send", notifHandlerRedis.Send)
		groupRedis.GET("/list", notifHandlerRedis.List)
	}
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	router.Run(":8000") // listen and serve on 0.0.0.0:8000
}
