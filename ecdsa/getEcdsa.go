package main

import (
    "fmt"
    "crypto"
    "crypto/rand"
    "crypto/elliptic"
    "crypto/ecdsa"
    "github.com/shengdoushi/base58"
)

func main() {
    var privateKey *ecdsa.PrivateKey
    var publicKey crypto.PublicKey
    privateKey, publicKey = newKeyPair()
//    privateKey, publicKey = 
    fmt.Println(base58.Encode(*privateKey)) 
    fmt.Println(base58.Encode(publicKey))
}



func newKeyPair() (*ecdsa.PrivateKey, crypto.PublicKey) {
    var privateKey *ecdsa.PrivateKey
    var publicKey crypto.PublicKey

    curve := elliptic.P256()
    privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
    if err != nil{
        fmt.Println("error is comming")
    }
    publicKey = privateKey.Public()
    

    return privateKey, publicKey
}
