/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test19.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 12:57:33 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 13:08:39 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func makeCh() chan string {
	return make(chan string)
}

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	vh := makeCh()
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			vh <- s.Text()
		}
		// TODO: チャネルを閉じる
		close(vh)
	}()
	// TODO: チャネルを返す
	return vh
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
