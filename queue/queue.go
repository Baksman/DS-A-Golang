package main

type Queue []*Order

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func(o *Order) NewQueue(priority int,quantity int,product string,customerName string){
	o.customerName = customerName
	o.priority =priority
	o.product = product
	o.quantity = quantity

}

func (q *Queue) Add(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	}else{
		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i,addedOrder = range *q {
			if order.priority > addedOrder.priority{
				*q = append((*q)[:i],append() )
			}
	}
}}
