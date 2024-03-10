# Rinha de Backend 2024 - API Crebitos

Esta é uma API HTTP desenvolvida em Go (Golang) para a Rinha de Backend 2024, crebitos

## Tecnologias Usadas

- [Go (Golang)](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [NGINX](https://www.nginx.com/)
- [Docker](https://www.docker.com/)

<div style="display: flex;">
    <img src="https://upload.wikimedia.org/wikipedia/commons/2/23/Golang.png" alt="logo go" width="30%" height="auto">
    <img src="https://upload.wikimedia.org/wikipedia/commons/2/29/Postgresql_elephant.svg" alt="logo postgres" width="30%" height="auto">
    <img src="https://upload.wikimedia.org/wikipedia/commons/c/c5/Nginx_logo.svg" alt="logo nginx" width="30%" height="auto">
</div>

## Pré-requisitos

Certifique-se de ter o seguinte instalado em sua máquina:

- Go 1.22.0 ou superior
- Docker, docker compose

## Instalação e Execução

1. Clone este repositório para o seu ambiente local.
2. Altere para o diretório do projeto:
    ```bash
    cd rinha-crebito
    ```
3. Instale as dependências do Go:
    ```bash
    go mod tidy
    ```
4. Execute o servidor da API:
    ```bash
    go run cmd/main.go
    ```
5. A API estará disponível em `http://localhost:9999` caso não seja inserido alguma porta na variável de ambiente `PORT`.

## Executando com NGINX, exigido pela rinha

1. Clone este repositório para o seu ambiente local.
2. Altere para o diretório do projeto:
    ```bash
    cd rinha-crebito
    ```
3. Execute o docker compose:
    ```bash
    docker compose up -d
    ```
5. A API estará disponível em `http://localhost:9999`, tendo 2 instâncias da API balanceadas pelo NGINX.

## Contribuindo

Se você quiser contribuir para este projeto, sinta-se à vontade para abrir uma issue ou enviar uma pull request.
