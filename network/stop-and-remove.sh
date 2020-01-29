docker-compose -f docker-compose-cli.yaml down --volumes --remove-orphans
docker rm -f $(docker ps -aq)
docker rmi -f $(docker images | awk '($1 ~ /dev-peer.*.mycc.*/) {print $3}')
rm -r ./crypto-config
rm -r ./channel-artifacts
rm ./mychannel.block
