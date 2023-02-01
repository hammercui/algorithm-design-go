package main

import "fmt"

/**
结构型设计模式：
装饰器模式，一种proxy模式的特殊实现，类似于责任链模式的链式模型。
但是component和decorator均实现相同的interface
以下以文件读取为例
 */

//公共接口
type DataSource interface {
	WriteData(data string)
	ReadData() string
}

//基础component
type FileDataSource struct {
	Filename string
}
func (p *FileDataSource) WriteData(data string) {
	fmt.Printf("FileDataSource WriteData %s \n",data)
}
func (p *FileDataSource) ReadData() string {
	fmt.Printf("FileDataSource ReadData \n")
	return "data"
}

//父decorator
type DataSourceDecorator struct {
	wrapper DataSource
}
//func (p *DataSourceDecorator) SetWrapper(wrapper DataSource) {
//	p.wrapper = wrapper
//}
func (p *DataSourceDecorator) WriteData(data string) {
	//fmt.Printf("Base Decorator WriteData %s \n",data)
	p.wrapper.WriteData(data)
}
func (p *DataSourceDecorator) ReadData() string {
	//fmt.Printf("Base Decorator ReadData \n")
	return p.wrapper.ReadData()
}

//加解密装饰器
type EncryptionDecorator struct {
	DataSourceDecorator
}
func NewEncryptionDecorator(source DataSource) *EncryptionDecorator{
	p :=  &EncryptionDecorator{}
	p.wrapper = source
	return p
}
func (p *EncryptionDecorator) WriteData(data string) {
	fmt.Printf("EncryptionDecorator WriteData 加密 %s \n",data)
	p.DataSourceDecorator.WriteData(data)
}
func (p *EncryptionDecorator) ReadData() string {
	data := p.DataSourceDecorator.ReadData()
	fmt.Printf("EncryptionDecorator ReadData 解密\n")
	return data
}

//压缩装饰器
type CompressionDecorator struct {
	DataSourceDecorator
}
func NewCompressionDecorator(source DataSource) *CompressionDecorator{
	p :=  &CompressionDecorator{}
	p.wrapper = source
	return p
}
func (p *CompressionDecorator) WriteData(data string) {
	fmt.Printf("EncryptionDecorator WriteData 压缩 %s \n",data)
	p.DataSourceDecorator.WriteData(data)
}
func (p *CompressionDecorator) ReadData() string {
	data := p.DataSourceDecorator.ReadData()
	fmt.Printf("EncryptionDecorator ReadData 解压缩\n")
	return data
}








