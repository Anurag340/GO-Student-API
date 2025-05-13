package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Anurag340/student-api/internal/config"
	"github.com/Anurag340/student-api/internal/http/handlers/student"
	"github.com/Anurag340/student-api/storage/mysql"
)

func main(){
	//load config
	cfg := config.MustLoad()


	//database setup
	storage ,errr := mysql.New(cfg)

	if errr!=nil{
		log.Fatal("Failed to connect to database",errr)
	}

	slog.Info("Connected to database" , )



	//setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("PUT /api/students/{id}" , student.UpdateStudent(storage))
	router.HandleFunc("DELETE /api/students/{id}",student.DeleteStudent(storage))

	//setup server
	server := http.Server{
		Addr:cfg.Addr,
		Handler: router,
	}
	slog.Info("Starting server...")
	//fmt.Printf("Server started on : %s" , cfg.HTTP.Addr)
	
	done:= make(chan os.Signal , 1)
	signal.Notify(done,os.Interrupt , syscall.SIGINT,syscall.SIGTERM)

	go func(){
		err := server.ListenAndServe()

		if err!=nil{
			log.Fatal("Failed tp start server",err)
		}
	}()

	<-done
	
	slog.Info("Shutting down server...")

	ctx , cancel := context.WithTimeout(context.Background(),5*time.Second)

	defer cancel()

	err:=server.Shutdown(ctx)

	if err!=nil{
		slog.Error("Failed to shutdown server",slog.String("error",err.Error()))
	
	}

	slog.Info("Server stopped")
	

}