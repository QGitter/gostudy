package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"os"
)

var (
	has                         hash.Hash
	err                         error
	data, rest, result, padtext []byte
	privateKey                  *rsa.PrivateKey
	publicKey                   *rsa.PublicKey
	privateFile, publicFile     *os.File
	block                       *pem.Block
	padding, bs                 int
	blockDes                    cipher.Block
)

func main() {
	/*src := "hello,world"
	key := "123456"
	fmt.Println(Md5Str(src))
	fmt.Println(HmacSha256(src, key))
	fmt.Println(HmacSha512(src, key))
	fmt.Println(HmacSha1(src, key))
	fmt.Println(Sha256str(src))
	fmt.Println(Sha512str(src))
	fmt.Println(Sha1str(src))
	fmt.Println(Base64EncodeStr(src))
	fmt.Println(Base64DecodeStr(Base64EncodeStr(src)))
	//GenRsaKey(1024, "privateKey.pem", "publicKey.pem")
	msg := "ddddd"
	cipherText, _ := RsaEncrypt([]byte(msg), "./publicKey.pem")
	plainText, _ := RsaDecrypt(cipherText, "./privateKey.pem")
	fmt.Println(string(plainText))*/
	key := []byte("2fa6c1e2")
	str := "I love this beautiful world!"
	strEncrypted, _ := DesEncrypt(str, key)

	fmt.Println("Encrypted:", strEncrypted)
	strDecrypted, _ := DesDecrypt(strEncrypted, key)

	fmt.Println("Decrypted:", strDecrypted)
}

/******************************签名算法***************************************/
/**
 * md5签名算法
 */
func Md5Str(src string) string {
	has = md5.New()
	has.Write([]byte(src))
	return hex.EncodeToString(has.Sum(nil))
}

/**
 * 哈希sha256签名算法
 */
func HmacSha256(src, key string) string {
	has = hmac.New(sha256.New, []byte(key))
	has.Write([]byte(src))
	return hex.EncodeToString(has.Sum(nil))
}

/**
 * 哈希sha512签名算法
 */
func HmacSha512(src, key string) string {
	has = hmac.New(sha512.New, []byte(key))
	has.Write([]byte(src))
	return hex.EncodeToString(has.Sum(nil))
}

/**
 * 哈希sha1签名算法
 */
func HmacSha1(src, key string) string {
	has = hmac.New(sha1.New, []byte(key))
	has.Write([]byte(src))
	return hex.EncodeToString(has.Sum(nil))
}

/**
 * sha256密码散列函数
 */
func Sha256str(src string) string {
	has = sha256.New()
	has.Write([]byte(src))
	return hex.EncodeToString(has.Sum(nil))
}

/**
 * sha512密码散列函数
 */
func Sha512str(src string) string {
	has = sha512.New()
	has.Write([]byte(src))
	return hex.EncodeToString(has.Sum(nil))
}

/**
 * sha1密码散列函数
 */
func Sha1str(src string) string {
	has = sha1.New()
	has.Write([]byte(src))
	return hex.EncodeToString(has.Sum(nil))
}

/**
 * base64 编码函数
 */
func Base64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

/**
 * base64 解码函数
 */
func Base64DecodeStr(src string) string {
	data, err = base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(data)
}

/******************************非对称加密算法（RSA）***************************************
* 简介：
*	1、乙方生成两把密钥（公钥和私钥）。公钥是公开的，任何人都可以获得，私钥则是保密的。
*	2、甲方获取乙方的公钥，然后用它对信息加密。
*	3、乙方得到加密后的信息，用私钥解密。
*****************************************************************************************/
/**
 * 生成公钥、私钥文件
 */
func GenRsaKey(bits int, privatePath, publicPath string) error {

	/*****生成私钥文件*******/
	privateKey, err = rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	data = x509.MarshalPKCS1PrivateKey(privateKey)
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: data,
	}
	privateFile, err = os.Create(privatePath)
	defer privateFile.Close()
	err = pem.Encode(privateFile, &block)
	if err != nil {
		return err
	}

	/***********生成公钥文件************/
	publicKey = &privateKey.PublicKey
	data = x509.MarshalPKCS1PublicKey(publicKey)
	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: data,
	}
	publicFile, err = os.Create(publicPath)
	defer publicFile.Close()
	if err != nil {
		return err
	}
	err = pem.Encode(publicFile, &block)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 公钥加密函数
 */
func RsaEncrypt(src []byte, filename string) ([]byte, error) {
	publicFile, err = os.Open(filename)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer publicFile.Close()
	data, err = ioutil.ReadAll(publicFile)
	if err != nil {
		panic(err)
		return nil, err
	}
	block, rest = pem.Decode(data)
	if block == nil {
		panic(err)
		return nil, err
	}
	publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)
}

/**
 * 私钥解密函数
 */
func RsaDecrypt(src []byte, filename string) ([]byte, error) {
	privateFile, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer privateFile.Close()
	data, err = ioutil.ReadAll(privateFile)
	if err != nil {
		return nil, err
	}
	block, rest = pem.Decode(data)
	if block == nil {
		return nil, err
	}
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
}

/******************************对称加密算法（DES）***************************************
* 简介：加密和解密都使用的是同一个长度为64的密钥，实际上只用到了其中的56位，密钥中的第8、16...64位用来作奇偶校验。
* DES算法的安全性很高，目前除了穷举搜索破解外， 尚无更好的的办法来破解。其密钥长度越长，破解难度就越大。
*****************************************************************************************/

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding = blockSize - len(ciphertext)%blockSize
	padtext = bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func DesEncrypt(text string, key []byte) (string, error) {
	data = []byte(text)
	blockDes, err = des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs = blockDes.BlockSize()
	data = ZeroPadding(data, bs)
	if len(data)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		blockDes.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func DesDecrypt(decrypted string, key []byte) (string, error) {
	data, err = hex.DecodeString(decrypted)
	if err != nil {
		return "", nil
	}
	blockDes, err = des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(data))
	dst := out
	bs = blockDes.BlockSize()
	if len(data)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(data) > 0 {
		blockDes.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}
