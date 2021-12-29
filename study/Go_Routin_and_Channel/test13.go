/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test13.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 10:58:09 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 11:57:22 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

type Intinger interface {
	Inting() int
	Add() int
	String() string
}

type Hex int

func (h Hex) Inting() int {
	return int(h)
}

func (h Hex) Add() int {
	return (int(h) + int(h))
}

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func F(s Intinger) {
	v := s.(type)
	fmt.Println(int(v), "I")
}

func main() {
	var s Intinger = Hex(100)
	fmt.Println(s.Inting())
	fmt.Println(s.Add())
	fmt.Println(s.String())
}
