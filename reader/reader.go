package reader

type Reader interface {

}

type reader struct{}

func New() Reader {
	return &reader{}
}