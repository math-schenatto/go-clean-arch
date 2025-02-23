# Order System - Go Application

Este é um sistema de pedidos desenvolvido em Go utilizando Clean Architecture. A aplicação expõe três métodos de interação: um webserver, um servidor gRPC e um servidor GraphQL. O projeto utiliza RabbitMQ para mensageria e MySQL para persistência de dados. O ambiente de execução do RabbitMQ e do MySQL está configurado via Docker Compose.

---

## Features

- **WebServer**: Interação através de um arquivo http.
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

## Execution

Para executar o projeto, siga os comandos abaixo:

1. **Run the application**:

   Para executar esse projeto é necessário apenas executar o comando docker:
   ```bash
   docker compose up --build
   ```

   Ele vai subir todos os containers necessários e a aplicação.

   O banco de dados já será criado com alguns registros na tabela de orders.

    **WebServer**: Porta **8000**
    **GraphQL**: Porta **8080**
    **gRPC**: Porta **50051**
    



   # Para testar o webserver:

   1. Acessar a pasta api do projeto e abrir o arquivo list_orders.http e executar a request. A request é feita para a porta 8000.

   ```bash

        ### Listar pedidos (GET /list)
        GET http://localhost:8000/list?page=1&limit=10 HTTP/1.1
        Host: localhost:8000
        Content-Type: application/json

   ```


   # Para testar o GraphQl:

   1. Acessar o playground na porta 8080 (http://localhost:8080/)
    ```bash

        query {
            ListOrders(page: 1, limit: 10) {
                Id
                Price
                Tax
                FinalPrice
            }
        }

    ```


   # Para testar o gRpc:

   1. Necessário acessar o evans:

   ```bash
        1. evans -r repl
        2. package pb
        3. service OrderService
        4. call ListOrders 
   ```

   2. Será solicitado page e limit, você pode adicionar 1 e 10 respectivamente.


