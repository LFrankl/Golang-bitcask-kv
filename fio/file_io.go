package fio

import "os"

// FileIO 标准系统文件IO
type FileIO struct {
	fd *os.File
}

// NewFileIOManager 初始化标准文件
func NewFileIOManager(fileName string) (*FileIO, error) {
	fd, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		DataFilePerm,
	)
	if err != nil {
		return nil, err
	}
	return &FileIO{fd: fd}, nil
}

// 从文件的给定位置读取对应的数据
func (fio *FileIO) Read([]byte, int64) (int, error) {

}

// 写入字节数组到文件中
func (fio *FileIO) Write([]byte) (int, error) {

}

// 持久化数据
func (fio *FileIO) Sync() error {

}

// 关闭文件
func (fio *FileIO) Close() error {

}
