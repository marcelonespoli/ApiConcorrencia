# APIs Concorrencia para consulta de CEP em Go

Este projeto é uma implementação em **Go (Golang)** que faz a consulta de um CEP em **duas APIs públicas** de forma concorrente:

- [ViaCEP](https://viacep.com.br/)
- [BrasilAPI](https://brasilapi.com.br/)


## Como funciona

1. Duas goroutines são disparadas ao mesmo tempo, uma para cada API.
2. O conteúdo retornado pela API vencedora é exibido no console.

## Como executar

```
go run main.go
