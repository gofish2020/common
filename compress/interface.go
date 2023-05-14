package compress

type DataCompression interface {
	Compress(data []byte) ([]byte, error)
	UnCompress(data []byte) ([]byte, error)
}
