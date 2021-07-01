package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

/*
	需求：实现RSA的数字签名和认证
	基本步骤：
	1.获得RSA的密匙对
	2.对明文的散列值采用rsa的私匙进行签名
	3.验证数字签名的正确性
*/
// ------------数字签名------------
func SignatureRSA(plainText []byte, fileName string) []byte {
	// ------1.获取私匙------
	// Step1:打开文件获取私匙
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)
	// Step2:将私匙反pem化
	block, _ := pem.Decode(buf)
	// Step3:将私匙反X509序列化
	privkey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// ------2.获取明文的散列值------
	// Step1:创建指定哈希函数的Hash接口
	myHash := sha256.New()
	// Step2:将明文写入myHash结构体
	myHash.Write(plainText)
	// Step3：获得明文的散列值
	hashText := myHash.Sum(nil)
	// ------3.将明文的散列值采用RSA私匙进行签名------
	/*
		Step1:将明文的散列值采用RSA私匙进行签名
		函数：func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) (s []byte, err error)
		作用：实现数字签名
		返回参数1：带有数字签名的密文
		返回参数2：error
		参数1：rand.Reader
		参数2：私匙
		参数3：采用的哈希函数
		参数4：散列值
	*/
	cipher, err := rsa.SignPKCS1v15(rand.Reader, privkey, crypto.SHA256, hashText)
	if err != nil {
		panic(err)
	}
	return cipher
}

// ------------验证数字签名------------
func VerifyRSA(plainText, sigText []byte, fileName string) bool {
	// ------1.获取公钥------
	// Step1:打开文件获取公匙
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileinfo.Size())
	file.Read(buf)
	// Step2：将公匙反pem码化
	block, _ := pem.Decode(buf)
	// Step3:将公匙反x509序列化
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	// Step4:执行公匙的类型断言
	pulicKey := pubInterface.(*rsa.PublicKey)
	// ------2.获取明文的散列值------
	// Step1:创建hash接口，指定采用的哈希函数
	myHash := sha256.New()
	// Step2:向myHash中写入内容
	myHash.Write(plainText)
	// Step3:生成明文的散列值
	hashText := myHash.Sum(nil)
	// ------3.对数字签名后的内容进行解密------
	/*
		Step1:验证数字签名准确性
		函数：:func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (err error)
		作用：验证数字签名准确性
		返回参数1：error
		参数1:公匙
		参数2：指明采用的哈希函数
		参数3：明文的散列值
		参数4：数字签名后的内容
	*/
	err = rsa.VerifyPKCS1v15(pulicKey, crypto.SHA256, hashText, sigText)
	if err != nil {
		return false
	} else {
		return true
	}
}

func main() {
	src := []byte("在消息认证码中，需要发送者和接收者之间共享密钥，而这个密钥不能被主动攻击者Mallory获取。如果这个密钥落入Mallory手中，则Mallory也可以计算出MAC值，从而就能够自由地进行篡改和伪装攻击，这样一来消息认证码就无法发挥作用了。")
	sigText := SignatureRSA(src, "tpri.pem")
	bl := VerifyRSA(src, sigText, "tpub.pem")
	fmt.Println(bl)
}
