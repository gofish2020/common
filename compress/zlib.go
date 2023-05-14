package compress

import (
	"bytes"
	"compress/zlib"
	"io"
)

type Zlib struct {
}

func GetZlibCompresser() Zlib {
	return Zlib{}
}

func (t Zlib) Compress(data []byte) ([]byte, error) {
	var in bytes.Buffer
	w, err := zlib.NewWriterLevel(&in, zlib.DefaultCompression)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return in.Bytes(), nil
}

/*

r, err := zlib.NewReader(&in)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	io.Copy(&out, r)
	r.Close()

	fmt.Println(out.String())
*/

func (t Zlib) UnCompress(data []byte) ([]byte, error) {
	var (
		in  bytes.Buffer
		out bytes.Buffer
	)

	_, err := in.Write(data)
	if err != nil {
		return nil, err
	}

	r, err := zlib.NewReader(&in)
	if err != nil {
		return nil, err
	}
	io.Copy(&out, r)

	err = r.Close()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func NewCompressZlib() Zlib {
	return Zlib{}
}
