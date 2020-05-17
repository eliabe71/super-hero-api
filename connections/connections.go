package connections

import (
	"api/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"api/db"
)

const (
	baseSearchApi string = "https://superheroapi.com/api/884949558596858/search/"
	baseUrl       string = "https://superheroapi.com/api/884949558596858/"
)
//Retorna true se pode ser inserido ou false caso já tenha sido  inserido 
//func searchId() bool{

//}
func POST() {

	var nameSuper string
	fmt.Println("please enter the name of the villain/super")
	// Digitar o nome do super/vilan para poder buscalo na API
	fmt.Scanln(&nameSuper)
	// Obtendo o pacote Json e salvando na váriavel response
	response, err := http.Get(baseSearchApi + nameSuper)

	if err != nil {
		panic(err.Error())
		return
	}

	// Lendo o Json e Retornando Conjunto do bits para a variável bodyJson

	bodyJson, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
		return
	}
	/// res abreviação para resposta
	var super types.SuperAndVillains
	//  atribuindo os valor de bodyJson a estrutura do tipo data
	json.Unmarshal(bodyJson, &super)
	/// Garante que o numero de ids e supers adquiridos da API online não sejam nulos
	if super.Response == "error" {
		// Fechando o pacote json
		response.Body.Close()

		fmt.Println("Super/Villain not allowed")

		return
	}
	fmt.Println("Villain/Super Available for registration")
	fmt.Print("Loading .")
	ids := make([]string,0)
	for i:=0; i<len(super.Results);i++{
		ids = append(ids,super.Results[i].Id)
	}
	fmt.Print(".")
	db.Saving(ids, &super)
	
	response.Body.Close()
}
func GET(who string) []types.Get {
	var nameOrId string
	if who == "search"{
		fmt.Println("type 1 if you are looking for the ID or 2 if you are looking for the name")
		fmt.Scanln(&nameOrId)
		if nameOrId == "1"{
			fmt.Println("type the id")
			fmt.Scanln(&nameOrId)
			
			return db.SearchSuperId(nameOrId)
		}
		if nameOrId == "2"{
			fmt.Println("type the Name")
			fmt.Scanln(&nameOrId)
			return	db.SearchSuperName(nameOrId)
		} 
	}
	supers := db.SearchSuperTable(who)
	for i:=0 ; i<len(supers);i++{
		fmt.Print(supers[i].Name+" Id: ")
		fmt.Println(supers[i].Id)
	}
	return supers
}
func DELETE(who string){
	if who == "all"{
		fmt.Print("Loading .")
		 db.RemoveDataBase()
		 fmt.Println("all tables destroyed")
	}
	if who=="name"{
		var nameSuper string
		fmt.Println("please enter the name of the villain/super")
		// Digitar o nome do super/vilan para poder buscalo na API
		fmt.Scanln(&nameSuper)
		fmt.Print("Loading .")
		if db.RemoveSuper(nameSuper){
			fmt.Println("complete removal")
			return
		}
		fmt.Println("there was an error, maybe the super does not exist")
		return
	}
	fmt.Println("Error")
}