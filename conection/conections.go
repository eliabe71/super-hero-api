package main
import(
	//_"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	//avl "github.com/emirpasic/gods/trees/avltree"
)
const(
	baseUrlApi string = "https://superheroapi.com/api/884949558596858/" 
)
func GetJson() (*work){
	//va := fmt.Sprintf("%v")
	res, err := http.Get(baseUrlApi+""+"1"+""+"/work")
	defer res.Body.Close()
	if err != nil {
			panic(err)
			//return nil
	}else{
			byte,err := ioutil.ReadAll(res.Body)
			if err !=nil{
					panic(err)
					//return nil
			
			}else{
					vareli *work= new(work)

					json.Unmarshal(byte,(*eli)) 
					return eli
			}	
	}
	
} 
func main (){
	R := GetJson()
	
	fmt.Println(R.Base)
	
	Avl := avl.NewWithIntComparator()
	
	Avl.Put(1,R)
	
	E,_:= Avl.Get(1)
	fmt.Println((*E).Base)
}