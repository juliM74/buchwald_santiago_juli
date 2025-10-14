package lista

// Estructura de nodo (privada)
type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

// Estructura de la lista enlazada
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

// Estructura de iterador
type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

// Metodos de listaEnlazada

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func crearNodoLista[T any](nuevoDato T, apuntado *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: nuevoDato, siguiente: apuntado}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil && l.ultimo == nil && l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(val T) {
	n := &nodo[T]{dato: val}
	if l.EstaVacia() {
		l.primero = n
		l.ultimo = n
	} else {
		n.prox = l.primero
		l.primero = n
	}
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(val T) {
	n := &nodo[T]{dato: val}
	if l.EstaVacia() {
		l.primero = n
		l.ultimo = n
	} else {
		l.ultimo.prox = n
		l.ultimo = n
	}
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := l.primero.dato
	l.primero = l.primero.siguiente
	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) VerPrimero() T {
}

func (l *listaEnlazada[T]) VerUltimo() T {}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	puntero := l.primero
	for puntero != nil {
		if !visitar(puntero.dato) {
			return
		}
		puntero = puntero.siguiente
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{
		actual:   l.primero,
		anterior: nil,
		lista:    l,
	}
}

// Metodos del iterador externo

func (i *iterListaEnlazada[T]) VerActual() T {}

func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iterListaEnlazada[T]) Siguiente() {}

func (i *iterListaEnlazada[T]) Insertar(nuevoElemento T) {}

func (i *iterListaEnlazada[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	borrado := i.actual.dato

	if i.anterior == nil {
		i.lista.primero = i.actual.siguiente
	} else {
		i.anterior.siguiente = i.actual.siguiente
	}
	if i.actual.siguiente == nil {
		i.lista.ultimo = i.anterior
	}

	i.actual = i.actual.siguiente
	i.lista.largo--

	return borrado
}

