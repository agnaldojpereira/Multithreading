# Corrida de APIs de CEP

Este projeto é um programa em Go que realiza uma "corrida" entre duas APIs de consulta de CEP (Código de Endereçamento Postal) no Brasil. O programa faz requisições simultâneas para as APIs BrasilAPI e ViaCEP, exibindo o resultado da API que responder mais rápido.

## Funcionalidades

- Consulta simultânea de CEP em duas APIs diferentes
- Exibição do resultado mais rápido
- Timeout de 1 segundo para as requisições
- Formatação amigável dos dados do endereço

## Pré-requisitos

- Go 1.15 ou superior

## Como usar

1. Clone este repositório:
   ```
   git clone https://github.com/seu-usuario/cep-api-race.git
   ```

2. Navegue até o diretório do projeto:
   ```
   cd cep-api-race
   ```

3. Execute o programa:
   ```
   go run main.go
   ```

4. Digite o CEP desejado quando solicitado.

## Estrutura do código

- `main.go`: Contém todo o código do programa
- `BrasilAPIResponse`: Estrutura para desserializar a resposta da API BrasilAPI
- `ViaCEPResponse`: Estrutura para desserializar a resposta da API ViaCEP
- `fetchCEP`: Função que realiza a requisição HTTP para uma API específica
- `main`: Função principal que coordena as goroutines e exibe o resultado

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

