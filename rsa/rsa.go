package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
func Encrypt(data []byte) (string ,error){
	pub := ReadPem("public.pem")
	block, _ := pem.Decode([]byte(pub))

	pubInterface, errs := x509.ParsePKIXPublicKey(block.Bytes)
	if errs != nil {
		return "",errors.New("pubInterface失败")
	}
	ciphertext, e := rsa.EncryptPKCS1v15(rand.Reader, pubInterface.(*rsa.PublicKey), data)
	if e != nil {
		return "",errors.New("ciphertext失败")
	}
	fmt.Println("加密：",ciphertext)
	return base64.StdEncoding.EncodeToString(ciphertext),nil
}
func Decrypt(data string) (string ,error) {
	dedata,errd := base64.StdEncoding.DecodeString(data)
	//fmt.Println("解码：",dedata)
	if errd != nil {
		return "",errors.New("base64解码失败")
	}
	pri := ReadPem("private.pem")
	block,_ := pem.Decode([]byte(pri))
	privs, errs := x509.ParsePKCS1PrivateKey(block.Bytes)
	if errs != nil {
		return "",errors.New("ParsePKCS1PrivateKey失败")
	}
	//解密得到明文
	plaintext1, e := rsa.DecryptPKCS1v15(rand.Reader, privs, dedata)
	if e != nil {
		fmt.Println(e)
	}
	return string(plaintext1),nil
}