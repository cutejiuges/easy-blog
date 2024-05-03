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

func newTbComment(db *gorm.DB, opts ...gen.DOOption) tbComment {
	_tbComment := tbComment{}

	_tbComment.tbCommentDo.UseDB(db, opts...)
	_tbComment.tbCommentDo.UseModel(&entities.TbComment{})

	tableName := _tbComment.tbCommentDo.TableName()
	_tbComment.ALL = field.NewAsterisk(tableName)
	_tbComment.ID = field.NewInt64(tableName, "id")
	_tbComment.CreatedAt = field.NewTime(tableName, "created_at")
	_tbComment.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tbComment.DeletedAt = field.NewField(tableName, "deleted_at")
	_tbComment.ParentCommentID = field.NewInt64(tableName, "parent_comment_id")
	_tbComment.Content = field.NewString(tableName, "content")
	_tbComment.DiggCount = field.NewInt64(tableName, "digg_count")
	_tbComment.CommentCount = field.NewInt64(tableName, "comment_count")
	_tbComment.ArticleID = field.NewInt64(tableName, "article_id")
	_tbComment.UserID = field.NewInt64(tableName, "user_id")

	_tbComment.fillFieldMap()

	return _tbComment
}

type tbComment struct {
	tbCommentDo tbCommentDo

	ALL             field.Asterisk
	ID              field.Int64
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field
	ParentCommentID field.Int64
	Content         field.String
	DiggCount       field.Int64
	CommentCount    field.Int64
	ArticleID       field.Int64
	UserID          field.Int64

	fieldMap map[string]field.Expr
}

func (t tbComment) Table(newTableName string) *tbComment {
	t.tbCommentDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbComment) As(alias string) *tbComment {
	t.tbCommentDo.DO = *(t.tbCommentDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbComment) updateTableName(table string) *tbComment {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.ParentCommentID = field.NewInt64(table, "parent_comment_id")
	t.Content = field.NewString(table, "content")
	t.DiggCount = field.NewInt64(table, "digg_count")
	t.CommentCount = field.NewInt64(table, "comment_count")
	t.ArticleID = field.NewInt64(table, "article_id")
	t.UserID = field.NewInt64(table, "user_id")

	t.fillFieldMap()

	return t
}

func (t *tbComment) WithContext(ctx context.Context) ITbCommentDo {
	return t.tbCommentDo.WithContext(ctx)
}

func (t tbComment) TableName() string { return t.tbCommentDo.TableName() }

func (t tbComment) Alias() string { return t.tbCommentDo.Alias() }

func (t tbComment) Columns(cols ...field.Expr) gen.Columns { return t.tbCommentDo.Columns(cols...) }

func (t *tbComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbComment) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["parent_comment_id"] = t.ParentCommentID
	t.fieldMap["content"] = t.Content
	t.fieldMap["digg_count"] = t.DiggCount
	t.fieldMap["comment_count"] = t.CommentCount
	t.fieldMap["article_id"] = t.ArticleID
	t.fieldMap["user_id"] = t.UserID
}

func (t tbComment) clone(db *gorm.DB) tbComment {
	t.tbCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbComment) replaceDB(db *gorm.DB) tbComment {
	t.tbCommentDo.ReplaceDB(db)
	return t
}

type tbCommentDo struct{ gen.DO }

type ITbCommentDo interface {
	gen.SubQuery
	Debug() ITbCommentDo
	WithContext(ctx context.Context) ITbCommentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITbCommentDo
	WriteDB() ITbCommentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITbCommentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITbCommentDo
	Not(conds ...gen.Condition) ITbCommentDo
	Or(conds ...gen.Condition) ITbCommentDo
	Select(conds ...field.Expr) ITbCommentDo
	Where(conds ...gen.Condition) ITbCommentDo
	Order(conds ...field.Expr) ITbCommentDo
	Distinct(cols ...field.Expr) ITbCommentDo
	Omit(cols ...field.Expr) ITbCommentDo
	Join(table schema.Tabler, on ...field.Expr) ITbCommentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITbCommentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITbCommentDo
	Group(cols ...field.Expr) ITbCommentDo
	Having(conds ...gen.Condition) ITbCommentDo
	Limit(limit int) ITbCommentDo
	Offset(offset int) ITbCommentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITbCommentDo
	Unscoped() ITbCommentDo
	Create(values ...*entities.TbComment) error
	CreateInBatches(values []*entities.TbComment, batchSize int) error
	Save(values ...*entities.TbComment) error
	First() (*entities.TbComment, error)
	Take() (*entities.TbComment, error)
	Last() (*entities.TbComment, error)
	Find() ([]*entities.TbComment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entities.TbComment, err error)
	FindInBatches(result *[]*entities.TbComment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entities.TbComment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITbCommentDo
	Assign(attrs ...field.AssignExpr) ITbCommentDo
	Joins(fields ...field.RelationField) ITbCommentDo
	Preload(fields ...field.RelationField) ITbCommentDo
	FirstOrInit() (*entities.TbComment, error)
	FirstOrCreate() (*entities.TbComment, error)
	FindByPage(offset int, limit int) (result []*entities.TbComment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITbCommentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tbCommentDo) Debug() ITbCommentDo {
	return t.withDO(t.DO.Debug())
}

func (t tbCommentDo) WithContext(ctx context.Context) ITbCommentDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbCommentDo) ReadDB() ITbCommentDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbCommentDo) WriteDB() ITbCommentDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbCommentDo) Session(config *gorm.Session) ITbCommentDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbCommentDo) Clauses(conds ...clause.Expression) ITbCommentDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbCommentDo) Returning(value interface{}, columns ...string) ITbCommentDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbCommentDo) Not(conds ...gen.Condition) ITbCommentDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbCommentDo) Or(conds ...gen.Condition) ITbCommentDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbCommentDo) Select(conds ...field.Expr) ITbCommentDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbCommentDo) Where(conds ...gen.Condition) ITbCommentDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbCommentDo) Order(conds ...field.Expr) ITbCommentDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbCommentDo) Distinct(cols ...field.Expr) ITbCommentDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbCommentDo) Omit(cols ...field.Expr) ITbCommentDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbCommentDo) Join(table schema.Tabler, on ...field.Expr) ITbCommentDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITbCommentDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) ITbCommentDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbCommentDo) Group(cols ...field.Expr) ITbCommentDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbCommentDo) Having(conds ...gen.Condition) ITbCommentDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbCommentDo) Limit(limit int) ITbCommentDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbCommentDo) Offset(offset int) ITbCommentDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITbCommentDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbCommentDo) Unscoped() ITbCommentDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbCommentDo) Create(values ...*entities.TbComment) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbCommentDo) CreateInBatches(values []*entities.TbComment, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbCommentDo) Save(values ...*entities.TbComment) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbCommentDo) First() (*entities.TbComment, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbComment), nil
	}
}

func (t tbCommentDo) Take() (*entities.TbComment, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbComment), nil
	}
}

func (t tbCommentDo) Last() (*entities.TbComment, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbComment), nil
	}
}

func (t tbCommentDo) Find() ([]*entities.TbComment, error) {
	result, err := t.DO.Find()
	return result.([]*entities.TbComment), err
}

func (t tbCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entities.TbComment, err error) {
	buf := make([]*entities.TbComment, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbCommentDo) FindInBatches(result *[]*entities.TbComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbCommentDo) Attrs(attrs ...field.AssignExpr) ITbCommentDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbCommentDo) Assign(attrs ...field.AssignExpr) ITbCommentDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbCommentDo) Joins(fields ...field.RelationField) ITbCommentDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbCommentDo) Preload(fields ...field.RelationField) ITbCommentDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbCommentDo) FirstOrInit() (*entities.TbComment, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbComment), nil
	}
}

func (t tbCommentDo) FirstOrCreate() (*entities.TbComment, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entities.TbComment), nil
	}
}

func (t tbCommentDo) FindByPage(offset int, limit int) (result []*entities.TbComment, count int64, err error) {
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

func (t tbCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbCommentDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbCommentDo) Delete(models ...*entities.TbComment) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbCommentDo) withDO(do gen.Dao) *tbCommentDo {
	t.DO = *do.(*gen.DO)
	return t
}
