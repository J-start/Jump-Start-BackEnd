# Endpoints da API 

<details>
  <summary>/datas/shares</summary>

  ## Método do tipo GET que retorna a útima inserção ao banco de dados de todas as ações. Em resumo, retorna o dado mais atual de cada ação.
 
  ### Exemplo de retorno:
  ```json
[
  {
    "Id": 99,
    "NameShare": "ABEV3.SA",
    "DateShare": "2024-11-30",
    "OpenShare": 13.09,
    "HighShare": 13.12,
    "LowShare": 12.84,
    "CloseShare": 12.88,
    "VolumeShare": 31193700
  },
  {
    "Id": 92,
    "NameShare": "BBAS3.SA",
    "DateShare": "2024-11-30",
    "OpenShare": 26.28,
    "HighShare": 26.46,
    "LowShare": 26.17,
    "CloseShare": 26.33,
    "VolumeShare": 12175400
  }
]
  ```
</details>

<details>
  <summary>/datas/shares/offset</summary>

  ## Método do tipo GET que retorna os dados das ações de forma páginada, retornando 10 registro por vez. Para páginação é preciso passar na url um valor inteiro maior que 0 incrementado de 1 em 1, a cada soma de 1 que ocorre no parametro de páginação mais 10 registros são retornados até que não haja mais registros(ações).




  ### Exemplo de chamada

  ```bash
http://localhost:8080/datas/shares/offset?offset=1
  ```
 
  ### Exemplo de retorno:
  ```json
[
  {
    "Id": 99,
    "NameShare": "ABEV3.SA",
    "DateShare": "2024-11-30",
    "OpenShare": 13.09,
    "HighShare": 13.12,
    "LowShare": 12.84,
    "CloseShare": 12.88,
    "VolumeShare": 31193700
  },
  {
    "Id": 92,
    "NameShare": "BBAS3.SA",
    "DateShare": "2024-11-30",
    "OpenShare": 26.28,
    "HighShare": 26.46,
    "LowShare": 26.17,
    "CloseShare": 26.33,
    "VolumeShare": 12175400
  }
]
  ```
</details>


<details>
  <summary>/data/share/</summary>

  ## Método do tipo GET que retorna os dados de uma ação específica, somente o útimo registro e não um histórico
   ### Exemplo de chamada

  ```bash
http://localhost:8080/data/share/offset?shareName=PETR4.SA
  ```
  ### Exemplo de retorno:
  ```json
{
  "Id": 91,
  "NameShare": "PETR4.SA",
  "DateShare": "2024-11-30",
  "OpenShare": 37.6,
  "HighShare": 37.65,
  "LowShare": 37.32,
  "CloseShare": 37.62,
  "VolumeShare": 16343000
}
  ```
</details>

<details>
  <summary>/datas/share/</summary>

  ## Método do tipo GET que retorna todos os dados de uma ação específica, um histórico da ação escolhida

  ```bash
http://localhost:8080/datas/share?shareName=PETR4.SA
  ```
  ### Exemplo de retorno:
  ```json
[
  {
    "Id": 71,
    "NameShare": "PETR4.SA",
    "DateShare": "2024-10-12",
    "OpenShare": 37.6,
    "HighShare": 37.65,
    "LowShare": 37.32,
    "CloseShare": 37.62,
    "VolumeShare": 16343000
  },
  {
    "Id": 101,
    "NameShare": "PETR4.SA",
    "DateShare": "2024-11-14",
    "OpenShare": 38,
    "HighShare": 38.4,
    "LowShare": 37.8,
    "CloseShare": 38.2,
    "VolumeShare": 16500000
  },
  {
    "Id": 102,
    "NameShare": "PETR4.SA",
    "DateShare": "2024-11-15",
    "OpenShare": 38.3,
    "HighShare": 38.6,
    "LowShare": 38,
    "CloseShare": 38.5,
    "VolumeShare": 16850000
  }
]
  ```
</details>

<details>
  <summary>/buy/</summary>

  ## Método do tipo POST para realizar a compra de um ativo

  ```bash
http://localhost:8080/buy/
  ```
  ### Corpo da requisição esperado:
  ```json
{
    "AssetName":"CHZ",
    "AssetCode":"CHZ-BRL",
    "AssetType":"CRYPTO",
    "AssetAmount":10,
    "OperationType":"BUY",
    "CodeInvestor":"12345"
    
  }
  ```

  ## Se tudo ocorrer bem o código 200 é retornado com uma mensagem de sucesso, caso contrário, um código específico e uma mensagem de erro é retornado .
</details>

</details>

<details>
  <summary>/sell/</summary>

  ## Método do tipo POST para realizar a venda de um ativo

  ```bash
http://localhost:8080/sell/
  ```
  ### Corpo da requisição esperado:
  ```json
{
    "AssetName":"CHZ",
    "AssetCode":"CHZ-BRL",
    "AssetType":"CRYPTO",
    "AssetAmount":10,
    "OperationType":"SELL",
    "CodeInvestor":"12345"
    
  }
  ```

  ## Se tudo ocorrer bem o código 200 é retornado com uma mensagem de sucesso, caso contrário, um código específico e uma mensagem de erro é retornado .
</details>