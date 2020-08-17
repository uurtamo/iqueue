// Copyright 2020 Steve Uurtamo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package "github.com/uurtamo/iqueue"

import (

	//	"fmt"
	"sync"
)

type Interval struct {

	Lower uint32
	Upper uint32
}

type node struct {

	value *Interval
	next *node
}


type Queue struct {

	hmux sync.Mutex
	tmux sync.Mutex
	head *node
	tail *node
}

func (Q *Queue) Init () *Queue {

	i_node := new(node)
	i_node.next=nil
	Q.head=i_node
	Q.tail=i_node
	return Q
}

func (Q *Queue) Enqueue (bounds *Interval) {

	t_node := new(node)
	t_node.value=bounds
	t_node.next=nil

	Q.tmux.Lock()

	Q.tail.next=t_node
	Q.tail=t_node

	Q.tmux.Unlock()
	return
}

func (Q *Queue) Dequeue () *Interval {

	Q.hmux.Lock()
	defer Q.hmux.Unlock()

	h_new := Q.head.next

	if h_new == nil {
		return nil
	}

	value := h_new.value

	Q.head = h_new

	return value
}
