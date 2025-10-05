# My Tasks
> Status: ✅ Finished

<br/>

## 💻 About The Project

A simple and functional task management application built with **Go** on the backend and **JavaScript** on the frontend. Users manage tasks with it.

<br/>

## 🎯 Features
- Expand Task
- Add New Tasks
- Edit Task
- Remove Task

<br/>

## ⚡Installation

```bash
# Clone the repository
git clone https://github.com/lucasmsaluno/my-tasks.git
cd url-shortener

# Build and start containers
docker-compose up --build
```
<br/>

## 📂 Folder Structure
```bash
backend/                     # Serviço backend em Go
├── cmd/                    # Ponto de entrada da aplicação
│   └── main.go            # Inicializa o servidor HTTP
├── internal/              # Lógica interna da aplicação
│   ├── database/          # Conexão e configuração do banco de dados
│   ├── handlers/          # Funções que lidam com requisições HTTP
│   ├── models/            # Definições das estruturas de dados
│   ├── repositories/      # Acesso aos dados (CRUD)
│   ├── routes/            # Definição das rotas da API
│   └── usecases/          # Regras de negócio e lógica da aplicação
├── go.mod                 # Definição do módulo Go
├── go.sum                 # Checksums das dependências
├── Dockerfile             # Instruções para construir a imagem Docker
└── docker-compose.yml     # Orquestração dos serviços com Docker Compose

frontend/                  # Interface do usuário
├── index.html             # Página principal da aplicação
├── app.js                 # Lógica do frontend (interações e chamadas à API)
└── style.css              # Estilos visuais da interface

.gitignore                 # Arquivos e pastas ignorados pelo Git
README.md                  # Documentação do projeto
   
```
<br/>

## ⚙️ Technologies

- GO
- Javascript
- Docker
- PostgreSQL
