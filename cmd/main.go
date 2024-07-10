package main

import (
	"fmt"
	"gym/pkg/api"
	"gym/pkg/api/handlers"
	"gym/pkg/chain"
	"gym/pkg/contract/store"
	"gym/pkg/db"
	"gym/pkg/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		log.Fatalf("error loading env")
	}
	db, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	chain.Init()
	store.Init()
	usecase := usecase.NewUseCase(db)
	handler := handlers.NewHandler(usecase)
	router := gin.Default()
	api.NewServer(router, handler)

	err1 := http.ListenAndServe(":8080", router)
	if err1 != nil {
		log.Fatal(err1)
	}
}
