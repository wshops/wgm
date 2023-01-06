package wgm

import (
    "context"

    "github.com/gookit/slog"
    "github.com/qiniu/qmgo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func (w *wgm) GetModelCollection(m IDefaultModel) *qmgo.Collection {
    return w.client.Database(w.dbName).Collection(m.ColName())
}

func (w *wgm) GetCollection(name string) *qmgo.Collection {
    return w.client.Database(w.dbName).Collection(name)
}

func (w *wgm) Ctx() context.Context {
    return w.newCtx()
}

func Col(name string) *qmgo.Collection {
    return instance.GetCollection(name)
}

func Ctx() context.Context {
    return instance.Ctx()
}

// IsNoResult 是否结果不存在
// err 数据库查询后返回的 err
// bool 结果，true 为未查询到数据，反之亦然
func IsNoResult(err error) bool {
    if err == mongo.ErrNoDocuments || err == qmgo.ErrNoSuchDocuments {
        return true
    }

    return false
}

// FindWithPage 数据库分页查询
// m           查询的合集
// filter      查询条件，查询全部文档使用 nil，查询条件使用 bson.M
// res         结果集指针，必须为指向切片的指针!!!
// pageSize    页面大小
// currentPage 当前页面
// totalDoc 总数据数量
// totalPage 总页面数量
func FindWithPage(m IDefaultModel, filter any, res any, pageSize int64, currentPage int64, sortFields ...string) (totalDoc int64, totalPage int64) {
    if instance == nil {
        slog.Fatal("must initialize WGM first, by calling InitWgm() method")
    }

    if filter == nil {
        filter = bson.D{}
    }

    countDoc, err := instance.GetModelCollection(m).Find(instance.Ctx(), filter).Count()
    if IsNoResult(err) {
        return 0, 0
    }
    if err != nil {
        slog.Error(err)
        res = nil
        return 0, 0
    }

    // 计算应该跳过的doc,
    offset := (currentPage - 1) * pageSize
    // 计算应该返回多少条记录
    var size int64
    if countDoc-offset < pageSize {
        size = countDoc - offset
    } else {
        size = pageSize
    }

    err = instance.GetModelCollection(m).Find(instance.Ctx(), filter).Sort(sortFields...).Limit(size).Skip(offset).All(res)

    if err != nil {
        slog.Error(err)
        return 0, 0
    }
    // 判断总页数totalPage
    if pageSize == 1 {
        totalPage = countDoc / pageSize
    } else {
        totalPage = 1 + countDoc/pageSize
    }
    return countDoc, totalPage
}

// FindOne 查询符合条件的第一条数据
// m      查询的合集，结果也会被绑定在这
// filter 查询条件，查询全部文档使用 nil，查询条件使用 bson.M
// hasResult 是否查询到结果
func FindOne(m IDefaultModel, filter map[string]any) (hasResult bool) {
    if instance == nil {
        slog.Fatal("must initialize WGM first, by calling InitWgm() method")
    }

    if filter == nil {
        filter = bson.M{}
    }
    err := instance.GetModelCollection(m).Find(instance.Ctx(), filter).One(m)

    if IsNoResult(err) {
        return false
    }

    if err != nil {
        slog.Error(err)
        return false
    }

    return true
}

func FindById(colName string, id string, res any) (bool, error) {
    if instance == nil {
        slog.Fatal("must initialize WGM first, by calling InitWgm() method")
    }
    err := instance.GetCollection(colName).Find(instance.Ctx(), bson.M{"_id": MustHexToObjectId(id)}).One(res)
    if IsNoResult(err) {
        return false, err
    }

    if err != nil {
        slog.Error(err)
        return false, err
    }

    return true, nil
}

func MustHexToObjectId(strId string) primitive.ObjectID {
    objId, err := primitive.ObjectIDFromHex(strId)
    if err != nil {
        slog.Error(err)
        return primitive.NilObjectID
    }
    return objId
}

func Insert(m IDefaultModel) (*qmgo.InsertOneResult, error) {
    if instance == nil {
        slog.Fatal("must initialize WGM first, by calling InitWgm() method")
    }
    result, err := instance.GetModelCollection(m).InsertOne(instance.Ctx(), m)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func Update(m IDefaultModel, filter ...map[string]any) error {
    if instance == nil {
        slog.Fatal("must initialize WGM first, by calling InitWgm() method")
    }
    f := bson.M{}
    if len(filter) > 0 {
        f = filter[0]
        f["_id"] = m.GetObjectID()
    } else {
        f["_id"] = m.GetObjectID()
    }

    m.setDefaultLastModifyTime()
    err := instance.GetModelCollection(m).UpdateOne(instance.Ctx(), f, bson.M{"$set": m})
    if err != nil {
        return err
    }
    return nil
}

func Delete(m IDefaultModel) error {
    if instance == nil {
        slog.Fatal("must initialize WGM first, by calling InitWgm() method")
    }

    err := instance.GetModelCollection(m).RemoveId(instance.Ctx(), m.GetObjectID())
    if err != nil {
        return err
    }
    return nil
}

// ExistInDB 查询是否存在数据库
// m 查询的合集
// filter 查询条件，查询全部文档使用 nil，查询条件使用 bson.M
// bool 是否存在
func ExistInDB(m IDefaultModel, filter any) bool {
    if instance == nil {
        slog.Fatal("must initialize WGM first, by calling InitWgm() method")
    }
    if filter == nil {
        filter = bson.D{}
    }
    err := instance.GetModelCollection(m).Find(instance.Ctx(), filter).One(nil)
    if IsNoResult(err) {
        return false
    }
    return true
}