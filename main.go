package main

import "fmt" //ninja 6.6, 6.7, 6.8

type countDown struct {
	current int
}

func (c *countDown) tick() int {
	if c.current > 0 {
		c.current--
	}
	return c.current
}

func newTicker(n int) func() int {
	return func() int {
		if n > 0 {
			n--
		}
		return n
	}
}

func upCount(n int) func() int {
	return func() int {
		n++
		return n
	}
}

func newFibSeq() func() int {
	a := 0
	b := 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}

func main() {
	count := countDown{13}
	for count.tick() > 0 {

		for i := 0; i < count.current; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("Done")

	tick := newTicker(10)
	for tick() > 0 {
		fmt.Println("tick")
	}
	fmt.Println()
	count2 := upCount(10)
	fmt.Println(count2())
	count2()

	fmt.Println(count2())
	fmt.Println(count2())
	fmt.Println(count2())
	count3 := upCount(-3)
	for i := 0; i < 8; i++ {
		fmt.Println(count3())
		fmt.Println(count2())
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	f := newFibSeq()
	for i := 0; i < 19; i++ {
		fmt.Println(f())
	}

}
