#!/bin/bash
#配置调用peer0.org1.github.com节点的环境变量
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.github.com/users/Admin@org1.github.com/msp
CORE_PEER_ADDRESS=peer0.org1.github.com:7051
CORE_PEER_LOCALMSPID="Org1MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.github.com/peers/peer0.org1.github.com/tls/ca.crt
export CHANNEL_NAME=mychannel
#安装链码至peer0.org1.github.com节点
peer chaincode install -n mycc -v 1 -p github.com/chaincode/go
sleep 3
#实例化链码
peer chaincode instantiate -C $CHANNEL_NAME -n mycc -c '{"Args":[]}' -v 1 -P "AND ('Org1MSP.peer')" -o orderer.github.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/github.com/orderers/orderer.github.com/msp/tlscacerts/tlsca.github.com-cert.pem
sleep 6
#IPO
peer chaincode invoke -C $CHANNEL_NAME -n mycc -c '{"Args":["IPO","A1305","华夏","2000-1-1","0","10000","1.0"]}' -o orderer.github.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/github.com/orderers/orderer.github.com/msp/tlscacerts/tlsca.github.com-cert.pem --peerAddresses peer0.org1.github.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.github.com/peers/peer0.org1.github.com/tls/ca.crt
sleep 3
#查询
peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["Query","A1305"]}'
