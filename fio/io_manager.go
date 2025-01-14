package fio

const DataFilePerm = 0644

// IOManager  抽象的IO管理器的接口
type IOManager interface {
	Read([]byte, int64) (int, error) //从文件的给定位置读取对应的数据
	Write([]byte) (int, error)       //写入字节数组到文件中
	Sync() error                     //持久化数据
	Close() error                    //关闭文件
}
