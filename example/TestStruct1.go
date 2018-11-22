/// -------------------------------------------------------------------------------
/// THIS FILE IS ORIGINALLY GENERATED BY redis2go.exe.
/// PLEASE DO NOT MODIFY THIS FILE.
/// -------------------------------------------------------------------------------

package main

import (
	"errors"
	"fmt"
	cstruct "github.com/fananchong/cstruct-go"
	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/gomodule/redigo/redis"
)

type TestStruct1 struct {
	Key   uint64
	myb   bool
	myf1  float32
	myf2  float64
	myi1  int8
	myi2  int16
	myi3  int32
	myi4  int64
	myi6  uint8
	myi7  uint16
	myi8  uint32
	myi9  uint64
	mys1  string
	mys2  []byte
	myst1 StrcutXX
	myst2 StrcutYY

	__dirtyData               map[string]interface{}
	__dirtyDataForStructFiled map[string]interface{}
	__isLoad                  bool
	__dbKey                   string
	__dbName                  string
	__expire                  uint
}

func NewTestStruct1(dbName string, key uint64) *TestStruct1 {
	return &TestStruct1{
		Key:                       key,
		__dbName:                  dbName,
		__dbKey:                   "TestStruct1:" + fmt.Sprintf("%d", key),
		__dirtyData:               make(map[string]interface{}),
		__dirtyDataForStructFiled: make(map[string]interface{}),
	}
}

// 若访问数据库失败返回-1；若 key 存在返回 1 ，否则返回 0 。
func (this *TestStruct1) HasKey() (int, error) {
	db := go_redis_orm.GetDB(this.__dbName)
	val, err := redis.Int(db.Do("EXISTS", this.__dbKey))
	if err != nil {
		return -1, err
	}
	return val, nil
}

func (this *TestStruct1) Load() error {
	if this.__isLoad == true {
		return errors.New("alreay load!")
	}
	db := go_redis_orm.GetDB(this.__dbName)
	val, err := redis.Values(db.Do("HGETALL", this.__dbKey))
	if err != nil {
		return err
	}
	if len(val) == 0 {
		return go_redis_orm.ERR_ISNOT_EXIST_KEY
	}
	var data struct {
		Myb   bool    `redis:"myb"`
		Myf1  float32 `redis:"myf1"`
		Myf2  float64 `redis:"myf2"`
		Myi1  int8    `redis:"myi1"`
		Myi2  int16   `redis:"myi2"`
		Myi3  int32   `redis:"myi3"`
		Myi4  int64   `redis:"myi4"`
		Myi6  uint8   `redis:"myi6"`
		Myi7  uint16  `redis:"myi7"`
		Myi8  uint32  `redis:"myi8"`
		Myi9  uint64  `redis:"myi9"`
		Mys1  string  `redis:"mys1"`
		Mys2  []byte  `redis:"mys2"`
		Myst1 []byte  `redis:"myst1"`
		Myst2 []byte  `redis:"myst2"`
	}
	if err := redis.ScanStruct(val, &data); err != nil {
		return err
	}
	this.myb = data.Myb
	this.myf1 = data.Myf1
	this.myf2 = data.Myf2
	this.myi1 = data.Myi1
	this.myi2 = data.Myi2
	this.myi3 = data.Myi3
	this.myi4 = data.Myi4
	this.myi6 = data.Myi6
	this.myi7 = data.Myi7
	this.myi8 = data.Myi8
	this.myi9 = data.Myi9
	this.mys1 = data.Mys1
	this.mys2 = data.Mys2
	if err := cstruct.Unmarshal(data.Myst1, &this.myst1); err != nil {
		return err
	}

	if err := cstruct.Unmarshal(data.Myst2, &this.myst2); err != nil {
		return err
	}

	this.__isLoad = true
	return nil
}

func (this *TestStruct1) Save() error {
	if len(this.__dirtyData) == 0 && len(this.__dirtyDataForStructFiled) == 0 {
		return nil
	}
	for k, _ := range this.__dirtyDataForStructFiled {
		_ = k
		if k == "myst1" {
			data, err := cstruct.Marshal(&this.myst1)
			if err != nil {
				return err
			}
			this.__dirtyData["myst1"] = data
		}
		if k == "myst2" {
			data, err := cstruct.Marshal(&this.myst2)
			if err != nil {
				return err
			}
			this.__dirtyData["myst2"] = data
		}
	}
	db := go_redis_orm.GetDB(this.__dbName)
	if _, err := db.Do("HMSET", redis.Args{}.Add(this.__dbKey).AddFlat(this.__dirtyData)...); err != nil {
		return err
	}
	if this.__expire != 0 {
		if _, err := db.Do("EXPIRE", this.__dbKey, this.__expire); err != nil {
			return err
		}
	}
	this.__dirtyData = make(map[string]interface{})
	this.__dirtyDataForStructFiled = make(map[string]interface{})
	return nil
}

func (this *TestStruct1) Delete() error {
	db := go_redis_orm.GetDB(this.__dbName)
	_, err := db.Do("DEL", this.__dbKey)
	if err == nil {
		this.__isLoad = false
		this.__dirtyData = make(map[string]interface{})
		this.__dirtyDataForStructFiled = make(map[string]interface{})
	}
	return err
}

func (this *TestStruct1) IsLoad() bool {
	return this.__isLoad
}

func (this *TestStruct1) Expire(v uint) {
	this.__expire = v
}

func (this *TestStruct1) DirtyData() (map[string]interface{}, error) {
	data := make(map[string]interface{})
	for k, v := range this.__dirtyData {
		data[k] = v
	}
	for k, _ := range this.__dirtyDataForStructFiled {
		_ = k
		if k == "myst1" {
			data, err := cstruct.Marshal(&this.myst1)
			if err != nil {
				return nil, err
			}
			this.__dirtyData["myst1"] = data
		}
		if k == "myst2" {
			data, err := cstruct.Marshal(&this.myst2)
			if err != nil {
				return nil, err
			}
			this.__dirtyData["myst2"] = data
		}
	}
	return data, nil
}

func (this *TestStruct1) Save2(dirtyData map[string]interface{}) error {
	if len(dirtyData) == 0 {
		return nil
	}
	db := go_redis_orm.GetDB(this.__dbName)
	if _, err := db.Do("HMSET", redis.Args{}.Add(this.__dbKey).AddFlat(dirtyData)...); err != nil {
		return err
	}
	if this.__expire != 0 {
		if _, err := db.Do("EXPIRE", this.__dbKey, this.__expire); err != nil {
			return err
		}
	}
	return nil
}

func (this *TestStruct1) GetMyb() bool {
	return this.myb
}

func (this *TestStruct1) GetMyf1() float32 {
	return this.myf1
}

func (this *TestStruct1) GetMyf2() float64 {
	return this.myf2
}

func (this *TestStruct1) GetMyi1() int8 {
	return this.myi1
}

func (this *TestStruct1) GetMyi2() int16 {
	return this.myi2
}

func (this *TestStruct1) GetMyi3() int32 {
	return this.myi3
}

func (this *TestStruct1) GetMyi4() int64 {
	return this.myi4
}

func (this *TestStruct1) GetMyi6() uint8 {
	return this.myi6
}

func (this *TestStruct1) GetMyi7() uint16 {
	return this.myi7
}

func (this *TestStruct1) GetMyi8() uint32 {
	return this.myi8
}

func (this *TestStruct1) GetMyi9() uint64 {
	return this.myi9
}

func (this *TestStruct1) GetMys1() string {
	return this.mys1
}

func (this *TestStruct1) GetMys2() []byte {
	return this.mys2
}

func (this *TestStruct1) GetMyst1(mutable bool) *StrcutXX {
	if mutable {
		this.__dirtyDataForStructFiled["myst1"] = nil
	}
	return &this.myst1
}

func (this *TestStruct1) GetMyst2(mutable bool) *StrcutYY {
	if mutable {
		this.__dirtyDataForStructFiled["myst2"] = nil
	}
	return &this.myst2
}

func (this *TestStruct1) SetMyb(value bool) {
	this.myb = value
	this.__dirtyData["myb"] = value
}

func (this *TestStruct1) SetMyf1(value float32) {
	this.myf1 = value
	this.__dirtyData["myf1"] = value
}

func (this *TestStruct1) SetMyf2(value float64) {
	this.myf2 = value
	this.__dirtyData["myf2"] = value
}

func (this *TestStruct1) SetMyi1(value int8) {
	this.myi1 = value
	this.__dirtyData["myi1"] = value
}

func (this *TestStruct1) SetMyi2(value int16) {
	this.myi2 = value
	this.__dirtyData["myi2"] = value
}

func (this *TestStruct1) SetMyi3(value int32) {
	this.myi3 = value
	this.__dirtyData["myi3"] = value
}

func (this *TestStruct1) SetMyi4(value int64) {
	this.myi4 = value
	this.__dirtyData["myi4"] = value
}

func (this *TestStruct1) SetMyi6(value uint8) {
	this.myi6 = value
	this.__dirtyData["myi6"] = value
}

func (this *TestStruct1) SetMyi7(value uint16) {
	this.myi7 = value
	this.__dirtyData["myi7"] = value
}

func (this *TestStruct1) SetMyi8(value uint32) {
	this.myi8 = value
	this.__dirtyData["myi8"] = value
}

func (this *TestStruct1) SetMyi9(value uint64) {
	this.myi9 = value
	this.__dirtyData["myi9"] = value
}

func (this *TestStruct1) SetMys1(value string) {
	this.mys1 = value
	this.__dirtyData["mys1"] = string([]byte(value))
}

func (this *TestStruct1) SetMys2(value []byte) {
	this.mys2 = value
	var tmp []byte = make([]byte, len(value))
	copy(tmp, value)
	this.__dirtyData["mys2"] = tmp
}
