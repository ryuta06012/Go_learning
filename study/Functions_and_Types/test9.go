/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test9.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/18 15:44:52 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/18 15:46:58 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

func add(x int, y int) int {
	return x + y
}

func main() {
	var x int

	x = add(10, 10)
	println(x)

}