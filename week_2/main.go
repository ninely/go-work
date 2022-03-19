package main

import (
	"fmt"
	"log"
)

// 题目：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 答：不应该Wrap抛给上层，sql.ErrNoRows 是 DAO 层错误码，指未查询到数据，应该降级处理，避免上层对 DAO 层的依赖。
func main() {
	data := NewData()
	repo := NewUserRepo(data)
	uc := NewUserUC(repo)
	queryName := "Tom"
	user, err := uc.FindUser(queryName)
	if err != nil {
		panic(err)
	}

	if user == nil {
		log.Printf("user not found for name: %v", queryName)
		return
	}

	fmt.Println("find the user!")
}
