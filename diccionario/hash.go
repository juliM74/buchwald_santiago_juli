package diccionario

import (
	"fmt"
	"hash/fnv"
	TDAlista "tdas/lista"
)

const (
	CAPACIDAD_INICIAL      = 50
	FACTOR_CARGA_AUMENTO   = 2.0
	FACTOR_CARGA_REDUCCION = 0.5
	FACTOR_AGRANDAR        = 2
	FACTOR_ACHICAR         = 2
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla     []TDAlista.Lista[parClaveValor[K, V]]
	cantidad  int // cantidad real de elementos guardados
	capacidad int // cantidad de posiciones en la tabla
}

type iteradorHash[K comparable, V any] struct {
	iterLista   TDAlista.IteradorLista[parClaveValor[K, V]]
	diccionario *hashAbierto[K, V]
	posDic      int
}

func crearTabla[K comparable, V any](capacidad int) []TDAlista.Lista[parClaveValor[K, V]] {}

func CrearHash[K comparable, V any]() Diccionario[K, V] {}

func (d *hashAbierto[K, V]) redimensionar(nuevaCapacidad int) {}

func (d *hashAbierto[K, V]) buscarNodoClave(clave K) TDAlista.IteradorLista[parClaveValor[K, V]] {}

func (d *hashAbierto[K, V]) Guardar(clave K, dato V) {}

func (d *hashAbierto[K, V]) Pertenece(clave K) bool {}

func (d *hashAbierto[K, V]) Obtener(clave K) V {}

func (d *hashAbierto[K, V]) Borrar(clave K) V {}

func (d *hashAbierto[K, V]) Cantidad() int {
	return d.cantidad
}

func (d *hashAbierto[K, V]) Iterar(visitar func(K, V) bool) {}

func (d *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {}

func (it *iteradorHash[K, V]) avanzarHastaElemento() {}

func (it *iteradorHash[K, V]) HaySiguiente() bool {}

func (it *iteradorHash[K, V]) Siguiente() {}

func funcionHashing[K comparable](clave K, tam int) int {
	h := fnv.New64a()
	h.Write([]byte(fmt.Sprintf("%v", clave)))
	return int(h.Sum64() % uint64(tam))
}
