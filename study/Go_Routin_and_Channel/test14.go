/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test14.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 11:04:43 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 11:04:47 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

type Func func() string

func (f Func) String() string { return f() }
func main() {
	var s fmt.Stringer = Func(func() string { return "hello" })
	fmt.Println(s)
}
