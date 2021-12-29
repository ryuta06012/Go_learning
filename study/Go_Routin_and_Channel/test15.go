/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test15.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 11:15:09 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 11:38:42 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

type Stringer interface {
	String() string
}

type I int

func (n I) String() string {
	return "I"
}

type B bool

func (b B) String() string {
	return "B"
}

type S string

func (s S) String() string {
	return "S"
}

func F(s Stringer) {
	switch v := s.(type) {
	case I:
		fmt.Println(int(v), "I")
	case B:
		fmt.Println(bool(v), "B")
	case S:
		fmt.Println(string(v), "S")
	}
}

func main() {
	var n I = I(100)
	F(n)
	F(B(true))
	F(S("hoge"))
}
