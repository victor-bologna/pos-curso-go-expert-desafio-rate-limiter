# Desafio Rate Limiter 

Essse desafio tem como criar um Middleware capaz de limitar o acesso de requisições por segundo de um usuário por IP ou Token via header, onde o limite de requisições, tempo de bloqueio podem ser configurados em um arquivo ``.env`` na pasta raíz do projeto.

## Requisitos

Para utilizar esse app é necessário ter instalado no seu computador o seguinte app:

- Docker

## Como Usar o app

1. Clone o repositório completo da seguinte URL: `https://github.com/victor-bologna/pos-curso-go-expert-desafio-rate-limiter`
2. (opcional) Crie um arquivo de configuração com nome ``.env`` na pasta raiz do projeto para configurar o número de requisições, configurações do Redis e o tempo que o usuário ficará bloqueado de acessar a URL desejada. Caso uma ou nenhuma informação será adicionada, será usado valores default que podem ser vistos no log. Mais abaixo segue um exemplo de como criar um arquivo ``.env``.
3. Na pasta root do projeto, usar o seguinte comando: `docker compose up --build -d` para executar o app.

### Template arquivo .env

#### .env

```ini
REDIS_ADDR=redis:6379
REDIS_PASS=
REDIS_DB=0
IP_MAX_REQ=1
TOKEN_MAX_REQ=2
TIMEOUT=10
```