package bimap

import (
	"fmt"

	"github.com/SallyKinoshita/go-rest-api/backend/pkg/cerror"
)

type (
	bimap[K, V comparable] struct {
		forwardMap map[K]V
		reverseMap map[V]K
	}
)

func New[K, V comparable](fmap map[K]V) (bimap[K, V], error) {
	rmap, err := reverse(fmap)
	if err != nil {
		return bimap[K, V]{}, cerror.Wrap(err, cerror.DetailInternalServerError)
	}
	return bimap[K, V]{
		forwardMap: fmap,
		reverseMap: rmap,
	}, nil
}

func NewMust[K, V comparable](forwardMap map[K]V) bimap[K, V] {
	bimap, err := New(forwardMap)
	if err != nil {
		panic(cerror.Wrap(err, cerror.DetailInternalServerError))
	}
	return bimap
}

func reverse[K, V comparable](forwardMap map[K]V) (map[V]K, error) {
	m := make(map[V]K, len(forwardMap))
	for k, v := range forwardMap {
		if existKey, ok := m[v]; ok {
			// valueの重複があった場合はエラーを返す
			return nil, cerror.New(fmt.Sprintf("failed to duplicate value: %v (New key = %v, Current key = %v)", v, k, existKey), cerror.DetailInternalServerError)
		} else {
			m[v] = k
		}
	}
	return m, nil
}

func (b *bimap[K, V]) GetForwardMap() map[K]V {
	return b.forwardMap
}

func (b *bimap[K, V]) GetReverseMap() map[V]K {
	return b.reverseMap
}
