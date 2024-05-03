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

func newTbArticleTag(db *gorm.DB, opts ...gen.DOOption) tbArticleTag {
	_tbArticleTag := tbArticleTag{}

	_tbArticleTag.tbArticleTagDo.UseDB(db, opts...)
	_tbArticleTag.tbArticleTagDo.UseModel(&entities.TbArticleTag{})

	tableName := _tbArticleTag.tbArticleTagDo.TableName()
	_tbArticleTag.ALL = field.NewAsterisk(tableName)
	_tbArticleTag.TagID = field.NewInt64(tableName, "tag_id")
	_tbArticleTag.ArticleID = field.NewInt64(tableName, "article_id")

	_tbArticleTag.fillFieldMap()

	return _tbArticleTag
}

type tbArticleTag struct {
	tbArticleTagDo tbArticleTagDo

	ALL       field.Asterisk
	TagID     field.Int64
	ArticleID field.Int64

	fieldMap map[string]field.Expr
}

func (t tbArticleTag) Table(newTableName string) *tbArticleTag {
	t.tbArticleTagDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbArticleTag) As(alias string) *tbArticleTag {
	t.tbArticleTagDo.DO = *(t.tbArticleTagDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbArticleTag) updateTableName(table string) *tbArticleTag {
	t.ALL = field.NewAsterisk(table)
	t.TagID = field.NewInt64(table, "tag_id")
	t.ArticleID = field.NewInt64(table, "article_id")

	t.fillFieldMap()

	return t
}

func (t *tbArticleTag) WithContext(ctx context.Context) ITbArticleTagDo {
	return t.tbArticleTagDo.WithContext(ctx)
}

func (t tbArticleTag) TableName() string { return t.tbArticleTagDo.TableName() }

func (t tbArticleTag) Alias() string { return t.tbArticleTagDo.Alias() }

func (t tbArticleTag) Columns(cols ...field.Expr) gen.Columns {
	return t.tbArticleTagDo.Columns(cols...)
}

func (t *tbArticleTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbArticleTag) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 2)
	t.fieldMap["tag_id"] = t.TagID
	t.fieldMap["article_id"] = t.ArticleID
}

func (t tbArticleTag) clone(db *gorm.DB) tbArticleTag {
	t.tbArticleTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbArticleTag) replaceDB(db *gorm.DB) tbArticleTag {
	t.tbArticleTagDo.ReplaceDB(db)
	return t
}

type tbArticleTagDo struct{ gen.DO }

type ITbArticleTagDo interface {
	gen.SubQuery
	Debug() ITbArticleTagDo
	WithContext(ctx context.Context) ITbArticleTagDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITbArticleTagDo
	WriteDB() ITbArticleTagDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITbArticleTagDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITbArticleTagDo
	Not(conds ...gen.Condition) ITbArticleTagDo
	Or(conds ...gen.Condition) ITbArticleTagDo
	Select(conds ...field.Expr) ITbArticleTagDo
	Where(conds ...gen.Condition) ITbArticleTagDo
	Order(conds ...field.Expr) ITbArticleTagDo
	Distinct(cols ...field.Expr) ITbArticleTagDo
	Omit(cols ...field.Expr) ITbArticleTagDo
	Join(table schema.Tabler, on ...field.Expr) ITbArticleTagDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITbArticleTagDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITbArticleTagDo
	Group(cols ...field.Expr) ITbArticleTagDo
	Having(conds ...gen.Condition) ITbArticleTagDo
	Limit(limit int) ITbArticleTagDo
	Offset(offset int) ITbArticleTagDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITbArticleTagDo
	Unscoped() ITbArticleTagDo
	Create(values ...*entities.TbArticleTag) error
	CreateInBatches(values []*entities.TbArticleTag, batchSize int) error
	Save(values ...*entities.TbArticleTag) error
	First() (*entities.TbArticleTag, error)
	Take() (*entities.TbArticleTag, error)
	Last() (*entities.TbArticleTag, error)
	Find() ([]*entities.TbArticleTag, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entities.TbArticleTag, err error)
	FindInBatches(result *[]*entities.TbArticleTag, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entities.TbArticleTag) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITbArticleTagDo
	Assign(attrs ...field.AssignExpr) ITbArticleTagDo
	Joins(fields ...field.RelationField) ITbArticleTagDo
	Preload(fields ...field.RelationField) ITbArticleTagDo
	FirstOrInit() (*entities.TbArticleTag, error)
	FirstOrCreate() (*entities.TbArticleTag, error)
	FindByPage(offset int, limit int) (result []*entities.TbArticleTag, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITbArticleTagDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tbArticleTagDo) Debug() ITbArticleTagDo {
	return t.withDO(t.DO.Debug())
}

func (t tbArticleTagDo) WithContext(ctx context.Context) ITbArticleTagDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbArticleTagDo) ReadDB() ITbArticleTagDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbArticleTagDo) WriteDB() ITbArticleTagDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbArticleTagDo) Session(config *gorm.Session) ITbArticleTagDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbArticleTagDo) Clauses(conds ...clause.Expression) ITbArticleTagDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbArticleTagDo) Returning(value interface{}, columns ...string) ITbArticleTagDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbArticleTagDo) Not(conds ...gen.Condition) ITbArticleTagDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbArticleTagDo) Or(conds ...gen.Condition) ITbArticleTagDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbArticleTagDo) Select(conds ...field.Expr) ITbArticleTagDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbArticleTagDo) Where(conds ...gen.Condition) ITbArticleTagDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbArticleTagDo) Order(conds ...field.Expr) ITbArticleTagDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbArticleTagDo) Distinct(cols ...field.Expr) ITbArticleTagDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbArticleTagDo) Omit(cols ...field.Expr) ITbArticleTagDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbArticleTagDo) Join(table schema.Tabler, on ...field.Expr) ITbArticleTagDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbArticleTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITbArticleTagDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbArticleTagDo) RightJoin(table schema.Tabler, on ...field.Expr) ITbArticleTagDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbArticleTagDo) Group(cols ...field.Expr) ITbArticleTagDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbArticleTagDo) Having(conds ...gen.Condition) ITbArticleTagDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbArticleTagDo) Limit(limit int) ITbArticleTagDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbArticleTagDo) Offset(offset int) ITbArticleTagDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbArticleTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITbArticleTagDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbArticleTagDo) Unscoped() ITbArticleTagDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbArticleTagDo) Create(values ...*entities.TbArticleTag) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbArticleTagDo) CreateInBatches(values []*entities.TbArticleTag, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbArticleTagDo) Save(values ...*entities.TbArticleTag) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbArticleTagDo) First() (*entities.TbArticleTag, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbArticleTag), nil
	}
}

func (t tbArticleTagDo) Take() (*entities.TbArticleTag, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbArticleTag), nil
	}
}

func (t tbArticleTagDo) Last() (*entities.TbArticleTag, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbArticleTag), nil
	}
}

func (t tbArticleTagDo) Find() ([]*entities.TbArticleTag, error) {
	result, err := t.DO.Find()
	return result.([]*entities.TbArticleTag), err
}

func (t tbArticleTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entities.TbArticleTag, err error) {
	buf := make([]*entities.TbArticleTag, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbArticleTagDo) FindInBatches(result *[]*entities.TbArticleTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbArticleTagDo) Attrs(attrs ...field.AssignExpr) ITbArticleTagDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbArticleTagDo) Assign(attrs ...field.AssignExpr) ITbArticleTagDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbArticleTagDo) Joins(fields ...field.RelationField) ITbArticleTagDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbArticleTagDo) Preload(fields ...field.RelationField) ITbArticleTagDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbArticleTagDo) FirstOrInit() (*entities.TbArticleTag, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbArticleTag), nil
	}
}

func (t tbArticleTagDo) FirstOrCreate() (*entities.TbArticleTag, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbArticleTag), nil
	}
}

func (t tbArticleTagDo) FindByPage(offset int, limit int) (result []*entities.TbArticleTag, count int64, err error) {
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

func (t tbArticleTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbArticleTagDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbArticleTagDo) Delete(models ...*entities.TbArticleTag) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbArticleTagDo) withDO(do gen.Dao) *tbArticleTagDo {
	t.DO = *do.(*gen.DO)
	return t
}
