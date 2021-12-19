/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test12.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/19 10:28:56 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/19 10:30:03 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
)

func main() {
	fs := make([]func() int, 2)
	fs[0] = func() int { return 2 }
	fs[1] = func() int { return 4 }
	for a, f := range fs {
		fmt.Println(a, f())
	}
}
