package main
import (
	//"api/controllers"
	"api/connections"
	_"api/types"
	_"fmt"
    "os"
    "os/exec"
    "runtime"
    _"time"
) 
const(
	allsupers   =  "allsupers"
	heroes  	=  "heroes"
	villains 	=  "villains" 

)
var clear map[string]func()
func init(){
 	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() { 
   	 	cmd := exec.Command("clear") //Linux example, its tested
    	cmd.Stdout = os.Stdout
    	cmd.Run()
    }
}
func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}
//Função CallClear() e init () retirados do site https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go 
// usada para deixar o terminal na execução mais amigável
func main (){
	// Cadrastar um Super
	//connections.POST()
	// todos os super 
	//_ = conections.GET(allsupers)
	// apenas os super herois 	
	//_= conections.GET(heroes)
	// apenas os super viloes
	//_= controllers.GET(villains)
	// Buscar por nome ou id
	//_= conections.GET("search")
	connections.DELETE("name")
	//connections.DELETE("all")
}