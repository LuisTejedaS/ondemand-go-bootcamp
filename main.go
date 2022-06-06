package main

import (
	"log"
	"net/http"
	"os"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/infraestructure/config"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/infraestructure/data"
	controller "github.com/LuisTejedaS/ondemand-go-bootcamp/interface/controllers"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/interface/repository"
	routerV1 "github.com/LuisTejedaS/ondemand-go-bootcamp/router"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/interactor"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/presenter"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/service"
)

// const port = ":8090"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	WD, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	configuration, err := config.NewConfiguration(WD)
	if err != nil {
		log.Fatalln(err)
	}

	csvDataSourcePokemon, err := data.NewCSVDataSource(configuration.ConfData.CSV)
	if err != nil {
		log.Fatalln(err)
	}

	csvDataStorePokemon, err := data.NewCSVDataStore(configuration.ConfData.CSV)
	if err != nil {
		log.Fatalln(err)
	}

	apiDataSourcePokemon, err := data.NewApiDataSource(configuration.ConfData.PokeBaseURL)
	if err != nil {
		log.Fatalln(err)
	}

	pokemonCSVRepository, err := repository.NewpokemonCSVRepository(csvDataSourcePokemon)
	if err != nil {
		log.Fatalln(err)
	}

	pokemonApiRepository, err := repository.NewPokemonApiLoader(apiDataSourcePokemon, csvDataStorePokemon)
	if err != nil {
		log.Fatalln(err)
	}

	pokemonService := service.NewPokemonService(pokemonCSVRepository, pokemonApiRepository)
	pokemonPresenter := presenter.NewPokemonPresenter()
	pokemonInteractor := interactor.NewPokemonInteractor(pokemonService, pokemonPresenter)

	controller := controller.NewPokemonController(pokemonInteractor)

	router := routerV1.CreateRouter()
	routerGroup := routerV1.CreateRouterGroup(router)

	routerV1.RegisterPokemonRoutes(routerGroup, controller)

	log.Printf("Listening on port %s", configuration.ConfData.Port)

	log.Fatal(http.ListenAndServe(configuration.ConfData.Port, router))
}
