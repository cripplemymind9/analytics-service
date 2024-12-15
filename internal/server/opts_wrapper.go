package server

type EntrypointOption interface {
	apply(*options)
}