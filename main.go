package main

import (

        // get version number of sol
        "context"
        "crypto/ed25519"
        //"encoding/json"
        "fmt"
        "github.com/portto/solana-go-sdk/client"
        "github.com/portto/solana-go-sdk/rpc"
        qrcode "github.com/skip2/go-qrcode"
        "log"
        "reflect"

        // create sol wallet
        "github.com/portto/solana-go-sdk/types"
)

type Address struct {
        KeyPublic    string
        KeyPrivate   ed25519.PrivateKey
        DonoString   string
        AmountToSend float64
        Arrived      bool
}

var addressSlice []Address

// Mainnet
//var c = client.NewClient(rpc.MainnetRPCEndpoint)

// Devnet
var c = client.NewClient(rpc.DevnetRPCEndpoint)

func main() {
        // create a RPC client
        fmt.Println(reflect.TypeOf(c))

        // get the current running Solana version
        response, err := c.GetVersion(context.TODO())
        if err != nil {
                panic(err)
        }

        fmt.Println("version", response.SolanaCore)

        createWalletSolana()

        checkBalance()

}

func createWalletSolana() {
        //create a new wallet using
        wallet := types.NewAccount()

        // display the wallet public and private keys
        //fmt.Println("Wallet Address:", wallet.PublicKey.ToBase58())
        //fmt.Println("Wallet Private Key:", wallet.PrivateKey)

        address := Address{}
        address.KeyPublic = wallet.PublicKey.ToBase58()
        address.KeyPrivate = wallet.PrivateKey
        address.AmountToSend = 5.2342
        address.DonoString = "Haha"
        address.Arrived = false
        //addressByte, _ := json.Marshal(address)

        //fmt.Println("Json OBJ: " + string(addressByte))
        addToAddressSlice(address)
        generateQR(address.KeyPublic, address.AmountToSend)

}

func setWallets() {

}

func checkBalance() {

        for _, wallet := range addressSlice {

                fmt.Println("Checking: " + wallet.KeyPublic)
                checkAddressBalance(wallet.KeyPublic)
        }

}

func generateQR(address string, amount float64) {
        var amountString = fmt.Sprintf("%f", amount)
        qrcode.WriteFile("solana:"+address+"?amount="+amountString, qrcode.Medium, 320, "qr.png")
}

func checkAddressBalance(w string) {

        balance, err := c.GetBalance(
                context.TODO(), // request context
                w,              // wallet to fetch balance for
        )

        if err != nil {
                log.Fatalln("get balance error", err)
        }
        // the smallest unit like lamports
        fmt.Println("balance", balance/1e9)

}

func addToAddressSlice(a Address) Address {
        addressSlice = append(addressSlice, a)
        fmt.Println(len(addressSlice))
        return a
}
