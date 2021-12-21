/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test4.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/21 10:18:11 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/21 11:05:42 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
)

func check_dir(test_dir string) error {
	var err error

	if _, err := os.Stat(test_dir); os.IsNotExist(err) {
		fmt.Println("error: nosuchdirectory: no such file or directory")
		return err
	}
	return err
}

func main() {
	test_dir := os.Args
	fmt.Printf("%s\n", test_dir[1])
	if err := check_dir(test_dir[1]) {
		println(err)
	}
}
