/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test1.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/28 11:49:19 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/28 12:01:14 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("別のゴールーチン")
	}()
	fmt.Println("mainゴールーチン")
	time.Sleep(100 * time.Millisecond)
}
