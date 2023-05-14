package compress

type Compress interface {
	Marshal(data []byte) ([]byte, error)
	Unmarshal(data []byte) ([]byte, error)
}
