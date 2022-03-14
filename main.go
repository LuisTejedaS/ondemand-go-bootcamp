package main

import (
	"log"
	"net/http"

	"github.com/LuisTejedaS/ondemand-go-bootcamp/infraestructure/data"
	controller "github.com/LuisTejedaS/ondemand-go-bootcamp/interface/controllers"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/interface/repository"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/router"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/interactor"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/presenter"
	"github.com/LuisTejedaS/ondemand-go-bootcamp/usecase/service"
)

const port = ":8090"

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	csvDataSourcePokemon := data.NewCSVDataSource("/Users/luis.tejeda/Documents/Source/GitHub/ondemand-go-bootcamp/poke.csv")

	pokemonCSVRepository, err := repository.NewpokemonCSVRepository(csvDataSourcePokemon)
	if err != nil {
		log.Fatalln(err)
	}

	pokemonService := service.NewPokemonService(pokemonCSVRepository)
	pokemonPresenter := presenter.NewPokemonPresenter()
	pokemonInteractor := interactor.NewPokemonInteractor(pokemonService, pokemonPresenter)

	controller := controller.NewPokemonController(pokemonInteractor)

	appRouter := router.NewRoute(controller)
	log.Printf("Listening on port %s", port)

	log.Fatal(http.ListenAndServe(port, appRouter.Router()))
}
