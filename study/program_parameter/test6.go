/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test6.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/21 12:21:41 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/21 12:23:12 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// パスを結合する
	path := filepath.Join("dir", "main.go")
	// 拡張子を取る
	fmt.Println(filepath.Ext(path))
	// ファイル名を取得
	str := filepath.Base(path)
	fmt.Println(filepath.Base(str))
	// ディレクトリ名を取得
	fmt.Println(filepath.Dir(path))
}
