package main

import (
	"github.com/joho/godotenv"
	"github.com/yoratyo/material-handler-webapp/app"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	engine, err := app.BuildRuntime()
	if err != nil {
		panic(err)
	}
	err = engine.Start()
	if err != nil {
		panic(err)
	}
}
