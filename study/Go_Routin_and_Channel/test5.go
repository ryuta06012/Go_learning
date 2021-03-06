/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test5.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/28 15:03:50 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/28 15:04:05 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // 容量0
	go func() {          // ゴールーチン-1
		ch <- 100
	}()
	go func() { // ゴールーチン-2
		v := <-ch
		fmt.Println(v)
	}()
	time.Sleep(2 * time.Second)
}
