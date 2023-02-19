package main

import "fmt"

// Queue array of Orders Type
type Queue []*Order

// Order class
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// New method initializes with Order with priority, quantity, product, customerName
func (order *Order) New(priority int, quantity int, product string, customerName string) {
	order.priority = priority
	order.product = product
	order.quantity = quantity
	order.customerName = customerName
}

// Add method adds the order to the queue
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false

		for i, addedOrder := range *queue {
			if order.priority > addedOrder.priority {
				fmt.Println("==== run ====")
				x := (*queue)[:i]
				y := Queue{order}
				z := (*queue)[i:]
				o := append(Queue{order}, (*queue)[i:]...)
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}

		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func main() {
	//var queue Queue
	//queue = make(Queue, 0)
	queue := make(Queue, 0)

	order1 := Order{
		priority:     2,
		quantity:     20,
		product:      "PC",
		customerName: "devtry",
	}

	order2 := Order{
		priority:     3,
		quantity:     10,
		product:      "Monitor",
		customerName: "iamdevtry",
	}

	order3 := Order{
		priority:     4,
		quantity:     10,
		product:      "Monitor",
		customerName: "iamdevtry",
	}

	queue.Add(&order1)
	queue.Add(&order2)
	queue.Add(&order3)

	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
