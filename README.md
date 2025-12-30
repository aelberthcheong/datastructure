# Go Data Structures
A collection of common data structures implemented in Go

### Features
- Stack - LIFO data structure 
- Queue - FIFO data structure
- Heap - Priority queue implementation
- Tree - Binary search tree and variants

### Installation
```bash
go get github.com/aelberthcheong/datastructure
```

### Usage
**Stack**
```go
    import (
        "fmt"
        "github.com/aelberthcheong/datastructure/stack"
    )
    func main() {
        s := stack.New[int]()

        s.Push(10)
        s.Push(20)
        s.PushMany([]int{30, 40, 50}...)

        // 50 40 30 20 10
        for !s.IsEmpty() {
            v, _ := s.PopN(5)
            fmt.Printf("%d ", v)
        }
    }
``` 

**Queue**
```go
    import (
        "fmt"
        "github.com/aelberthcheong/datastructure/queue"
    )

    func main() {
        q := queue.New[int](5)

        q.Enqueue(10)
        q.Enqueue(20)
        q.Enqueue(30)

        // 10 20 30
        for !q.IsEmpty() {
            v, _ := q.Dequeue()
            fmt.Printf("%d ", v)
        }
    }
```

### Requirements
- Go 1.18 or higher (for generics support)

### LICENSE
MIT License - see [LICENSE](LICENSE) for details.