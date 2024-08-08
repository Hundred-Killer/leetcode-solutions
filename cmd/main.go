package main

import (
	"fmt"
)

func main() {
	fmt.Println(mergeNodes(&ListNode{
		Val: 0, Next: &ListNode{
			Val: 3, Next: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 0, Next: &ListNode{
						Val: 4, Next: &ListNode{
							Val: 5, Next: &ListNode{
								Val: 2, Next: &ListNode{
									Val: 0, Next: &ListNode{},
								},
							},
						},
					},
				},
			},
		},
	}).Val)
}
