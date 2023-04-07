package main

import "fmt"

func main() {
    // 创建一个空 map
    m := make(map[string]int)

    // 添加键值对
    m["apple"] = 1
    m["banana"] = 2
    m["orange"] = 3

    // 获取键值对
    fmt.Println(m["apple"]) // 输出：1

    // 删除键值对
    delete(m, "banana")

    // 遍历 map
    for k, v := range m {
        fmt.Printf("%s -> %d\n", k, v)
    }
}