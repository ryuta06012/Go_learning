/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test17.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 12:36:45 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 12:48:52 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func makeCh() chan int {
	return make(chan int)
}
func recvCh(recv chan int) int {
	go func() {
		recv <- 200
	}()
	return <-recv
}
func main() {
	ch := makeCh()
	go func() { ch <- 100 }()
	fmt.Println(recvCh(ch))
	fmt.Println(recvCh(ch))
	fmt.Println(recvCh(ch))
}
