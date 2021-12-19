/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test13.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/19 11:12:08 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/19 11:15:55 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func main() {
	x := 0
	fs := make([]func(), 4)
	for i := range fs {
		println(i)
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs {
		fmt.Printf("2回目%d\n", x)
		f()
		x++
	}
}
