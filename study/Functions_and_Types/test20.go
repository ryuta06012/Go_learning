/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test20.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/19 11:46:21 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/19 12:21:23 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

type Hex int

func (h Hex) String() string {
	println(h)
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// 100をHex型として代入
	var hex Hex = 100
	// Stringメソッドを呼び出す
	fmt.Println(hex.String())
}
