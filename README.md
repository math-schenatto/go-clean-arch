# Order System - Go Application

Este é um sistema de pedidos desenvolvido em Go utilizando Clean Architecture. A aplicação expõe três métodos de interação: um webserver, um servidor gRPC e um servidor GraphQL. O projeto utiliza RabbitMQ para mensageria e MySQL para persistência de dados. O ambiente de execução do RabbitMQ e do MySQL está configurado via Docker Compose.

---

## Features

- **WebServer**: Interação através de um arquivo HTML.
- **gRPC**: Interação via protocolo gRPC.
- **GraphQL**: Interação através de uma API GraphQL.
- **Mensageria**: Utiliza RabbitMQ para comunicação assíncrona.
- **Banco de Dados**: MySQL para persistência de ordens.

---

## APIs Used

- **gRPC**: Endpoints para criação e listagem de ordens.
- **GraphQL**: Queries e Mutations para interação com o sistema de pedidos.
- **WebServer**: Arquivo http para criação e listagem de ordens.

---



## Requirements

- Go version 1.20 

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/math-schenatto/go-clean-arch
   cd go-clean-arch
    ```

## Execution

Para executar o projeto, siga os comandos abaixo:

1. **Run the application**:
   Navegue até o diretório do projeto e execute o seguinte comando:
   ```bash
   go run cmd/ordersystem/main.go wire_gen.go


## Test WebServer

Basta utilizar o arquivo create_order.http dentro da pasta api

## Test Graphql

Execute o playground na porta 8080
```bash
query {
    orders(page: 1, limit: 10) {
        Id
        Price
        Tax
        FinalPrice
    }
}
```

## Test gRpc

Execute o evans.

1. evans -r repl
2. package pb
3. service OrderService
4. call ListOrders