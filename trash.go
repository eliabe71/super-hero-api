func searchId (name string, nameAux types.NameAux, i int) (string) {
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

//////////////////////////////////////////////////

//USADA EM ALGUM MOMENTO DO db.go
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
}
