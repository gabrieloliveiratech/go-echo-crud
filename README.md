# Go Echo CRUD API

Este é um projeto de API REST que utiliza Go e o framework Echo para simular um CRUD de produtos. A API permite que você gerencie Produtos e suas informações, como "name", "price" e "stock".

## Pré-requisitos

- Docker
- Go

## Como executar o projeto

1. Copie o arquivo `.env.example` para `.env` e altere as variáveis de ambiente, se necessário.

```bash
cp .env.example .env
```

2. Suba o Docker Compose, que criará um banco de dados MySQL para o projeto.

```bash
docker-compose up
```

3. Execute o projeto.

```bash
go run main.go
```

## Comandos opcionais

- Para executar os testes:

```bash
go test ./handlers_test -v
```

- Para recriar a documentação:

```bash
swag i
```

## Documentação

A documentação da API está disponível em(se você não modificar a porta): [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)