/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test3.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/20 13:39:07 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/20 14:05:41 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"flag"
	"fmt"
)

func main() {
	boolPtr := flag.Bool("b", false, "真偽値の値の例")
	intPtr := flag.Int("i", 0, "数値の場合の例")
	strPtr := flag.String("s", "", "文字列の場合の例")

	flag.Parse()
	src := flag.Arg(1)

	fmt.Printf("aaaaaaa\n")
	fmt.Println(src)
	fmt.Printf("aaaaaaa\n")

	fmt.Println(*boolPtr)
	fmt.Println(*intPtr)
	fmt.Println(*strPtr)
}
