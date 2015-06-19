package common

import (
    "fmt"
    "container/heap"
)

type element struct {
    next_time int  // priority
    item_id   int  // val
}


type vector []*element

func (v *vector)Len() int {
    return len(*v)
}

func (v *vector)Less(i, j int) bool {
    return (*v)[i].next_time < (*v)[j].next_time
}

func (v *vector)Swap(i, j int) {
    (*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

func (v *vector)Push(el interface{}) {
    *v = append(*v, el.(*element))
}

func (v *vector)Pop() interface{} {
    old := *v
    n := len(old)
    if n < 1 {
	return nil
    }
    el := old[n-1]
    *v = old[0:n-1]
    return el
}


type PQueue struct {
    v *vector
}

func NewPQueue(items map[int]int) *PQueue {
    pq := &PQueue{}
    vec := make(vector, len(items))
    i := 0
    for key, value := range items {
	vec[i] = &element{
	    next_time:key,
	    item_id:value,
	}
	i++
    }
    pq.v = &vec
    heap.Init(pq.v)
    return pq
}

func (this *PQueue)Push(next_time int, item_id int) {
    el := &element{}
    el.next_time = next_time
    el.item_id = item_id
    heap.Push(this.v, el)
}

func (this *PQueue)Pop() (int, error) {
    if len(*this.v) < 1 {
	return -1, fmt.Errorf("no element")
    }
    el := heap.Pop(this.v).(*element)
    if el != nil {
	return el.item_id, nil
    }
    return -1, nil
}
