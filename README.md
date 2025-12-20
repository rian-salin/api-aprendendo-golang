# API — Gerenciamento de Biblioteca

API RESTful desenvolvida como Projeto de Extensão, com foco em backend, segurança e boas práticas utilizando Go (Golang).

O sistema permite que usuários se cadastrem, realizem login e gerenciem sua própria biblioteca pessoal, garantindo isolamento de dados por meio de autenticação JWT.

---

## Objetivo

Desenvolver uma aplicação backend robusta aplicando conceitos de:
- Arquitetura REST
- Autenticação stateless com JWT
- Integração com banco de dados relacional
- Criptografia de senhas
- Boas práticas de organização de código

---

## Tecnologias Utilizadas

- Go (Golang) v1.24+
- Gorilla Mux
- MySQL
- JWT (JSON Web Token)
- BCrypt
- Go-SQL-Driver

---

## Funcionalidades

- Cadastro de usuários com senha criptografada
- Autenticação com geração de token JWT
- CRUD completo de livros
- Associação de livros ao usuário autenticado
- Proteção de rotas via middleware

---

## Autenticação

Após o login, a API retorna um token JWT, que deve ser enviado no header das requisições protegidas:

Authorization: Bearer SEU_TOKEN_AQUI

---

## Rotas da API

Usuários e Autenticação

Método: POST  
Endpoint: /users  
Descrição: Cadastro de novo usuário  
Auth: Não  

Método: POST  
Endpoint: /login  
Descrição: Login e geração do token JWT  
Auth: Não  

Livros (Rotas Protegidas)

Método: POST  
Endpoint: /books  
Descrição: Cadastra um livro  

Método: GET  
Endpoint: /users/{id}/books  
Descrição: Lista os livros do usuário  

Método: PUT  
Endpoint: /books/{id}  
Descrição: Atualiza um livro  

Método: DELETE  
Endpoint: /books/{id}  
Descrição: Remove um livro  

---

## Guia de Instalação e Execução

Requisitos:
- Go 1.24+
- MySQL em execução
- Acesso ao terminal na raiz do projeto

Instalação das dependências Go:

    go mod download

Criação do banco de dados:

    mysql -u root -p < sql/ddl.sql

Esse script cria o schema treehousedb e as tabelas necessárias.

Configuração das variáveis de ambiente:

Crie o arquivo config/.env com o conteúdo abaixo:

    API_PORT=":8080"
    DB_USER="root"
    DB_PASSWORD="sua_senha_mysql"
    DB_ADDR="127.0.0.1:3306"
    DB_DATABASE="treehousedb"
    SECRET_KEY="uma_chave_secreta_forte_para_o_jwt"

Ajuste os valores conforme seu ambiente.

Execução da aplicação:

    go run ./main.go

A API ficará disponível em:

    http://localhost:8080

Build do executável (opcional):

    go build -o bin/api ./...
    ./bin/api

---

## Membros do Grupo

Isabela Vianna  
Rian Salin  

---

Projeto de Extensão — Engenharia de Software
