package Golang_bitcask_kv

import (
	"go-bitcask/data"
	"sync"
)

// DB 存储引擎实例
type DB struct {
	options    Options
	mu         sync.RWMutex
	activeFile *data.DataFile            //当前活跃数据文件，可用于写入
	olderFiles map[uint32]*data.DataFile //旧的数据文件，只能用于读
}

func (db *DB) Put(key []byte, value []byte) error {
	if len(key) == 0 {
		//判断key是否有效
		return ErrKeyIsEmpty
	}
	//构造LogRecord结构体
	log_record := data.LogRecord{
		Key:   key,
		Value: value,
		Type:  data.LogRecordNormal,
	}
}

func (db *DB) appendLogRecord(logRecord *data.LogRecord) (*data.LogRecordPos, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	//判断当前活跃数据文件是否存在，db在没有写入的时候是没有文件生成的
	//如果为空则初始化数据文件
	if db.activeFile != nil {
		if err := db.setActiveFile(); err != nil {
			return nil, err
		}
	}
	//写入数据编码
	encRecord, size := data.EncodeLogRecord(logRecord)
	//如果写入的数据已经达到了活跃文件的阈值，则关闭活跃文件，并打开新的文件
	if db.activeFile.WriteOff+size > db.options.DataFileSize {
		//先持久化数据文件夹，保证已有的数据持久化
		if err := db.activeFile.Sync(); err != nil {
			return nil, err
		}
		//当前活跃文件在转会旧的数据文件
		db.olderFiles[db.activeFile.FileId] = db.activeFile

		//打开新的数据文件夹
		if err := db.setActiveFile(); err != nil {
			return nil, err
		}
	}

	writeOff :=
}

// 设置当前活跃文件
// setActiveFile 在访问此方法前把必须加锁
func (db *DB) setActiveFile() error {
	var initialFileId uint32 = 0
	if db.activeFile != nil {
		initialFileId = db.activeFile.FileId + 1
	}
	//	打开新的数据文件

	dataFile, err := data.OpenDataFile(db.options.DirPath, initialFileId)
	if err != nil {
		return err
	}
	db.activeFile = dataFile
	return nil
}
