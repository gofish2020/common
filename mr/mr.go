package mr

import (
	"fmt"
	"log"

	"github.com/kevwan/mapreduce/v2"
)

func Example() {
	val, err := mapreduce.MapReduce(func(source chan<- int) {
		// generator
		for i := 0; i < 10; i++ {
			source <- i
		}
	}, func(i int, writer mapreduce.Writer[int], cancel func(error)) {
		// mapper
		fmt.Println(i)
		writer.Write(i * i)
	}, func(pipe <-chan int, writer mapreduce.Writer[int], cancel func(error)) {
		// reducer
		var sum int
		for i := range pipe {
			sum += i
		}
		writer.Write(sum)
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
}
