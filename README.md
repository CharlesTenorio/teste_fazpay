# FAZ PAY  PRODUCT API

Api Crud de produtos em Golang usando Go-chi

> Requisitos do projeto:

- Go Lang >= 1.18

As demais dependências estão no arquivo go.mod

- https://go.dev/dl/

> Build do Back-End Go:
```bash
# Baixando as dependências
$ go mod tidy

# Compilar servidor HTTP
$ go build -o main cmd/api/main.go

# Ou compilar para outra plataforma ex: windows
$ GOOS=windows GOARCH=amd64 go build -o main64.exe cmd/api/main.go


$ go build -ldflags "-s -w" .
# Ou
$ go build -ldflags "-s -w" cmd/api/main.go
# Ou
$ go build -ldflags "-s -w" -o main cmd/api/main.go
```
## Tem quer seta esssa vars de ambiente antes de rodar a aplicação
roda dentro do diretorio cmd/api
export SRV_DB_DRIVE="postgres"
export SRV_DB_HOST="localhost"
export SRV_DB_PORT="5432"
export SRV_DB_USER="postgres"
export SRV_DB_PASS="supersenha"
export SRV_DB_NAME="fazpay_db_dev"
## Opções de execução
- SRV_PORT (Porta padrão 8080)
- SRV_MODE (developer, homologation ou production / padrão production)

> Exemplo de Uso:
```bash
$ ./main.exe
# Ou
$ SRV_PORT=8080 SRV_MODE=developer ./main.exe
# Ou
$ SRV_PORT=9090 SRV_MODE=production ./main.exe
```

> Acesse:
- http://localhost:8080/
### RODAR postgresql docker para dev
docker-compose up -d dbpostgre 
docker-compose down

#### Para rodar os endpoints utilize o Insommia
https://insomnia.rest/
apos baixa e instalar insomnia abra e va na opção import escolha o arquivo
 ApiProduct.json na pasta docs/Insomnia so executar os endpoint 

### Documentação com Swagger
http://localhost:8080/docs/index.html

### OBS com relação a testes tentei usar o pacotes:
   "github.com/DATA-DOG/go-sqlmock"
  "github.com/stretchr/testify/require"
   Embora eu já tenha utilizado bibliotecas como github.com/DATA-DOG/go-sqlmock e github.com/stretchr/testify/require em meus testes, enfrentei alguns problemas com go-sqlmock, acrdito qeu seja a versão do Go, teria q ver mais afundo para nao acabar atrasando a entrega do teste. Optei por prosseguir sem a realização dos testes para não impactar na entrega . No entanto, é importante ressaltar que sou um defensor dos testes e geralmente os incorporo em meu fluxo de trabalho.

Na arquitetura do meu projeto, priorizei a simplicidade, mantendo, ao mesmo tempo, uma estrutura organizada. Utilizei práticas como injeção de dependência, interfaces e segui os padrões  recomendados pela linguagem. 
