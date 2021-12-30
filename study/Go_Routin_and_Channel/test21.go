/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test21.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/29 14:47:37 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/29 14:49:19 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	bc := context.Background()
	t := 5 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()
	select {
	case <-time.After(6 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	println("fda")
}
