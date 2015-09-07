package gocache

import (
	"sync"
	"time"
)

type DataObject struct {
	Data      string
	Timestamp time.Time
	Expiry    int64
	sync.Mutex
}

var inmem map[string]DataObject

func init() {
	inmem = make(map[string]DataObject)
}
func SetData(key, data string, exp ...int64) {
	temp := DataObject{}
	temp.Data = data
	temp.Timestamp = time.Now()
	if len(exp) > 0 {
		temp.Expiry = exp[0]
	}
	inmem[key] = temp
}

func GetData(key string) (string, bool) {
	retstr := ""
	retbool := false
	if v, ok := inmem[key]; ok {

		v.Lock()
		retstr = v.Data
		v.Unlock()

		retbool = true
		if inmem[key].Expiry > 0 && time.Now().Sub(inmem[key].Timestamp).Nanoseconds() >= inmem[key].Expiry*1000000 {

			v.Lock()
			delete(inmem, key)
			v.Unlock()

			retstr = ""
			retbool = false
		}

	}
	return retstr, retbool
}
