package IO

import (
	"io"
	"log"
	"os"
)

// WriterReaderSolution 需求:
//
//	接受 Writer 的函数，并且我们对该函数将写入 Writer 的内容感兴趣。我们希望拦截内容并将其表示为 Reader，以便以流方式进一步处理它。
//
// 通过 io.Pipe(管道) 方式
type WriterReaderSolution interface {
	Pipe() (*io.PipeReader, *io.PipeWriter)
}

// Writer 可用于传递给接受 Writer 的函数。已写入其中的所有内容都可以通过读取器访问，即创建一个同步内存管道，可用于连接需要 io.Reader 的代码与需要 io.Writer
// Writer ---> PipeWriter ----|
//                            | ---> 内存中的管道 ---> PipeReader ---> Reader
// Reader <-------------------|

// Example1:
/*
对于用于代码导航的 LSIF 文件转换，我们需要：
	1.读取 zip 文件的内容。
	2.转换内容并将其序列化为 zip.Writer 。
	3.将新的压缩内容表示为 Reader，以便以流方式进一步处理。
	zip.writer ---> PipeWriter ----> pipeReader ---> 压缩内容表示为reader
*/

func PipeExample() {
	// 创建管道
	pr, pw := io.Pipe()

	// 启动一个goroutine写入数据到PipeWriter
	go func() {
		defer pw.Close()
		_, err := pw.Write([]byte("Hello, Pipe!"))
		if err != nil {
			log.Fatal(err)
		}
	}()

	// 在主goroutine中从PipeReader读取数据
	if _, err := io.Copy(os.Stdout, pr); err != nil {
		log.Fatal(err)
	}
}
