package main

import "fmt"

func main() {
	q := &Queue{}
	q.Enqueue("Alex")
	q.Enqueue("John")
	q.Enqueue("Mary")

	f := &Fifo{}
	l := &Lifo{}

	i := q.GetIterator()

	fmt.Println(fmt.Sprintf("%d totals", i.Count()))
	fmt.Println("=============================")

	for i.Valid() {
		fmt.Println(i.Current())
		i.Next()
	}

	fmt.Println("======== DEQUEUEING =============")

	fmt.Println(l.Dequeue(q))
	fmt.Println(i.Count())
	fmt.Println(f.Dequeue(q))
	fmt.Println(i.Count())

	fmt.Println("========= HAS ==========")
	i.Rewind()

	for i.Valid() {
		fmt.Println(i.Current())
		i.Next()
	}
	fmt.Println("========= DEQUEUEING =====")

	fmt.Println(q.Dequeue(l))
	fmt.Println(i.Count())
	fmt.Println(q.Dequeue(l))
}
