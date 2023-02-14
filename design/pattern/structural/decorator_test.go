package structural

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	//基础文件源
	fileSource := &FileDataSource{Filename: "test.txt"}

	compression := NewCompressionDecorator(fileSource)

	encryption := NewEncryptionDecorator(compression)

	fmt.Println("*******write data***********")
	encryption.WriteData("data")

	fmt.Println("*******read data***********")
	fmt.Println(encryption.ReadData())
}
