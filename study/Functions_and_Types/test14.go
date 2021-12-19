/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test14.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/19 11:17:45 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/19 11:18:16 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

func main() {
	p := struct {
		age  int
		name string
	}{age: 10, name: "Gopher"}
	p2 := p // コピー
	p2.age = 20
	println(p.age, p.name)
	println(p2.age, p2.name)
}
