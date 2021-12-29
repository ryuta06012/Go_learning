/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test3.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/28 12:26:27 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/28 12:31:32 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"time"
)

func main() {
	var v int
	go func() { // ゴールーチン-1
		//time.Sleep(1 * time.Second)
		v = 100
	}()
	go func() { // ゴールーチン-2
		//time.Sleep(1 * time.Second)
		fmt.Println(v)
		v = 200
		fmt.Println(v)
	}()
	time.Sleep(2 * time.Second)
}
