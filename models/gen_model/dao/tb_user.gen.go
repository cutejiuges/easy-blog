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

func newTbUser(db *gorm.DB, opts ...gen.DOOption) tbUser {
	_tbUser := tbUser{}

	_tbUser.tbUserDo.UseDB(db, opts...)
	_tbUser.tbUserDo.UseModel(&domain.TbUser{})

	tableName := _tbUser.tbUserDo.TableName()
	_tbUser.ALL = field.NewAsterisk(tableName)
	_tbUser.ID = field.NewInt64(tableName, "id")
	_tbUser.CreatedAt = field.NewTime(tableName, "created_at")
	_tbUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tbUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_tbUser.NickName = field.NewString(tableName, "nick_name")
	_tbUser.UserName = field.NewString(tableName, "user_name")
	_tbUser.Password = field.NewString(tableName, "password")
	_tbUser.Avatar = field.NewString(tableName, "avatar")
	_tbUser.Email = field.NewString(tableName, "email")
	_tbUser.Tel = field.NewString(tableName, "tel")
	_tbUser.Address = field.NewString(tableName, "address")
	_tbUser.Token = field.NewString(tableName, "token")
	_tbUser.IP = field.NewString(tableName, "ip")
	_tbUser.Role = field.NewInt64(tableName, "role")
	_tbUser.SignOrigin = field.NewInt64(tableName, "sign_origin")

	_tbUser.fillFieldMap()

	return _tbUser
}

type tbUser struct {
	tbUserDo tbUserDo

	ALL        field.Asterisk
	ID         field.Int64
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	NickName   field.String
	UserName   field.String
	Password   field.String
	Avatar     field.String
	Email      field.String
	Tel        field.String
	Address    field.String
	Token      field.String
	IP         field.String
	Role       field.Int64
	SignOrigin field.Int64

	fieldMap map[string]field.Expr
}

func (t tbUser) Table(newTableName string) *tbUser {
	t.tbUserDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbUser) As(alias string) *tbUser {
	t.tbUserDo.DO = *(t.tbUserDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbUser) updateTableName(table string) *tbUser {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")
	t.NickName = field.NewString(table, "nick_name")
	t.UserName = field.NewString(table, "user_name")
	t.Password = field.NewString(table, "password")
	t.Avatar = field.NewString(table, "avatar")
	t.Email = field.NewString(table, "email")
	t.Tel = field.NewString(table, "tel")
	t.Address = field.NewString(table, "address")
	t.Token = field.NewString(table, "token")
	t.IP = field.NewString(table, "ip")
	t.Role = field.NewInt64(table, "role")
	t.SignOrigin = field.NewInt64(table, "sign_origin")

	t.fillFieldMap()

	return t
}

func (t *tbUser) WithContext(ctx context.Context) ITbUserDo { return t.tbUserDo.WithContext(ctx) }

func (t tbUser) TableName() string { return t.tbUserDo.TableName() }

func (t tbUser) Alias() string { return t.tbUserDo.Alias() }

func (t tbUser) Columns(cols ...field.Expr) gen.Columns { return t.tbUserDo.Columns(cols...) }

func (t *tbUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbUser) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 15)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
	t.fieldMap["nick_name"] = t.NickName
	t.fieldMap["user_name"] = t.UserName
	t.fieldMap["password"] = t.Password
	t.fieldMap["avatar"] = t.Avatar
	t.fieldMap["email"] = t.Email
	t.fieldMap["tel"] = t.Tel
	t.fieldMap["address"] = t.Address
	t.fieldMap["token"] = t.Token
	t.fieldMap["ip"] = t.IP
	t.fieldMap["role"] = t.Role
	t.fieldMap["sign_origin"] = t.SignOrigin
}

func (t tbUser) clone(db *gorm.DB) tbUser {
	t.tbUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbUser) replaceDB(db *gorm.DB) tbUser {
	t.tbUserDo.ReplaceDB(db)
	return t
}

type tbUserDo struct{ gen.DO }

type ITbUserDo interface {
	gen.SubQuery
	Debug() ITbUserDo
	WithContext(ctx context.Context) ITbUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITbUserDo
	WriteDB() ITbUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITbUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITbUserDo
	Not(conds ...gen.Condition) ITbUserDo
	Or(conds ...gen.Condition) ITbUserDo
	Select(conds ...field.Expr) ITbUserDo
	Where(conds ...gen.Condition) ITbUserDo
	Order(conds ...field.Expr) ITbUserDo
	Distinct(cols ...field.Expr) ITbUserDo
	Omit(cols ...field.Expr) ITbUserDo
	Join(table schema.Tabler, on ...field.Expr) ITbUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITbUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITbUserDo
	Group(cols ...field.Expr) ITbUserDo
	Having(conds ...gen.Condition) ITbUserDo
	Limit(limit int) ITbUserDo
	Offset(offset int) ITbUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITbUserDo
	Unscoped() ITbUserDo
	Create(values ...*domain.TbUser) error
	CreateInBatches(values []*domain.TbUser, batchSize int) error
	Save(values ...*domain.TbUser) error
	First() (*domain.TbUser, error)
	Take() (*domain.TbUser, error)
	Last() (*domain.TbUser, error)
	Find() ([]*domain.TbUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.TbUser, err error)
	FindInBatches(result *[]*domain.TbUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.TbUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITbUserDo
	Assign(attrs ...field.AssignExpr) ITbUserDo
	Joins(fields ...field.RelationField) ITbUserDo
	Preload(fields ...field.RelationField) ITbUserDo
	FirstOrInit() (*domain.TbUser, error)
	FirstOrCreate() (*domain.TbUser, error)
	FindByPage(offset int, limit int) (result []*domain.TbUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITbUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tbUserDo) Debug() ITbUserDo {
	return t.withDO(t.DO.Debug())
}

func (t tbUserDo) WithContext(ctx context.Context) ITbUserDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbUserDo) ReadDB() ITbUserDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbUserDo) WriteDB() ITbUserDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbUserDo) Session(config *gorm.Session) ITbUserDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbUserDo) Clauses(conds ...clause.Expression) ITbUserDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbUserDo) Returning(value interface{}, columns ...string) ITbUserDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbUserDo) Not(conds ...gen.Condition) ITbUserDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbUserDo) Or(conds ...gen.Condition) ITbUserDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbUserDo) Select(conds ...field.Expr) ITbUserDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbUserDo) Where(conds ...gen.Condition) ITbUserDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbUserDo) Order(conds ...field.Expr) ITbUserDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbUserDo) Distinct(cols ...field.Expr) ITbUserDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbUserDo) Omit(cols ...field.Expr) ITbUserDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbUserDo) Join(table schema.Tabler, on ...field.Expr) ITbUserDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITbUserDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbUserDo) RightJoin(table schema.Tabler, on ...field.Expr) ITbUserDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbUserDo) Group(cols ...field.Expr) ITbUserDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbUserDo) Having(conds ...gen.Condition) ITbUserDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbUserDo) Limit(limit int) ITbUserDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbUserDo) Offset(offset int) ITbUserDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITbUserDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbUserDo) Unscoped() ITbUserDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbUserDo) Create(values ...*domain.TbUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbUserDo) CreateInBatches(values []*domain.TbUser, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbUserDo) Save(values ...*domain.TbUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbUserDo) First() (*domain.TbUser, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbUser), nil
	}
}

func (t tbUserDo) Take() (*domain.TbUser, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbUser), nil
	}
}

func (t tbUserDo) Last() (*domain.TbUser, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbUser), nil
	}
}

func (t tbUserDo) Find() ([]*domain.TbUser, error) {
	result, err := t.DO.Find()
	return result.([]*domain.TbUser), err
}

func (t tbUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.TbUser, err error) {
	buf := make([]*domain.TbUser, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbUserDo) FindInBatches(result *[]*domain.TbUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbUserDo) Attrs(attrs ...field.AssignExpr) ITbUserDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbUserDo) Assign(attrs ...field.AssignExpr) ITbUserDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbUserDo) Joins(fields ...field.RelationField) ITbUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbUserDo) Preload(fields ...field.RelationField) ITbUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbUserDo) FirstOrInit() (*domain.TbUser, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbUser), nil
	}
}

func (t tbUserDo) FirstOrCreate() (*domain.TbUser, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.TbUser), nil
	}
}

func (t tbUserDo) FindByPage(offset int, limit int) (result []*domain.TbUser, count int64, err error) {
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

func (t tbUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbUserDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbUserDo) Delete(models ...*domain.TbUser) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbUserDo) withDO(do gen.Dao) *tbUserDo {
	t.DO = *do.(*gen.DO)
	return t
}
