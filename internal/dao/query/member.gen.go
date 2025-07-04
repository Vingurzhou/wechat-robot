// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"bot/internal/dao/model"
)

func newMember(db *gorm.DB, opts ...gen.DOOption) member {
	_member := member{}

	_member.memberDo.UseDB(db, opts...)
	_member.memberDo.UseModel(&model.Member{})

	tableName := _member.memberDo.TableName()
	_member.ALL = field.NewAsterisk(tableName)
	_member.ID = field.NewInt32(tableName, "id")
	_member.Wxid = field.NewString(tableName, "wxid")
	_member.NickName = field.NewString(tableName, "nick_name")
	_member.InviterUserName = field.NewString(tableName, "inviter_user_name")
	_member.MemberFlag = field.NewInt32(tableName, "member_flag")
	_member.DisplayName = field.NewString(tableName, "display_name")
	_member.BigHeadImgURL = field.NewString(tableName, "big_head_img_url")
	_member.SmallHeadImgURL = field.NewString(tableName, "small_head_img_url")
	_member.ChatroomID = field.NewString(tableName, "chatroom_id")
	_member.IsOwner = field.NewBool(tableName, "is_owner")
	_member.IsAdmin = field.NewBool(tableName, "is_admin")

	_member.fillFieldMap()

	return _member
}

type member struct {
	memberDo memberDo

	ALL             field.Asterisk
	ID              field.Int32
	Wxid            field.String
	NickName        field.String
	InviterUserName field.String
	MemberFlag      field.Int32
	DisplayName     field.String
	BigHeadImgURL   field.String
	SmallHeadImgURL field.String
	ChatroomID      field.String
	IsOwner         field.Bool
	IsAdmin         field.Bool

	fieldMap map[string]field.Expr
}

func (m member) Table(newTableName string) *member {
	m.memberDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m member) As(alias string) *member {
	m.memberDo.DO = *(m.memberDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *member) updateTableName(table string) *member {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.Wxid = field.NewString(table, "wxid")
	m.NickName = field.NewString(table, "nick_name")
	m.InviterUserName = field.NewString(table, "inviter_user_name")
	m.MemberFlag = field.NewInt32(table, "member_flag")
	m.DisplayName = field.NewString(table, "display_name")
	m.BigHeadImgURL = field.NewString(table, "big_head_img_url")
	m.SmallHeadImgURL = field.NewString(table, "small_head_img_url")
	m.ChatroomID = field.NewString(table, "chatroom_id")
	m.IsOwner = field.NewBool(table, "is_owner")
	m.IsAdmin = field.NewBool(table, "is_admin")

	m.fillFieldMap()

	return m
}

func (m *member) WithContext(ctx context.Context) *memberDo { return m.memberDo.WithContext(ctx) }

func (m member) TableName() string { return m.memberDo.TableName() }

func (m member) Alias() string { return m.memberDo.Alias() }

func (m member) Columns(cols ...field.Expr) gen.Columns { return m.memberDo.Columns(cols...) }

func (m *member) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *member) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 11)
	m.fieldMap["id"] = m.ID
	m.fieldMap["wxid"] = m.Wxid
	m.fieldMap["nick_name"] = m.NickName
	m.fieldMap["inviter_user_name"] = m.InviterUserName
	m.fieldMap["member_flag"] = m.MemberFlag
	m.fieldMap["display_name"] = m.DisplayName
	m.fieldMap["big_head_img_url"] = m.BigHeadImgURL
	m.fieldMap["small_head_img_url"] = m.SmallHeadImgURL
	m.fieldMap["chatroom_id"] = m.ChatroomID
	m.fieldMap["is_owner"] = m.IsOwner
	m.fieldMap["is_admin"] = m.IsAdmin
}

func (m member) clone(db *gorm.DB) member {
	m.memberDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m member) replaceDB(db *gorm.DB) member {
	m.memberDo.ReplaceDB(db)
	return m
}

type memberDo struct{ gen.DO }

func (m memberDo) Debug() *memberDo {
	return m.withDO(m.DO.Debug())
}

func (m memberDo) WithContext(ctx context.Context) *memberDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m memberDo) ReadDB() *memberDo {
	return m.Clauses(dbresolver.Read)
}

func (m memberDo) WriteDB() *memberDo {
	return m.Clauses(dbresolver.Write)
}

func (m memberDo) Session(config *gorm.Session) *memberDo {
	return m.withDO(m.DO.Session(config))
}

func (m memberDo) Clauses(conds ...clause.Expression) *memberDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m memberDo) Returning(value interface{}, columns ...string) *memberDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m memberDo) Not(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m memberDo) Or(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m memberDo) Select(conds ...field.Expr) *memberDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m memberDo) Where(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m memberDo) Order(conds ...field.Expr) *memberDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m memberDo) Distinct(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m memberDo) Omit(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m memberDo) Join(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m memberDo) LeftJoin(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m memberDo) RightJoin(table schema.Tabler, on ...field.Expr) *memberDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m memberDo) Group(cols ...field.Expr) *memberDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m memberDo) Having(conds ...gen.Condition) *memberDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m memberDo) Limit(limit int) *memberDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m memberDo) Offset(offset int) *memberDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m memberDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *memberDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m memberDo) Unscoped() *memberDo {
	return m.withDO(m.DO.Unscoped())
}

func (m memberDo) Create(values ...*model.Member) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m memberDo) CreateInBatches(values []*model.Member, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m memberDo) Save(values ...*model.Member) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m memberDo) First() (*model.Member, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) Take() (*model.Member, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) Last() (*model.Member, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) Find() ([]*model.Member, error) {
	result, err := m.DO.Find()
	return result.([]*model.Member), err
}

func (m memberDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Member, err error) {
	buf := make([]*model.Member, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m memberDo) FindInBatches(result *[]*model.Member, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m memberDo) Attrs(attrs ...field.AssignExpr) *memberDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m memberDo) Assign(attrs ...field.AssignExpr) *memberDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m memberDo) Joins(fields ...field.RelationField) *memberDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m memberDo) Preload(fields ...field.RelationField) *memberDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m memberDo) FirstOrInit() (*model.Member, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) FirstOrCreate() (*model.Member, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Member), nil
	}
}

func (m memberDo) FindByPage(offset int, limit int) (result []*model.Member, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m memberDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m memberDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m memberDo) Delete(models ...*model.Member) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *memberDo) withDO(do gen.Dao) *memberDo {
	m.DO = *do.(*gen.DO)
	return m
}
