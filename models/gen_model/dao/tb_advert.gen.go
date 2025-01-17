// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"cutejiuges/easy-blog/models/gen_model/domain"
)

func newTbAdvert(db *gorm.DB, opts ...gen.DOOption) tbAdvert {
	_tbAdvert := tbAdvert{}

	_tbAdvert.tbAdvertDo.UseDB(db, opts...)
	_tbAdvert.tbAdvertDo.UseModel(&domain.TbAdvert{})

	tableName := _tbAdvert.tbAdvertDo.TableName()
	_tbAdvert.ALL = field.NewAsterisk(tableName)
	_tbAdvert.ID = field.NewInt64(tableName, "id")
	_tbAdvert.CreatedAt = field.NewTime(tableName, "created_at")
	_tbAdvert.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tbAdvert.DeletedAt = field.NewField(tableName, "deleted_at")
	_tbAdvert.Title = field.NewString(tableName, "title")
	_tbAdvert.Href = field.NewString(tableName, "href")
	_tbAdvert.Images = field.NewString(tableName, "images")
	_tbAdvert.IsShow = field.NewInt64(tableName, "is_show")

	_tbAdvert.fillFieldMap()

	return _tbAdvert
}

type tbAdvert struct {
	tbAdvertDo tbAdvertDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Title     field.String
	Href      field.String
	Images    field.String
	IsShow    field.Int64

	fieldMap map[string]field.Expr
}

func (t tbAdvert) Table(newTableName string) *tbAdvert {
	t.tbAdvertDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbAdvert) As(alias string) *tbAdvert {
	t.tbAdvertDo.DO = *(t.tbAdvertDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbAdvert) updateTableName(table string) *tbAdvert {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.Title = field.NewString(table, "title")
	t.Href = field.NewString(table, "href")
	t.Images = field.NewString(table, "images")
	t.IsShow = field.NewInt64(table, "is_show")

	t.fillFieldMap()

	return t
}

func (t *tbAdvert) WithContext(ctx context.Context) ITbAdvertDo { return t.tbAdvertDo.WithContext(ctx) }

func (t tbAdvert) TableName() string { return t.tbAdvertDo.TableName() }

func (t tbAdvert) Alias() string { return t.tbAdvertDo.Alias() }

func (t tbAdvert) Columns(cols ...field.Expr) gen.Columns { return t.tbAdvertDo.Columns(cols...) }

func (t *tbAdvert) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbAdvert) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["title"] = t.Title
	t.fieldMap["href"] = t.Href
	t.fieldMap["images"] = t.Images
	t.fieldMap["is_show"] = t.IsShow
}

func (t tbAdvert) clone(db *gorm.DB) tbAdvert {
	t.tbAdvertDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbAdvert) replaceDB(db *gorm.DB) tbAdvert {
	t.tbAdvertDo.ReplaceDB(db)
	return t
}

type tbAdvertDo struct{ gen.DO }

type ITbAdvertDo interface {
	gen.SubQuery
	Debug() ITbAdvertDo
	WithContext(ctx context.Context) ITbAdvertDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITbAdvertDo
	WriteDB() ITbAdvertDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITbAdvertDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITbAdvertDo
	Not(conds ...gen.Condition) ITbAdvertDo
	Or(conds ...gen.Condition) ITbAdvertDo
	Select(conds ...field.Expr) ITbAdvertDo
	Where(conds ...gen.Condition) ITbAdvertDo
	Order(conds ...field.Expr) ITbAdvertDo
	Distinct(cols ...field.Expr) ITbAdvertDo
	Omit(cols ...field.Expr) ITbAdvertDo
	Join(table schema.Tabler, on ...field.Expr) ITbAdvertDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITbAdvertDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITbAdvertDo
	Group(cols ...field.Expr) ITbAdvertDo
	Having(conds ...gen.Condition) ITbAdvertDo
	Limit(limit int) ITbAdvertDo
	Offset(offset int) ITbAdvertDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITbAdvertDo
	Unscoped() ITbAdvertDo
	Create(values ...*domain.TbAdvert) error
	CreateInBatches(values []*domain.TbAdvert, batchSize int) error
	Save(values ...*domain.TbAdvert) error
	First() (*domain.TbAdvert, error)
	Take() (*domain.TbAdvert, error)
	Last() (*domain.TbAdvert, error)
	Find() ([]*domain.TbAdvert, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.TbAdvert, err error)
	FindInBatches(result *[]*domain.TbAdvert, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.TbAdvert) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITbAdvertDo
	Assign(attrs ...field.AssignExpr) ITbAdvertDo
	Joins(fields ...field.RelationField) ITbAdvertDo
	Preload(fields ...field.RelationField) ITbAdvertDo
	FirstOrInit() (*domain.TbAdvert, error)
	FirstOrCreate() (*domain.TbAdvert, error)
	FindByPage(offset int, limit int) (result []*domain.TbAdvert, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITbAdvertDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tbAdvertDo) Debug() ITbAdvertDo {
	return t.withDO(t.DO.Debug())
}

func (t tbAdvertDo) WithContext(ctx context.Context) ITbAdvertDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbAdvertDo) ReadDB() ITbAdvertDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbAdvertDo) WriteDB() ITbAdvertDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbAdvertDo) Session(config *gorm.Session) ITbAdvertDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbAdvertDo) Clauses(conds ...clause.Expression) ITbAdvertDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbAdvertDo) Returning(value interface{}, columns ...string) ITbAdvertDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbAdvertDo) Not(conds ...gen.Condition) ITbAdvertDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbAdvertDo) Or(conds ...gen.Condition) ITbAdvertDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbAdvertDo) Select(conds ...field.Expr) ITbAdvertDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbAdvertDo) Where(conds ...gen.Condition) ITbAdvertDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbAdvertDo) Order(conds ...field.Expr) ITbAdvertDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbAdvertDo) Distinct(cols ...field.Expr) ITbAdvertDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbAdvertDo) Omit(cols ...field.Expr) ITbAdvertDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbAdvertDo) Join(table schema.Tabler, on ...field.Expr) ITbAdvertDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbAdvertDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITbAdvertDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbAdvertDo) RightJoin(table schema.Tabler, on ...field.Expr) ITbAdvertDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbAdvertDo) Group(cols ...field.Expr) ITbAdvertDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbAdvertDo) Having(conds ...gen.Condition) ITbAdvertDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbAdvertDo) Limit(limit int) ITbAdvertDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbAdvertDo) Offset(offset int) ITbAdvertDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbAdvertDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITbAdvertDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbAdvertDo) Unscoped() ITbAdvertDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbAdvertDo) Create(values ...*domain.TbAdvert) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbAdvertDo) CreateInBatches(values []*domain.TbAdvert, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbAdvertDo) Save(values ...*domain.TbAdvert) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbAdvertDo) First() (*domain.TbAdvert, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbAdvert), nil
	}
}

func (t tbAdvertDo) Take() (*domain.TbAdvert, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbAdvert), nil
	}
}

func (t tbAdvertDo) Last() (*domain.TbAdvert, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbAdvert), nil
	}
}

func (t tbAdvertDo) Find() ([]*domain.TbAdvert, error) {
	result, err := t.DO.Find()
	return result.([]*domain.TbAdvert), err
}

func (t tbAdvertDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.TbAdvert, err error) {
	buf := make([]*domain.TbAdvert, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbAdvertDo) FindInBatches(result *[]*domain.TbAdvert, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbAdvertDo) Attrs(attrs ...field.AssignExpr) ITbAdvertDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbAdvertDo) Assign(attrs ...field.AssignExpr) ITbAdvertDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbAdvertDo) Joins(fields ...field.RelationField) ITbAdvertDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbAdvertDo) Preload(fields ...field.RelationField) ITbAdvertDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbAdvertDo) FirstOrInit() (*domain.TbAdvert, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbAdvert), nil
	}
}

func (t tbAdvertDo) FirstOrCreate() (*domain.TbAdvert, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbAdvert), nil
	}
}

func (t tbAdvertDo) FindByPage(offset int, limit int) (result []*domain.TbAdvert, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tbAdvertDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbAdvertDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbAdvertDo) Delete(models ...*domain.TbAdvert) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbAdvertDo) withDO(do gen.Dao) *tbAdvertDo {
	t.DO = *do.(*gen.DO)
	return t
}
