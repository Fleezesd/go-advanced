package IO

import (
	"errors"
	"io"
	"log"
	"os"
)

// reference: https://about.gitlab.com/blog/2024/02/15/compose-readers-and-writers-in-golang-applications/
// 核心: 内容从 Reader 传递到 Writer (可能为多个)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// BlockWrite 文件以块的形式读入到os.stdout 最高32KB
func BlockWrite() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p := make([]byte, 32*1024) // 文件以块的形式 最高32KB
	// 读取文件并将其内容写入 os.Stdout
	for {
		n, err := file.Read(p)

		_, errW := os.Stdout.Write(p[:n])
		if errW != nil {
			log.Fatal(errW)
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatal(err)
		}
	}
}

// CopyWrite 以copy方式读取数据 上述方案的简化方式
func CopyWrite() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
	}
}
