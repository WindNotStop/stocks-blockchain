#!/bin/bash
#配置环境变量
export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=mychannel
export IMAGE_TAG=1.4.4

COMPOSE_FILE="-f docker-compose-cli.yaml"
#加密配置
../bin/cryptogen generate --config=./crypto-config.yaml
#生成创世块
mkdir channel-artifacts
../bin/configtxgen -profile TwoOrgsOrdererGenesis -channelID byfn-sys-channel -outputBlock ./channel-artifacts/genesis.block
#通道配置交易
../bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
#定义锚节点
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
../bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
#启动所有容器
docker-compose $COMPOSE_FILE up -d 2>&1
#docker exec cli scripts/script.sh
#将配置文件拷贝至2号主机
#IP=192.168.226.129
#NAME=ytz
#CATALOG=/home/ytz/stocks-blockchain/network
#scp -r ./channel-artifacts ${NAME}@${IP}:${CATALOG}
#scp -r ./crypto-config ${NAME}@${IP}:${CATALOG}
#CLI_ID=$(docker ps | grep cli | awk '{print $1}')
#docker cp $CLI_ID:/opt/gopath/src/github.com/hyperledger/fabric/peer/mychannel.block ./
#scp  ./mychannel.block ${NAME}@${IP}:${CATALOG}

