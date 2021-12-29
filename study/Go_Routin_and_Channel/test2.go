/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test2.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/28 12:01:32 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/28 12:11:28 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		//time.Sleep(3 * time.Second)
	}()
	go func() {
		defer fmt.Println("goroutine2 done")
		//time.Sleep(1 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}
