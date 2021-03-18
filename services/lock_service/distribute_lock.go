package lock_service

import (
	_ "context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/maybaby/gscheduler/pkg/setting"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanketplus/go-mysql-lock"
)

type Locker interface {
	InitLocker()
	GetLock(name string) *gomysqllock.Lock
	ReleaseLock(*gomysqllock.Lock) error
}
type MysqlLocker struct {
	db     *sql.DB
	locker *gomysqllock.MysqlLocker
}

func (m *MysqlLocker) InitLocker() {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	fmt.Println("初始化Locker中", url)
	db, _ := sql.Open(setting.DatabaseSetting.Type, url)
	locker := gomysqllock.NewMysqlLocker(db)
	m.db = db
	m.locker = locker
}

func (m *MysqlLocker) GetLock(name string) *gomysqllock.Lock {
	lock, _ := m.locker.Obtain(name)

	return lock
}

func (m *MysqlLocker) ReleaseLock(l *gomysqllock.Lock) error {
	if l == nil {
		return errors.New("Null")

	}
	l.Release()
	return nil
}

func GetAndInitLocker() Locker {
	var l Locker
	switch setting.DSlockerSetting.Type {
	case "mysql":
		l = &MysqlLocker{}
		l.InitLocker()
	default:
		l = &MysqlLocker{}
		l.InitLocker()
	}
	return l

}

func main() {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	fmt.Println(url)
	db, _ := sql.Open(setting.DatabaseSetting.Type, url)
	locker := gomysqllock.NewMysqlLocker(db)

	lock, _ := locker.Obtain("foo")
	fmt.Println("我拿到了锁", os.Args[0])
	time.Sleep(10 * time.Second)
	lock.Release()
	fmt.Println("我释放了锁", os.Args[0])

}
