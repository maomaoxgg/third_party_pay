package alipay

import (
	"crypto/sha1"
	"io"
	"fmt"
	"crypto/sha256"
	"crypto/rsa"
	"crypto"
	"encoding/base64"
	"encoding/pem"
	"crypto/x509"
)

/*
SHA1加密
 */
func stringSHA1(data string) []byte {
	t := sha1.New()
	io.WriteString(t,data)
	return t.Sum(nil)
}

/*
SHA256加密
 */
func stringSHA256(data string) []byte {
	t := sha256.New()
	io.WriteString(t,data)
	return t.Sum(nil)
}

func createPrivateKey(PRK string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(PRK))
	if block == nil {
		fmt.Println("create private_key error")
		return nil
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKCS1PrivateKey-------privateKey----- error : %v\n", err)
		return nil
	} else {
		return privateKey
	}
}

func createPublicKey(PUK string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(PUK))
	if block == nil {
		fmt.Println("create private_key error")
		return nil
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKIXPublicKey-------publicKey----- error : %v\n", err)
		return nil
	} else {
		return publicKey.(*rsa.PublicKey)
	}
}

/*
RSA1签名
 */
func RSA1Sign(Data string, privateKey *rsa.PrivateKey) ([]byte, error) {
	digest := stringSHA1(Data)
	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, digest)
	if err != nil {
		fmt.Errorf("rsaSign SignPKCS1v15 error")
		return nil, err
	}
	return s, nil
}

/*
RSA256签名
 */
func RSA2Sign(Data string, privateKey *rsa.PrivateKey) ([]byte, error) {
	digest := stringSHA256(Data)
	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, digest)
	if err != nil {
		fmt.Errorf("SignPKCS1v15 error")
		return nil, err
	}
	return s, nil
}

/*
base64加码
 */
func Base64(b []byte) string{
	return base64.StdEncoding.EncodeToString(b)
}

/*
base64解码
 */
func deBase64(s string) []byte{
	b,e := base64.StdEncoding.DecodeString(s)
	if e != nil{
		b = []byte{}
	}
	return b
}

/*
RSA1验签
 */
func RSA1Verify(Data,sign string,publicKey *rsa.PublicKey) error {
	digest := stringSHA1(Data)
	signByte := deBase64(sign)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, digest,signByte)
	if err != nil {
		fmt.Errorf("rsaVerify VerifyPKCS1v15 error")
	}
	return err
}

/*
RSA256验签
 */
func RSA2Verify(Data,sign string,publicKey *rsa.PublicKey)  error {
	digest := stringSHA256(Data)
	signByte := deBase64(sign)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest,signByte)
	if err != nil {
		fmt.Errorf("rsaVerify VerifyPKCS1v15 error")
	}
	return err
}




