package wgm

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultModel struct {
    Id             primitive.ObjectID `bson:"_id" json:"id"`
    CreateTime     int64              `bson:"create_time" json:"create_time"`
    LastModifyTime int64              `bson:"last_modify_time" json:"last_modify_time"`
}

type IDefaultModel interface {
    ColName() string
    GetId() string
    GetObjectID() primitive.ObjectID
    PutId(id string)
    setDefaultCreateTime()
    setDefaultLastModifyTime()
    setDefaultId()
}

// ColName 获取所对应结构体的 collection 名称
func (m *DefaultModel) ColName() string {
    return ""
}

// GetId 获取当前结构体的 id 字段内容Hex
func (m *DefaultModel) GetId() string {
    return m.Id.Hex()
}

// PutId 讲 Hex id 转为 ObjectId 然后填充至结构体
func (m *DefaultModel) PutId(id string) {
    hex, _ := primitive.ObjectIDFromHex(id)
    m.Id = hex
}

func (m *DefaultModel) setDefaultCreateTime() {
    m.CreateTime = time.Now().UnixMilli()
}

func (m *DefaultModel) setDefaultLastModifyTime() {
    m.LastModifyTime = time.Now().UnixMilli()
}

func (m *DefaultModel) setDefaultId() {
    if m.Id.IsZero() {
        m.Id = primitive.NewObjectID()
    }
}

func (m *DefaultModel) BeforeInsert(ctx context.Context) error {
    m.setDefaultId()
    m.setDefaultCreateTime()
    m.setDefaultLastModifyTime()
    return nil
}

func (m *DefaultModel) BeforeUpdate(ctx context.Context) error {
    m.setDefaultLastModifyTime()
    return nil
}

func (m *DefaultModel) BeforeUpsert(ctx context.Context) error {
    m.setDefaultId()
    m.setDefaultCreateTime()
    m.setDefaultLastModifyTime()
    return nil
}

func (m *DefaultModel) GetObjectID() primitive.ObjectID {
    return m.Id
}
