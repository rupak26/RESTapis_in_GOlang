package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rupak26/RESTapis_in_GOlang/internal/config"
	"github.com/rupak26/RESTapis_in_GOlang/internal/http/handlers/students"
	// "github.com/rupak26/RESTapis_in_GOlang/internal/storage"
	"github.com/rupak26/RESTapis_in_GOlang/internal/storage/sqlite"
)



func main() {
	// load config 
	cfg := config.MustLoad() 
	// database setup
    storage,err := sqlite.New(*cfg)
	fmt.Print(storage)
    if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initialize" , slog.String("env" , cfg.Env) , slog.String("version" , "1.0.0"))
	// setup router 
    router := http.NewServeMux() 
	
	router.HandleFunc("/api/students" , students.New(storage)) 
	router.HandleFunc("/api/students/{id}" , students.GetByID(storage))

	// setup server 
	server := http.Server{
		Addr : cfg.Addr ,
		Handler: router,
	}
	slog.Info("Server is running on port ", slog.String("adress",cfg.Addr))
   
	done := make(chan os.Signal , 1) 

	signal.Notify(done , os.Interrupt , syscall.SIGINT , syscall.SIGTERM)

    go func () {
        err := server.ListenAndServe()
    
		if err != nil {
			log.Fatal("Failed to start server" , err)
		}
	}()

	<- done

	slog.Info("Shuting down the server")

	ctx , cancle := context.WithTimeout(context.Background() , 5 * time.Second) 
	defer cancle() 
	err = server.Shutdown(ctx) 

	if err != nil {
		slog.Error("Failed to shut down server" , slog.String("error" , err.Error()))
	}

	slog.Info("Server shutdown successfully")
}


