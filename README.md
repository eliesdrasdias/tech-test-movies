# 🎬 Movies API

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-4285F4?style=for-the-badge&logo=google&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)

> Uma API robusta para gerenciamento de catálogo de filmes, construída com foco em escalabilidade, performance e baixo acoplamento.

---

## 📝 Visão Geral

O **Movies API** é uma solução de backend desenvolvida em Go, projetada para gerenciar informações de filmes. O projeto evoluiu de uma estrutura monolítica para uma **arquitetura de microsserviços**, visando simular um ambiente de produção real onde escalabilidade horizontal e independência de serviços são fundamentais.

O foco principal deste projeto é demonstrar o domínio de conceitos avançados de backend, como comunicação inter-serviços via gRPC, orquestração com Docker e organização de workspaces em Go.

## 🏗️ Arquitetura e Decisões Técnicas

A aplicação foi decomposta para garantir que cada componente tenha uma responsabilidade única:

- **API Gateway:** O ponto de entrada único para o cliente. Ele recebe requisições HTTP RESTful e as traduz para chamadas gRPC de alta performance para os serviços internos. Isso protege a infraestrutura interna e simplifica a interface para o front-end.
- **Movie Service:** Responsável pela lógica de negócio, persistência e manipulação dos dados de filmes.
- **gRPC & Protobuf (`/pb`):** Utilizados para a comunicação entre o Gateway e o Service. A escolha do gRPC se deve à sua eficiência (binário), contratos tipados estritamente e performance superior ao JSON/HTTP tradicional em redes internas.
- **Go Workspaces (`go.work`):** Utilizado para gerenciar múltiplos módulos localmente de forma fluida, facilitando o desenvolvimento sem a necessidade de publicar módulos privados.

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Comunicação:** gRPC e Protocol Buffers (Protobuf)
- **Gateway:** HTTP/REST
- **Banco de Dados:** PostgreSQL
- **Orquestração:** Docker & Docker Compose
- **Gerenciamento de Dependências:** Go Modules & Workspaces

## 🚀 Como Executar

O projeto foi totalmente "dockerizado" para garantir que o ambiente suba de forma idêntica em qualquer máquina.

### Pré-requisitos

- [Docker](https://www.docker.com/get-started) instalado.
- [Docker Compose](https://docs.docker.com/compose/install/) instalado.

### Passo a Passo

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/tech-test-movies.git
   cd tech-test-movies
   ```

2. Suba todo o ecossistema (Gateway, Microserviço e Banco de Dados) com um único comando:

   ```bash
   docker-compose up --build
   ```

3. A API estará disponível em: `http://localhost:8080` (ou a porta configurada no seu gateway).

## 🛣️ Endpoints Principais (API Gateway)

| Método   | Endpoint      | Descrição                                |
| :------- | :------------ | :--------------------------------------- |
| `GET`    | `/movies`     | Lista todos os filmes cadastrados.       |
| `GET`    | `/movies/:id` | Retorna detalhes de um filme específico. |
| `POST`   | `/movies`     | Cadastra um novo filme.                  |
| `PUT`    | `/movies/:id` | Atualiza os dados de um filme.           |
| `DELETE` | `/movies/:id` | Remove um filme do catálogo.             |

---

Desenvolvido com ☕ e Go por Eliesdras Dias.
