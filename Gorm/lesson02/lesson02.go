package lesson02

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string //`gorm:"type:varchar(100);uniqueIndex:idx_name"`
	Age      int
	Birthday time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) { //BeforeCreate 是 GORM 框架提供的创建前钩子（Hook）方法，当调用 GORM 的 Create 方法插入新记录时，会自动触发该方法
	u.Name = u.Name + "_12123"
	return
}

func (u *BankUser) BeforeCreate(tx *gorm.DB) (err error) { //BeforeCreate 是 GORM 框架提供的创建前钩子（Hook）方法，当调用 GORM 的 Create 方法插入新记录时，会自动触发该方法
	u.Name = u.Name + "_12123"
	return
}

type CreditCard struct {
	gorm.Model
	Number     string
	BankUserID uint
}

type BankUser struct {
	gorm.Model //嵌套 gorm.Model 后，结构体自动拥有以上4个字段，无需手动定义： 其中包含ID
	Name       string
	CreditCard CreditCard
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&BankUser{})
	db.AutoMigrate(&CreditCard{})

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}  //初始化类型，后续见下一行用指针也可以
	//
	//result := db.Create(&user) // 通过数据的指针来创建
	//fmt.Println(result.RowsAffected)
	//// 检查保存是否成功
	//if result.Error != nil {
	//	panic("保存失败: " + result.Error.Error()) // 打印错误原因
	//}
	//user.ID             // 返回插入数据的主键
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数
	//
	//users := []*User{ //果需要存储多个 User 实例，应该使用切片 []*User{} 而非单个指针
	//	{Name: "xingwen", Age: 18, Birthday: time.Now()},
	//	{Name: "wangyuxin", Age: 19, Birthday: time.Now()},
	//}
	//
	//db.Create(users)
	//fmt.Println(users[0])

	//db.Create(&BankUser{
	//	Name:       "test",
	//	CreditCard: CreditCard{Number: "42088119920114"},
	//})

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user) //当插入记录时，如果违反唯一键约束（如重复的 name 字段），数据库会忽略这条插入操作，不会报错也不会修改现有数据，现实是只有主键ID是唯一的，会不断插入数据
	//user := User{Name: "Jinzhu", Age: 22, Birthday: time.Now()}
	//user.ID = 1  // 这里是定死了，只有去update，此时名字一样，年纪才去更新。
	//db.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "name"}},           //两行一起用 ① 冲突检测字段
	//	DoUpdates: clause.AssignmentColumns([]string{"age"}), // ② 冲突时更新的字段
	//}).Create(&user) // ③ 插入的对象

	//var user User
	//db.Debug().First(&user)             // 查询表中第一条记录（等价于LIMIT 1 ORDER BY id ASC），输出会包含debug模式
	//err := db.Debug().Take(&user).Error //查询表中任意一条记录（不指定排序，效率略高于First）
	//if err != nil {
	//	panic(err)
	//}

	//user.ID = 1
	//db.Debug().First(&user, []int{1, 2, 3})
	//db.Debug().Where("id IN (?)", []int{1, 2, 3}).First(&user)
	//db.Debug().Where("id IN (?)", []int{1, 2, 3}).First(&user)

	//var users []User
	//db.Debug().Find(&users)

	var user User
	////user.ID = 1
	//db.Debug().First(&user, "name = ?", "Jinzhu_12123") //查询name="123123"的第一条记录
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	//db.Save(&user)
	//db.Unscoped().Delete(&user)
}
