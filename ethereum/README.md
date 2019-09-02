# Generate Go Bindings for contracts
In this folder we have Ethereum standard contracts copied in
./zeppelin-contracts. To generated go bindings run the following scrips for the
intended contract and it will auto generate them

```sh
solc --bin --abi -o ./compiled --overwrite --allow-paths . ./contracts/token/ERC20/ERC20.sol
abigen --abi compiled/ERC20.abi --pkg generated --type ERC20 --out ./generated/ERC20.go --bin compiled/ERC20.bin
```
