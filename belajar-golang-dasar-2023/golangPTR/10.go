package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// Jawaban soal 10 - cara yang benar
func createList() *Node {
	// Metode 1: Buat dari belakang ke depan
	// node3 := &Node{Value: 3, Next: nil}
	// node2 := &Node{Value: 2, Next: node3}
	// node1 := &Node{Value: 1, Next: node2}
	// return node1

	// Metode 2: Buat dari depan, sambung secara manual
	head := &Node{Value: 1, Next: nil}
	second := &Node{Value: 2, Next: nil}
	third := &Node{Value: 3, Next: nil}

	// Sambungkan
	head.Next = second
	second.Next = third

	return head
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// BONUS: Implementasi LinkedList yang lebih praktis
type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, size: 0}
}

func (ll *LinkedList) Add(value int) {
	newNode := &Node{Value: value, Next: ll.head}
	ll.head = newNode
	ll.size++
}

func (ll *LinkedList) AddLast(value int) {
	newNode := &Node{Value: value, Next: nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.size++
}

func (ll *LinkedList) Remove(value int) bool {
	if ll.head == nil {
		return false
	}

	// Jika yang dihapus adalah head
	if ll.head.Value == value {
		ll.head = ll.head.Next
		ll.size--
		return true
	}

	current := ll.head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkedList) Print() {
	printList(ll.head)
}

func (ll *LinkedList) Size() int {
	return ll.size
}

// DEMO: Bagaimana pointer bekerja di linked list
func demonstratePointers() {
	fmt.Println("=== DEMO: Bagaimana Pointer Bekerja ===")

	// Buat node pertama
	node1 := &Node{Value: 10, Next: nil}
	fmt.Printf("node1 address: %p, value: %d\n", node1, node1.Value)

	// Buat node kedua
	node2 := &Node{Value: 20, Next: nil}
	fmt.Printf("node2 address: %p, value: %d\n", node2, node2.Value)

	// Sambungkan - node1.Next menunjuk ke alamat node2
	node1.Next = node2
	fmt.Printf("node1.Next points to: %p (same as node2 address)\n", node1.Next)

	// Akses nilai node2 melalui node1
	fmt.Printf("Access node2.Value via node1: %d\n", node1.Next.Value)

	fmt.Println()
}

// POLA UMUM: Traversal dengan pointer
func findValue(head *Node, target int) *Node {
	current := head
	for current != nil {
		if current.Value == target {
			return current
		}
		current = current.Next // Pindah ke node berikutnya
	}
	return nil // Tidak ditemukan
}

// POLA UMUM: Modifikasi via pointer
func updateValue(head *Node, oldVal, newVal int) {
	current := head
	for current != nil {
		if current.Value == oldVal {
			current.Value = newVal
			return
		}
		current = current.Next
	}
}

func main() {
	fmt.Println("=== SOAL 10: Basic Implementation ===")
	list := createList()
	printList(list)

	demonstratePointers()

	fmt.Println("=== PRACTICAL LinkedList ===")
	ll := NewLinkedList()
	ll.AddLast(1)
	ll.AddLast(2)
	ll.AddLast(3)
	ll.AddLast(4)

	fmt.Print("Original: ")
	ll.Print()

	ll.Add(0) // Add di depan
	fmt.Print("After adding 0 at front: ")
	ll.Print()

	ll.Remove(2)
	fmt.Print("After removing 2: ")
	ll.Print()

	fmt.Printf("Size: %d\n", ll.Size())

	fmt.Println("\n=== POINTER PATTERNS ===")

	// Find value
	found := findValue(ll.head, 3)
	if found != nil {
		fmt.Printf("Found value 3 at address: %p\n", found)
	}

	// Update value
	updateValue(ll.head, 3, 99)
	fmt.Print("After updating 3 to 99: ")
	ll.Print()
}
