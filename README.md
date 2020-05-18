# Documentação super-hero-api
  Projeto Api Levpay
# Organização dos Diretórios:
Pasta Principal é Super-Hero-Api/ , lá contém o arquivo (main.go) que é o arquivo para executar as chamadas de metódos (POST, GET, DELETE)

-> Diretório Super-Hero-Api/Connections 
  Dentro do diretório contém o arquivos connections.go  que vai conectar a API(superheroapi.com) com o Banco de Dados (Postresql)

-> Diretório Super-Hero-Api/Trash 
  Alguns pedaços de códigos que forma algum momento usados no projeto

-> Diretório Super-Hero-Api/db 
  contém o arquivo db.go que tem funções que de alguma forma vão manipular o Banco de Dados de alguma Maneira

-> Diretório Super-Hero-Api/types
   Contém o arquivo types.go que são os tipos usados ex:(types.Get usado para receber os dados do Banco)

# Funcionamento dos Arquivos

# Caminho: connections/connections.go 

imports de fora do projeto:

"io/ioutill" -> usado para ler o json da API(superheroapi.com) 

"fmt" -> padrão para input/output

"net/http" -> para fazer requisições do http da API(superheroapi.com)

FUNÇõES:
caminho :/main.go
POST()-> Usadas para criar cadrastos não recebe nada como parâmetros. Recebe entradas de dados para dizer qual o nome do super para ser buscada na API(superheroapi.com) e ser cadrastado no banco chama a função Saving() passando os ids com o nome encontrados.

GET(string) -> Recebe uma String com um parâmetro passada na main.go , retorna uma slice do tipo (type.Get) que estão contidos no (packages types no caminho types/types.go), caso for requisição for das tabelas (allsuper, heroes, villains)
os únicos valores válidos na estrutura (type.Get) são Id e Name.

GET(villains) retornar apenas os vilões 

GET(heroes) retornar apenas os heróis 

GET(allsupers) retonar todos os Supers  cadrastados

DELETE()-> Recebe uma String já definido na main.go caso string recebida seja "all" todas as tabelas do Banco de Dados serão Destruídas. Caso for "name", então será necessária entrada de dados para dizer qual Super será excluir. 

DELETE("all") -> destrói tudo

# Caminho: db/db.go
importações de fora do projeto 
"database/sql" -> para uso no postresql.
"github.com/lib/pq" -> Além de ter implementações necessárias para "database/sql" e foi para colocar array em tipo Json para ser enviados para o Banco de Dados.

Constantes
Nelas contem as informações do Servidor/Banco de Dados e nomes de algumas tabelas.

Funções 
initServer() -> Usada para termos acesso ao servido. Retorna  um ponteiro *sql.DB, para que outras funções tenham acesso  ao servidor.

Saving()-> Recebe dois parâmetros uma slice de (id's) que foram passsados no arquivos (connections.go) esses id's são resultados obtidos na consulta http com o nome digitados pelo usuário.

existTable() -> recebe dois parâmetros (db *sql.DB , nameTable string) db é para ter acesso ao server/Banco de dados e nameTable é nome da tabela a ser pesquisada para ver se ela existe, retorna true caso exista e false caso contrário.

createTables() -> verfica se as tabelas que irá criar já existe. Caso não exista cria. Recebe um ponteiro *sql.DB para a função ter acesso ao banco e poder executar fuções específicas do Postresql.
 
 createSpecificTables(db *sql.DB) cria 3 tabelas específicas (allsuper, heroes, villiains ) essas serão usadas para fazer requisições de apenas heróis ou todos os Super cadrastados ou apenas os vilões.

checkValidityId() -> recebe dois parâmetros (idConection []string, arrayId []types.Id) dois arrays idConection fazendo menção que vem do arquivo do connections.go, nela há todos os id's adquiridos pelo nome digitado para cadrasto, arrayId contém os id de Super Já cadrastados no banco, essa função irá avaliar  pelo id se o super correspondente poderá ser cadrastado. Retorna dois valores, true para sinalizar que algum super pode ser cadrastado e a posição do no array idConection que é a mesma sequência obtida na requisição http.

addInDb() -> recebe três parâmetros (db *sql.DB, super *types.SuperAndVillains, posResponse int) db para acesso as funções do banco , super que é a estrutura que contém todos os dados do SUPER a ser cadrastado e posResponse que a posição obtida pela função checkValidityId(). A função simplemente irá passar os dados da estrutura para as tabelas existentes.

standardDatabaseString() -> retorna uma string com o nome do super que será pesquisado com padrão que está salvo no banco de dados.

SearchSuperTable() -> recebe uma string como parâmetro nome da tabela ,que já está pré definido na main.go (allsuper, heroes, villiains ) retonar uma slice do tipo type.Get más apenas o campos Id e Name são válidos 

SearchSuperId() -> recebe um Id da sua chamada em connections.go e vai pesquisar o Super por esse id retornando um tipo type.Get  todos os campos estando válidos.

SearchSuperName() -> recebe o nome na sua chamada em connections.go e vai pesquisar o super por esse  nome retornando um tipo type.Get todos os campos sendo válidos.

 RemoveDataBase() -> detrois todas as tabelas do Banco de dados.
 
 RemoveSuper() -> recebe o nome do super deixa no padrão do Banco e remove todas as aparições do Super com aquele dterminado nome, caso o super não exista não se faz nada.
 
erasingTheSuper() recebe dois parâmetros db *sql.DB, id int) db para ter o acesso as funções de manipulção do banco e o id para remover todos os atributos do super com aquele id, esssa função elimina cada atributo de todas as tabelas.

# Teste 
Para os teste alterei a função POST na main.go para que ela receba parâmetros e criei um package testes criei uma nova pasta para não alterá a ideia original caminho dos teste é Super-Hero-Api-Test, nele estou adicionando 200 Super's No Banco de Dados, logo após dou os todos os tipos de GET e por fim apago tudo.

Execução dos Testes 
Requistos : Necessário criar um server/banco de dados com essas especificações
  host          =  "localhost"
  port          =  5432
  user          =  "postgres"
  password      =  "eliabe1106"
  dbname        =  "super-hero"
Para executar os testes basta ir na pasta Super-Hero-Api-Test 
digite os comandos no terminal linux go run main.go ou go build main.go && ./main.go 

# Execução Normal 
Va na pasta Super-Hero-Api abra o arquivo main.go e escolha qual função irá executar e comente as outras, logo após digite os comandos no terminal linux go run main.go ou go build main.go && ./main.go.

# Especificações
go version go1.14.2 linux/amd64

Postresql 12

Sistema Operacional: linux
