/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test2.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/20 11:43:12 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/20 11:43:21 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"flag"
	"fmt"
	"strings"
)

// 設定される変数のポインタを取得
var msg = flag.String("msg", "デフォルト値", "説明")
var n int

func init() {
	// ポインタを指定して設定を予約
	flag.IntVar(&n, "n", 1, "回数")
}
func main() {
	// ここで実際に設定される
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))
}
