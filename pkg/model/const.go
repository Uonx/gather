package model

const (
	Unprocessed = 1 << iota //未处理
	Processing              //处理中
	Completed               //已完成
	Failed                  //失败
)

const (
	NotUse = 1 << iota //未使用
	InUse
)
