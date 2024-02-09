package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

func main() {
	fmt.Println("Test")

	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	test7()
	test8()
	test9()
	test10()
	test11()
	test12()
	test12a()
	test13()
	test14()
	test16()
	test16a()
	test17()

	fmt.Println()

	slices.Reverse(fs)
	for _, f := range fs {
		f()
	}
}

var fs = []func(){
	func() {

	},

	func() {
		// Readers
		r := strings.NewReader("Hello, Reader!")

		b := make([]byte, 8)
		for {
			n, err := r.Read(b)
			fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
			fmt.Printf("b[:n] = %q\n", b[:n])
			if err == io.EOF {
				break
			}
		}
	},

	func() {
		// Error handling
		i, err := strconv.Atoi("42o")
		if err != nil {
			fmt.Printf("couldn't convert number: %v\n", err)
			return
		}
		fmt.Println("Converted integer:", i)
	},

	func() {
		// Type assertions
		var u interface{} = "hello"
		fmt.Println(u.(string))
		v, ok := u.(int)
		fmt.Println(ok, v)

		switch u.(type) { // only allowed in a switch
		case int:
			fmt.Println("int")
		case float32:
			fmt.Println("float32")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("<unknown>")
		}
	},

	func() {
		// Empty interface
		type any interface{}
		i := 5
		var a any = i // any can hold anything
		fmt.Println(a)
	},

	func() {
		// Implicit closures
		a := func(in int) func(int) int {
			x := in
			return func(y int) int {
				return x * y
			}
		}
		f := a(2)
		g := a(4)
		fmt.Println(f(3), g(3))
	},

	func() {
		// Maps
		// Maps literals short declaration
		type Vert struct {
			x, y float32
		}
		a := map[int]Vert{
			1: {0.1, 0.2},
			2: {0.3, 0.4},
		}
		fmt.Println(a)
		// Insert
		a[3] = Vert{0.5, 0.6} // Type needed
		delete(a, 2)
		fmt.Println(a)
		// Test existence
		elem, exists := a[2]
		fmt.Println(exists, (map[bool]Vert{true: elem, false: {}})[exists])
		elem, exists = a[1]
		fmt.Println(exists, (map[bool]Vert{true: elem, false: {}})[exists])
	},

	func() {
		// Empty slices and maps equal to nil
		var a []int
		fmt.Println(a == nil)
		var m map[int]int
		fmt.Println(m == nil)
	},

	func() {
		// Maps
		m := map[int]string{0: "bla", 1: "blubb"}
		for i, s := range m {
			fmt.Println(i, s)
		}
	},
}

// 18 Testing

func Fooer(input int) string {
    isfoo := (input % 3) == 0
    if isfoo {
        return "Foo"
    }
    return strconv.Itoa(input)
}

// 17 Web crawler example

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, urlCache *SafeUrlCache, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	if depth <= 0 {
		return
	}
	urlCache.mu.Lock()
	defer urlCache.mu.Unlock()
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	urlCache.urls[url] = true
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if ! urlCache.urls[u] {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, urlCache, wg)
		}
	}
}

type SafeUrlCache struct {
	urls map[string]bool
	mu sync.Mutex
}

func test17() {
	urlCache := SafeUrlCache{urls: make(map[string]bool)}
	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &urlCache, &wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/cmd/": &fakeResult{
		"Command",
		[]string{},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}


// 16 GENERICS / Type parameters

// generic function

func Map[T, V any](ts []T, fn func(T) V) []V {
	res := make([]V, len(ts))
	for i := range ts {
		res[i] = fn(ts[i])
	}
	return res
}

func test16() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(Map(a, func(x int) float64 { return math.Sqrt(float64(x)) }))
}

// generic type

type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](a []T) *List[T] {
	if len(a) == 0 {
		return nil
	}
	l := new(List[T])
	l.val = a[0]
	p := l
	for i:= 1; i < len(a); i++ {
		p.next = new(List[T])
		p = p.next
		p.val = a[i]
	}
	return l
}

func (l *List[T]) Len() int {
	i := 0
	for p := l; p != nil; p = p.next {
		i++
	}
	return i
}

func test16a() {
	l := NewList([]int{1, 2, 3})
	fmt.Println(l.Len())
	for p := l; p != nil; p = p.next {
		fmt.Println(p.val)
	}
}

// 15 Images

type Image struct {
	w, h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

// 14 Wrapping readers

type UpperCaseReader struct {
	r io.Reader
}

func (s UpperCaseReader) Read(a []byte) (n int, e error) {
	n, e = s.r.Read(a)
	for i := 0; i < n; i++ {
		a[i] = byte(unicode.ToUpper(rune(a[i])))
	}
	return
}

func test14() {
	sr := strings.NewReader("Bla")
	s := UpperCaseReader{sr}
	io.Copy(os.Stdout, &s)
	fmt.Println()
}

// 13 Readers example

type MyReader struct{}

func (s MyReader) Read(a []byte) (int, error) {
	for i := 0; i < cap(a); i++ {
		a[i] = byte('A')
	}
	return cap(a), nil
}

func test13() {
	a := make([]byte, 8)
	s := MyReader{}
	fmt.Println(s.Read(a))
	fmt.Println(a)
}

// 12 User defined errors

// simple

func test12() {
	a := func(x int) (int, error) {
		if x < 0 {
			return 0, errors.New("negative not allowed")
		}
		return x * x * x, nil
	}
	fmt.Println(a(3))
	fmt.Println(a(-3))
}

// custom

type CustomError struct {
	description string
	severity    int
}

func (e CustomError) Error() string {
	return fmt.Sprintf("error level %d: %s", e.severity, e.description)
}

func test12a() {
	a := func(x int) (int, error) {
		if x < 0 {
			return 0, CustomError{"negative", 3}
		}
		return x * x * x, nil
	}
	fmt.Println(a(3))
	fmt.Println(a(-3))
}

// 11 Method call on nil value is allowed (!)

type In interface {
	Stretch(float64)
}

type Cl struct {
	a, b float64
}

func (v *Cl) Stretch(f float64) {
	if v == nil {
		return
	}
	v.a *= f
	v.b *= f
}

func test11() {
	var x In
	var v *Cl
	v.Stretch(5)
	fmt.Println(v)
	x = v
	x.Stretch(5)
	fmt.Println(x)
	// but:
	//var y In
	//y.Stretch(5) // error!
}

// 10 "Classes"

type Vector struct {
	a, b float64
}

func (v *Vector) Stretch(f float64) {
	v.a *= f
	v.b *= f
}

func test10() {
	x := Vector{2.0, 3.0}
	x.Stretch(5)
	fmt.Println(x)
}

// 9 defer -> call at end of surrounding function (last-in-first-out)

func test9() {
	f := func(immediately, then string) string {
		fmt.Println(immediately)
		return then
	}

	defer fmt.Println(f("1", "4"))
	defer fmt.Println("3")

	fmt.Println("2")
}

// 8 Switch statemant: no break; short declaration; possible values

func test8() {
	linux := func() string {
		return "linux"
	}
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case linux(): // values can be other types then int and non constants
		// only values up to match will be evaluated
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	switch /*"true"*/ { // empty switch instead of if-else-chain
	case 12 > 15:
		fmt.Println("Huh?!")
	case 17 > 15:
		fmt.Println("Ah!")
	default:
		fmt.Println("Oh.")
	}
}

// 7 Loops and if statement with declaration

func test7() {
	for 1 > 0 { // this is a "while" loop
		fmt.Println("Run 1")
		break
	}

	for { // "forever"
		fmt.Println("Run 1")
		break
	}

	if a := 1; a > 0 {
		fmt.Println("True")
	} else if a < 0 {
		fmt.Println("False")
	} else {
		fmt.Println("Whatever")
	}
}

// 6 Constants vs. variables (and their types)

func test6() {
	var a = 100
	const b = 150
	fmt.Println(a, b)

	f := func(x float64) {
		fmt.Printf("%T %v\n", x, x)
	}
	f(float64(a)) // variable is int -> need conversion
	f(b)          // constant can be any number type -> ok
	g := func(x int) {
		fmt.Printf("%T %v\n", x, x)
	}
	g(a) // variable is int -> ok
	g(b) // constant can be any number type -> ok
}

// 5 Function objects and type inference

func test5() {
	f := func() string {
		return "bla"
	}
	var a = f()
	fmt.Println(a)
}

// 4 Multi assign at initialization + short declaration

func test4() {
	var a, b, c bool = true, false, true
	fmt.Println(a, b, c)

	d, e, bla := false, true, "bla"
	fmt.Println(d, e, bla)
}

// 3 Named return values and naked return

func test3() {
	fmt.Println(split(17))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 2 Multiple return values

func test2() {
	fmt.Println(shift(1, 2, 3))
}

func shift(a, b, c int) (int, int, int, string) {
	return c, a, b, fmt.Sprintf("%d %d %d", c, a, b)
}

// 1 Interface vs Struct

func test1() {
	a := Stru{bla: "bla"}
	test1func(a)
}

func test1func(a Interf) {
	fmt.Println(a.Bla())
}

type Interf interface {
	Bla() string
}

type Stru struct { // interface is implemented implicitly
	bla string
}

func (stru Stru) Bla() string {
	return stru.bla
}
