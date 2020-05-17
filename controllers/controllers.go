package controllers
import(
	"api/conections"
	"api/types"
)
func POST(){
	conections.POST()
}
func GET(who string) []types.SuperSearch {
	return conections.GET(who)
}