CONFIG_ROOT=/opt/gopath/src/github.com/hyperledger/fabric/peer
ORDERER_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/ordererOrganizations/supplychain.com/orderers/orderer.supplychain.com/msp/tlscacerts/tlsca.supplychain.com-cert.pem

#FarmInspector Environment Variables
FARMINSPECTOR_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/farminspector.supplychain.com/users/Admin@farminspector.supplychain.com/msp
FARMINSPECTOR_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/farminspector.supplychain.com/peers/peer0.farminspector.supplychain.com/tls/ca.crt

#Harvester Environment Variables
HARVESTER_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/harvester.supplychain.com/users/Admin@harvester.supplychain.com/msp
HARVESTER_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/harvester.supplychain.com/peers/peer0.harvester.supplychain.com/tls/ca.crt

#Exporter Environment Variables
EXPORTER_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/exporter.supplychain.com/users/Admin@exporter.supplychain.com/msp
EXPORTER_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/exporter.supplychain.com/peers/peer0.exporter.supplychain.com/tls/ca.crt

#Importer Environment Variables
IMPORTER_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/importer.supplychain.com/users/Admin@importer.supplychain.com/msp
IMPORTER_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/importer.supplychain.com/peers/peer0.importer.supplychain.com/tls/ca.crt

#Processor Environment Variables
PROCESSOR_MSPCONFIGPATH=${CONFIG_ROOT}/crypto/peerOrganizations/processor.supplychain.com/users/Admin@processor.supplychain.com/msp
PROCESSOR_TLS_ROOTCERT_FILE=${CONFIG_ROOT}/crypto/peerOrganizations/processor.supplychain.com/peers/peer0.processor.supplychain.com/tls/ca.crt

CAFILE_PATH=${CONFIG_ROOT}/crypto/ordererOrganizations/supplychain.com/msp/tlscacerts/tlsca.supplychain.com-cert.pem

CC_VERSION_NO=0


MODE=$1
shift

if [ "$MODE" == "invoke" ]; then
    echo
    echo "invoking store"
    set -x
    docker exec \
    cli peer chaincode invoke \
        -C scchannel \
        -n supplycc \
        -c '{"Args":["storeTestBlock","{}"]}' \
        -o orderer.supplychain.com:7050 --tls --cafile=${CAFILE_PATH}
    set +x
elif [ "$MODE" == "query" ]; then
    arg=$1
    shift
    stringarg="$arg"
    prefix='{"Args":["queryTestBlock", "'
    suffix='"]}'
    comm=$prefix$stringarg$suffix
    
    echo
    echo "querying store"
    set -x
    docker exec \
    cli peer chaincode query \
        -C scchannel \
        -n supplycc \
        -c $comm \
        -o orderer.supplychain.com:7050 --tls --cafile=${CAFILE_PATH}
    set +x
fi

# 1d8Jp8A23oZ9EZkAahpIM223VM4