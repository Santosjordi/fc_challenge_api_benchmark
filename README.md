# Load Tester

Projeto apresentado como desafio técnico para conclusão do curso GOEXPERT.

## How to Run
---

## Running the Load Test CLI

### 1. Build the Docker image

From the project root (where `Dockerfile` and `go.mod` are located), run:

```bash
docker build -t loadtest .
```

This builds a statically linked Go binary from the `cmd/benchmark` folder and packages it into a small Docker image.

---

### 2. Run the load test

Use the following command, replacing the URL, number of requests, and concurrency level as needed:

```bash
docker run --rm loadtest --url=<TARGET_URL> --requests=<TOTAL_REQUESTS> --concurrency=<CONCURRENCY>
```

**Example:**

```bash
docker run --rm loadtest --url=https://pkg.go.dev/std --requests=100 --concurrency=10
```

---

### 3. Output

The CLI prints a summary after the test finishes, including:

* Total requests sent
* Total wall-clock time
* Average time per request
* Requests per second
* HTTP status code distribution

---

## ASSIGNMENT

Objetivo: Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.


O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.


Execução do Teste:

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.
- Geração de Relatório:

Apresentar um relatório ao final dos testes contendo:

- [X] Tempo total gasto na execução
- [X] Quantidade total de requests realizados.
- [X] Quantidade de requests com status HTTP 200.
- [X] Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

Execução da aplicação:
Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:

```bash
docker run <sua imagem docker> —url=http://google.com —requests=1000 —concurrency=10
```