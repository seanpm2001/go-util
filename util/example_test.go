package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

func ExampleBitMask() {
	var b BitMask
	b.Add("Justin")
	b.Add("Naren|Mashiat|Derek|Dieter")
	v, err := b.Parse("Justin|Derek|Naren")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("0x%08b\n", v)
	fmt.Printf("%t %t %t\n", b.IsSet(v, "Justin"), b.IsSet(v, "Naren|Derek"), b.IsSet(v, "Mashiat"))
	s, err := b.Format(v)
	fmt.Println(s)
	// Output: 0x00001011
	// true true false
	// justin|naren|derek
}

func ExampleCopyFile() {
	s := "This is a test of the CopyFile() function."
	t1 := path.Join(path.Dir(os.Args[0]), "test1.txt")
	t2 := path.Join(path.Dir(os.Args[0]), "test2.txt")
	err := ioutil.WriteFile(t1, []byte(s), os.FileMode(0666))
	if err != nil {
		log.Fatal(err)
	}
	err = CopyFile(t2, t1)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile(t2)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(t1)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(t2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(b))
	// Output: This is a test of the CopyFile() function.
}

func ExampleStopWatch() {
	var s StopWatch
	s.Start()
	time.Sleep(time.Duration(1 * time.Second))
	s.Stop()
	d := s.GetElapsed()
	fmt.Printf("%d\n", Lrint(d.Seconds()))
	s.Reset()
	d = s.GetElapsed()
	fmt.Printf("%d\n", Lrint(d.Seconds()))
	// Output: 1
	// 0
}

func ExampleGaussian() {
	fmt.Println(Gaussian(0.0, 1.0))
	fmt.Println(Gaussian(1.0, 1.0))
	fmt.Println(Gaussian(-1.0, 0.5))
	// Output: 1
	// 0.6065306597126334
	// 0.1353352832366127
}

func ExampleLrint() {
	fmt.Println(Lrint(5.01))
	fmt.Println(Lrint(4.99))
	fmt.Println(Lrint(4.50))
	fmt.Println(Lrint(5.50))
	fmt.Println(Lrint(5.499999999999999))
	fmt.Println(Lrint(5.4999999999999999))
	fmt.Println(Lrint(-1.4))
	// Output: 5
	// 5
	// 5
	// 6
	// 5
	// 6
	// -1
}

func ExampleClipDuration() {
	min := time.Duration(5 * time.Second)
	max := time.Duration(30 * time.Second)
	fmt.Println(ClipDuration(3*time.Second, min, max))
	fmt.Println(ClipDuration(10*time.Second, min, max))
	fmt.Println(ClipDuration(50*time.Second, min, max))
	// Output: 5s
	// 10s
	// 30s
}

func ExampleClipInt() {
	min := 0
	max := 10
	fmt.Println(ClipInt(-1, min, max))
	fmt.Println(ClipInt(5, min, max))
	fmt.Println(ClipInt(15, min, max))
	// Output: 0
	// 5
	// 10
}

func ExampleMinInt() {
	fmt.Println(MinInt(5, 6))
	fmt.Println(MinInt(100, -1))
	fmt.Println(MinInt(10, 10))
	// Output: 5
	// -1
	// 10
}

func ExampleMaxInt() {
	fmt.Println(MaxInt(5, 6))
	fmt.Println(MaxInt(-1, -100))
	fmt.Println(MaxInt(10, 10))
	// Output: 6
	// -1
	// 10
}

func ExampleTwoDimSplit() {
	opts := TwoDimSplit("FirstName=Justin:LastName=Ruggles:EyeColor=Blue", ":", "=")
	b, err := json.MarshalIndent(opts, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	// Output: {
	//     "EyeColor": "Blue",
	//     "FirstName": "Justin",
	//     "LastName": "Ruggles"
	// }
}

func ExampleQueue() {
	q := NewQueue()
	q.Add(1)
	q.Add(2)
	q.Add(3)
	fmt.Printf("len=%d\n", q.Len())
	v := q.Remove().(int)
	fmt.Printf("%d\n", v)
	fmt.Printf("len=%d\n", q.Len())
	q.Clear()
	fmt.Printf("len=%d\n", q.Len())
	c := time.After(100 * time.Millisecond)
	d := make(chan int)
	go func() {
		v := q.RemoveWait()
		fmt.Printf("%v\n", v)
		d <- 0
	}()
	<-c
	q.Close()
	<-d
	// Output: len=3
	// 1
	// len=2
	// len=0
	// <nil>
}

func ExamplePriorityQueue() {
	pq := NewPriorityQueueWithWaitLimit(2, 1)
	pq.Add(5, 1)
	pq.Add(1, 0)
	pq.Add(2, 2)
	pq.Add(3, 1)
	pq.Add(4, 2)
	fmt.Printf("len=%d\n", pq.Len())
	v := pq.PeekP(0).(int)
	fmt.Printf("%d\n", v)
	v = pq.RemoveP(0).(int)
	fmt.Printf("%d\n", v)
	v = pq.Remove().(int)
	fmt.Printf("%d\n", v)
	v = pq.Remove().(int)
	fmt.Printf("%d\n", v)
	v = pq.Remove().(int)
	fmt.Printf("%d\n", v)
	v = pq.Peek().(int)
	fmt.Printf("%d\n", v)
	fmt.Printf("len=%d\n", pq.Len())
	pq.Clear()
	fmt.Printf("len=%d\n", pq.Len())
	c := time.After(100 * time.Millisecond)
	d := make(chan int)
	go func() {
		v := pq.RemoveWait()
		fmt.Printf("%v\n", v)
		d <- 0
	}()
	<-c
	pq.Close()
	<-d
	// Output: len=5
	// 1
	// 1
	// 2
	// 5
	// 4
	// 3
	// len=1
	// len=0
	// <nil>
}
