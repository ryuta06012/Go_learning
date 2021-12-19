/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test21.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/19 12:25:05 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/19 12:26:12 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

type MyInt int

func (n *MyInt) Inc() {
	*n++
}
func main() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
}
