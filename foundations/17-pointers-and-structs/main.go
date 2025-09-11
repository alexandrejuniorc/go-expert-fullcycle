package main

import "fmt"

/*
===============================================
PONTEIROS E STRUCTS EM GO
===============================================
Este exemplo demonstra a diferença fundamental entre:
- Passagem por valor (cria uma cópia)
- Passagem por ponteiro (referencia o objeto original)
*/

// =============================================
// EXEMPLO 1: STRUCT SEM PONTEIRO (Por Valor)
// =============================================
// Quando um método recebe uma struct diretamente,
// Go faz uma CÓPIA da struct original
type Cliente struct {
	nome string
}

// Método que recebe Cliente por VALOR (sem *)
// Qualquer mudança aqui NÃO afeta a struct original
func (c Cliente) andou() {
	c.nome = "Outro nome" // Esta mudança só existe na CÓPIA
	fmt.Printf("Dentro do método - cliente %v andou\n", c.nome)
}

// =============================================
// EXEMPLO 2: STRUCT COM PONTEIRO (Por Referência)
// =============================================
// Para modificar a struct original, usamos PONTEIRO (*)
type Conta struct {
	saldo int
}

/*
MÉTODO COM PONTEIRO: func (c *Conta)
- O asterisco (*) indica que recebemos um PONTEIRO
- Isso significa que temos acesso à struct ORIGINAL
- Mudanças aqui AFETAM diretamente o objeto original
- Sintaxe: (c *Conta) = "c é um ponteiro para Conta"
*/
func (c *Conta) simular(valor int) int {
	c.saldo += valor // Modifica a struct ORIGINAL
	fmt.Printf("Novo saldo: %d\n", c.saldo)
	return c.saldo
}

// =============================================
// EXEMPLO 3: FUNÇÃO CONSTRUTORA COM PONTEIRO
// =============================================
/*
FUNÇÃO CONSTRUTORA: NewConta() *Conta
- Retorna um PONTEIRO para uma nova instância de Conta
- O símbolo & cria um ponteiro para a struct
- Quando você usa esta função, está trabalhando com o ORIGINAL
- Não há cópia sendo feita, apenas uma referência

Por que usar?
✅ Economia de memória (não cria cópias)
✅ Modificações afetam diretamente o objeto
✅ Padrão comum em Go para "construtores"
*/
func NewConta() *Conta {
	return &Conta{saldo: 0} // & = "endereço de" (cria ponteiro)
}

func main() {
	fmt.Println("=== DEMONSTRAÇÃO DE PONTEIROS E STRUCTS ===")
	fmt.Println()

	// -----------------------------------------------
	// TESTE 1: Struct por VALOR (sem ponteiro)
	// -----------------------------------------------
	fmt.Println("📝 EXEMPLO 1 - Passagem por VALOR:")
	alexandre := Cliente{nome: "Alexandre"}
	fmt.Printf("Nome original: %s\n", alexandre.nome)

	alexandre.andou() // Chama método que recebe cópia

	// ⚠️  IMPORTANTE: O nome original NÃO foi alterado!
	fmt.Printf("Nome após método: %s\n", alexandre.nome)
	fmt.Println("👆 Viu? O nome original não mudou!")
	fmt.Println()

	// -----------------------------------------------
	// TESTE 2: Struct por PONTEIRO (com *)
	// -----------------------------------------------
	fmt.Println("📝 EXEMPLO 2 - Passagem por PONTEIRO:")
	conta := Conta{saldo: 1000}
	fmt.Printf("Saldo inicial: %d\n", conta.saldo)

	novoSaldo := conta.simular(500) // Chama método com ponteiro

	// ✅ SUCESSO: O saldo original FOI alterado!
	fmt.Printf("Saldo retornado: %d\n", novoSaldo)
	fmt.Printf("Saldo da conta original: %d\n", conta.saldo)
	fmt.Println("👆 O saldo original foi modificado!")

	// -----------------------------------------------
	// TESTE 3: FUNÇÃO CONSTRUTORA COM PONTEIRO
	// -----------------------------------------------
	fmt.Println()
	fmt.Println("📝 EXEMPLO 3 - Função Construtora com PONTEIRO:")

	// Usando a função construtora NewConta()
	// Esta função retorna um PONTEIRO (*Conta), não uma cópia
	contaNova := NewConta() // contaNova é do tipo *Conta (ponteiro)
	fmt.Printf("Tipo da variável: %T\n", contaNova)
	fmt.Printf("Saldo inicial da conta nova: %d\n", contaNova.saldo)

	// Como contaNova já é um ponteiro, podemos usar métodos com ponteiro
	contaNova.simular(200) // Modifica diretamente o objeto original
	fmt.Printf("Saldo após simular: %d\n", contaNova.saldo)

	// Comparação: criar sem função construtora
	contaSemPonteiro := Conta{saldo: 100} // Cria uma struct normal
	fmt.Printf("Tipo sem ponteiro: %T\n", contaSemPonteiro)

	// Para usar métodos com ponteiro em struct normal, Go converte automaticamente
	contaSemPonteiro.simular(300) // Go faz: (&contaSemPonteiro).simular(300)
	fmt.Printf("Saldo da conta sem ponteiro: %d\n", contaSemPonteiro.saldo)

	// -----------------------------------------------
	// RESUMO PARA ESTUDOS
	// -----------------------------------------------
	fmt.Println()
	fmt.Println("🎯 RESUMO COMPLETO:")
	fmt.Println("• Método SEM *: trabalha com CÓPIA (não altera original)")
	fmt.Println("• Método COM *: trabalha com PONTEIRO (altera original)")
	fmt.Println("• Use ponteiros quando quiser MODIFICAR a struct original")
	fmt.Println("• Função construtora com *: retorna ponteiro (mais eficiente)")
	fmt.Println("• Símbolo &: cria ponteiro para struct (&Conta{...})")
	fmt.Println("• Go converte automaticamente struct para ponteiro quando necessário")
}
