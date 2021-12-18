/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test4.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/18 14:08:31 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/18 14:12:35 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func main() {
	//ns := make([]int, 3, 10)
	var array [10]int
	ns := array[0:3]
	fmt.Println(ns)
	//ms := []int{10, 20, 30, 40, 50}
	var array2 = [...]int{10, 20, 30, 40, 50}
	ms := array2[0:1]
	fmt.Println(ms)
}
