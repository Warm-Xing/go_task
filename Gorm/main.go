// 数据库驱动作业参考链接：https://github.com/Avrinkong/web3Study/blob/main/taskThree/main.go
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

type Account struct {
	ID      int
	Balance float64
}
type Transaction struct {
	ID            int
	FromAccountId int
	ToAccountId   int
	Amount        float64
}

var db *gorm.DB

func InitDB() *gorm.DB {
	var err error
	db, err = gorm.Open(mysql.Open("root:Wam1992-7414@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("sorry，failed to connect database")
		panic(err)
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	return db
}

func One() {
	db.AutoMigrate(&Students{})
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//student := Students{Name: "张1", Age: 14, Grade: "三年级"}
	//result := db.Create(&student)
	//fmt.Printf("插入记录ID: %d, 影响行数: %d, 错误: %v, 姓名: %s", student.ID, result.RowsAffected, result.Error, student.Name)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	//var student18 []Students
	//db.Where("age > ?", 18).Find(&student18)
	//fmt.Printf("年龄大于18的学生: %+v", student18)

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//updateResult := db.Model(&Students{}).Where("name = ?", "张三").Update("grade", "四年级")
	//fmt.Printf("更新影响行数: %d, 错误: %v\n", updateResult.RowsAffected, updateResult.Error)
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	deleteResult := db.Unscoped().Where("age < ?", 15).Delete(&Students{})
	fmt.Printf("更新影响行数: %d, 错误: %v\n", deleteResult.RowsAffected, deleteResult.Error)
}

//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求 ：
//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

//func Two() {
//
//	err := db.AutoMigrate(&Account{}, &Transaction{}) //一次调用多表迁移，调用有误整体多表回滚。
//	if err != nil {
//		return
//	}
//	// 2. 事务处理（使用tx对象）
//	err1 := db.Transaction(func(tx *gorm.DB) error {
//		var accountA, accountB Account
//		tx.Create(&Account{ID: 1, Balance: 50})
//		tx.Create(&Account{ID: 2, Balance: 1000})
//		tx.Where("id = ?", 1).First(&accountA)
//		tx.Where("id = ?", 2).First(&accountB)
//		if accountA.Balance < 100 {
//			return fmt.Errorf("账户余额不足")
//		}
//		accountA.Balance -= 100
//		accountB.Balance += 100
//		tx.Save(&accountA)
//		tx.Save(&accountB)
//		return nil
//	})
//	if err1 != nil {
//		fmt.Println("事务错误:", err1) // 打印错误信息
//		return
//	}
//
//}

// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

//func Three() {
//
//	err := db.AutoMigrate(&Employee{}) //一次调用多表迁移，调用有误整体多表回滚。
//	if err != nil {
//		return
//	}
//
//	//employees := Employee{Name: "张3", Department: "技术部", Salary: 500.3}
//	//result := db.Create(&employees)
//	//fmt.Printf("影响行数: %d, 错误: %v", result.RowsAffected, result.Error)
//	var employees []Employee
//
//	db.Where("department = ?", "技术部").Find(&employees)
//	if len(employees) == 0 {
//		fmt.Println("技术部暂无员工")
//	}
//	fmt.Printf("技术部员工信息: %+v\n", employees)
//
//	var employee Employee
//	db.Order("salary DESC").Limit(1).Find(&employee)
//	fmt.Printf("工资最高的员工信息: %+v\n", employee)
//}

// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

//func Four() {
//	db.AutoMigrate(&Book{})
//	err := db.AutoMigrate(&Employee{}) //一次调用多表迁移，调用有误整体多表回滚。
//	if err != nil {
//		return
//	}
//
//	//books := Book{Title: "清华大学出版社", Author: "刘二小", Price: 65}
//	//result := db.Create(&books)
//	//fmt.Printf("影响行数: %d, 错误: %v", result.RowsAffected, result.Error)
//	var books []Book
//
//	db.Where("Price > ?", "50").Find(&books)
//	if len(books) == 0 {
//		fmt.Println("走吧没有你要的书")
//		return
//	}
//	fmt.Printf("恭喜你有你要的书: %+v\n", books)
//}

// User 模型定义
type User struct {
	gorm.Model
	Username      string `gorm:"size:50;not null;unique"`  // 用户名  //unique避免出现重复的用户名或邮箱
	Email         string `gorm:"size:100;not null;unique"` // 邮箱
	Posts         []Post `gorm:"foreignKey:UserID"`        // 一对多关联Post
	ArticleCount  int    `gorm:"default:0"`                // 新增：文章数量统计
	ArticleCount2 int    `gorm:"default:0"`                // 新增：文章数量统计测试
}

// Post 模型定义
type Post struct {
	gorm.Model
	Title         string    `gorm:"size:200;not null"` // 文章标题
	Content       string    `gorm:"type:text"`         // 文章内容
	UserID        uint      `gorm:"not null"`          // 外键关联User
	Comments      []Comment `gorm:"foreignKey:PostID"` // 一对多关联Comment
	CommentStatus string    `gorm:"default:'有评论'"`     // 新增：评论状态
}

// Comment 模型定义
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"` // 评论内容
	PostID  uint   `gorm:"not null"`           // 外键关联Post
}

func Five() {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{}) //一次调用多表迁移，调用有误整体多表回滚。
	if err != nil {
		panic("failed to connect database")
	}

	// 创建用户（补充Email字段，确保用户名唯一）
	db.Create(&User{
		Username: "小王1",
		Email:    "wang1@example.com", // 补充非空Email
	})
	db.Create(&User{
		Username: "小王2",
		Email:    "wang2@example.com",
	})
	db.Create(&User{
		Username: "小王3",
		Email:    "wang3@example.com",
	})
	db.Create(&User{
		Username: "小王4", // 修改重复用户名
		Email:    "wang4@example.com",
	})
	db.Create(&User{
		Username: "小王5", // 修改重复用户名
		Email:    "wang5@example.com",
	})
	db.Create(&User{
		Username: "小王6", // 修改重复用户名
		Email:    "wang6@example.com",
	})

	// 创建文章（关联正确的UserID）
	db.Create(&Post{
		Title:   "《Go 语言基础》",
		Content: "Go 语言基础",
		UserID:  1, // 关联小王1
	})
	db.Create(&Post{
		Title:   "《Go 语言进阶》",
		Content: "Go 语言进阶",
		UserID:  1, // 关联小王1
	})
	db.Create(&Post{
		Title:   "《Go 语言实战》",
		Content: "Go 语言实战",
		UserID:  1, // 关联小王1
	})
	db.Create(&Post{
		Title:   "《Go 语言微服务》",
		Content: "Go 语言微服务",
		UserID:  2, // 关联小王2
	})
	db.Create(&Post{
		Title:   "《Go 语言微服务》",
		Content: "Go 语言微服务",
		UserID:  3, // 关联小王3
	})
	db.Create(&Post{
		Title:   "《Go 语言微服务》",
		Content: "Go 语言微服务",
		UserID:  4, // 关联小王4
	})

	// 创建评论（修复PostID关联错误）
	db.Create(&Comment{
		Content: "Go 语言基础不错",
		PostID:  1, // 关联文章1
	})
	db.Create(&Comment{
		Content: "Go 语言进阶不错",
		PostID:  2, // 关联文章2
	})
	db.Create(&Comment{
		Content: "Go 语言实战不错",
		PostID:  3, // 关联文章3
	})
	db.Create(&Comment{
		Content: "Go 语言微服务不错",
		PostID:  4, // 关联文章4
	})
	db.Create(&Comment{
		Content: "Go 语言微服务不错",
		PostID:  4, // 修正原PostID=5的错误，关联文章4
	})
	db.Create(&Comment{
		Content: "Go 语言微服务不错",
		PostID:  4, // 关联文章4
	})

}

// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func Six() {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{}) //一次调用多表迁移，调用有误整体多表回滚。
	if err != nil {
		panic("failed to connect database")
	}

	//var user User
	//db.Preload("Posts").Preload("Posts.Comments").Where("id = ?", 1).Find(&user)
	////db.Where("id = ?", 1).Find(&user)像这种情况就打印不出来
	//fmt.Printf("用户%s的博客文章和评论信息: %v\n", user.Username, user.Posts)
	var postStats Post
	db.Model(&Post{}). // 操作主体：Post表
				Select("posts.*, COUNT(comments.id) as comments_count").    // 返回 posts 表的所有字段（文章ID、标题、内容等），COUNT(comments.id) as comments_count：统计每个文章的评论数量，并重命名为 comments_count
				Joins("left join comments on posts.id = comments.post_id"). // 关联表，左连接，值保post表，comment只匹配
				Group("posts.id").                                          // 分组条件
				Order("comments_count DESC").                               // 排序规则
				Limit(1).                                                   // 结果限制
				Scan(&postStats)                                            // 结果映射

	var postMost Post
	db.Preload("Comments").First(&postMost, postStats.ID) // 根据ID查询完整数据
	fmt.Printf("评论数量最多的文章信息: %+v\n", postMost)
}

// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
// AfterCreate 钩子：文章创建后更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户表的文章计数（+1）
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("article_count", gorm.Expr("article_count + ?", 1)).Error
}

// AfterDelete 钩子：评论删除后检查文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 1. 查询当前文章的剩余评论数
	var count int64
	tx.Debug().Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	// 2. 若评论数为0，更新文章状态
	if count == 0 {
		return tx.Debug().Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

func DeleteComment(db *gorm.DB, commentID uint) error {
	// 步骤1：查询完整评论对象（填充c.PostID）
	var comment Comment
	if err := db.Unscoped().First(&comment, commentID).Error; err != nil {
		return err
	}
	// 步骤2：删除评论（此时钩子中的c.PostID已被正确赋值）
	return db.Unscoped().Delete(&comment).Error
}

func main() {
	InitDB()
	//One()
	//Two()
	//Three()
	//Four()
	//Five()
	//Six()
	// 创建文章后，用户article_count会自动+1
	//db.Create(&Post{Title: "测试1", Content: "内容1", UserID: 1})

	// 删除最后一条评论后，文章comment_status会变为"无评论"
	//db.Delete(&Comment{}, 1) // 假设这是文章的最后一条评论
	//db.Unscoped().Delete(&Comment{}, 3) //永久删除
	//db.Preload("Posts").Unscoped().Delete(&Comment{}, 6)
	DeleteComment(db, 8) // 这样才能正确出发

}
