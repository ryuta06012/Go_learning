/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test3.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/18 13:42:31 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/18 14:04:57 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func main() {
	// ゼロ値で初期化
	var ns1 [5]int
	// 配列リテラルで初期化
	var ns2 = [5]int{10, 20, 30, 40, 50}
	// 要素数を値から推論
	ns3 := [...]int{10, 20, 30, 40, 50}
	// 5番目が50、10番目が100で他が0の要素数11の配列
	ns4 := [...]int{5: 50, 10: 100}
	fmt.Println(ns1[1:2])
	fmt.Println(ns2[1:2])
	println(len(ns2[1:2]))
	fmt.Println(ns3[1:2])
	fmt.Println(ns4[1:2])
}