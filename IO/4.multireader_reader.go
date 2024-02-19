package IO

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// MultiReaderSolution 简化从多个数据源顺序读取数据的过程
type MultiReaderSolution interface {
	MultiReader(readers ...Reader) Reader
}

// reader  ------
// reader  ______| -----> reader

// 通过 io.MultiReader, 将多个Reader转换为单Reader   MultiReader -----> Reader
/*
func NewReader(r io.Reader) (io.Reader, error) {
	magicBytes, err := readMagic(r)
	if err != nil {
		return nil, err
	}

	if string(magicBytes) != pngMagic {
		debug("Not a PNG - read file unchanged")
		return io.MultiReader(bytes.NewReader(magicBytes), r), nil
	}

	return io.MultiReader(bytes.NewReader(magicBytes), &Reader{underlying: r}), nil
}
*/

// MultiReaderExample 暂时测试 go 官方multi代码存在问题 后续看情况修改
func MultiReaderExample() {
	// 创建三个字符串读取器，模拟多个数据源
	reader1 := strings.NewReader("Hello")
	reader2 := strings.NewReader(", ")
	reader3 := strings.NewReader("MultiReader!")

	// 使用io.MultiReader将这三个读取器合并为一个
	multiReader := io.MultiReader(reader1, reader2, reader3)

	// 从multiReader中读取数据到一个足够大的缓冲区
	// 注意：在实际应用中，应根据数据大小适当选择缓冲区的大小
	buf := make([]byte, 1024)
	n, err := multiReader.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatalf("Failed to read: %v", err)
	}

	// 打印从multiReader读取到的数据
	fmt.Printf("Read from multiReader: %s", string(buf[:n]))
}
