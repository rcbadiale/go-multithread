# Go Expert Multithread Requests

Desafio de aplicaçào dos conhecimentos sobre multithreading em Go para a pós
graduação Go Expert.

## Especificações

Deverão ser feitas duas requisições em paralelo consultando os dados de um CEP utilizando as duas APIs abaixo.
- https://brasilapi.com.br/api/cep/v1/<cep>
- http://viacep.com.br/ws/<cep>/json/

A API que responder mais rápido deverá ter o resultado apresentado na linha de
comando indicando qual a URL utilizada.

Caso as APIs demorem mais que 1s deverá ser retornado um timeout na linha de
comando.

## Execução

Necessário utilizar Go >= 1.22.0.

Ao executar o CEP deverá ser fornecido na linha de comando de acordo com o
exemplo abaixo.
```shell
go run main.go 70150904
```

A resposta esperada segue o formato abaixo.
```shell
URL: <url que respondeu primeiro>
Response: <dados do endereço em formato json>
```

Em caso de erro será apresentada a mensagem de timeout.
```shell
timed out after 1s
```
