package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)
 
func scanInt() int       { sc.Scan(); return parseInt(sc.Bytes()) }
func scanString() string { sc.Scan(); return string(sc.Bytes()) }
func scanFloat() float64 { sc.Scan(); return parseFloat(sc.Bytes()) }
func scanInts(n int) (ints []int) {
	ints = make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = scanInt()
	}
	return
}
 
func scanPairInts(n int) (a, b []int) {
	a = make([]int, n)
	b = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		b[i] = scanInt()
	}
	return
}
 
func scanStrings(n int) (st []string) {
	st = make([]string, n)
	for i := 0; i < n; i++ {
		st[i] = scanString()
	}
	return
}
func scanFloats(n int) (fs []float64) {
	fs = make([]float64, n)
	for i := 0; i < n; i++ {
		fs[i] = scanFloat()
	}
	return fs
}
 
func parseInt(b []byte) (n int) {
	for _, ch := range b {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
		}
	}
	if b[0] == '-' {
		return -n
	}
	return
}
 
var float64pow10 = []float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6,
	1e7, 1e8, 1e9, 1e10, 1e11, 1e12,
	1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
	1e19, 1e20, 1e21, 1e22,
}
 
func parseFloat(b []byte) float64 {
	f, dot := 0.0, 0
	for i, ch := range b {
		if ch == '.' {
			dot = i + 1
			continue
		}
		if ch -= '0'; ch <= 9 {
			f = f*10 + float64(ch)
		}
	}
	if dot != 0 {
		f /= float64pow10[len(b)-dot]
	}
	if b[0] == '-' {
		return -f
	}
	return f
}
 
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
 
type queue []int
 
func (q *queue) push(n int) {
	*q = append(*q, n)
}
 
func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}
 
func (q *queue) front() int {
	return (*q)[0]
}
 
func (q *queue) empty() bool {
	return len(*q) == 0
}
 
func topoSort(g [][]edge) (sorted []int) {
	n := len(g)
	index := make([]int, n)
	for i := 0; i < n; i++ {
		for _, e := range g[i] {
			index[e.to]++
		}
	}
 
	q := queue{}
	for i := 0; i < n; i++ {
		if index[i] == 0 {
			q.push(i)
		}
	}
 
	for !q.empty() {
		now := q.pop()
		sorted = append(sorted, now)
		for _, e := range g[now] {
			index[e.to]--
			if index[e.to] == 0 {
				q.push(e.to)
			}
		}
	}
	return
}
 
func chmax(x *int, y int) {
	*x = max(*x, y)
}
 
type edge struct {
	to   int
	cost int
}
 
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)
 
	n := scanInt()
	m := scanInt()
 
	g := make([][]edge, n)
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = scanInt()
		b[i] = scanInt()
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], edge{b[i], 0})
	}
	fmt.Println(g)

	dp := make([]int, n)
	so := topoSort(g)
	fmt.Println(so)
	for _, i := range topoSort(g) {
		for _, j := range g[i] {
			chmax(&dp[j.to], dp[i]+1)
		}
	}
	ans := 0
	for i := range dp {
		chmax(&ans, dp[i])
	}
	fmt.Fprintf(wr, "%v\n", ans)
 
}