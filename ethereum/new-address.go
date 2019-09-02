package ethereum

// CreateNewAddress creates a new Ethereum address in the server wallet and
// returns it
func (b *Ethereum) CreateNewAddress() (string, error) {
	acc, err := b.keystore.NewAccount("server")
	return acc.Address.String(), err
}
