/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test20.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 13:47:57 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 14:06:13 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("done1")
	}()
	wg.Add(2)
	go func() {
		defer wg.Done()
		//defer wg.Done()
		fmt.Println("done2")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("done2")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("done10")
	}()
	wg.Wait()
	fmt.Println("done all")
}
