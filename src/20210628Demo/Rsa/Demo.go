package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

// yes
func RSAGenKey(bits int) error {
	/*
		生成私钥
	*/
	// 1、使用RSA中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	// 2、通过X509标准将得到的RAS私钥序列化为：ASN.1 的DER编码字符串
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// privateStream = Pkcs1ToPkcs8(privateStream)
	// 3、将私钥字符串设置到pem格式块中
	block1 := pem.Block{
		Type:  "private key",
		Bytes: privateStream,
	}
	// 4、通过pem将设置的数据进行编码，并写入磁盘文件
	fPrivate, err := os.Create("privateKey.pem")
	if err != nil {
		return err
	}
	defer fPrivate.Close()
	err = pem.Encode(fPrivate, &block1)
	if err != nil {
		return err
	}

	/*
		生成公钥
	*/
	publicKey := privateKey.PublicKey
	publicStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	// publicStream:=x509.MarshalPKCS1PublicKey(&publicKey)
	block2 := pem.Block{
		Type:  "public key",
		Bytes: publicStream,
	}
	fPublic, err := os.Create("publicKey.pem")
	if err != nil {
		return err
	}
	defer fPublic.Close()
	pem.Encode(fPublic, &block2)
	return nil
}

// EncyptogRSA 对数据进行加密操作
func EncyptogRSA(src []byte, path string) (res []byte, pubkey *rsa.PublicKey, err error) {
	// 1.获取秘钥（从本地磁盘读取）
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	// 2、将得到的字符串解码
	block, _ := pem.Decode(b)

	// 使用X509将解码之后的数据 解析出来
	// x509.MarshalPKCS1PublicKey(block):解析之后无法用，所以采用以下方法：ParsePKIXPublicKey
	keyInit, err := x509.ParsePKIXPublicKey(block.Bytes) // 对应于生成秘钥的x509.MarshalPKIXPublicKey(&publicKey)
	// keyInit1,err:=x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return
	}
	// 4.使用公钥加密数据
	pubKey := keyInit.(*rsa.PublicKey)
	res, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	return
}

// DecrptogRSA 对数据进行解密操作
func DecrptogRSA(src []byte, path string) (res []byte, privateKey *rsa.PrivateKey, err error) {
	// 1.获取秘钥（从本地磁盘读取）
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	block, _ := pem.Decode(b)                                // 解码
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes) // 还原数据
	res, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	return
}
func main() {
	err := RSAGenKey(2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("秘钥生成成功！")

	str := "山重水复疑无路，柳暗花明又一村！"
	fmt.Println("加密之前的数据为：", string(str))
	data, _, err := EncyptogRSA([]byte(str), "publicKey.pem")
	if err != nil {
		fmt.Printf("err1:%v\n", err)
	}
	fmt.Println("加密之后的数据为：", string(data))
	fmt.Println("加密之后16进制编码的数据为：", hex.EncodeToString(data))

	data2, _, er := DecrptogRSA(data, "privateKey.pem")
	if er != nil {
		fmt.Printf("err2:%v\n", er)
	}
	fmt.Println("解密之后的数据为：", string(data2))

	// ---------------------------------------------------
	sign := SignatureRSA([]byte("data"), "tpri.pem")
	fmt.Println("sign:", sign)
	verify := VerifyRSA([]byte("data"), sign, "tpub.pem")
	fmt.Println("verify:", verify)

}

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
