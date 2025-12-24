package main 
import (
	"os"
	"fmt"
	"bufio"
	"container/heap"
	
)

type Node struct{
	freq int
	char rune
	left *Node
	right *Node 
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int{
	return len(pq)
}

func(pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int){
	pq[i] , pq[j] = pq[j], pq[i]
}

func(pq *PriorityQueue) Push(x interface{}){
	*pq = append(*pq, x.(*Node))
}

func(pq *PriorityQueue) Pop() interface{}{
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}


func BuildHuffmanTree(freq map[rune]int) *Node{
	pq := &PriorityQueue{}
	heap.Init(pq)

	for ch , f := range freq{
		heap.Push(pq, &Node{
			char : ch,
			freq: f,
		})
	}

	//combine until one node remains 
	for pq.Len() > 1{
		left := heap.Pop(pq).(*Node)
		right := heap.Pop(pq).(*Node)

		parent := &Node{
			freq : left.freq + right.freq,
			left : left,
			right : right,
		}

		heap.Push(pq, parent)
	}

	return heap.Pop(pq).(*Node)
}

func PrintTree(n *Node, indent string){
	if n == nil{
		return 
	}

	fmt.Printf("%sFreq: %d Char: %q \n", indent, n.freq, n.char)
	PrintTree(n.left, indent+"  ")
	PrintTree(n.right, indent+"  ")

}




func GenerateFreqMap(filename string) (map[rune]int,error){
	data, err := os.Open(filename)
	if err != nil{
		fmt.Println("Error while opening file: ", err)
	}
	defer data.Close()
	result := make(map[rune]int)

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan(){
		ch := []rune(scanner.Text())[0] //get teh ccharacter
		result[ch]++
	}

	if err := scanner.Err(); err != nil{
		fmt.Println("Scanner error:", err)
		return nil, err
	}
	return result,err
}

func main(){

	if len(os.Args) < 2{
		fmt.Println("Usage: go run main.go <filename>")
		return 
	}

	filename := os.Args[1]

	freqMap,err := GenerateFreqMap(filename)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for ch, count := range freqMap {
		fmt.Printf("%q : %d\n", ch, count)
	}

	// Example checks
	fmt.Println("Occurrences of 'X':", freqMap['X'])
	fmt.Println("Occurrences of 't':", freqMap['t'])

	root := BuildHuffmanTree(freqMap)
	fmt.Println("Root Frequency: ", root.freq)
	PrintTree(root, "")
}