package data

type LogRecordPos struct {
	Fid    uint32
	Offset int64
	//文件id和偏移,数据存到文件的名字，数据在文件中的位置
	//首字母大写表示可被外部访问
}
