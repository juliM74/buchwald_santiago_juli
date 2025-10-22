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

func crearTabla[K comparable, V any](capacidad int) []TDAlista.Lista[parClaveValor[K, V]] {
	tabla := make([]TDAlista.Lista[parClaveValor[K, V]], capacidad)
	for i := range tabla {
		tabla[i] = TDAlista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return tabla
}

func CrearHash[K any, V any](cmp func(K, K) bool) Diccionario[K, V]
	return &hashAbierto[K, V]{
	tabla: crearTabla[K, V](CAPACIDAD_INICIAL),
	cantidad: 0,
	capacidad: CAPACIDAD_INICIAL,
	}
}

func (d *hashAbierto[K, V]) redimensionar(nuevaCapacidad int) {
	nuevaTabla := crearTabla[K, V](nuevaCapacidad)

	for i := range nuevaTabla {
		nuevaTabla[i] = TDAlista.CrearListaEnlazada[parClaveValor[K, V]]()
	}

	for _, lista := range d.tabla {
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			campo := iter.VerActual()
			indice := funcionHashing(campo.clave, nuevaCapacidad)
			nuevaTabla[indice].InsertarUltimo(campo)
			iter.Siguiente()
		}
	}

	d.tabla = nuevaTabla
	d.capacidad = nuevaCapacidad
}

func (d *hashAbierto[K, V]) buscarNodoClave(clave K) TDAlista.IteradorLista[parClaveValor[K, V]] {
	indice := funcionHashing(clave, d.capacidad)
	lista := d.tabla[indice]
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		campo := iter.VerActual()
		if campo.clave == clave {
			return iter
		}
		iter.Siguiente()
	}
	return nil
}

func (d *hashAbierto[K, V]) Guardar(clave K, dato V) {
	factorCarga := float64(d.cantidad+1) / float64(d.capacidad)
	if factorCarga > FACTOR_CARGA_AUMENTO {
		d.redimensionar(d.capacidad * FACTOR_AGRANDAR)
	}
	iter := d.buscarNodoClave(clave)
	indice := funcionHashing(clave, d.capacidad)

	if iter == nil {
		campoNuevo := parClaveValor[K, V]{clave, dato}
		d.tabla[indice].InsertarUltimo(campoNuevo)
		d.cantidad++
		return
	}
	campo := iter.Borrar()
	campo.dato = dato
	iter.Insertar(campo)
}

func (d *hashAbierto[K, V]) Pertenece(clave K) bool {
	return d.buscarNodoClave(clave) != nil
}

func (d *hashAbierto[K, V]) Obtener(clave K) V {
	iter := d.buscarNodoClave(clave)
	if iter == nil {
		panic("La clave no pertenece al diccionario")
	}
	campo := iter.VerActual()
	return campo.dato
}

func (d *hashAbierto[K, V]) Borrar(clave K) V {
	indice := funcionHashing(clave, d.capacidad)
	lista := d.tabla[indice]
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		campo := iter.VerActual()
		if campo.clave == clave {
			iter.Borrar()
			d.cantidad--
			factorCarga := float64(d.cantidad) / float64(d.capacidad)
		if factorCarga < FACTOR_CARGA_REDUCCION && d.capacidad > CAPACIDAD_INICIAL {
			d.redimensionar(d.capacidad / FACTOR_ACHICAR)
	}
	return campo.dato
	}
	iter.Siguiente()
	}
	panic("La clave no pertenece al diccionario")
}

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
