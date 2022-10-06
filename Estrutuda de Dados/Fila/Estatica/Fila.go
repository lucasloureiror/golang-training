package Fila

type Elemento struct{
	chave int;
	nome string;

}

type apontador *No;

type No struct{

	proximo apontador;
	item Elemento;
}

type Fila struct{
	inicio apontador;
	fim apontador;
}


func criar(novaFila *Fila){

	novaFila.inicio = nil;
	novaFila.fim = nil;

}

func enfileirar (fila *Fila, novoItem Elemento){

	novo *No;

	novo.item = novoItem;

	novo.proximo = nil;

	if (fila.inicio == nil){
		fila.inicio = novo;
	}else{
		fila.ultimo.proximo = novo;
	}

	fila.ultimo = novo;

}

func remover (fila *Fila){

}
