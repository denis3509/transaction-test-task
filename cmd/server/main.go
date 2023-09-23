package main

import (
	"fmt"
	"transaction/internal/config"
	"transaction/internal/transaction"
	"transaction/pkg/db"
	"transaction/pkg/log"

	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func gracefullShutdown(log log.Logger) {
	log.Info("gracefully shutdown")
}

func main() {
	cfg := config.GetConfig()

	log := log.NewLogger(os.Stdout)

	db, err := db.MustOpen("postgres", cfg.DB.ConnString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	transaction.CreateRoutes(router, &db)
	go func() {
		connStr := fmt.Sprint("localhost:", strconv.Itoa(cfg.Port))
		router.Run(connStr)

	}()
	log.Infof("server listening on localhost:%d", cfg.Port)

	// signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %s \n", sig)
		gracefullShutdown(log)
		done <- true
	}()
	<-done
	fmt.Println("exiting")

}
