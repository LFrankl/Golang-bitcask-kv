package data

type DataFile struct {
	FileId    uint32 //文件ID
	WriteOff  int64  //写到了什么位置
	IoManager fio.Manager
}

// OpenDataFile 打开新的数据文件
func OpenDataFile(dirPath string, fileId uint32) (*Datafile, error) {
	return nil, nil
}

func (df *DataFile) Sync() error {
	return nil
}
