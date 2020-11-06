# challenge 1

## Regra de negócio

 - O cliente cadastra um cartão de credito e valida se o numero é valido
 - O cliente cadastra um coupon e verifica se ele é valido
   - Verifica se o cupon é igual a `abc` ou igual a `abcd`
   - Se o coupon for igual a `abc` ele retorna como removed coupon, caso não ele aprova o desconto do coupon 

## Funcionamento
  - O cliente faz uma request para :9090 com intuito tanto de validar o cartão como o coupon, caso algum dos microserviços esteja fora ele possui uma politica de try até que o serviço volte novamente - 5 tentativas são feitas antes de retornar error