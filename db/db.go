package db
import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
  //"api/types"
  "strconv"

)
const(
	   host      =  "localhost"
  	 port      =   5432
  	 user      =  "postgres"
 	   password  =   "eliabe1106"
 	   dbname    =   "super-hero"
)

// usado para receber json dos id's apenas uso interno
type id struct{
  id int `json:"id"`
}
/*
  TABLES QUE SERÃO CRIADA
*/
//funcao para verifcar se a table existe
// usada internamente 
// retorna true se a tabela existe e false caso contrário
func existTable(db *sql.DB,nameTable string) bool{
    _,error := db.Exec("SELECT FROM "+nameTable)
    if(error != nil){
       //fmt.Println(error.Error())
        return false
    }
    return true
}
// Criar tabela 
func createTable(db *sql.DB, nameTable string){
  //não precisamos  da resposta estamos utilzaremos apenas a variável err para ver se houve algum error
  stringTables  := [7]string{"powerstats","biography","appearance","work","connections","image","super"}
   switch nameTable {
        case stringTables[0]: 
              if existTable(db, stringTables[0]) == false { 
                    _,err := db.Exec("CREATE TABLE "+nameTable+" ("+
                           "idpower integer primary key,"+
                          "intelligence varchar(10000) not null,"+ 
                          "strength varchar(10000) not null,"+    
                          "speed varchar(10000) not null,"+
                          "durability varchar(10000) not null,"+
                          "power varchar(10000) not null,"+  
                          "combat varchar(10000) not null"+")")
                          if err != nil{
                            fmt.Println(stringTables[0])
                                panic(err)
                          }
              }
        case stringTables[1]:
             if existTable(db, stringTables[1]) == false { 
                _,err := db.Exec("CREATE TABLE "+nameTable+" ("+
                      "idbio integer primary key,"+
                      "fullname varchar(10000) not null,"+            
                      "alteregos varchar(10000) not null,"+ 
                      "aliases varchar[10000] not null,"+
                      "placeofbirth varchar(10000) not null,"+ 
                      "firstappearance varchar(10000) not null,"+   
                      "publisher varchar(10000) not null,"+     
                      "alignment varchar(10000) not null"+")")  
                      if err != nil{
                        fmt.Println(stringTables[1])
                                panic(err)
                      }
              }      
        case stringTables[2]:
              if existTable(db, stringTables[2]) == false { 
                 _,err := db.Exec("CREATE TABLE "+nameTable+"("+
                      "idappearance integer primary key,"+
                      "gender varchar(10000) not null,"+   
                      "race varchar(10000) not null,"+  
                      "height varchar[10000] not null,"+  
                      "weight varchar[10000] not null,"+ 
                      "eyecolor varchar(10000) not null,"+  
                      "haircolor varchar(10000) not null"+")")   
                      if err != nil{
                        fmt.Println(stringTables[2])
                                panic(err)
                      }
              }
        case stringTables[3]:
             if existTable(db, stringTables[3]) == false { 
                _,err := db.Exec("CREATE TABLE "+nameTable+" ("+
                       "idwork integer primary key not null,"+
                        "occupation varchar(10000) not null,"+    
                        "baseofoperation varchar(10000) not null"+")")
                      if err != nil{
                                panic(err)
                      }
              }
        case stringTables[4]:
              if existTable(db, stringTables[4]) == false {   
                _,err := db.Exec("CREATE TABLE "+nameTable+" ("+
                        "idcon integer primary key not null,"+
                        "groupaffiliation varchar(10000) not null,"+
                        "relatives varchar(10000) not null"+")")
                      if err != nil{
                                fmt.Println(stringTables[4])
                                panic(err)
                      }
              }  
        case stringTables[5]:
           if existTable(db, stringTables[5]) == false { 
                 _,err := db.Exec("CREATE TABLE "+nameTable+" ("+
                  "idurl integer primary key not null,"+
                  "url varchar(100000) not null )")  
                      if err != nil{
                        fmt.Println(stringTables[5])
                                panic(err)
                      }
                  }
        case stringTables[6]:
            if existTable(db, stringTables[6]) == false {   
                _,err := db.Exec("CREATE TABLE "+nameTable+" ("+ 
                    "name varchar(10000) not null,"+             
                    "id integer primary key not null)")  
                    if err != nil{
                      fmt.Println(stringTables[6])
                              panic(err)
                    }
              }
        default:
          fmt.Println("Error")
    }
}
// funcao interna par a iniciar o servidor para  busca de id's
func initServer() *sql.DB{
      psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
      db, err := sql.Open("postgres", psqlInfo)
      
      if err != nil {
        panic(err.Error())
      }

      err = db.Ping()
      if err != nil {
        panic(err)
      }
      return db
}
//L
func VerifyId(idConection string) bool{
    db := initServer()
    fmt.Println("Verify your server")
    defer db.Close()
    records,err := db.Query("SELECT id FROM super")
    if err != nil{
        panic(err)
    }
    for records.Next(){
      var id id 
      erroScan:= records.Scan(&id.id)
      if erroScan != nil{
        fmt.Println("erro no VerifyId")
        fmt.Println(erroScan.Error())
        continue
      }
      s := strconv.Itoa(id.id)
      fmt.Print("Valor do Recebido do servidor ")
      fmt.Println(s+" is string"+ idConection)

      if idConection == s{
        fmt.Println("Entrou")
        db.Close()
        records.Close()
        return true
      } 
    }
    return false
}
func InitServer() {
      psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
      db, err := sql.Open("postgres", psqlInfo)
      
      if err != nil {
        panic(err.Error())
        return
      }
      defer db.Close()

      err = db.Ping()
      if err != nil {
        panic(err)
        return
      }
      stringTables  := []string{"powerstats","biography","appearance","work","connections","image","super"}
      /// USADO PARA CRIAR TODAS AS TABELAS
      for i:=6 ; i>=0;i--{createTable(db, stringTables[i])}
      fmt.Println("SUCESS")
}