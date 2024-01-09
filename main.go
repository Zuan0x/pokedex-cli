package main                                                                              
 
import (
	"net/http"
	"log"
	"io/ioutil"
)



type locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func mapf() {
	// displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on. The idea is that the map command will let us explore the world of Pokemon.using the PokeAPI
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
	   log.Fatalln(err)
	}
	//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
	//Convert the body to type string
   sb := string(body)
   log.Printf(sb)
}

func main() {
	startREPL()
}
