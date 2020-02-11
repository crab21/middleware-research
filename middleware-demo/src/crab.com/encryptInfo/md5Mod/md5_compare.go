package md5Mod

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)
//8a0b6a70f78e8cbdc1949eccac65c5d7
//2a411cb70eac79be63dfd906d68c0b71be8e3cfa63549b5ac5600e9ecd05987c
//2a411cb70eac79be63dfd906d68c0b71be8e3cfa63549b5ac5600e9ecd05987c
//8a0b6a70f78e8cbdc1949eccac65c5d7
//d13e6eae16fc8a16cd9fb741a308365b925b6260a5d520d7bca8cdaa6abb8ad0
//4c77e3ecff3c8b57f3fbc14b2c7afd82
func fileMd5Comp() {
	open, err := os.Open("/Users/wangpeiyuan/MergetSort.java")
	if err!= nil {

	}
	defer open.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, open);err!= nil {
		log.Print("复制过程出错......")
		return
	}

	sum := hash.Sum(nil)
	fmt.Printf("%x",sum)

}