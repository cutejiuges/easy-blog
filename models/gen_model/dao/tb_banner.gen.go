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

func newTbBanner(db *gorm.DB, opts ...gen.DOOption) tbBanner {
	_tbBanner := tbBanner{}

	_tbBanner.tbBannerDo.UseDB(db, opts...)
	_tbBanner.tbBannerDo.UseModel(&domain.TbBanner{})

	tableName := _tbBanner.tbBannerDo.TableName()
	_tbBanner.ALL = field.NewAsterisk(tableName)
	_tbBanner.ID = field.NewInt64(tableName, "id")
	_tbBanner.CreatedAt = field.NewTime(tableName, "created_at")
	_tbBanner.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tbBanner.DeletedAt = field.NewField(tableName, "deleted_at")
	_tbBanner.Path = field.NewString(tableName, "path")
	_tbBanner.Hash = field.NewString(tableName, "hash")
	_tbBanner.Name = field.NewString(tableName, "name")

	_tbBanner.fillFieldMap()

	return _tbBanner
}

type tbBanner struct {
	tbBannerDo tbBannerDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Path      field.String
	Hash      field.String
	Name      field.String

	fieldMap map[string]field.Expr
}

func (t tbBanner) Table(newTableName string) *tbBanner {
	t.tbBannerDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbBanner) As(alias string) *tbBanner {
	t.tbBannerDo.DO = *(t.tbBannerDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbBanner) updateTableName(table string) *tbBanner {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.Path = field.NewString(table, "path")
	t.Hash = field.NewString(table, "hash")
	t.Name = field.NewString(table, "name")

	t.fillFieldMap()

	return t
}

func (t *tbBanner) WithContext(ctx context.Context) ITbBannerDo { return t.tbBannerDo.WithContext(ctx) }

func (t tbBanner) TableName() string { return t.tbBannerDo.TableName() }

func (t tbBanner) Alias() string { return t.tbBannerDo.Alias() }

func (t tbBanner) Columns(cols ...field.Expr) gen.Columns { return t.tbBannerDo.Columns(cols...) }

func (t *tbBanner) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbBanner) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["path"] = t.Path
	t.fieldMap["hash"] = t.Hash
	t.fieldMap["name"] = t.Name
}

func (t tbBanner) clone(db *gorm.DB) tbBanner {
	t.tbBannerDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbBanner) replaceDB(db *gorm.DB) tbBanner {
	t.tbBannerDo.ReplaceDB(db)
	return t
}

type tbBannerDo struct{ gen.DO }

type ITbBannerDo interface {
	gen.SubQuery
	Debug() ITbBannerDo
	WithContext(ctx context.Context) ITbBannerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITbBannerDo
	WriteDB() ITbBannerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITbBannerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITbBannerDo
	Not(conds ...gen.Condition) ITbBannerDo
	Or(conds ...gen.Condition) ITbBannerDo
	Select(conds ...field.Expr) ITbBannerDo
	Where(conds ...gen.Condition) ITbBannerDo
	Order(conds ...field.Expr) ITbBannerDo
	Distinct(cols ...field.Expr) ITbBannerDo
	Omit(cols ...field.Expr) ITbBannerDo
	Join(table schema.Tabler, on ...field.Expr) ITbBannerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITbBannerDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITbBannerDo
	Group(cols ...field.Expr) ITbBannerDo
	Having(conds ...gen.Condition) ITbBannerDo
	Limit(limit int) ITbBannerDo
	Offset(offset int) ITbBannerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITbBannerDo
	Unscoped() ITbBannerDo
	Create(values ...*domain.TbBanner) error
	CreateInBatches(values []*domain.TbBanner, batchSize int) error
	Save(values ...*domain.TbBanner) error
	First() (*domain.TbBanner, error)
	Take() (*domain.TbBanner, error)
	Last() (*domain.TbBanner, error)
	Find() ([]*domain.TbBanner, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.TbBanner, err error)
	FindInBatches(result *[]*domain.TbBanner, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.TbBanner) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITbBannerDo
	Assign(attrs ...field.AssignExpr) ITbBannerDo
	Joins(fields ...field.RelationField) ITbBannerDo
	Preload(fields ...field.RelationField) ITbBannerDo
	FirstOrInit() (*domain.TbBanner, error)
	FirstOrCreate() (*domain.TbBanner, error)
	FindByPage(offset int, limit int) (result []*domain.TbBanner, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITbBannerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tbBannerDo) Debug() ITbBannerDo {
	return t.withDO(t.DO.Debug())
}

func (t tbBannerDo) WithContext(ctx context.Context) ITbBannerDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbBannerDo) ReadDB() ITbBannerDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbBannerDo) WriteDB() ITbBannerDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbBannerDo) Session(config *gorm.Session) ITbBannerDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbBannerDo) Clauses(conds ...clause.Expression) ITbBannerDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbBannerDo) Returning(value interface{}, columns ...string) ITbBannerDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbBannerDo) Not(conds ...gen.Condition) ITbBannerDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbBannerDo) Or(conds ...gen.Condition) ITbBannerDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbBannerDo) Select(conds ...field.Expr) ITbBannerDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbBannerDo) Where(conds ...gen.Condition) ITbBannerDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbBannerDo) Order(conds ...field.Expr) ITbBannerDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbBannerDo) Distinct(cols ...field.Expr) ITbBannerDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbBannerDo) Omit(cols ...field.Expr) ITbBannerDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbBannerDo) Join(table schema.Tabler, on ...field.Expr) ITbBannerDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbBannerDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITbBannerDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbBannerDo) RightJoin(table schema.Tabler, on ...field.Expr) ITbBannerDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbBannerDo) Group(cols ...field.Expr) ITbBannerDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbBannerDo) Having(conds ...gen.Condition) ITbBannerDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbBannerDo) Limit(limit int) ITbBannerDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbBannerDo) Offset(offset int) ITbBannerDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbBannerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITbBannerDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbBannerDo) Unscoped() ITbBannerDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbBannerDo) Create(values ...*domain.TbBanner) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbBannerDo) CreateInBatches(values []*domain.TbBanner, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbBannerDo) Save(values ...*domain.TbBanner) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbBannerDo) First() (*domain.TbBanner, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbBanner), nil
	}
}

func (t tbBannerDo) Take() (*domain.TbBanner, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbBanner), nil
	}
}

func (t tbBannerDo) Last() (*domain.TbBanner, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbBanner), nil
	}
}

func (t tbBannerDo) Find() ([]*domain.TbBanner, error) {
	result, err := t.DO.Find()
	return result.([]*domain.TbBanner), err
}

func (t tbBannerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.TbBanner, err error) {
	buf := make([]*domain.TbBanner, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbBannerDo) FindInBatches(result *[]*domain.TbBanner, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbBannerDo) Attrs(attrs ...field.AssignExpr) ITbBannerDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbBannerDo) Assign(attrs ...field.AssignExpr) ITbBannerDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbBannerDo) Joins(fields ...field.RelationField) ITbBannerDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbBannerDo) Preload(fields ...field.RelationField) ITbBannerDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbBannerDo) FirstOrInit() (*domain.TbBanner, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbBanner), nil
	}
}

func (t tbBannerDo) FirstOrCreate() (*domain.TbBanner, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbBanner), nil
	}
}

func (t tbBannerDo) FindByPage(offset int, limit int) (result []*domain.TbBanner, count int64, err error) {
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

func (t tbBannerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbBannerDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbBannerDo) Delete(models ...*domain.TbBanner) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbBannerDo) withDO(do gen.Dao) *tbBannerDo {
	t.DO = *do.(*gen.DO)
	return t
}