package main

import (
	"fmt"
	"transaction/pkg/log"
	"transaction/internal/config"
 
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

func gracefullShutdown(log log.Logger){
 	log.Info("gracefully shutdown")
}

func main(){
	cfg := config.GetConfig()

	log := log.NewLogger(os.Stdout)

	db, err := dbx.MustOpen("postgres", cfg.DB.DSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
 
 
	go func () {
		err  = http.ListenAndServe(":" + strconv.Itoa(cfg.Port) , nil)
		if err != nil {
			log.Fatal("Can't start server: ", err)
		} 
	}()
	log.Infof("server listening on localhost:%d",cfg.Port)
	
	// signal 
    sigs := make(chan os.Signal, 1) 
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) 
    done := make(chan bool, 1) 
    go func() {
        sig := <-sigs 
        fmt.Printf("Received signal: %s \n",sig)
		gracefullShutdown(log)
        done <- true
    }() 
    <-done
    fmt.Println("exiting")
 
}