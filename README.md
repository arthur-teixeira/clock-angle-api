# Clock Angle API
Essa API foi desenvolvida para calcular o ângulo entre os ponteiros de um relógio analógico.
Ela retorna sempre o menor ângulo entre os ponteiros, ou seja, se o ângulo entre os ponteiros for maior que 180°, ela retorna o ângulo complementar.

## Como usar
Para usar a API, basta fazer uma requisição GET para o endpoint /api/clock, passando os parâmetros hours e minutes, como no exemplo abaixo:
```
curl --request GET \
  --url 'http://localhost:3000/api/clock?hours=13&minutes=30'
```
O retorno será um JSON com o ângulo entre os ponteiros, como no exemplo abaixo:
```json
{
  "angle": 75
}
```
O parâmetro hours deve ser um número inteiro entre 0 e 23,
e o parâmetro minutes deve ser um número inteiro entre 0 e 59.
Minutes é opcional, caso não seja passado, será considerado 0.


## Como executar
Após clonar o repositório, basta rodar o comando abaixo:
```sh 
$ sudo docker compose up
```

## Como executar os testes
Para executar os testes, basta rodar o comando abaixo:
```sh 
$ go test -v ./...
```
