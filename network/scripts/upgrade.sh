#!/bin/bash
#配置调用peer0.org1.github.com节点的环境变量
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.github.com/users/Admin@org1.github.com/msp
CORE_PEER_ADDRESS=peer0.org1.github.com:7051
CORE_PEER_LOCALMSPID="Org1MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.github.com/peers/peer0.org1.github.com/tls/ca.crt
export CHANNEL_NAME=mychannel
#安装链码至peer0.org1.github.com节点
peer chaincode install -n mycc -v 3 -p github.com/chaincode/go
sleep 3
#升级链码
peer chaincode upgrade -C $CHANNEL_NAME -n mycc -v 3 -c '{"Args":[]}' -P "AND ('Org1MSP.peer')" -o orderer.github.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/github.com/orderers/orderer.github.com/msp/tlscacerts/tlsca.github.com-cert.pem