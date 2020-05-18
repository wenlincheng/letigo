package util

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// 生成随机编号
func GenRandNo(uid int64) string {
	t := time.Now()
	// 年月日
	p1 := t.Format("20060102")
	// linux环境下，time.Parse()的默认时区是UTC，time.Format()的时区默认是本地
	location, _ := time.LoadLocation("UTC")
	utct := t.UTC().Format("20060102")
	t1, _ := time.ParseInLocation("20060102", utct, location)

	timeNumber := t1.Unix() * 1e3
	p2 := t.UnixNano()/1e6 - timeNumber
	// 年月日+
	prefix := p1 + strconv.FormatInt(p2, 10) + strconv.FormatInt(uid, 10)
	seed := time.Now().UnixNano() + uid*1e12
	source := rand.NewSource(seed)
	r := rand.New(source)
	randomIntLen := 26 - len(prefix)
	pow10 := math.Pow10(randomIntLen)
	n := r.Int63n(int64(pow10))
	strFormat := fmt.Sprintf("%%0%dd", randomIntLen)
	return prefix + fmt.Sprintf(strFormat, n)
}

// 生成随机数字
func GenRandNumber() int64 {

	return 0
}
