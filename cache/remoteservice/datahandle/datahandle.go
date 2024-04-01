package datahandle

// DataHandle 返回的数据处理
type DataHandle[T, K any] interface {
	HandleData(data T) K
}
