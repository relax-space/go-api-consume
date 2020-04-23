package main

import (
	"fmt"
	"github.com/relax-space/go-api-consumer/adapters"
	"github.com/relax-space/go-api-consumer/config"
	"github.com/relax-space/go-api-consumer/models"
	"log"
	"net/http"
	"github.com/hublabs/common/eventconsume"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
)

func main() {

	c := config.Init(os.Getenv("APP_ENV"))
	fmt.Println(c)
	db, err := models.InitDB(c.Database.Driver, c.Database.Connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := models.InitTable(db); err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	notifyKill()
	if err := adapters.Consume(c.ServiceName, c.EventBroker.Kafka,
		adapters.EventFruit,
		eventconsume.Recover(),
		eventconsume.ContextDB(c.ServiceName, db, c.Database.Logger.Kafka),
	); err != nil {
		panic(err)
	}

	if err := e.Start(":8091"); err != nil {
		log.Println(err)
	}
}

func notifyKill() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)
	go func() {
		for s := range signals {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				os.Exit(0)
			}
		}
	}()
}
