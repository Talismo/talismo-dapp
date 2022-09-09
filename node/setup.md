# Setup Tutorial

## - Create SSH pub/prv keys
```
ssh-keygen -t rsa -b 4096 -f ~/.ssh/cardano_keys -C cardano
```

## - Copy to Instance
Add public key to GCE VM

_Advanced options > Security > Management > Add manually generated SSH keys_

## - Update and Upgrade
```
sudo apt update -y && sudo apt upgrade -y
```

## -- Install Cardano (haskell, cabal, libsodium, cardano-node and cardano-cli)


### - Install dependencies
```
sudo apt install automake build-essential pkg-config libffi-dev libgmp-dev libssl-dev libtinfo-dev libsystemd-dev zlib1g-dev make g++ tmux git jq wget libncursesw5 libtool autoconf -y
```

### - Install GHC and Cabal
```
curl --proto '=https' --tlsv1.2 -sSf https://get-ghcup.haskell.org | sh
```

* Please follow the instructions and provide the necessary input to the installer.
* Do you want ghcup to automatically add the required PATH variable to "/home/ubuntu/.bashrc"? - (P or enter)
* Do you want to install haskell-language-server (HLS)? - (N or enter)
* Do you want to install stack? - (N or enter)
* Press ENTER to proceed or ctrl-c to abort. (enter)

```
source ~/.bashrc
```

check ghcup version

```
ghcup --version
The GHCup Haskell installer, version v0.1.18.0
```

**Install ghc and Set the version**
```
ghcup install ghc 8.10.7
```
```
ghcup set ghc 8.10.7
```
```
ghc --version
The Glorious Glasgow Haskell Compilation System, version 8.10.7
```

**Install cabal and set the version**
```
ghcup install cabal 3.6.2.0
```
```
ghcup set cabal 3.6.2.0
```
```
cabal --version
cabal-install version 3.6.2.0
compiled using version 3.6.2.0 of the Cabal library
```

### - Libsodium

Sodium is a new, easy-to-use software library for encryption, decryption, signatures, password hashing and more.

```
git clone https://github.com/input-output-hk/libsodium
```
```
cd libsodium
```
```
git checkout 66f017f1
```
```
./autogen.sh
```
```
./configure
```
```
make
```
```
sudo make install
```

add libsodium to the bash profile

```
export LD_LIBRARY_PATH="/usr/local/lib:$LD_LIBRARY_PATH"
export PKG_CONFIG_PATH="/usr/local/lib/pkgconfig:$PKG_CONFIG_PATH"
```

Add symbolic link for libsodium

```
sudo ln -s /usr/local/lib/libsodium.so.23.3.0 /usr/lib/libsodium.so.23
```
```
source ~/.bashrc
```

### - Secp256k1
Secp256k1 is the name of the elliptic curve used by Bitcoin to implement its public key cryptography

```
git clone https://github.com/bitcoin-core/secp256k1
```
```
cd secp256k1
```
```
git checkout ac83be33
```
```
./autogen.sh
```
```
./configure --enable-module-schnorrsig --enable-experimental
```
```
make
```
```
make check
```
```
sudo make install
```

### - cardano-node

```
git clone https://github.com/input-output-hk/cardano-node.git
```
```
cd cardano-node
```
```
git fetch --all --recurse-submodules --tags
```
```
git checkout $(curl -s https://api.github.com/repos/input-output-hk/cardano-node/releases/latest | jq -r .tag_name)
```

### - Set Cabal as Compiler

```
cabal configure --with-compiler=ghc-8.10.7
```

### - cardano-wallet

```
git clone https://github.com/input-output-hk/cardano-wallet.git 
```
```
cd ./cardano-wallet/ 
```
```
TAG=$(git describe --tags --abbrev=0) && echo latest tag $TAG 
```
```
git checkout $TAG
```

### - Install cardano-node and cardano-cli

```
cabal build cardano-node cardano-cli
```

### - Install cardano-wallet

```
cabal build all
```
```
cabal install cardano-wallet
```

### - Add to local bin

```
mkdir -p ~/.local/bin
```
```
cd ~/cardano-node
```
```
cp -p "$(./scripts/bin-path.sh cardano-node)" ~/.local/bin/
```
```
cp -p "$(./scripts/bin-path.sh cardano-cli)" ~/.local/bin/
```
```
cp -p "$(./scripts/bin-path.sh cardano-wallet)" ~/.local/bin/
```

Add to bashrc
```
export PATH="$HOME/.local/bin/:$PATH"
```

### - check versions

```
cardano-node --version
cardano-node 1.35.3 - linux-x86_64 - ghc-8.10
git rev ea6d78c775d0f70dde979b52de022db749a2cc32
```

```
cardano-cli --version
cardano-cli 1.35.3 - linux-x86_64 - ghc-8.10
git rev ea6d78c775d0f70dde979b52de022db749a2cc32
```

```
cardano-wallet version
v2022-08-16 (git revision: afe575663a866c612b4a4dc3a90a8a700e387a86)
```

## -- Run Node
### - Download Config Files

```
wget https://raw.githubusercontent.com/input-output-hk/cardano-world/master/docs/environments/preprod/alonzo-genesis.json
```
```
wget https://raw.githubusercontent.com/input-output-hk/cardano-world/master/docs/environments/preprod/byron-genesis.json
```
```
wget https://raw.githubusercontent.com/input-output-hk/cardano-world/master/docs/environments/preprod/config.json
```
```
wget https://raw.githubusercontent.com/input-output-hk/cardano-world/master/docs/environments/preprod/shelley-genesis.json
```
```
wget https://raw.githubusercontent.com/input-output-hk/cardano-world/master/docs/environments/preprod/topology.json
```

### - Create node database folder

```
mkdir path/to/db
```

### - Run node

```
cardano-node run \
--config /path/to/config.json \
--database-path /path/to/db/ \
--socket-path /path/to/db/node.socket \
--host-addr 127.0.0.1 \
--port 1337 \
--topology /path/to/topology.json
```

### - Run wallet api

Listening to port 5090

Create a folder to save wallets db
```
mkdir -p /path/to/wallets/db
```

```
/home/cardano/.local/bin/cardano-wallet serve +RTS -N2 -RTS \
--node-socket /path/to/db/node.socket \
--testnet /path/to/byron-genesis.json \
--database /path/to/wallets/db/ \
--port 5090 \
--listen-address 0.0.0.0
```

### - Add to systemd to make it run as a service

**node executable**

```
vim node.sh
```

inside node.sh
```
#!/bin/bash

/home/cardano/.local/bin/cardano-node run \
--config /path/to/config.json \
--database-path /path/to/db/ \
--socket-path /path/to/db/node.socket \
--host-addr 127.0.0.1 \
--port 1337 \
--topology /path/to/topology.json
```

```
sudo chmod +x node.sh
```

**wallet executable**

```
vim wallet.sh
```

inside wallet.sh
```
#!/bin/bash

/home/cardano/.local/bin/cardano-wallet serve +RTS -N2 -RTS \
--node-socket /path/to/db/node.socket \
--testnet /path/to/byron-genesis.json \
--database /path/to/wallets/db/ \
--port 5090 \
--listen-address 0.0.0.0
```

```
sudo chmod +x wallet.sh
```

**add cardano-node as a service using systemd**

```
vim cardano-node.service
```
```
# talismo cardano-node service
# file: /etc/systemd/system/cardano-node.service

[Unit]
Description	= Talismo Cardano Node Service
Wants		= network-online.target
After		= network-online.target

[Service]
User		= ${USER}
Type		= simple
WorkingDirectory= /home/cardano/path/to/cardano-node
ExecStart	= /bin/bash -c '/home/cardano/path/to/cardano-node/node-service.sh'
KillSignal	= SIGINT
RestartKillSignal=SIGINT
TimeoutStopSec	= 300
LimitNOFILE	= 32768
Restart		= always
EstartSec	= 5
SyslogIdentifier= cardano-node

[Install]
WantedBy	= multi-user.target
```
```
mv cardano-node.service /etc/systemd/system/
```
```
sudo chmod 644 /etc/systemd/system/cardano-node.service
```

**add cardano-wallet as a service using systemd**

```
vim cardano-wallet.service
```
```
# talismo cardano-wallet service
# file: /etc/systemd/system/cardano-wallet.service

[Unit]
Description	= Talismo Cardano Wallet Service
Wants		= network-online.target
After		= network-online.target

[Service]
User		= cardano
Type		= simple
WorkingDirectory= /home/cardano/path/to/cardano-wallet
ExecStart	= /bin/bash -c '/home/cardano/path/to/cardano-wallet/wallet-service.sh'
KillSignal	= SIGINT
RestartKillSignal=SIGINT
TimeoutStopSec	= 300
LimitNOFILE	= 32768
Restart		= always
EstartSec	= 5
SyslogIdentifier= cardano-wallet

[Install]
WantedBy	= multi-user.target
```
```
mv cardano-wallet.service /etc/systemd/system/
```
```
sudo chmod 644 /etc/systemd/system/cardano-wallet.service
```

**enable services**

```
sudo systemctl enable cardano-node.service
```
```
sudo systemctl enable cardano-wallet.service
```
```
sudo systectl daemon-reload
```

**start services**
```
sudo systemctl start cardano-node
```
```
sudo systemctl start cardano-wallet
```