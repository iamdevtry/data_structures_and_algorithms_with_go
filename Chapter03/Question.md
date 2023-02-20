1. Where can you use double linked list? Please provide an example.

Double linked lists can be used in many situations where a data structure is required to provide fast insertion and 
deletion of elements, as well as easy traversal of the list in both forward and backward directions.

One example of where a double linked list can be used is in implementing a cache. In a cache, items that are frequently 
accessed are stored in memory to reduce the number of times they need to be loaded from slower storage such as a database. 
When a new item is added to the cache, it can be added to the front of the list. When the cache is full and a new item 
needs to be added, the least recently used item can be removed from the back of the list.

```
type Node struct {
    key   string
    value string
    prev  *Node
    next  *Node
}

type LRUCache struct {
    capacity int
    head     *Node
    tail     *Node
    cache    map[string]*Node
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        cache:    make(map[string]*Node),
    }
}

func (lru *LRUCache) Get(key string) string {
    if node, ok := lru.cache[key]; ok {
        lru.moveToFront(node)
        return node.value
    }
    return ""
}

func (lru *LRUCache) Put(key, value string) {
    if node, ok := lru.cache[key]; ok {
        node.value = value
        lru.moveToFront(node)
    } else {
        if len(lru.cache) == lru.capacity {
            delete(lru.cache, lru.tail.key)
            lru.removeTail()
        }
        node := &Node{
            key:   key,
            value: value,
        }
        lru.cache[key] = node
        lru.addToFront(node)
    }
}

func (lru *LRUCache) addToFront(node *Node) {
    if lru.head == nil {
        lru.head = node
        lru.tail = node
    } else {
        node.next = lru.head
        lru.head.prev = node
        lru.head = node
    }
}

func (lru *LRUCache) removeTail() {
    if lru.tail == nil {
        return
    }
    if lru.tail.prev == nil {
        lru.tail = nil
        lru.head = nil
    } else {
        lru.tail = lru.tail.prev
        lru.tail.next = nil
    }
}

func (lru *LRUCache) moveToFront(node *Node) {
    if node == lru.head {
        return
    }
    if node == lru.tail {
        lru.tail = node.prev
        lru.tail.next = nil
    } else {
        node.prev.next = node.next
        node.next.prev = node.prev
    }
    node.prev = nil
    node.next = lru.head
    lru.head.prev = node
    lru.head = node
}

```
2. Which method on linked list can be used for printing out node values?

In Go, to print out the values of a linked list, you would typically iterate through the linked list and print out the
value of each node. One common way to do this is to define a Print method on the linked list type that iterates through
the list and prints out the value of each node.

Here's an example implementation of a Print method for a singly linked list in Go:
```
type ListNode struct {
    Val int
    Next *ListNode
}

type LinkedList struct {
    Head *ListNode
}

func (l *LinkedList) Print() {
    current := l.Head
    for current != nil {
        fmt.Printf("%d ", current.Val)
        current = current.Next
    }
    fmt.Println()
}

```
In this implementation, the Print method iterates through the linked list starting at the head node (l.Head), printing 
out the value of each node (current.Val) until it reaches the end of the list (when current becomes nil). The fmt.Printf 
function is used to print out the value of each node, and fmt.Println is called at the end to add a newline character.

3. Which queue was shown with channels from the Go language?

In Go, a simple way to implement a queue using channels is by using a buffered channel with a fixed capacity. By sending
and receiving elements from the channel, we can add and remove elements from the queue. Here's an example implementation
of a queue using a buffered channel in Go: 

```
type Queue struct {
    data chan int
}

func NewQueue(size int) *Queue {
    return &Queue{data: make(chan int, size)}
}

func (q *Queue) Enqueue(val int) {
    q.data <- val
}

func (q *Queue) Dequeue() int {
    return <-q.data
}

func (q *Queue) Len() int {
    return len(q.data)
}

func (q *Queue) IsEmpty() bool {
    return len(q.data) == 0
}

```

In this implementation, we define a Queue struct with a single field data which is a buffered channel of integers. 
The Enqueue method adds an element to the back of the queue by sending it over the channel, while the Dequeue method 
removes and returns the element at the front of the queue by receiving it from the channel. The Len method returns the
current length of the queue, while the IsEmpty method returns a boolean indicating whether the queue is empty or not.

To use this queue, we can create a new instance with a fixed size using the NewQueue function, and then add and remove 
elements using the Enqueue and Dequeue methods. For example:
```
q := NewQueue(5)
q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)
fmt.Println(q.Len()) // prints 3
fmt.Println(q.Dequeue()) // prints 1
fmt.Println(q.Dequeue()) // prints 2
fmt.Println(q.IsEmpty()) // prints false

```
4. Write a method that returns multiple values. What data structure can be used for
   returning multiple values?

```
func tuples(x, y int){
    return x, y
}
```
