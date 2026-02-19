package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	unix := timeObj.Unix()        // 1970年1月1日到现在的秒数
	milli := timeObj.UnixMilli()  // 毫秒数
	nano := timeObj.UnixNano()    // 纳秒数
	timeNow := time.Unix(0, nano) // 反向推导正常日期
	timeNow.Format("2006/01/02 03:04:05")
	str := timeObj.Format("2006-01-02 03:04:05")
	fmt.Println(str, unix, milli, nano, timeNow)
}

/*

	timeObj.Year()
	timeObj.Month()
	timeObj.Day()
	timeObj.Hour()
	timeObj.Minute()
	timeObj.Second()

*/
