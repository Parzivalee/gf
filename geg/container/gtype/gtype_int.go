package main

import (
    "gitee.com/johng/gf/g/container/gtype"
    "fmt"
)

func main() {
    // 创建一个Int型的并发安全基本类型对象
    i := gtype.NewInt()

    // 设置值
    i.Set(10)

    // 获取值
    fmt.Println(i.Val())

    // (整型/浮点型有效)数值 增加/删除 delta
    // 返回修改之前的数值
    fmt.Println(i.Add(-1))

    fmt.Println(i.Val())
}
