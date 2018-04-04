package term

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"sync"
)

type ImageWriter struct {
	once   sync.Once
	b64enc io.WriteCloser
	buf    *bytes.Buffer
}

func (w *ImageWriter) init() {
	w.buf = &bytes.Buffer{}
	w.b64enc = base64.NewEncoder(base64.StdEncoding, w.buf)
}

func (w *ImageWriter) Write(p []byte) (n int, err error) {
	w.once.Do(w.init)

	// base64でエンコード
	return w.b64enc.Write(p)
}

func (w *ImageWriter) Close() error {
	w.once.Do(w.init)
	// w.buf.Bytes() でバッファの中身を取り出す.
	fmt.Printf(ecsi+"]1337;File=preserveAspectRatio=1;inline=1:%s%s", w.buf.Bytes(), st)
	return w.b64enc.Close()
}
