package analyzer

type Analyzer interface {

}

type analyzer struct{}



func New() Analyzer{
	return &analyzer{}
}