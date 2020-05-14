package conections
import (
	"fmt"
 	"net/http"
	"io/ioutil"
	"encoding/json"
	"api/types"	
	"strings"
	//"github.com/emirpasic/gods/trees/avltree"
	"api/db"

)

const (
	baseSearchApi string = "https://superheroapi.com/api/884949558596858/search/"
	baseUrl = "https://superheroapi.com/api/884949558596858/"

)
func searchId (name string, nameAux types.NameAux) (string) {
	i:= 68
	id := fmt.Sprintf("%v",i);
	for{
		response,err := http.Get(baseUrl+id)
		if 	err!=nil{
			panic(err.Error())
		}
		// lendo o pacote json coloando  os bits na variável bodyjson	
		bodyJson,err := ioutil.ReadAll(response.Body)
	
		if err!= nil{
			panic(err.Error())
		}

		// atribuindo os valores do json no endereço passado 
		json.Unmarshal(bodyJson,&nameAux)	
		if strings.ToLower(nameAux.NameAux) == "error"{
			fmt.Println("super already registered")
			return "error"
		}
		//VERIFICA SE OS NOMES SÃO IGUAIS 
		if strings.ToLower(nameAux.NameAux) == strings.ToLower(name) {
			//fmt.Println(strings.ToLower(nameAux.NameAux)+" = "+ strings.ToLower(name))
			if db.VerifyId(id){
				fmt.Println("voltando")
				i++
				id = fmt.Sprintf("%v",i);
				continue
			}
			response.Body.Close()
			if i%2 ==0 {fmt.Println(".")}
			fmt.Println("super Available for registration")
			return id
		}
		response.Body.Close()
		fmt.Println(id)
		i++
		id = fmt.Sprintf("%v",i)
		//if i%2 ==0 {fmt.Println(".")}

	}
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
	//fmt.Println("Villain/Super Available for registration")
	var nameAux types.NameAux
	fmt.Println("Loading .")
	fmt.Println(searchId(nameSuper,nameAux))
	response.Body.Close()
}