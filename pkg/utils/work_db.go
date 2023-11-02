package utils

/*
此处的功能主要是做work链接的存储
原始数据的结构可能随时会发生变化 所以需要增加索引模块
数据增加、删除、检索都要对结构进行处理 所以不能使用读写锁
*/
import (
	"sync"

	v1 "github.com/Uonx/gather/internal/proto/v1"
	"github.com/Uonx/gather/pkg/model"
)

type WorkData struct {
	Datas  map[int]Data
	mutex  sync.Mutex //根据业务要求 每次操作都有写的操作 所以用互斥锁
	Id     int
	indexs map[string]RoundRobinBalance
}

type RoundRobinBalance struct {
	curIndex int
	index    []int
}

type Data struct {
	id        int //这个id是内部控制，不对外
	Name      string
	Ip        string
	Templates []string
	Stream    *v1.SchedulerEvent_RegistryServer
	status    int //是否在使用,内部控制，不对外
}

var (
	WorkDB  *WorkData
	workOne sync.Once
)

// 初始化workdb 并赋予初始Id从几开始
func NewWorkDb(id int) {
	workOne.Do(func() {
		WorkDB = &WorkData{
			Id:     id,
			indexs: make(map[string]RoundRobinBalance),
			Datas:  make(map[int]Data),
		}
	})
}

// 添加数据流链接
func (w *WorkData) AddData(data *Data) {
	//写入时加写锁
	w.mutex.Lock()
	defer w.mutex.Unlock()
	data.id = w.Id
	data.status = model.NotUse
	w.Datas[data.id] = *data
	for _, d := range data.Templates {
		index := w.indexs[d]
		index.index = append(index.index, data.id)
		w.indexs[d] = index
	}
	w.Id++
}

// 删除数据流链接
func (w *WorkData) DeleteData(data Data) {
	//删除时加写锁
	w.mutex.Lock()
	defer w.mutex.Unlock()

	//先删除索引再删除数据 原因在于查找的时候是用的索引去定位对应的数据
	for _, t := range data.Templates {
		if inx, ok := w.indexs[t]; ok {
			tmp := make([]int, 0, len(inx.index))
			for _, i := range inx.index {
				if i != data.id {
					tmp = append(tmp, i)
				}
			}
			inx.index = tmp
			w.indexs[t] = inx
		}
	}

	delete(w.Datas, data.id)
}

// 暂借
func (w *WorkData) Borrow(template string) (Data, bool) {
	//借stream的时候需要修改stream的状态 所以要加写锁
	w.mutex.Lock()
	defer w.mutex.Unlock()
	var ids RoundRobinBalance
	var data Data
	var ok bool
	ids, ok = w.indexs[template]
	if !ok {
		return data, ok
	}
	len := len(ids.index)
	if ids.curIndex >= len {
		ids.curIndex = 0
	}
	i := ids.index[ids.curIndex]
	if w.Datas[i].status == model.NotUse {
		data = w.Datas[i]
		data.status = model.InUse
		w.Datas[i] = data
		ok = true
	} else {
		ok = false
	}
	ids.curIndex = (ids.curIndex + 1) % len
	w.indexs[template] = ids
	return data, ok
}

// 归还
func (w *WorkData) Return(data Data) {
	// 还data的时候需要修改stream的状态 所以要加写锁 此处还只需要做实际数据的状态修改就行
	w.mutex.Lock()
	defer w.mutex.Unlock()
	d := w.Datas[data.id]
	d.status = model.NotUse
	w.Datas[data.id] = d
}
