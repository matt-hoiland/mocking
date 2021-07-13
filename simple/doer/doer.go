package doer

type Doer interface {
	Do(s string, i int) error
}
