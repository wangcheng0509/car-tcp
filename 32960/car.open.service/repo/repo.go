package repo

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"car.open.service/conf"
	"car.open.service/ginx/common"

	"gorm.io/driver/clickhouse"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	ckDB, mysqlDB *gorm.DB
)

// Init 初始化数据库
func Init() {
	// getCkDB()
	getMySQLDB()
}

// getCkDB 获取ClickHouse链接
func getCkDB() *gorm.DB {
	if ckDB != nil {
		return ckDB
	}

	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%s/%s?dial_timeout=10s&read_timeout=20s&debug=%s",
		conf.Conf.CK.User,
		conf.Conf.CK.Password,
		conf.Conf.CK.Host,
		conf.Conf.CK.PortTCP,
		conf.Conf.CK.DBName,
		conf.Conf.CK.DeBug,
	)

	var err error
	ckDB, err = openDB(clickhouse.New(clickhouse.Config{
		DSN: dsn,
	}))
	// ckDB, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot open clickhouse:%v\n", err)
	}

	return ckDB
}

func getMySQLDB() *gorm.DB {
	if mysqlDB != nil {
		return mysqlDB
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Conf.MySQL.User,
		conf.Conf.MySQL.Password,
		conf.Conf.MySQL.Host,
		conf.Conf.MySQL.DBName,
	)

	var err error
	mysqlDB, err = openDB(mysql.New(mysql.Config{
		DSN: dsn,
	}))
	if err != nil {
		log.Fatalf("cannot open clickhouse:%v\n", err)
	}

	return mysqlDB
}

func openDB(dialector gorm.Dialector) (db *gorm.DB, err error) {
	db, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("cannot open db:%v\n", err)
		return
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Fatalf("cannot open db:%v\n", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// 设置可空闲的最大连接数，随时等待调用
	sqldb.SetMaxIdleConns(2)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	// 设置连接池的最大连接数，不配置，默认为 0，就是不限制
	sqldb.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// 连接的最长存活期，超过这个时间，连接将会重置，不再被复用，不配置默认就是永不过期
	sqldb.SetConnMaxLifetime(time.Second * 600)

	return
}

// WrapPageQuery 包装成带有分页的查询
func WrapPageQuery(db *gorm.DB, pp common.PaginationParam, out interface{}) (*common.PaginationResult, error) {
	if pp.OnlyCount {
		var count int64
		err := db.Count(&count).Error
		if err != nil {
			return nil, err
		}
		return &common.PaginationResult{Total: int(count)}, nil
	} else if !pp.Pagination {
		err := db.Find(out).Error
		return nil, err
	}
	total, err := findPage(db, pp, out)
	if err != nil {
		return nil, err
	}
	return &common.PaginationResult{
		Total:    total,
		Current:  pp.Current,
		PageSize: pp.PageSize,
	}, nil
}

func findPage(db *gorm.DB, pp common.PaginationParam, out interface{}) (int, error) {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	} else if count == 0 {
		return int(count), nil
	}
	current, pageSize := int(pp.Current), int(pp.PageSize)
	if current > 0 && pageSize > 0 {
		db = db.Offset((current - 1) * pageSize).Limit(pageSize)
	} else if pageSize > 0 {
		db = db.Limit(pageSize)
	}

	err = db.Find(out).Error
	return int(count), err
}

// TransRepo 事务
type TransRepo struct {
	DB *gorm.DB
}

type (
	transCtx struct{}
)

// NewTrans 新建包含事务的上下文
func NewTrans(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, transCtx{}, db)
}

// FromTrans 从上下文获取事务
func FromTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

// Exec 事务执行
func (a *TransRepo) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := FromTrans(ctx); ok {
		return fn(ctx)
	}

	return a.DB.Transaction(func(db *gorm.DB) error {
		return fn(NewTrans(ctx, db))
	})
}

func getDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := FromTrans(ctx)
	if ok {
		db, ok := trans.(*gorm.DB)
		if ok {
			return db
		}
	}

	return defDB
}

// getDBWithTable 获取指定struct表名的db
func getDBWithTable(ctx context.Context, taber schema.Tabler) *gorm.DB {
	return getDB(ctx, mysqlDB).Table(taber.TableName())
}

// getCkDBWithTable 获取指定struct表名的db
func getCkDBWithTable(ctx context.Context, taber schema.Tabler) *gorm.DB {
	return getDB(ctx, ckDB).Table(taber.TableName())
}

// NewTransRepo 事务repo
func NewTransRepo() *TransRepo {
	return &TransRepo{
		DB: mysqlDB,
	}
}

// WrapCreate wrap create
func WrapCreate(db *gorm.DB, v interface{}) error {
	if l, ok := sliceLen(v); ok && l == 0 {
		return nil
	}

	return db.Create(v).Error
}

// sliceLen ..
func sliceLen(v interface{}) (int, bool) {
	rv := reflect.ValueOf(v)
	val, ok := isslice(rv)
	if ok {
		return val.Len(), ok
	}

	return 0, ok
}

func isslice(rv reflect.Value) (reflect.Value, bool) {
	if k := rv.Kind(); k == reflect.Array || k == reflect.Slice {
		return rv, true
	}

	if rv.Kind() == reflect.Ptr && rv.CanAddr() {
		return isslice(rv.Addr())
	} else if rv.Kind() == reflect.Ptr && !rv.CanAddr() {
		return isslice(rv.Elem())
	}

	if rv.Kind() == reflect.Pointer && rv.CanAddr() {
		return isslice(rv.Addr())
	}

	return rv, false
}
