package main
import (
	//"api/controllers"
	"api/conections"
	_"api/types"
) 
const(
	allsupers   =  "allsupers"
	heroes  	=  "heroes"
	villains 	=  "villains" 

)
func main (){
	// Cadrastar um Super
	conections.POST()
	// todos os super 
	//_ = conections.GET(allsupers)
	// apenas os super herois 	
	//_= conections.GET(heroes)
	// apenas os super viloes
	//_= controllers.GET(villains)
	// Buscar por nome ou id
	//_= conections.GET("search")
}