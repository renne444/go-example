package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main()  {
	//创建私钥
	priv, _ := rsa.GenerateKey(rand.Reader, 512)
	priv2, _ := rsa.GenerateKey(rand.Reader, 512)

	pub := priv.PublicKey
	pub2:= priv2.PublicKey

	plainText := []byte("fuckyou")
	h := sha256.New()
	h.Write(plainText)
	hashed := h.Sum(nil)
	
	opt := rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	}
	sig,err := rsa.SignPSS(rand.Reader,priv,crypto.SHA256,hashed,&opt)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(sig)

	err = rsa.VerifyPSS(&pub,crypto.SHA256,hashed,sig,&opt)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("yes")
	}

	err = rsa.VerifyPSS(&pub2,crypto.SHA256,hashed,sig,&opt)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("yes")
	}

}
