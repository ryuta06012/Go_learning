/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test7.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/12/21 12:30:24 by hryuuta           #+#    #+#             */
/*   Updated: 2021/12/22 14:03:18 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

var from = flag.String("f", ".jpg", "Extension before conversion")
var to = flag.String("t", ".png", "Extension after conversion")
var rm = flag.Bool("r", false, "Remove file before conversion")

func removeSrc(src string) error {
	err := os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}

func Convert(src, dst string, rmsrc bool) error {
	// 変換前のファイルを読み取り専用で開く
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	// image.Imageへとデコード
	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	// 読み書き用ファイルを作成
	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer df.Close()

	// .以降を取り出して条件分岐 その後image.Imageからエンコード
	switch filepath.Ext(dst) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpg", ".jpeg":
		// 第三引数は画質で1~100まで。（今回は定数を使用:75)
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	case ".gif":
		err = gif.Encode(df, img, nil)
	}
	if err != nil {
		return err
	}

	// 変換前のファイル削除
	if rmsrc {
		if err = removeSrc(src); err != nil {
			return err
		}
	}
	return nil
}

func check_dir(test_dir string) error {
	var err error

	if _, err := os.Stat(test_dir); os.IsNotExist(err) {
		fmt.Println("error: nosuchdirectory: no such file or directory")
		return err
	}
	return err
}

func find_dir(dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == *from {
				fmt.Println(path)
				dst := path[:len(path)-len(*from)] + *to
				println(dst)
				err := Convert(path, dst, *rm)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err.Error())
					return err
				}
			}
			return nil
		})
	println(err)
	if err != nil {
		fmt.Println("Error on filepath.Walk : ", err)
	}
	return err
}

func main() {

	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s\n", "Not enough arguments")
		os.Exit(1)
	}
	if err := check_dir(os.Args[1]); err != nil {
		os.Exit(1)
	}
	if err := find_dir(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
