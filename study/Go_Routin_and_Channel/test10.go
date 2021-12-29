/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test10.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/28 15:45:20 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/28 16:13:47 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func makeCh() chan int {
	return make(chan int, 2)
}
func recvCh(recv chan int) int {
	go func() { recv <- 200 }()
	/* recv <- 200
	recv <- 300 */
	return <-recv
}
func main() {
	ch := makeCh()
	go func() { ch <- 100 }()
	fmt.Println(recvCh(ch))
	//fmt.Println(recvCh(ch))
	go func() { ch <- 500 }()
	ch <- 500
	fmt.Println(recvCh(ch))
	fmt.Println(<-ch)
}
