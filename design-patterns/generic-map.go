package designpatterns

// ~ imples all aliases of int
type SupportedGenrics interface {
	~int | string
}
type GenericMap[K SupportedGenrics, V SupportedGenrics] map[K]V

func NewGenericMap[K SupportedGenrics, V SupportedGenrics](size int) GenericMap[K, V] {
	return make(GenericMap[K, V], size)
}
