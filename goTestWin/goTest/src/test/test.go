
package main

//import ("fmt";"os")
//import "fmt"

import (
	"fmt"
	"os"
	"math/rand"
	"math"
	"time"
	"strings"
	"runtime"
	"strconv"
	"unsafe"
	"reflect"
	"sort"
	"bufio"
	"flag"
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
	fmt.Printf("===========================\n")
	for pos2, ch := range str {
		fmt.Printf("%d, %c \n", pos2, ch)
	}
	
	fmt.Printf("pos3===========================\n")
	for pos3 := 0; pos3 < len(str); pos3++ {
		ch := str[pos3]
		fmt.Printf("%d, %c \n", pos3, ch)
	}
}

func t_strconv() {
	fmt.Printf("the size of int is: %d \n", strconv.IntSize)
	//var s1 string = "a88"
	var s1 string = "1666"
	n_s1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Printf("n_s1=%d, err=%s \n", n_s1, err)
		return
	} else {
		fmt.Println(err)
		fmt.Printf("err=%v\n", err)
		fmt.Printf("n_s1=%d \n", n_s1)

		s2 := strconv.Itoa(n_s1)
		fmt.Printf("s2=%s \n", s2)
	}
	
	s2 := "a33"
	if n_s2, err := strconv.Atoi(s2); err != nil {
		fmt.Printf("n_s2=%d, err=%s \n", n_s2, err)
		return
	} else {
		fmt.Printf("err=%v\n", err)
	}
}

func t_math() {
	var a = int64(math.Abs(-3))
	fmt.Printf("a=%d \n", a)
	var pow1 = math.Pow(3, 3)
	fmt.Println(pow1)
	//fmt.Printf("pow1=%d \n", pow1)
	fmt.Println(math.Sqrt(8)) //开根号
	fmt.Println(math.Cbrt(8)) //开立方根
	fmt.Println(math.Mod(11,7)) //取模
}

func t_pointer() {
	var a int = 100
	var p = &a
	fmt.Printf("a:%d 's address is %p \n", a, p)
	var ptr *int
	fmt.Printf("ptr=%p \n", ptr)
	ptr = p
	*ptr = 88
	fmt.Printf("a:%d 's address is p=%p, ptr=%p \n", a, p, ptr)

	var list[10] int
	fmt.Printf("list[10] is: \n")
	for i:=0; i<10; i++ {
		fmt.Printf("[%d]=%d \n", i, list[i])
	}

	for i:=5; i<8; i++ {
		list[i] = i + 1
	}
	fmt.Printf("list[10] is: \n")
	for i:=0; i<10; i++ {
		fmt.Printf("[%d]=%d \n", i, list[i])
	}

	var p2 [10]*int
	p2[0] = &list[0]
	*p2[0] = 30
	fmt.Printf("list[10] is: \n")
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
	
	a = 88
	
	if a > b {
	//if (a > b) {
		flag = true
		fmt.Printf(" a is bigger than b.\n")
	} else if a == b {
		fmt.Printf(" a is equal b.\n")
	} else {
		flag = false
	}

	fmt.Printf("a=%d, b=%d, flag=%t \n", a, b, flag)
}

func t_loop() {
	
	count := 0
	//for ;; {
	for {
		count ++
		fmt.Println("hello ", count)
		time.Sleep(time.Duration(1)*time.Second)
	}
	
	/*
	for i:=0; i<10; i++ {
		fmt.Printf("i=%d\n", i)
	}
	*/
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
	
	min := 100
	max := 1000
	r3 := rand.Int63n(int64(max - min))
	fmt.Printf("r3=%d \n", r3)
	
	for i:=0; i<20; i++ {
		r := rand.Int63n(int64(max - min))
		fmt.Printf("%d, r=%d \n", i, r)
	}
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
	
	var c6 = c5 * c5
	fmt.Printf("c6=%v", c6)
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
	
	if b2 {
		fmt.Println(" b2 true")
	} else {
		fmt.Println(" b2 false")
	}
	
	if b1 && b2 {
		fmt.Printf("b1=%t, b2=%t \n", b1, b2)
		fmt.Println(" b1 && b2 true")
	} else {
		fmt.Printf("b1=%t, b2=%t \n", b1, b2)
		fmt.Println(" b1 && b2 false")
	}
	
	if b1 || b2 {
		fmt.Printf("b1=%t, b2=%t \n", b1, b2)
		fmt.Println(" b1 || b2 true")
	} else {
		fmt.Printf("b1=%t, b2=%t \n", b1, b2)
		fmt.Println(" b1 || b2 false")
	}
	
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
	
	// 四舍五入处理，浮点转整型
	b := math.Trunc(float64(f1+0.5));
	fmt.Printf("b=%v, f1=%f\n", b, f1)
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

func add_float2(pf1 *float, pf2 *float) float {
	fmt.Printf("f1=%f, f2=%f\n", *pf1, *pf2)
	*pf1 += 1.0 
	f3 := *pf1 + *pf2
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

func t_float2() {
	var f1 float = 0.1
	var f2 float = 2.3
	f3 := add_float2(&f1, &f2)
	fmt.Printf("f1=%f, f2=%f, f3=%f\n", f1, f2, f3)
}

func test_int() {
	fmt.Println("test_int()")
	var i int = 888
	fmt.Println(reflect.TypeOf(i))
	
	
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

func t_version() {
	fmt.Printf("runtime.Version=%s", runtime.Version())
}

func t_chars() {
	for i := 0; i < 256; i++ {
		fmt.Printf("%d, 0x%x, %c \n", i, i, i)
	}
}

func t_for() {
	/*
	for i := 0; ; i++ {
		fmt.Printf("%d \n", i)
	}
	
	for i := 0; i < 3; {
		fmt.Printf("%d \n", i)
		time.Sleep(time.Duration(1)*time.Second)
	}
	
	
	s := ""
	for ; s != "aaaa"; {
		fmt.Printf("%s \n", s)
		time.Sleep(time.Duration(1)*time.Second)
		s += "a"
	}
	
	for i := 0; i < 7;  i++ {
		for j := 5; j < 100; j++ {
			for s := "a"; s != "aaaaaaaaa"; s += "a" {
				fmt.Println("Value of i, j, s:", i, j, s)
			} 
		}
	}
	*/
	
	for i, j, s := 0, 5, "a"; i < 7 && j < 100 && s != "aaaaaaaaa"; i, j, s = i+1, j+1, s + "a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
	
	
	j := 5
	s := "a"
	for i := 0; i < 7;  i++ {
		if j < 100 {
			if s != "aaaa" {
				 fmt.Println("Value of i, j, s:", i, j, s)
				 s += "a"
			} else {
				fmt.Println("else Value of i, j, s:", i, j, s)
			}
			j++
		}
	}
}

func t_break() {
	i := 10
	for {
		i--
		fmt.Printf("Value of i=%d\n", i)
		if i < 6 {
			fmt.Printf("break!\n")
			break
		}
	}
}

func t_continue() {
	i := 5
	for {
		i--
		if i == 3 {
			fmt.Printf("continue!\n")
			continue
		}
		fmt.Printf("Value of i=%d\n", i)
		
		if i < -10 {
			fmt.Printf("break!\n")
			break
		}
	}
}

func t_reflect() {
	funcName, file, line, ok := runtime.Caller(0)
	if ok { fmt.Printf("FUN=%s, file=%s, line=%d\n", runtime.FuncForPC(funcName).Name(), file, line) }
	
	var str string = "hello"
	fmt.Println(reflect.TypeOf(str))
	
	var a int = 10
	fmt.Println(reflect.TypeOf(a))
	
	var ary [10]int
	fmt.Println(reflect.TypeOf(ary))
	
	var b int64 = 33333
	fmt.Println(reflect.TypeOf(b))
	
	var map1 = map[int]string {3:"xxxx", 4:"ssss"}
	map1[1] = "heeee"
	map1[2] = "oooo"
	fmt.Printf("map1=%v \n", map1)
	fmt.Println(reflect.TypeOf(map1[1]))
	fmt.Println(reflect.TypeOf(map1))
}

func g() {
	fmt.Println("no name function!\n")
}

//g redeclared in this block
/*
func g(a int) {
	fmt.Println("no name function!\n")
}
*/

func t_fo_reference(pA *int) int {
	*pA += 10
	return *pA
}

func t_fo_value(a int) int {
	var pA *int = &a
	*pA = a + 10
	fmt.Printf("t_fo_value, a=%d \n", a)
	return *pA
}

func t_fun() {
	a := 8
	fmt.Printf("a=%d \n", a)
	t_fo_reference(&a)
	fmt.Printf("after t_fo_reference a=%d \n", a)
	t_fo_value(a)
	fmt.Printf("after t_fo_value a=%d \n", a)
}

func t_blank() int {
	//_ := 10 //cannot use _ as value
	//var _ int = 10
	a := 8
	b := t_fo_value(a)
	
	fmt.Printf("a=%d, b=%d \n", a, b)
	_ = t_fo_value(a) //对齐
	
	//fmt.Printf("_=%d \n", _) //cannot use _ as value
	//fmt.Println(_)	//cannot use _ as value
	//return _
	return 0
}

func Min(a ...int) int{
	fmt.Printf("len(a)=%d", len(a))
	//[]int{2,8,7,5} --> 4[]int
	fmt.Println(reflect.TypeOf(a))
	
	min := a[0]
	/*
	for _, v := range a {
		fmt.Printf("min=%d\n", min)
		if v < min {
			min = v
			fmt.Printf("v=%d\n", v)
		}
	}
	*/
	for pos, v := range a {
		fmt.Printf("pos=%d\n", pos)
		if v < min {
			min = v
			fmt.Printf("v=%d\n", v)
		}
	}
	
	return min
}

func t_nums_para() {
	n_min := Min(3,1,5)
	fmt.Printf("Min(1,3,5) n_min=%d\n", n_min)
	
	arr := []int{8,7,2,5}
	//Min(arr)//cannot use arr (type []int) as type int in argument to Min
	n_min = Min(arr...)
	fmt.Printf("n_min=%d\n", n_min)
}

func t_defer() {
	/*
	fmt.Println("t_defer 111111111")
	defer test()
	fmt.Println("t_defer 222222222")
	*/
	
	//越是放在前面的defer语句，越是最后才执行（类似，后进先出，栈的操作）
	i := 0
	defer fmt.Println("t_defer i=", i)
	i++
	defer fmt.Println("t_defer i=", i)
	
	for i := 0; i < 5; i++ {
		defer fmt.Println("2 t_defer i=", i)
	}
}

func trace(s string) { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func c() {
	trace("c")
	defer untrace("c")
	fmt.Println("in c")
	
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func t_fabonacci(n int) int {
	if n <= 1 {
		return 1
	} else {
		return t_fabonacci(n - 1) + t_fabonacci(n - 2)
	}
}

func t_fab() {
	/*
	for i := 0; i < 100; i++ {
		n := t_fabonacci(i)
		fmt.Printf("t_fab, i=%d, n=%d \n", i, n)
	}
	*/
	//var arry [100]int64;
	var arry [100]float;
	fmt.Println(reflect.TypeOf(arry))
	
	arry[0] = 1
	arry[1] = 1
	for i:=2; i<100; i++ {
		arry[i] = arry[i-1] + arry[i-2]
	}
	
	for pos, v := range arry {
		fmt.Printf("[%d] %f \n", pos, v)
	}


}

func t_check_map(map_in map[string]string, key string) bool {
	var value, isExist = map_in[key]
	if (isExist) {
		fmt.Printf("key=%s, value=%s\n", key, value)
	} else {
		fmt.Printf("key=%s not exist!\n", key)
	}
	return isExist
}

func t_map() {
	var map1 = make(map[string]string)
	map1["aaa"] = "hello"
	map1["bbb"] = "world"
	//fmt.Println(map1["aaa"])
	fmt.Println(map1["aaa"], map1["bbb"])
	
	var value, isExist = map1["aaa"]
	if (isExist) {
		fmt.Printf("value=%s\n", value)
	}
	
	t_check_map(map1, "bbb")
	t_check_map(map1, "bbbx")
}

func t_map_range() {
	var map1 = map[int]string {1:"aaa", 2:"bbb", 3:"hhhh", 4:"tttt"}
	//map1[1] = "aaa"
	//fmt.Println(map1)
	
	for key, value := range map1 {
		fmt.Printf("key=%d, value=%s\n", key, value)
	}
	
	for key := range map1 {
		fmt.Printf("222 key=%d, value=%s\n", key, map1[key])
	}
	
	for key, _ := range map1 {
		fmt.Printf("333 key=%d, value=%s\n", key, map1[key])
	}
	
	for _, value := range map1 {
		fmt.Printf("444 value=%s\n", value)
	}
}

func t_map_sort() {
	var map1 = map[int]string {2:"bbb",  4:"tttt",1:"aaa",3:"hhhh"}
	for key, value := range map1 {
		fmt.Printf("key=%d, value=%s\n", key, value)
	}
	int_keys := make([]int, len(map1))
	i := 0
	for key_elem, _ := range map1 {
		int_keys[i] = key_elem
		i++
	}
	var map2 = make(map[int]string, len(map1))
	sort.Ints(int_keys)
	for _, int_keys_value := range int_keys {
		//fmt.Printf("int_keys_value=%d\n", int_keys_value)
		value := map1[int_keys_value]
		fmt.Printf("%d:%s\n", int_keys_value, value)
		map2[int_keys_value] = value
	}
	
	for key, value := range map2 {
		fmt.Printf("key=%d, value=%s\n", key, value)
	}

}

func test_1() {
	//test()
	//test_int()
	//test_float()
	//t_float2()
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
	//t_slice()
	//t_version()
	//t_chars()
	//t_for()
	//t_break()
	//t_continue()
	//t_reflect()
	//t_fun()
	//t_blank()
	//t_nums_para()
	//t_defer()
	//b()
	//c()
	//t_fab()
	//t_map()
	//t_map_range()
	//t_map_sort()
}

func t_memstatus() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d \n", m.Alloc)
	fmt.Printf("m.Sys=%v \n", m.Sys)
}

func t_error_exit() {
	/*
	var err = nil //use of untyped nil
	if err != nil {
		fmt.Println("OK")
	} else {
		fmt.Printf("Program stopping with error %v", err)
		os.Exit(1)
	}
	*/
}

func t_file_open() {
	fileName := "a.txt"
	file, err := os.Open(fileName)
	if err != nil { 
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		fmt.Printf("Program stopping with error %v", err)
		os.Exit(1)
		return
	} else {
		iReader := bufio.NewReader(file)
		for {
			//str, err := iReader.ReadString('\n')
			str, _, err := iReader.ReadLine()
			if err != nil {
				fmt.Printf("Program stopping with error %v", err)
				return //error or EOF
			} else {
				fmt.Printf("Read: %s\n", str)
			}
		}
	}
	defer file.Close()
}

func t_fun_t(a int) {
	fmt.Printf("t_fun_t(int a) \n")
}

//t_fun_t redeclared in this block
/*
func t_fun_t() {
	fmt.Printf("t_fun_t() \n")
}
*/

func t_io_input() {
	context := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Printf("input.Text()=%s \n", input.Text())
		if input.Text() == "#" {
			break
		} else {
			context[input.Text()]++
		}
	}
	
	fmt.Printf("context=%v \n", context)
	
	for line, n := range context {
		if n > 1 {
			fmt.Printf("line=%s, n=%d \n", line, n)
		} else {
			fmt.Printf("line=%s, n=%d \n", line, n)
		}
	}
}

type Person struct {
	Name 	string
	Age 	int
}

//type

func t_person() {
	//Person per1 //#ERROR# unexpected per1 at end of statement
	var per1 Person
	per1.Name = "aaa"
	per1.Age = 30
	fmt.Printf("Name=%s, Age=%d \n", per1.Name, per1.Age)
	
	per2 :=  Person {"bbb", 32}
	fmt.Printf("Name=%s, Age=%d \n", per2.Name, per2.Age)
	
	//per3 := new Person{"bbb", 32}
}

func t_flag() {
	name := "NNNN"
	//flag.StringVar(&name, "name NNName", "value everyone", "useage xxx")
	pName := flag.String("name NNName", "value everyone", "useage xxx")
	flag.Parse()
	fmt.Printf("*pName=%v \n", *pName)
	
	//fmt.Printf("flag=%v \n", flag) //use of package flag without selector
	//fmt.Printf("name=%v \n", name)
	//fmt.Printf("useage=%v \n", flag.Usage)
	
	fmt.Printf("name=%v \n", name)
	{
		var name = "aaa" //name redeclared in this block
		fmt.Printf("in {} name=%v \n", name)
	}
	fmt.Printf("name=%v \n", name)
}

func t_getTheFlag() *string {
	return flag.String("name", "everybody", "The greeting object.")
}

func t_flag2() {
	var get_flag = t_getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v \n", *get_flag)
}

func test_2() {
	//t_memstatus()
	//t_error_exit()
	//t_file_open()
	//t_io_input()
	//t_person()
	//t_flag()
	t_flag2()
}

func main() {
	fmt.Println("main()")
	//test_1()
	test_2()
}

