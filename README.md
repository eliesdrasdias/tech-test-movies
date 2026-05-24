# Movie API

Uma API RESTful para gerenciamento de filmes, construída com foco em simplicidade e performance.

## Como Inicializar a Aplicação

Para subir o ambiente completo com apenas um comando, certifique-se de ter o **Docker** e o **Docker Compose** instalados na sua máquina.

Na raiz do projeto, execute:

```bash
docker compose up --build
```

Aguarde a inicialização dos contêineres. A API estará rodando em: http://localhost:8080.

Para confirmar que a aplicação subiu corretamente, acesse diretamente a interface do Swagger:
**http://localhost:8080/swagger/index.html**

## Principais Endpoints Mapeados
- GET /movies/ - Lista os filmes disponíveis (com limite de performance de 100 registros).

- GET /movies/{id} - Busca os detalhes de um filme específico pelo ID.

- POST /movies/ - Cadastra um novo filme na base de dados.

- DELETE /movies/{id} - Remove um filme pelo ID.

## Exemplos de Uso via cURL
Caso prefira testar a API diretamente pelo terminal, utilize os exemplos abaixo:

### Listar Filmes (GET)
```bash
curl -v http://localhost:8080/movies/
```
### Buscar Filme por ID (GET)
```bash
curl -v http://localhost:8080/movies/8
```