
Catsd
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/catspa3/catspad)

Catsd is the reference full node Cats implementation written in Go (golang).

Catsd fork from kaspad

## quick start

#### LINUX

1. Install the official Catspa Mining Software(Linux)

[download](https://github.com/catspa3/catspad/releases)

2. Make sure your harddisk has at least 50 GB free space, create a directory named 'Cats', and enter the directory with root permissions.

3. Download and unzip the latest Catspa Miner Software from Github

4. Grant executable permission to 'cats' programs (chmod +x *)

5. Run the command :

```
 chmod +x * 
 ./genesis --testnet
```

a. If the daemon processes Catspd and CatsMINER continue to encounter errors, please clear the cache with the following command :

```
 rm -rf ~/.catsd
```

and re-run step 5

6. Create a file firstrun.sh and grant with executable permission

a. Copy the following script to firstrun.sh#!/bin/bashnohup

```
./catsd --testnet --utxoindex >firstpad.log 2>&1 &
```

7. run firstrun.sh

8. Create a Catspa wallet address

```
./catswallet --testnet create
```

9. Create a file walletrpc.sh and grant with executable permission

a. Copy the following script to walletrpc.sh

```
#!/bin/bash
nohup ./catswallet --testnet start-daemon  > walletrpc.log 2>&1 &
```

10. run walletrpc.sh

11. Create a sub-wallet address with the following command:

```
./catswallet --testnet new-address
```

Please be sure to save the newly generated sub-wallet address, as it will be used for the mining program.

12. Create a file worker01.sh and grant with executable permission

a. Copy the following script to worker01.sh

```
#!/bin/bash
nohup ./catsminer --testnet --miningaddr YOUR_SUB_WALLET_ADDRESS > w01.log 2>&1 &
```

#### WINDOWS

1. Install the official Catspa Mining Software (Windows version)

[download](https://github.com/catspa3/catspad/releases)


2. Make sure your harddisk has at least 50 GB free space, create a directory named 'Cats', and enter the directory with root permissions.

3. Download and unzip the latest Catspa Miner Software from Github

4. Run the command :

```
./genesis.exe --testnet
```

a. If the daemon processes Catspd and CatsMINER continue to encounter errors, please clear the cache with the following command : 

```
rm -rf ~/.catsd
```

and re-run step 5

5. Run catsd.exe:  

```
./catsd.exe --testnet --utxoindex
```

6. Create a Catspa wallet address

```
./catswallet.exe --testnet create
```

7. Run catswallet.exe:

```
./catswallet.exe --testnet start-daemon
```

8. Create a sub-wallet address with the following command:

```
./catswallet.exe --testnet new-address
```

Please be sure to save the newly generated sub-wallet address, as it will be used for the mining program.

9. Run catsminer.exe with miningaddr parameter:

```
./catsminer.exe --testnet --miningaddr YOUR_SUB_WALLET_ADDRESS
```


## Documentation

The cats's document is the same as kaspa's

The [documentation](https://github.com/catspa3/docs) is a work-in-progress

## License

Catsd is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
