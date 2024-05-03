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

	"cutejiuges/easy-blog/models/dal/entities"
)

func newTbMenuBanner(db *gorm.DB, opts ...gen.DOOption) tbMenuBanner {
	_tbMenuBanner := tbMenuBanner{}

	_tbMenuBanner.tbMenuBannerDo.UseDB(db, opts...)
	_tbMenuBanner.tbMenuBannerDo.UseModel(&entities.TbMenuBanner{})

	tableName := _tbMenuBanner.tbMenuBannerDo.TableName()
	_tbMenuBanner.ALL = field.NewAsterisk(tableName)
	_tbMenuBanner.MenuID = field.NewInt64(tableName, "menu_id")
	_tbMenuBanner.BannerID = field.NewInt64(tableName, "banner_id")
	_tbMenuBanner.Sort = field.NewInt64(tableName, "sort")

	_tbMenuBanner.fillFieldMap()

	return _tbMenuBanner
}

type tbMenuBanner struct {
	tbMenuBannerDo tbMenuBannerDo

	ALL      field.Asterisk
	MenuID   field.Int64
	BannerID field.Int64
	Sort     field.Int64

	fieldMap map[string]field.Expr
}

func (t tbMenuBanner) Table(newTableName string) *tbMenuBanner {
	t.tbMenuBannerDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbMenuBanner) As(alias string) *tbMenuBanner {
	t.tbMenuBannerDo.DO = *(t.tbMenuBannerDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbMenuBanner) updateTableName(table string) *tbMenuBanner {
	t.ALL = field.NewAsterisk(table)
	t.MenuID = field.NewInt64(table, "menu_id")
	t.BannerID = field.NewInt64(table, "banner_id")
	t.Sort = field.NewInt64(table, "sort")

	t.fillFieldMap()

	return t
}

func (t *tbMenuBanner) WithContext(ctx context.Context) ITbMenuBannerDo {
	return t.tbMenuBannerDo.WithContext(ctx)
}

func (t tbMenuBanner) TableName() string { return t.tbMenuBannerDo.TableName() }

func (t tbMenuBanner) Alias() string { return t.tbMenuBannerDo.Alias() }

func (t tbMenuBanner) Columns(cols ...field.Expr) gen.Columns {
	return t.tbMenuBannerDo.Columns(cols...)
}

func (t *tbMenuBanner) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbMenuBanner) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["menu_id"] = t.MenuID
	t.fieldMap["banner_id"] = t.BannerID
	t.fieldMap["sort"] = t.Sort
}

func (t tbMenuBanner) clone(db *gorm.DB) tbMenuBanner {
	t.tbMenuBannerDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbMenuBanner) replaceDB(db *gorm.DB) tbMenuBanner {
	t.tbMenuBannerDo.ReplaceDB(db)
	return t
}

type tbMenuBannerDo struct{ gen.DO }

type ITbMenuBannerDo interface {
	gen.SubQuery
	Debug() ITbMenuBannerDo
	WithContext(ctx context.Context) ITbMenuBannerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITbMenuBannerDo
	WriteDB() ITbMenuBannerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITbMenuBannerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITbMenuBannerDo
	Not(conds ...gen.Condition) ITbMenuBannerDo
	Or(conds ...gen.Condition) ITbMenuBannerDo
	Select(conds ...field.Expr) ITbMenuBannerDo
	Where(conds ...gen.Condition) ITbMenuBannerDo
	Order(conds ...field.Expr) ITbMenuBannerDo
	Distinct(cols ...field.Expr) ITbMenuBannerDo
	Omit(cols ...field.Expr) ITbMenuBannerDo
	Join(table schema.Tabler, on ...field.Expr) ITbMenuBannerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITbMenuBannerDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITbMenuBannerDo
	Group(cols ...field.Expr) ITbMenuBannerDo
	Having(conds ...gen.Condition) ITbMenuBannerDo
	Limit(limit int) ITbMenuBannerDo
	Offset(offset int) ITbMenuBannerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITbMenuBannerDo
	Unscoped() ITbMenuBannerDo
	Create(values ...*entities.TbMenuBanner) error
	CreateInBatches(values []*entities.TbMenuBanner, batchSize int) error
	Save(values ...*entities.TbMenuBanner) error
	First() (*entities.TbMenuBanner, error)
	Take() (*entities.TbMenuBanner, error)
	Last() (*entities.TbMenuBanner, error)
	Find() ([]*entities.TbMenuBanner, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entities.TbMenuBanner, err error)
	FindInBatches(result *[]*entities.TbMenuBanner, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entities.TbMenuBanner) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITbMenuBannerDo
	Assign(attrs ...field.AssignExpr) ITbMenuBannerDo
	Joins(fields ...field.RelationField) ITbMenuBannerDo
	Preload(fields ...field.RelationField) ITbMenuBannerDo
	FirstOrInit() (*entities.TbMenuBanner, error)
	FirstOrCreate() (*entities.TbMenuBanner, error)
	FindByPage(offset int, limit int) (result []*entities.TbMenuBanner, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITbMenuBannerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tbMenuBannerDo) Debug() ITbMenuBannerDo {
	return t.withDO(t.DO.Debug())
}

func (t tbMenuBannerDo) WithContext(ctx context.Context) ITbMenuBannerDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbMenuBannerDo) ReadDB() ITbMenuBannerDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbMenuBannerDo) WriteDB() ITbMenuBannerDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbMenuBannerDo) Session(config *gorm.Session) ITbMenuBannerDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbMenuBannerDo) Clauses(conds ...clause.Expression) ITbMenuBannerDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbMenuBannerDo) Returning(value interface{}, columns ...string) ITbMenuBannerDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbMenuBannerDo) Not(conds ...gen.Condition) ITbMenuBannerDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbMenuBannerDo) Or(conds ...gen.Condition) ITbMenuBannerDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbMenuBannerDo) Select(conds ...field.Expr) ITbMenuBannerDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbMenuBannerDo) Where(conds ...gen.Condition) ITbMenuBannerDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbMenuBannerDo) Order(conds ...field.Expr) ITbMenuBannerDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbMenuBannerDo) Distinct(cols ...field.Expr) ITbMenuBannerDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbMenuBannerDo) Omit(cols ...field.Expr) ITbMenuBannerDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbMenuBannerDo) Join(table schema.Tabler, on ...field.Expr) ITbMenuBannerDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbMenuBannerDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITbMenuBannerDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbMenuBannerDo) RightJoin(table schema.Tabler, on ...field.Expr) ITbMenuBannerDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbMenuBannerDo) Group(cols ...field.Expr) ITbMenuBannerDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbMenuBannerDo) Having(conds ...gen.Condition) ITbMenuBannerDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbMenuBannerDo) Limit(limit int) ITbMenuBannerDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbMenuBannerDo) Offset(offset int) ITbMenuBannerDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbMenuBannerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITbMenuBannerDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbMenuBannerDo) Unscoped() ITbMenuBannerDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbMenuBannerDo) Create(values ...*entities.TbMenuBanner) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbMenuBannerDo) CreateInBatches(values []*entities.TbMenuBanner, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbMenuBannerDo) Save(values ...*entities.TbMenuBanner) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbMenuBannerDo) First() (*entities.TbMenuBanner, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbMenuBanner), nil
	}
}

func (t tbMenuBannerDo) Take() (*entities.TbMenuBanner, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbMenuBanner), nil
	}
}

func (t tbMenuBannerDo) Last() (*entities.TbMenuBanner, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbMenuBanner), nil
	}
}

func (t tbMenuBannerDo) Find() ([]*entities.TbMenuBanner, error) {
	result, err := t.DO.Find()
	return result.([]*entities.TbMenuBanner), err
}

func (t tbMenuBannerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entities.TbMenuBanner, err error) {
	buf := make([]*entities.TbMenuBanner, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbMenuBannerDo) FindInBatches(result *[]*entities.TbMenuBanner, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbMenuBannerDo) Attrs(attrs ...field.AssignExpr) ITbMenuBannerDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbMenuBannerDo) Assign(attrs ...field.AssignExpr) ITbMenuBannerDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbMenuBannerDo) Joins(fields ...field.RelationField) ITbMenuBannerDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbMenuBannerDo) Preload(fields ...field.RelationField) ITbMenuBannerDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbMenuBannerDo) FirstOrInit() (*entities.TbMenuBanner, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbMenuBanner), nil
	}
}

func (t tbMenuBannerDo) FirstOrCreate() (*entities.TbMenuBanner, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbMenuBanner), nil
	}
}

func (t tbMenuBannerDo) FindByPage(offset int, limit int) (result []*entities.TbMenuBanner, count int64, err error) {
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

func (t tbMenuBannerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbMenuBannerDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbMenuBannerDo) Delete(models ...*entities.TbMenuBanner) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbMenuBannerDo) withDO(do gen.Dao) *tbMenuBannerDo {
	t.DO = *do.(*gen.DO)
	return t
}