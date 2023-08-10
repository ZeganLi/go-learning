package main

import (
	"fmt"
	"time"
)

const GoTime = "2006-01-02 15:04:05"

func main() {

	time1 := time.Now()
	testTime()
	time2 := time.Now()
	fmt.Println(time2.Sub(time1).Seconds())
}

func testTime() {
	t := time.Now()

	fmt.Println("1. ", t)
	fmt.Println("2. ", t.Local())
	fmt.Println("2. ", t.UTC())

	t = time.Date(2023, time.December, 13, 9, 23, 0, 0, time.Local)
	fmt.Printf("4. 本地时间 %s ，国际统一时间：%s \n", t, t.UTC())

	t, _ = time.Parse(GoTime, "2023-04-13 09:26:53")
	fmt.Println("5. ", t)

	fmt.Println("6. ", time.Now().Format(GoTime))
	fmt.Println("7. ", time.Now().String())
	fmt.Println("8. ", time.Now().Unix())
	fmt.Println("9. ", time.Now().UnixNano())
	fmt.Println("10. ", t.Equal(time.Now()))
	fmt.Println("11. ", t.Before(time.Now()))
	fmt.Println("12. ", t.After(time.Now()))

	year, month, day := time.Now().Date()
	fmt.Println("13. ", year, month, day)
	fmt.Println("14. ", time.Now().Year())
	fmt.Println("15. ", time.Now().Month())
	fmt.Println("16. ", time.Now().Day())
	fmt.Println("17. ", time.Now().Weekday())

	hour, min, sec := time.Now().Clock()
	fmt.Println("18. ", hour, min, sec)
	fmt.Println("19. ", time.Now().Hour())
	fmt.Println("20. ", time.Now().Minute())
	fmt.Println("21. ", time.Now().Second())
	fmt.Println("22. ", time.Now().Nanosecond())

	fmt.Println("23. ", time.Now().Sub(time.Now()))
	fmt.Println("24. ", time.Now().Sub(time.Now()).Hours())
	fmt.Println("25. ", time.Now().Sub(time.Now()).Minutes())
	fmt.Println("26. ", time.Now().Sub(time.Now()).Seconds())
	fmt.Println("27. ", time.Now().Sub(time.Now()).Nanoseconds())

	fmt.Println("28. ", "时间间距：", t.Sub(time.Now()).String())
	d, _ := time.ParseDuration("1h30m")
	fmt.Println("29. ", d)
	fmt.Println("30. ", "交卷时间", time.Now().Add(d))
	fmt.Println("31. ", "一年一个月一天之后的日期：", time.Now().AddDate(1, 1, 1))
}
