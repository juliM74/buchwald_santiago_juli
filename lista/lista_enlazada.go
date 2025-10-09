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

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {}

func (l *listaEnlazada[T]) BorrarPrimero() T {}

func (l *listaEnlazada[T]) VerPrimero() T {
}

func (l *listaEnlazada[T]) VerUltimo() T {}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {}

// Metodos del iterador externo

func (i *iterListaEnlazada[T]) VerActual() T {}

func (i *iterListaEnlazada[T]) HaySiguiente() bool {}

func (i *iterListaEnlazada[T]) Siguiente() {}

func (i *iterListaEnlazada[T]) Insertar(nuevoElemento T) {}

func (i *iterListaEnlazada[T]) Borrar() T {}

