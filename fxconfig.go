package fxconfig

import "schneider.vip/config"

// New returns an constructor of a Dynamic Config and a parsed config of T.
func New[T any](opts ...config.Option[T]) func() (config.Dynamic[T], T) {
	return func() (config.Dynamic[T], T) {
		return config.NewDynamic(opts...)
	}
}
