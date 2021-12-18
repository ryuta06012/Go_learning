/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   fortuneSlip.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/17 11:42:43 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/17 11:52:21 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	t := time.Now().UnixNano()
	rand.Seed(t)
	for i := 0; i < 10; i++ {
		n := rand.Intn(6)
		switch n + 1 {
		case 6:
			fmt.Printf("大吉\n")
		case 5, 4:
			fmt.Printf("中吉\n")
		case 3, 2:
			fmt.Printf("小吉\n")
		default:
			fmt.Printf("凶\n")
		}
	}
}
