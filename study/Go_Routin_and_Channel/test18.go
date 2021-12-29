/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test18.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 12:44:07 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 12:44:30 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func makeCh() chan int {
	return make(chan int)
}
func recvCh(recv <-chan int) int {
	return <-recv
}
func main() {
	ch := makeCh()
	go func(ch chan<- int) { ch <- 100 }(ch)
	fmt.Println(recvCh(ch))
	fmt.Println(recvCh(ch))
}
