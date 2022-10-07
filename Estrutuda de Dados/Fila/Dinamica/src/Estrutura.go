package src

type Elemento struct {
	Chave int
	Nome  string
}

type apontador *Elemento

type Fila struct {
	Item   []Elemento
	Inicio apontador
	Final  apontador
}

func Criar(novaFila *Fila) {

	novaFila.Inicio = nil
	novaFila.Final = nil

}

func Enfileirar(fila *Fila, novoItem Elemento) {

	fila.Item = append(fila.Item, novoItem)

	if fila.Inicio == nil {
		fila.Inicio = &fila.Item[len(fila.Item)-1]
	} else {
		fila.Final = &fila.Item[len(fila.Item)-1]
	}

}

func Remover(fila *Fila) {

	var nullifier = Elemento{0, ""} //Aqui eu deixo nulo para o GC coletar na memoria.

	fila.Item[0] = nullifier

	fila.Item = fila.Item[1:]

	fila.Inicio = &fila.Item[0]

}

func PrimeiroDaFila(fila *Fila) *Elemento {

	return fila.Inicio

}

func TamanhoDaFila(fila *Fila) int {

	return len(fila.Item)
}

func ChecaNaoVazia(fila *Fila) bool {

	if fila.Inicio == nil {
		return false
	} else {
		return true
	}
}
