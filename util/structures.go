package util

type ComponentHandle func(ctx ComponentCtx) error

type ComponentCtx struct {
}
type ModalHandle func(ctx ModalCtx) error

type ModalCtx struct {
}
type AutocompleteHandle func(ctx AutocompleteCtx) error

type AutocompleteCtx struct {
}

type JComponent struct{}

type JModal struct{}

type JAutocomplete struct{}
