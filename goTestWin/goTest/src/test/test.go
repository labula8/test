
package main

//import ("fmt";"os")
//import "fmt"

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"strings"
	"runtime"
	"math"
	"strconv"
	"unsafe"
)

//type float float64
type float float32
type T struct{}
type Color int

const (
	Red Color = iota 	//0
	Green				//1
	Blue				//2
)

const MAX_SIZE int = 100

func init() {
	var s1 = runtime.GOOS
	fmt.Printf("s1=%s \n", s1)
	funcName, file, line, ok := runtime.Caller(0)
	if (ok) {
		fmt.Printf("func name: %s, file=%s, line=%d\n", runtime.FuncForPC(funcName).Name(), file, line)
	}
}

func t_slice() {
	var list1 = [5] int{1, 2, 3, 4, 8}
	fmt.Println("list1=", list1)
	var list2 = list1[1:]
	fmt.Println("list1[1:]=",list2)
	var list3 =list1[:2]
	fmt.Println("list1[:2]=", list3)
	fmt.Println("list1[1:3]=", list1[1:3])
}

func t_sizeof() {
	var a32 int32 = 32
	var a64 int64 = 64
	var a int = 100
	fmt.Printf("sizeof(a32)=%d, sizeof(a64)=%d\n", unsafe.Sizeof(a32), unsafe.Sizeof(a64))
	fmt.Println("sizeof(a)=", unsafe.Sizeof(a))
}

func t_range() {
	str := "go hello world!"
	fmt.Printf("str=%s, len=%d \n", str, len(str))
	for pos, ch := range str {
		fmt.Printf("[%d]=%c\n", pos, ch)
		ch = 'o'
	}
	fmt.Printf("str=%s, len=%d \n", str, len(str))
}

func t_strconv() {
	fmt.Printf("the size of int is: %d \n", strconv.IntSize)
	var s1 string = "a88"
	n_s1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Printf("n_s1=%d, err=%s \n", n_s1, err)
		return
	} else {
		fmt.Println(err)
		fmt.Printf("n_s1=%d \n", n_s1)

		s2 := strconv.Itoa(n_s1)
		fmt.Printf("s2=%s \n", s2)
	}
}

func t_math() {
	var a = int64(math.Abs(-3))
	fmt.Printf("a=%d \n", a)
	var pow1 = math.Pow(3, 3)
	fmt.Println(pow1)
	//fmt.Printf("pow1=%d \n", pow1)
	fmt.Println(math.Sqrt(8))
	fmt.Println(math.Cbrt(8))
	fmt.Println(math.Mod(11,3))
}

func t_pointer() {
	var a int = 100
	var p = &a
	fmt.Printf("%d 's address is %p \n", a, p)
	var ptr *int
	fmt.Printf("ptr=%p \n", ptr)
	ptr = p
	*ptr = 88
	fmt.Printf("%d 's address is p=%p, ptr=%p \n", a, p, ptr)

	var list[10] int
	for i:=0; i<10; i++ {
		fmt.Printf("[%d]=%d \n", i, list[i])
	}

	for i:=5; i<8; i++ {
		list[i] = i + 1
	}

	var p2 [10]*int
	p2[0] = &list[0]
	*p2[0] = 30
	for i:=0; i<10; i++ {
		fmt.Printf("[%d]=%d \n", i, list[i])
	}
}

func t_byte() {
	var ch byte = 65
	fmt.Printf("ch=%d, %c, 0x%x \n", ch, ch, ch)
}

func t_if() {
	var a = 100
	var b = 88
	var flag = false
	if (a > b) {
		flag = true
		fmt.Printf(" a is bigger than b.\n")
	} else {
		flag = false
	}

	fmt.Printf("a=%d, b=%d, flag=%t \n", a, b, flag)
}

func t_loop() {
	for i:=0; i<10; i++ {
		fmt.Printf("i=%d\n", i)
	}
/*
	j := 0
	while (j<10) {
		j++
		fmt.Printf("j=%d\n", j)
	}

	k := 0
	do {
		k++
		fmt.Printf("k=%d \n", k)
	}while(k < 10)
*/
}

func t_time() {
	var t1 = time.Now()
	fmt.Printf("t1=%d \n", t1)
	var t2 = time.Now().Nanosecond()
	fmt.Printf("t2=%d \n", t2)

	var day = t1.Day()
	var hour = t1.Hour()
	var min = t1.Minute()
	var year = t1.Year()
	var mon = t1.Month()
	fmt.Printf("year=%d, mon=%d, day=%d, hour=%d, min=%d \n", year, mon, day, hour, min)
	fmt.Printf("%d-%02d-%02d %02d:%02d\n", year, mon, day, hour, min)
	var t_utc = t1.UTC()
	fmt.Printf("t_utc=%d \n", t_utc)
	fmt.Printf("%d-%02d-%02d %02d:%02d\n", t_utc.Year(), t_utc.Month(), t_utc.Day(), t_utc.Hour(), t_utc.Minute())
}

func t_rand() {
	r1 := rand.Int()
	fmt.Printf("r1=%d \n", r1)

	var t2 = int64(time.Now().Nanosecond())
	rand.Seed(t2)
	r2 := rand.Int()
	fmt.Printf("rt2=%d \n", r2)
}

func t_complex() {
	var c1 complex64 = 5 + 10i
	fmt.Printf("c1=%v \n", c1)
	var c2 complex64 = 6 + 8i
	var c3 = c1 + c2
	var c4 = c1 * c2
	fmt.Printf("c1=%v, c2=%v, c3=%v \n", c1, c2, c3)
	fmt.Printf("c4=%v \n", c4)

	var re float32 = 10.8
	var im float32 = 7.5
	var c5 = complex(re, im)
	fmt.Printf("re=%f, im=%f, c5=%v \n", re, im, c5)
}

func t_print() {
	var a int = 88
	fmt.Printf("a=%d, [3d]=%03d, [16]=%x, [e]=%e, [f]=%f \n", a, a, a, int64(a*1000000000000000), float(a))
}

func t_bool() {
	var b1 bool = true
	fmt.Printf("b1=%t \n", b1)
	var b2 bool = !b1
	fmt.Printf("b2=%t \n", b2)
}

func test_getenv() {
	var s1 string = os.Getenv("GOROOT")
	fmt.Printf("s1=%s\n", s1)
	var path1 string = os.Getenv("PATH")
	fmt.Printf("path1=%s\n", path1)
}

func test_color() {
	c1 := Red
	c2 := Blue
	fmt.Printf("c1=%d, c2=%d\n", c1, c2)
}

func test_string() {
	var s1 string = "hello"
	fmt.Printf("s1=%s\n", s1)
	var len = len(s1)
	fmt.Printf("len=%d \n", len)

	for i:=0; i<len; i++ {
		fmt.Printf("[%d]=%c\t", i, s1[i])
	}
}

func test_str_contain() {
	var s1 string = "hello"
	var s2 string = "ll"
	var s3 string = "oo"
	var b1 = strings.Contains(s1, s2)
	var b2 = strings.Contains(s1, s3)
	fmt.Printf("s1=%s,s2=%s,s3=%s,b1=%t,b2=%t\n", s1, s2, s3, b1,b2)

	var index1 = strings.Index(s1, s2)
	fmt.Printf("index1=%d\n", index1)
	var index2 = strings.Index(s1, s3)
	fmt.Printf("index2=%d\n", index2)
}

func test_str_replace() {
	var s1 string = "hellohello"
	var s_new string = "OOOOOKKKK"
	var s2 = strings.Replace(s1, "ll", s_new, 1)
	fmt.Printf("s1=%s, s_new=%s, s2=%s \n", s1, s_new, s2)
	var s3 = strings.Replace(s1, "ll", s_new, -1)
	fmt.Printf("s3=%s \n", s3)
}

func test_convert() {
	var a int;
	var f1 float;
	f1 = 20.8
	a = int(f1)
	var f2 float
	f2 = float(a)

	fmt.Printf("a=%d, f1=%f\n", a, f1)
	fmt.Printf("f2=%f\n", f2)
	fmt.Printf("MAX_SIZE=%d\n", MAX_SIZE)
}

func test_T() {
	//var t1 T;
	fmt.Println("t1")
}

func add_float(f1 float, f2 float) float {
	fmt.Printf("f1=%f, f2=%f\n", f1, f2)
	f3 := f1 + f2
	return f3
}

func add(a int, b int) int {
	//fmt.Print("a=%d, b=%d\n", a, b)
	fmt.Printf("a=%d, b=%d\n", a, b)
	c := a + b
	return c
}

func test_float() {
	var f1 float = 0.1
	var f2 float = 2.3
	f3 := add_float(f1, f2)
	fmt.Printf("f1=%f, f2=%f, f3=%f\n", f1, f2, f3)
}

func test_int() {
	fmt.Println("test_int()")
	a := 3
	fmt.Println("a =", a)
	//_ := a
	//fmt.Println("-=", _)

	b := 8
	c := add(a, b)
	fmt.Println("c=", c)

	var d int32
	//d = a + b //cannot use a + b (type int) as type int32 in assignment
	d = int32(a+b)
	fmt.Printf("d=%d\n", d)
}

func test() {
	fmt.Println("test..")
}

func main() {
	fmt.Println("hello")
	//test()
	//test_int()
	//test_float()
	//test_convert()
	//test_string()
	//test_str_contain()
	//test_str_replace()
	//test_color()
	//test_getenv()
	//t_bool()
	//t_print()
	//t_complex()
	//t_rand()
	//t_time()
	//t_loop()
	//t_if()
	//t_byte()
	//t_pointer()
	//t_math()
	//t_strconv()
	//t_range()
	//t_sizeof()
	t_slice()
}

