/*SHA1散列经常用于计算二进制或文本块的短标识。
例如，git版本控制系统广泛使用SHA1s来标识版本化的文件和目录。
下面是如何在Go中计算SHA1哈希值。
Go在各种crypto/*包中实现了几个哈希函数。*/

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)

	fmt.Printf("%x\n", bs)
}
