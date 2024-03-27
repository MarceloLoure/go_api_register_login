# Use a imagem oficial do Golang como base
FROM golang:1.22.1-alpine3.19 as builder

# Define o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copia o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixa as dependências do módulo Go
RUN go mod download

# Copia todo o código fonte para o diretório de trabalho
COPY . .

# Compila o aplicativo Go e gera o binário
RUN go build -o /go/bin/app

# Segundo estágio para uma imagem mínima
FROM alpine:3.19.1

# Copia o binário gerado no estágio anterior para o contêiner final
COPY --from=builder /go/bin/app /usr/local/bin/app

# Define o comando padrão a ser executado quando o contêiner iniciar
CMD ["app"]
