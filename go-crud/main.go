package main

import (
	"golang-crud-init/config"
	"golang-crud-init/controller"
	"golang-crud-init/helper"
	"golang-crud-init/model"
	"golang-crud-init/repository"
	"golang-crud-init/router"
	"golang-crud-init/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")
	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagsRepository := repository.NewTagsRepositoryImpl(db)

	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888", // Tambahkan ':' sebelum port
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
