package IO

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

type MultiWriterSolution interface {
	MultiWriter(writers ...Writer) Writer // io 包也提供了io.MultiWriter
}

// writer ----> writer
//        ----> writer

func MultiWrite() {
	s1 := sha1.New()
	s256 := sha256.New()

	// 将写入数据, 并同时分发
	w := io.MultiWriter(s1, s256)
	if _, err := w.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}

	fmt.Println(s1.Sum(nil))   // content"的SHA-1哈希值。
	fmt.Println(s256.Sum(nil)) // content"的SHA-256哈希值。

	// 当内容来自 reader时, 也可以通过 io.copy 方式   reader ----> MultiWriter
	file, _ := os.Open("data.txt")
	_, _ = io.Copy(io.MultiWriter(s1, s256), file)
}
