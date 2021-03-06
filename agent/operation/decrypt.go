package operation

import (
	"github.com/ksang/blackbox/agent/cli"
	"github.com/ksang/blackbox/agent/connect"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// Decrypt a single file, write output to opath.
func AesDecryptFile(cipherdata []byte,
	hash []byte,
	key []byte,
	opath string) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	if len(cipherdata) < aes.BlockSize {
		return errors.New("ciphertext too short")
	}
	iv := cipherdata[md5.Size : md5.Size+aes.BlockSize]
	plain := make([]byte, len(cipherdata)-md5.Size-aes.BlockSize)
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(plain, cipherdata[md5.Size+aes.BlockSize:])
	fhash := md5.Sum(plain)
	if hex.EncodeToString(fhash[:]) != hex.EncodeToString(hash) {
		log.Println("Fhash:", hex.EncodeToString(fhash[:]))
		log.Println("Hash:", hex.EncodeToString(hash))
		return errors.New("md5 value not match")
	}
	err = ioutil.WriteFile(opath, plain, 0644)
	if err != nil {
		return err
	}
	return nil
}

// A wrapper of AesDecryptFile, including get file hash and connects
// to blackbox server for getting encryption key.
func AesDecryptFileAuto(arg cli.Args, path string) error {
	cipherdata, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	hash := cipherdata[:md5.Size]
	key, err := connect.RequestSecret(arg, hex.EncodeToString(hash))
	if err != nil {
		return err
	}
	_, filename := filepath.Split(path)
	sep := strings.Split(filename, ".")
	suf := ""
	if len(sep) >= 2 {
		suf = sep[len(sep)-1]
	}
	if suf != arg.Suffix {
		return errors.New("File: " + path + " doesn't have suffix.")
	}
	opath := path[:len(path)-len(arg.Suffix)-1]

	err = AesDecryptFile(cipherdata, hash, key, opath)
	if err != nil {
		return err
	}
	log.Println("Decrypted file saved to:", opath)
	return nil
}

// Recursively decrypt files within the root directory.
func AesDecryptFolderAuto(arg cli.Args, path string) error {
	return folderOp(AesDecryptFileAuto, arg, path)
}
