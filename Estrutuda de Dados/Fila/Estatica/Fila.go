package Estrutura

type Elemento struct{
	chave int;
	nome string;

}

type apontador *Elemento;

type Fila struct{
	item []Elemento;
	inicio apontador;
	final apontador;
}


func criar(novaFila *Fila){

	novaFila.inicio = nil;
	novaFila.final = nil;

}

func enfileirar (fila *Fila, novoItem Elemento){

	fila.item = append(fila.item, novoItem);

	if(fila.inicio == nil){
		fila.inicio = &fila.item[len(fila.item) - 1];
	} else{
		fila.final = &fila.item[len(fila.item) -1];
	}

}

func remover (fila *Fila){

	var nullifier = Elemento{0, ""}; //Aqui eu deixo nulo para o GC coletar na memoria.

	fila.item[0] = nullifier;

	fila.item = fila.item[1: ];

	fila.inicio = &fila.item[0];

}

func primeiroDaFila(fila *Fila) *Elemento{

	return fila.inicio; 

}

func tamanhoDaFila(fila *Fila) int{

	return len(fila.item);
}
