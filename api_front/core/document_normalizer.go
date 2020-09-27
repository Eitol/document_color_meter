package core

type DocumentNormalizer interface {
	Normalize(doc []byte) ([]byte, error)
}
