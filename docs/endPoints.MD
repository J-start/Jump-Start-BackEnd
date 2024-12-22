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
<details>
<summary>/details/asset/</summary>

  ## Método do tipo GET que retorna a lista de ativos disponíveis de cada tipo.

  ```bash
  /details/asset/?type=SHARE
  ```
 ### O campo type pode receber 3 tipos, sendo eles: COIN ou CRYPTO ou SHARE
  ### Exemplo de retorno:
  ```json
[
  {
    "id": 31,
    "name": "AMBEV S/A ON (ABEV3.SA)",
    "acronym": "ABEV3.SA",
    "urlImage": "https://investidor10.com.br/storage/companies/66b65af53af6c.jpg",
    "typeAsset": "SHARE"
  },
  {
    "id": 32,
    "name": "Banco do Brasil (BBAS3.SA)",
    "acronym": "BBAS3.SA",
    "urlImage": "https://investidor10.com.br/storage/companies/66b65b3de91ca.jpg",
    "typeAsset": "SHARE"
  },
  {
    "id": 33,
    "name": "Cemig (CMIG4.SA)",
    "acronym": "CMIG4.SA",
    "urlImage": "https://investidor10.com.br/storage/companies/5ea0b6985411c.jpeg",
    "typeAsset": "SHARE"
  }
]
  ```
</details>


<details>
<summary>/history/assets/</summary>

  ## Método do tipo POST que retorna o histórico de compra e venda de ativos.

  ```bash
  /history/assets/
  ```
 ### O endpoint espera receber um json no body da requisição, o body possui a seguinte estrutura:
```json
  {
      "TokenUser": "string",
      "OffSet": 0     
  } 

```

- TokenUser: token jwt do usuário
- offset: páginação do histórico, a cada incremento de 1 no offset é retornado 10 dados

  ### Exemplo de retorno:
```json
[
  {
    "AssetName": "BITCOIN",
    "AssetType": "CRYPTOMOEDA",
    "AssetQuantity": 0.0003,
    "AssetValue": 320.89,
    "OperationType": "BUY",
    "OperationDate": "2024-10-27"
  },
  {
    "AssetName": "BBAS3.SA",
    "AssetType": "AÇÃO",
    "AssetQuantity": 2,
    "AssetValue": 14.78,
    "OperationType": "BUY",
    "OperationDate": "2024-10-26"
  },
]
  ```
</details>

<details>
<summary>/wallet/datas/</summary>

  ## Método do tipo POST que retorna quais ativos/quantidade o usuário possui em carteira e o seu respectivo saldo atual.

  ```bash
  /wallet/datas/
  ```
 ### O endpoint espera receber um json no body da requisição, o body possui a seguinte estrutura:
```json
  {
    "tokenInvestor": "string"
  }

```

- tokenInvestor: token jwt do usuário

  ### Exemplo de retorno:
```json
{
  "InvestorBalance": 855.4034,
  "Assets": [
    {
      "AssetName": "SEK-BRL",
      "AssetType": "COIN",
      "AssetQuantity": 12
    },
    {
      "AssetName": "LINK-BRL",
      "AssetType": "CRYPTO",
      "AssetQuantity": 1
    }
  ]
}
  ```
</details>


<details>

<summary>/history/operations/</summary>

  ## Método do tipo POST que retorna o histórico de depósito e saque.

  ```bash
  /history/operations/
  ```
 ### O endpoint espera receber um json no body da requisição, o body possui a seguinte estrutura:
```json
{
  "TokenUser  ": "string",
  "offset":   0
}

```

- TokenUser: token jwt do usuário
- offset: páginação do histórico, a cada incremento de 1 no offset é retornado 10 dados

  ### Exemplo de retorno:
```json
[
  {
    "OperationType": "WITHDRAW",
    "OperationValue": 150,
    "DperationDate": "2024-12-22"
  },
  {
    "OperationType": "DEPOSIT",
    "OperationValue": 50,
    "DperationDate": "2024-12-22"
  },
  {
    "OperationType": "DEPOSIT",
    "OperationValue": 50,
    "DperationDate": "2024-12-22"
  }
]
  ```
</details>

<details>

<summary>/withdraw/</summary>

  ## Método do tipo POST que realiza um saque em carteira do usuário.

  ```bash
  /withdraw/
  ```
 ### O endpoint espera receber um json no body da requisição, o body possui a seguinte estrutura:
```json
{
  "TokenInvestor ": "string",
  "value":   50
}

```

- TokenUser: token jwt do usuário
- value: valor para realizar o saque, deve ser maior que 0

  ### Exemplo de retorno:
```json
{
  "code": 200,
  "message": "operação realizada com sucesso"
}
  ```
</details>


<details>

<summary>/deposit/</summary>

  ## Método do tipo POST que realiza um depósito em carteira do usuário.

  ```bash
  /withdraw/
  ```
 ### O endpoint espera receber um json no body da requisição, o body possui a seguinte estrutura:
```json
{
  "TokenInvestor ": "string",
  "value":   50
}

```

- TokenUser: token jwt do usuário
- value: valor para realizar o saque, deve ser maior que 0

  ### Exemplo de retorno:
```json
{
  "code": 200,
  "message": "operação realizada com sucesso"
}
  ```
</details>

  
