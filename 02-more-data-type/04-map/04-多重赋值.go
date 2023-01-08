package main

import (
	"fmt"
	"strconv"
)

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func main() {
	offset := timeZone["EST"]

	println(offset)

	//若试图通过映射中不存在的键来取值，就会返回与该映射中项的类型对应的零值。
	//例如，若某个映射包含整数，当查找一个不存在的键时会返回 0。
	//集合可实现成一个值类型为 bool 的映射。将该映射中的项置为 true 可将该值放入集合中，此后通过简单的索引操作即可判断是否存在。
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	person := "Mari"
	if !attended[person] { // 若某人不在此映射中，则为 false.此false实际为bool类型的零值.
		fmt.Println(person, "正在开会")
	}

	//有时你需要区分某项是不存在还是其值为零值。如对于一个值本应为零的 "UTC" 条目，也可能是由于不存在该项而得到零值。
	//你可以使用多重赋值的形式来分辨这种情况。
	seconds, ok := timeZone["ENC"]
	if ok {
		println(seconds)
	} else {
		println("未找到时间格式", "零值: "+strconv.Itoa(seconds))
	}

	//	若仅需判断映射中是否存在某项而不关心实际的值，可使用空白标识符 （_）来代替该值的一般变量。
	//_, present := timeZone[tz]

	//要删除映射中的某项，可使用内建函数 delete，它以映射及要被删除的键为实参。 即便对应的键不在该映射中，此操作也是安全的。
	delete(timeZone, "PDT") // 现在用标准时间
}
