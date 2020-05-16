package db

import (
  "api/types"
  "database/sql"
  "fmt"
  "github.com/lib/pq"
  //"errors"
  "strconv"
)

const (
  host          = "localhost"
  port          = 5432
  user          = "postgres"
  password      = "eliabe1106"
  dbname        = "super-hero"
  tableMain     = "super"
  tableHeroes   = "heroes"
  tableVillains = "villains"
)

// usado para receber json dos id's apenas uso interno
/*
  TABLES QUE SERÃO CRIADA
*/
//funcao para verifcar se a table existe
// usada internamente
// retorna true se a tabela existe e false caso contrário
func existTable(db *sql.DB, nameTable string) bool {
  _, error := db.Exec("SELECT FROM " + nameTable)
  if error != nil {
    //fmt.Println(error.Error())
    return false
  }
  return true
}

type id struct {
  Id int `json:"idsuper"`
}

func createSpecificTables(db *sql.DB) {
  if !existTable(db, tableHeroes) {
    _, err := db.Exec("CREATE TABLE " + tableHeroes + " (" +
      "heroid serial primary key," +
      "name varchar(10000) not null," +
      "idsuper integer references super(idsuper)" + ")")
    if err != nil {
      panic(err)
    }
  }
  if !existTable(db, tableVillains) {
    _, err := db.Exec("CREATE TABLE " + tableVillains + " (" +
      "villainid serial primary key," +
      "name varchar(10000) not null," +
      "idsuper integer references super(idsuper)" + ")")
    if err != nil {
      panic(err)
    }
  }
}

// Criar tabela
func createTables(db *sql.DB) {
  //não precisamos  da resposta estamos utilzaremos apenas a variável err para ver se houve algum error
  stringTables := [7]string{"powerstats", "biography", "appearance", "work", "connections", "image", "super"}
  //switch stringTables[i] {
  //case stringTables[6]:
  if existTable(db, stringTables[6]) == false {
    _, err := db.Exec("CREATE TABLE " + stringTables[6] + " (" +
      "name varchar(10000) not null," +
      "idsuper int primary key )")
    fmt.Print("Create ")
    fmt.Println(stringTables[6])
    if err != nil {
      fmt.Println(stringTables[6])
      panic(err)
    }
    fmt.Println("Ja feito ")
    fmt.Println(stringTables[6])
  }
  // case stringTables[0]:
  if existTable(db, stringTables[0]) == false {
    _, err := db.Exec("CREATE TABLE " + stringTables[0] + " (" +
      "idsuper int references super(idsuper)," +
      "idpower serial  primary key," +
      "intelligence varchar(10000) not null," +
      "strength varchar(10000) not null," +
      "speed varchar(10000) not null," +
      "durability varchar(10000) not null," +
      "power varchar(10000) not null," +
      "combat varchar(10000) not null" + ")")
    if err != nil {
      fmt.Println(stringTables[0])
      panic(err)
    }
  }
  // case stringTables[1]:
  if existTable(db, stringTables[1]) == false {
    _, err := db.Exec("CREATE TABLE " + stringTables[1] + " (" +
      "idsuper int references super(idsuper)," +
      "idbio serial primary key," +
      "fullname varchar(10000) not null," +
      "alteregos varchar(10000) not null," +
      "aliases varchar[10000] not null," +
      "placeofbirth varchar(10000) not null," +
      "firstappearance varchar(10000) not null," +
      "publisher varchar(10000) not null," +
      "alignment varchar(10000) not null" + ")")
    if err != nil {
      fmt.Println(stringTables[1])
      panic(err)
    }
  }
  //case stringTables[2]:
  if existTable(db, stringTables[2]) == false {
    _, err := db.Exec("CREATE TABLE " + stringTables[2] + "(" +
      "idsuper int references super(idsuper)," +
      "idappearance serial primary key," +
      "gender varchar(10000) not null," +
      "race varchar(10000) not null," +
      "height varchar[10000] not null," +
      "weight varchar[10000] not null," +
      "eyecolor varchar(10000) not null," +
      "haircolor varchar(10000) not null" + ")")
    if err != nil {
      fmt.Println(stringTables[2])
      panic(err)
    }
  }
  //case stringTables[3]:
  if existTable(db, stringTables[3]) == false {
    _, err := db.Exec("CREATE TABLE " + stringTables[3] + " (" +
      "idsuper int references super(idsuper)," +
      "idwork serial primary key not null," +
      "occupation varchar(10000) not null," +
      "baseofoperation varchar(10000) not null" + ")")
    fmt.Println(stringTables[3])
    if err != nil {
      panic(err)
    }
  }
  // case stringTables[4]:
  if existTable(db, stringTables[4]) == false {
    _, err := db.Exec("CREATE TABLE " + stringTables[4] + " (" +
      "idsuper int references super(idsuper)," +
      "idcon serial primary key not null," +
      "groupaffiliation varchar(10000) not null," +
      "relatives varchar(10000) not null" + ")")
    if err != nil {
      fmt.Println(stringTables[4])
      panic(err)
    }
  }
  //case stringTables[5]:
  if existTable(db, stringTables[5]) == false {
    _, err := db.Exec("CREATE TABLE " + stringTables[5] + " (" +
      "idsuper INT references super(idsuper)," +
      "idurl serial primary key not null," +
      "url varchar(100000) not null )")
    if err != nil {
      fmt.Println(stringTables[5])
      panic(err)
    }
  }
  createSpecificTables(db)
}

// funcao interna par a iniciar o servidor para  busca de id's
func addInDb(db *sql.DB, super *types.SuperAndVillains, posResponse int) {

  // convertendo o id pois ele vem formato de string e é salvo do banco de dados com integer
  convertIntId, _ := strconv.Atoi(super.Results[posResponse].Id)
  //"powerstats", "biography", "appearance", "work", "connections", "image", "super"
  _, err := db.Exec(`INSERT INTO super (idsuper,name) VALUES ($1,$2)`, convertIntId, super.Results[posResponse].Name)
  if err != nil {
    panic(err)
  }
  fmt.Println("Super")
  //-------------------------------------------------------------------------------------------------------
  _, err = db.Exec(`INSERT INTO powerstats (idsuper,intelligence, strength, speed,durability,power,combat) VALUES ($1,$2, $3,$4,$5,$6,$7)`,
    convertIntId, super.Results[posResponse].Powerstats.Intelligence, super.Results[posResponse].Powerstats.Strength,
    super.Results[posResponse].Powerstats.Speed, super.Results[posResponse].Powerstats.Durability,
    super.Results[posResponse].Powerstats.Power, super.Results[posResponse].Powerstats.Combat)
  if err != nil {
    panic(err)
  }
  fmt.Println("Powerstats")
  _, err = db.Exec(`INSERT INTO  biography (idsuper,fullname, alteregos, aliases, placeofbirth,firstappearance, publisher,alignment)`+
            `VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
    convertIntId, super.Results[posResponse].Biography.FullName, super.Results[posResponse].Biography.AlterEgos,
    pq.Array(super.Results[posResponse].Biography.Aliases), super.Results[posResponse].Biography.PlaceOfBirth,
    super.Results[posResponse].Biography.FirstAppearance, super.Results[posResponse].Biography.Publisher,
    super.Results[posResponse].Biography.Alignment)
  if err != nil {
    panic(err)
  }
  fmt.Println("Biography")
  _, err = db.Exec(`INSERT INTO appearance (idsuper, gender, race, height,weight,eyecolor,haircolor) VALUES ($1,$2, $3,$4,$5,$6,$7)`,
    convertIntId, super.Results[posResponse].Appearance.Gender, super.Results[posResponse].Appearance.Race,
    pq.Array(super.Results[posResponse].Appearance.Height), pq.Array(super.Results[posResponse].Appearance.Weight),
   super.Results[posResponse].Appearance.Eyecolor, super.Results[posResponse].Appearance.HairColor)
  
  if err != nil {
    panic(err)
  }
  fmt.Println("Appearance")
  
  _, err = db.Exec(`INSERT INTO work (idsuper, occupation,baseofoperation) VALUES ($1,$2,$3)`,
    convertIntId, super.Results[posResponse].Work.Occupation, super.Results[posResponse].Work.BaseOfOperation)
  if err != nil {
    panic(err)
  }
  
  fmt.Println("Work")
  _, err = db.Exec(`INSERT INTO connections (idsuper,groupaffiliation,relatives) VALUES ($1,$2,$3)`,
    convertIntId, super.Results[posResponse].Connections.GroupAffiliation, super.Results[posResponse].Connections.Relatives)
  if err != nil {
    panic(err)
  }
  fmt.Println("Connections")
  _, err = db.Exec(`INSERT INTO image (idsuper,url) VALUES ($1,$2)`, convertIntId, super.Results[posResponse].Image.Url)
  if err != nil {
    panic(err)
  }
  fmt.Println("Image")
  switch super.Results[posResponse].Biography.Alignment {
  case "good":
    _, err = db.Exec(`INSERT INTO heroes (idsuper,name) VALUES ($1,$2)`, convertIntId, super.Results[posResponse].Name)
    if err != nil {
      panic(err)
    }
    fmt.Println("heroes")
  case "bad":
    _, err = db.Exec(`INSERT INTO villains (idsuper,name) VALUES ($1,$2)`, convertIntId, super.Results[posResponse].Name)

    if err != nil {
      panic(err)
    }
    fmt.Println("villains")
  default:
  }
}

func initServer() *sql.DB {
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

// Idconection fazendo alusao que a variável vem do arquivo conection.gp
// Retorna true se foi possível adicionar
func Saving(idConection []string, super *types.SuperAndVillains) bool {
  if len(idConection) == 0 {
    err := fmt.Errorf("there is no id's")
    panic(err)
  }
  db := initServer()
  fmt.Println("Verify your server")
  //se a tablea não existir cria tabela
  //stringTables := [7]string{"powerstats", "biography", "appearance", "work", "connections", "image", "super"}
  //flags para garantir que ainda não tinha sido feito a table main ou alguma outras tabelas
  // cASO A TABELA MAIN NÃO EXISTa então significa que não ten cono hAVER CADRASTTOS
  // COMEO DE 6 VRIFICANDO SE ENTRA PELO FATO DO MEUS VETOR DE STRING COM OS NOMES DAS TABLES TEM O NOME DA tableMain na posicao 6

  // SE A TABLE MAIN NÃO EXISTE NENHUMA OUTRA EXISTE
  // nesse caso o primeiro registo será passsado 
  if !existTable(db, tableMain) {

    createTables(db)
    addInDb(db,super,0)
    fmt.Println("Cadrastado com Sucesso")
    return true
  }
  fmt.Println("table exist")
  defer db.Close()
  // Vendo os id gravados para ver se poderemos salvar o super com id especificado
  records, err := db.Query("SELECT idsuper FROM super")
  if err != nil {
    panic(err)
  }
  if !records.Next(){
       fmt.Println("Next() equal null")
      addInDb(db,super,0)
  }
  if err != nil {
    panic(err)
  }
  records, err = db.Query("SELECT idsuper FROM super")
  var id id
  fmt.Println("Next() different null")
  for records.Next(){
     fmt.Println("Aqui")
    erroScan := records.Scan(&id.Id)
    if erroScan != nil {
      fmt.Println("erro no VerifyId")
      fmt.Println(erroScan.Error())
      continue
    }
    //convertende de int para string
    s := strconv.Itoa(id.Id)
    fmt.Println("Aqui")
    for i := 0; i < len(idConection); i++ {
      fmt.Println(s+" "+idConection[i])
      if s == idConection[i] {
        continue
      }
      ///Cadrasto AQUi !!!!!!!
      addInDb(db,super,i)
      fmt.Println("Cadrastado com Sucesso")
      return true
    }
  }
  return false
}
