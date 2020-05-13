package controllers
import (
	"fmt"
 	"net/http"
	"io/ioutil"
	"encoding/json"
	"api/types"	
	"strings"
)

const (
	baseSearchApi string = "https://superheroapi.com/api/884949558596858/search/"
	baseUrl = "https://superheroapi.com/api/884949558596858/"

)
func searchId (name string, nameAux types.NameAux) (id string) {
	i:= 1
	id = fmt.Sprintf("%v",i);
	for nameAux.NameAux != name{
		response,err := http.Get(baseUrl+id)
		
		if 	err!=nil{
			panic(err.Error())
			return
		}	
		bodyJson,err := ioutil.ReadAll(response.Body)
	
		if err!= nil{
			panic(err.Error())
			return
		}
		json.Unmarshal(bodyJson,&nameAux)	
		if strings.ToLower(nameAux.NameAux) == strings.ToLower(name){
			response.Body.Close()
			return id
		}
		response.Body.Close()
		i++
		id = fmt.Sprintf("%v",i);
		fmt.Println(nameAux.NameAux)
		fmt.Println(id)
	}
	return id
}
func POST(){
	
	var nameSuper string 
	fmt.Println("please enter the name of the villain/super")
	// Digitar o nome do super/vilan para poder buscalo na API
	fmt.Scanln(&nameSuper)
	// Obtendo o pacote Json e salvando na váriavel response
	response,err := http.Get(baseSearchApi+nameSuper)
	
	if err!=nil{
		panic(err.Error())
		return
	}
	
	// Lendo o Json e Retornando Conjunto do bits para a variável bodyJson

	bodyJson,err := ioutil.ReadAll(response.Body)
	
	if err!= nil{
		panic(err.Error())
		return
	}
	/// res abreviação para resposta
	var res types.Response
	//  atribuindo os valor de bodyJson a estrutura do tipo data
	json.Unmarshal(bodyJson,&res)
	

	if res.Response == "error"{ 
		// Fechando o pacote json
		response.Body.Close()

		fmt.Println("Super/Villain not allowed")
		
		return
	}
	fmt.Println("Villain/Super Available for registration")
	var nameAux types.NameAux
	fmt.Println(searchId(nameSuper,nameAux ))
	response.Body.Close()
}