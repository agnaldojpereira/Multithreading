package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BrasilAPIResponse define a estrutura da resposta da API BrasilAPI
type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

// ViaCEPResponse define a estrutura da resposta da API ViaCEP
type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

// fetchCEP realiza uma requisição HTTP para a API especificada e envia o resultado formatado para o canal
func fetchCEP(url string, apiName string, ch chan<- string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- fmt.Sprintf("Erro ao criar requisição para %s: %v", apiName, err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Erro na requisição para %s: %v", apiName, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Erro ao ler resposta de %s: %v", apiName, err)
		return
	}

	var result string
	switch apiName {
	case "BrasilAPI":
		var data BrasilAPIResponse
		json.Unmarshal(body, &data)
		result = fmt.Sprintf("BrasilAPI: %s, %s - %s, %s", data.Street, data.Neighborhood, data.City, data.State)
	case "ViaCEP":
		var data ViaCEPResponse
		json.Unmarshal(body, &data)
		result = fmt.Sprintf("ViaCEP: %s, %s - %s, %s", data.Logradouro, data.Bairro, data.Localidade, data.Uf)
	}

	ch <- result
}

func main() {
	var cep string
	fmt.Print("Digite o CEP (apenas números): ")
	fmt.Scanln(&cep)

	brasilAPIURL := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	viaCEPURL := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	// Cria um canal para receber os resultados das goroutines
	ch := make(chan string, 2)

	// Inicia as goroutines para fazer as requisições concorrentemente
	go fetchCEP(brasilAPIURL, "BrasilAPI", ch)
	go fetchCEP(viaCEPURL, "ViaCEP", ch)

	// Espera pelo primeiro resultado ou pelo timeout de 1 segundo
	select {
	case result := <-ch:
		fmt.Println("Resultado mais rápido:")
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: Timeout - As APIs demoraram mais de 1 segundo para responder.")
	}
}
