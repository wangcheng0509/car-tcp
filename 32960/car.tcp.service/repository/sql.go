package repository

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"car.tcp.service/conf"
	"gorm.io/driver/clickhouse"

	"github.com/gohouse/gorose/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/wangcheng0509/gpkg/mysqlconn"
)

var (
	ServiceDBA       *Dba
	ClickHouseDbConf DataBase
	err              error
)

type (
	DataBase struct {
		mysqlconn.Database
		PortTcp  string
		PortHttp string
		DeBug    string
	}

	Dba struct {
		Mysql *gorm.DB
	}
)

var clickhouseEngine *gorose.Engin
var clickhouseOnce sync.Once

var ckDB *gorm.DB

func init() {
	ServiceDBA = &Dba{}
}

func RegistryMySQL() {
	// db configs
	mysqlDB := &DataBase{
		Database: mysqlconn.Database{
			Type:     conf.Conf.MySQL.Type,
			User:     conf.Conf.MySQL.User,
			Password: conf.Conf.MySQL.Password,
			Host:     conf.Conf.MySQL.Host,
			DBName:   conf.Conf.MySQL.DBName,
		},
		DeBug: conf.Conf.MySQL.Debug,
	}
	registry(*mysqlDB)
}

func registry(database DataBase) {
	switch database.Type {
	case `mysql`:
		ServiceDBA.mysqlConn(database)
	case `clickhouse`:
		ClickHouseDbConf = database
	default:

	}
}

// GetClickHouseDb 获取ClickHouse链接
func (d *Dba) GetClickHouseDb() *gorm.DB {

	clickhouseOnce.Do(func() {
		var err error
		ckDB, err = openClickDB()
		if err != nil {
			log.Fatalf("cannot open clickhouse:%v\n", err)
		}
		db, err := ckDB.DB()
		if err != nil {
			log.Fatalf("cannot open clickhouse:%v\n", err)
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		// 设置可空闲的最大连接数，随时等待调用
		db.SetMaxIdleConns(2)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		// 设置连接池的最大连接数，不配置，默认为 0，就是不限制
		db.SetMaxOpenConns(10)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		// 连接的最长存活期，超过这个时间，连接将会重置，不再被复用，不配置默认就是永不过期
		db.SetConnMaxLifetime(time.Second * 600)

	})
	return ckDB
}

func openClickDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("clickhouse://%s:%s@%s:%s/%s?debug=%s",
		ClickHouseDbConf.User,
		ClickHouseDbConf.Password,
		ClickHouseDbConf.Host,
		ClickHouseDbConf.PortTcp,
		ClickHouseDbConf.DBName,
		ClickHouseDbConf.DeBug,
	)
	conn, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func openDB(conf DataBase) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DBName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("cannot open mysql:%v\n", err)
	}

	return db
}

func (d *Dba) mysqlConn(database DataBase) *gorm.DB {
	if database.DeBug == "true" {
		d.Mysql = openDB(database).Debug()
	} else {
		d.Mysql = openDB(database)
	}
	sqlDB, err := d.Mysql.DB()
	if err != nil {
		log.Fatalf("can not use sql set:%v", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
		d.Mysql = openDB(database)
	}

	return d.Mysql
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

// Tabler 获取表名
type Tabler interface {
	TableName() string
}

// getDBWithTable 获取指定struct表名的db
func getDBWithTable(ctx context.Context, taber Tabler) *gorm.DB {
	return getDB(ctx, ServiceDBA.Mysql).Table(taber.TableName())
}

// NewTransRepo 事务repo
func NewTransRepo() *TransRepo {
	return &TransRepo{
		DB: ServiceDBA.Mysql,
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
