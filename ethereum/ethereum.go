package ethereum

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lytics/logrus"
	"github.com/qjouda/dignity-platform/backend/env"
)

// Ethereum ethereum struct to implement crypto interface
type Ethereum struct {
	rpc      *ethclient.Client
	auth     *bind.TransactOpts
	keystore *keystore.KeyStore
}

// MustNewEthereum create new Ethereum interface, panic if no connection
func MustNewEthereum(ethereumHost string) *Ethereum {
	conn, err := ethclient.Dial(ethereumHost)
	if err != nil {
		logrus.Fatal("Failed to connect to the Ethereum client: %v", err)
		return nil
	}
	// Loaderver key
	rawKey := env.MustGetenv("ETH_KEY_RAW")
	rawKeyECDSA, err := crypto.HexToECDSA(rawKey)
	if err != nil {
		logrus.Fatal("Something wrong with server private key.", err)
	}
	ks := keystore.NewKeyStore(
		env.MustGetenv("ETH_KEY_STORE_DIR"),
		keystore.LightScryptN,
		keystore.LightScryptP)
	ks.ImportECDSA(rawKeyECDSA, "passphrase")
	// Create an authorized transactor
	auth := bind.NewKeyedTransactor(rawKeyECDSA)
	if err != nil {
		logrus.Fatal("Failed to create transactor: %v", err)
	}
	return &Ethereum{
		rpc:      conn,
		auth:     auth,
		keystore: ks,
	}
}
