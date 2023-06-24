# Clock Angle API
Essa API foi desenvolvida para calcular o ângulo entre os ponteiros de um relógio analógico.
Ela retorna sempre o menor ângulo entre os ponteiros, ou seja, se o ângulo entre os ponteiros for maior que 180°, ela retorna o ângulo complementar.

## Como usar
Para usar a API, basta fazer uma requisição GET para o endpoint /api/clock, passando os parâmetros hours e minutes, como no exemplo abaixo:
```
http://localhost:3000/api/clock?hours=3&minutes=30
```
O retorno será um JSON com o ângulo entre os ponteiros, como no exemplo abaixo:
```json
{
  "angle": 75
}
```

## Como executar
Após clonar o repositório, basta rodar o comando abaixo:
```sh 
sudo docker compose up
```

