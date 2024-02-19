package IO

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

// ReaderToReaderWriterSolution 需求: 依赖代理功能
type ReaderToReaderWriterSolution interface {
	TeeReader(r Reader, w Writer) Reader // io 包提供了 io.TeeReader 函数，该函数接受要读取的 Reader、要写入的 Writer，并返回可以进一步处理的 Reader。
}

// 官方 io.teeReader 将传递的Reader和Writer结合下来，将 Read 委托给传递的 Reader ，并对传递的 Writer 操作
// 即  原始Reader ---> TeeReader ---> 调用者
//        \
//         \---> Writer

/*
func TeeReader(r Reader, w Writer) Reader {
	return &teeReader{r, w}
}

type teeReader struct {
	r Reader
	w Writer
}

func (t *teeReader) Read(p []byte) (n int, err error) {
	n, err = t.r.Read(p)
	if n > 0 {
		if n, err := t.w.Write(p[:n]); err != nil {
			return n, err
		}
	}
	return
}
*/

// Example: 使用 io.TeeReader 来实现依赖代理功能。依赖代理将请求的上游图像缓存在对象存储中。
/*
// 主要功能是从一个上游服务获取数据，同时将这些数据转发给客户端和对象存储
func (p *Injector) Inject(w http.ResponseWriter, r *http.Request, sendData string) {
	// 通过HTTP获取上游数据
	dependencyResponse, err := p.fetchUrl(r.Context(), sendData)
	...

	// 创建一个tee阅读器。每次Read操作都会从dependencyResponse.Body读取数据，并同时向w writer写入数据
	teeReader := io.TeeReader(dependencyResponse.Body, w)			----> Reader

	// 将teeReader作为HTTP请求的body部分，以将其上传到对象存储
	saveFileRequest, err := http.NewRequestWithContext(r.Context(), "POST", r.URL.String()+"/upload", teeReader)
	...

	// 创建一个空的响应写入器，用于uploadHandler的调用，避免向客户端发送实际的HTTP响应
	nrw := &nullResponseWriter{header: make(http.Header)}

	// 使用创建的saveFileRequest向上传处理程序发送请求，实际的HTTP响应由nullResponseWriter接收，客户端不会接收到任何响应
	p.uploadHandler.ServeHTTP(nrw, saveFileRequest)
	...
}

*/

func TeeReaderExample() {
	// 创建一个字符串读取器，模拟从中读取数据
	sourceReader := strings.NewReader("Hello, TeeReader!")

	// 创建一个Buffer，用于作为io.TeeReader的第二个参数，即Writer
	var buf bytes.Buffer

	// 使用io.TeeReader包装sourceReader和buf。当从teeReader读取数据时，
	// 读取的数据也会写入到buf中。
	teeReader := io.TeeReader(sourceReader, &buf)

	// 读取数据到data中（通过teeReader），这样做会同时将数据写入buf
	data := make([]byte, sourceReader.Len())
	if _, err := teeReader.Read(data); err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

	// 打印从teeReader读取到的数据
	fmt.Printf("Read from teeReader: %s\n", data)

	// 打印被写入buf的数据，这应该和从teeReader读取的数据相同
	fmt.Printf("Written to buffer: %s\n", buf.String())
}
