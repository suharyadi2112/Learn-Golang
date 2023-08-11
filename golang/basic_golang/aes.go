package main

import(
  "fmt"
  "crypto/aes"
	"crypto/cipher"
	"encoding/base64"
  "errors"
  "bytes"
)
// Credential ...
// Credential ...
type Credential struct {
	Key string
	Iv  string
}
func main(){

  var s1 = Credential{"AFd6N3v1ebLw711zxpZjxZ7iq4fYpNYa", "MesA7nqIVa23b167"}
  // dec,_ := s1.Decrypt("XF9ks/0jDS31ZpQAaoCemw==")
  encb,_ := s1.Decrypt("TTcpw695o2JdlUb3GGA4ZQ==")
  encbb,_ := s1.Encrypt("054421")

  // dncb,_ := s1.Decrypt("JmZKAkabdFp9FCVJzzBEXQ==")
  fmt.Println(encb, encbb)
}
func tes (tes string)(res string){
  res = tes
  return res
}

// Encrypt ...
func (cred *Credential) Encrypt(src string) (string, error) {
	if src == "" {
		return "", errors.New("empty_parameter")
	}

	key := []byte(cred.Key)
	iv := []byte(cred.Iv)


	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return base64.StdEncoding.EncodeToString(crypted), nil
}
// Decrypt ...
func (cred *Credential) Decrypt(data string) (res string, err error) {
	if data == "" {
		return res, nil
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return res, err
	}

	block, err := aes.NewCipher([]byte(cred.Key))
	if err != nil {
		return res, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return res, err
	}

	mode := cipher.NewCBCDecrypter(block, []byte(cred.Iv))
	mode.CryptBlocks(ciphertext, ciphertext)

	return string(PKCS5Trimming(ciphertext)), err
}
// PKCS5Trimming ...
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// PKCS5Padding ...
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
