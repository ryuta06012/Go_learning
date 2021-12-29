/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test11.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/28 16:07:47 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/28 16:08:21 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func makeCh() chan int {
	return make(chan int)
}
func recvCh(recv chan int) int {
	go func() { recv <- 200 }()
	return <-recv
}
func main() {
	ch := makeCh()
	go func() { ch <- 100 }()
	fmt.Println(recvCh(ch))
}
