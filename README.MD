# Jump-Start

## Como executar o projeto localmente

- Verificar se o docker está instalado
``` bash
        docker --version
```
- Caso não tenha, um link de ajuda para instalação

    [Tutorial instalação docker](https://www.youtube.com/watch?v=wpdcGgRY5kk)

- Parar containers que possam gerar conflito de portas

``` bash
        docker compose down -v
```
- Subir container que possui a aplicação e banco de dados

``` shell
        docker compose up --build -d 
```
- Se tudo ocorrer com sucesso, basta acessar `http://localhost:8080` que a aplicação irá estar disponível para acesso aos endpoints



