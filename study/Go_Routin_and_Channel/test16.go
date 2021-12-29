/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test16.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 12:33:26 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 12:35:44 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func makeCh() chan int {
	return make(chan int)
}
func recvCh(recv chan int) int {
	return <-recv
}
func main() {
	ch := makeCh()
	sh := makeCh()
	go func() {
		ch <- 100
		sh <- 200
	}()
	fmt.Println(recvCh(ch))
	fmt.Println(recvCh(sh))
}
