package db

import (
  "api/types"
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
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
func createTables(db *sql.DB, i int) {
  //não precisamos  da resposta estamos utilzaremos apenas a variável err para ver se houve algum error
  stringTables := [7]string{"powerstats", "biography", "appearance", "work", "connections", "image", "super"}
  switch stringTables[i] {
  case stringTables[6]:
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
  case stringTables[0]:
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
  case stringTables[1]:
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
  case stringTables[2]:
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
  case stringTables[3]:
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
  case stringTables[4]:
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
  case stringTables[5]:
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
  default:
    err := fmt.Errorf("there is no id's")
    panic(err)

  }
  createSpecificTables(db)
}

// funcao interna par a iniciar o servidor para  busca de id's
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
  stringTables := [7]string{"powerstats", "biography", "appearance", "work", "connections", "image", "super"}
  //flags para garantir que ainda não tinha sido feito a table main
  // cASO A TABELA MAIN NÃO EXISTa então significa que não ten cono hAVER CADRASTTOS
  // COMEO DE 6 VRIFICANDO SE ENTRA PELO FATO DO MEUS VETOR DE STRING COM OS NOMES DAS TABLES TEM O NOME DA tableMain na posicao 6
  flags := false
  for i := 6; i >= 0; i-- {
    if !existTable(db, stringTables[i]) {
      createTables(db, i)
      if i == 6 {
        flags = true
      }
    }
  }
  //ADCIONAR NA TABELA
  if flags {
    fmt.Println("SERIA ADCIONADO")
    return true
  }
  fmt.Println(super.Results[1].Id)
  defer db.Close()
  // Vendo os id gravados para ver se poderemos salvar o super com id especificado
  records, err := db.Query("SELECT idsuper FROM super")
  if err != nil {
    panic(err)
  }
  var id id
  for records.Next() {
    erroScan := records.Scan(&id.Id)
    if erroScan != nil {
      fmt.Println("erro no VerifyId")
      fmt.Println(erroScan.Error())
      continue
    }
    //convertende de int para string
    s := strconv.Itoa(id.Id)
    for i := 0; i < len(idConection); i++ {
      if s == idConection[i] {
        continue
      }
      ///Cadrasto AQUi !!!!!!!
      fmt.Println("SERIA ADCIONADO")
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
  // stringTables := []string{"powerstats", "biography", "appearance", "work", "connections", "image", "super"}
  /// USADO PARA CRIAR TODAS AS TABELAS
  for i := 6; i >= 0; i-- {
    createTables(db, i)
  }
  fmt.Println("SUCESS")
}
