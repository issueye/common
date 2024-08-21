package service

import (
	"errors"
	"reflect"

	"golang.corp.yxkj.com/guava/common/model"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type ServiceContext struct {
	db       *gorm.DB
	tx       *gorm.DB
	isOpenTx bool
}

type BaseService struct {
	ctx ServiceContext
}

func (base *BaseService) OpenTx() {
	base.ctx.isOpenTx = true
	base.ctx.tx = base.ctx.db.Begin()
}

func NewBaseService(args ...ServiceContext) *BaseService {
	base := &BaseService{
		ctx: ServiceContext{},
	}

	if len(args) > 0 {
		base.SetContext(args[0])
	} else {
		base.ctx.db = DB
	}

	return base
}

type BaseIntf interface {
	SetContext(ctx ServiceContext)
	GetContext() ServiceContext
	OpenTx()
	Rollback() *gorm.DB
	Commit() *gorm.DB
	GetDB() *gorm.DB
	DataFilter(tableName string, req, list interface{}, f ListFilter) error
	SetBase(BaseService)
}

func NewService[T BaseIntf](obj T, args ...ServiceContext) T {
	if len(args) > 0 {
		obj.SetContext(args[0])
	} else {
		obj.SetBase(NewBase())
	}

	return obj
}

type ListFilter func(db *gorm.DB) (*gorm.DB, error)

func (srv *BaseService) SetContext(ctx ServiceContext) {
	srv.ctx = ctx
}

func (srv *BaseService) GetContext() ServiceContext {
	return srv.ctx
}

func NewBase() BaseService {
	return BaseService{
		ctx: ServiceContext{
			db: DB,
		},
	}
}

func (srv *BaseService) Rollback() *gorm.DB {
	if srv.ctx.isOpenTx && srv.ctx.tx != nil {
		return srv.ctx.tx.Rollback()
	}

	return nil
}

func (srv *BaseService) Commit() *gorm.DB {
	if srv.ctx.isOpenTx && srv.ctx.tx != nil {
		return srv.ctx.tx.Commit()
	}

	return nil
}

func (srv *BaseService) GetDB() *gorm.DB {
	if srv.ctx.isOpenTx && srv.ctx.tx != nil {
		return srv.ctx.tx
	}

	return srv.ctx.db
}

// DataFilter 数据过滤
func (srv *BaseService) DataFilter(tableName string, req, list interface{}, f ListFilter) error {
	query := srv.GetDB().Table(tableName)
	db, err := f(query)
	if err != nil {
		return err
	}
	count := int64(0)
	err = db.Count(&count).Error
	if err != nil {
		return err
	}
	ref := reflect.ValueOf(req)
	// 判断 ref 是否是 ptr 类型
	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}

	// 判断 req 是否有 Total 字段
	// 如果有则将count 赋值给 req.Total
	if !ref.FieldByName("Total").IsValid() {
		return errors.New("请求参数错误，传入参数中没有 Total 字段")
	}
	ref.FieldByName("Total").SetInt(count)

	// 判断是否需要进行分页
	pageNum := ref.FieldByName("PageNum").Int()
	pageSize := ref.FieldByName("PageSize").Int()
	// 如果 页码、每页大小小于0 则也不分页
	if pageNum == 0 || pageSize == 0 {
		err = db.Find(list).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		// 从 ref 中获取 分页参数 PageNum PageSize
		err = db.Offset(int(pageNum-1) * int(pageSize)).Limit(int(pageSize)).Find(list).Error
		if err != nil {
			return err
		}
		return nil
	}
}

// DataFilter 数据过滤
func Filter[T any](srv BaseIntf, tableName string, req *model.Page[T], list interface{}, f ListFilter) error {
	query := srv.GetDB().Table(tableName)
	db, err := f(query)
	if err != nil {
		return err
	}
	count := int64(0)
	err = db.Count(&count).Error
	if err != nil {
		return err
	}
	ref := reflect.ValueOf(req)
	// 判断 ref 是否是 ptr 类型
	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}

	// 判断 req 是否有 Total 字段
	// 如果有则将count 赋值给 req.Total
	if !ref.FieldByName("Total").IsValid() {
		return errors.New("请求参数错误，传入参数中没有 Total 字段")
	}
	ref.FieldByName("Total").SetInt(count)

	req.Total = count
	pageNum := req.PageNum
	pageSize := req.PageSize

	// 判断是否需要进行分页
	// 如果 页码、每页大小小于0 则也不分页
	if pageNum == 0 || pageSize == 0 {
		err = db.Find(list).Error
		if err != nil {
			return err
		}
		return nil
	} else {
		// 从 ref 中获取 分页参数 PageNum PageSize
		err = db.Offset(int(pageNum-1) * int(pageSize)).Limit(int(pageSize)).Find(list).Error
		if err != nil {
			return err
		}
		return nil
	}
}
