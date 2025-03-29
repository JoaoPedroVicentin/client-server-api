# Desafio Go Expert - Client Server Api

## Server

- [x] Consumir a API contendo o câmbio de Dólar e Real;
- [x] Retornar no formato JSON o resultado para o cliente;
- [x] Registrar no banco de dados SQLite cada cotação recebida;
- [x] O timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms;
- [x] O timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms;
- [x] Enviar para o Client apenas o valor atual do câmbio (campo "bid" do JSON);
- [x] O endpoint deverá ser: /cotacao;
- [x] A porta a ser utilizada pelo servidor HTTP será a 8080;

## Client

- [x] Realizar uma requisição HTTP no Server solicitando a cotação do dólar;
- [x] O timeout máximo para receber o resultado do Server deverá ser de 300ms;
- [x] Salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

## Server and Client

- [] Os 3 contextos devem retornar erro nos logs caso o tempo de execução seja insuficiente;