# Commands
Generate crypto materials and channel artifacts
```bash
cryptogen generate --config=./crypto-config.yaml

configtxgen -profile SupplyChainOrdererGenesis -channelID syschannel -outputBlock channel-artifacts/genesis.block

configtxgen -profile SupplyChainChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID scchannel

configtxgen -profile SupplyChainChannel -outputAnchorPeersUpdate ./channel-artifacts/FarmInspectorMSPanchors.tx -channelID scchannel -asOrg FarmInspectorMSP

configtxgen -profile SupplyChainChannel -outputAnchorPeersUpdate ./channel-artifacts/HarvesterMSPanchors.tx -channelID scchannel -asOrg HarvesterMSP

configtxgen -profile SupplyChainChannel -outputAnchorPeersUpdate ./channel-artifacts/ExporterMSPanchors.tx -channelID scchannel -asOrg ExporterMSP

configtxgen -profile SupplyChainChannel -outputAnchorPeersUpdate ./channel-artifacts/ImporterMSPanchors.tx -channelID scchannel -asOrg ImporterMSP

configtxgen -profile SupplyChainChannel -outputAnchorPeersUpdate ./channel-artifacts/ProcessorMSPanchors.tx -channelID scchannel -asOrg ProcessorMSP
```


# Network Design Details

<u>Channel Name</u> : scchannel

<u>Domain</u> : *.supplychain.com

1. Farm-Inspector
2. Harvester
3. Exporter
4. Importer
5. Processor

<u>Admin</u> : Admin creates new batch which is initial stage of coffee batch.

<u>Farm-Inspector</u> : Farm-inspectors are responsible for inspecting coffee farms and updating the information like coffee family, type of seed and fertilizers used for growing coffee.

<u>Harvester</u> : Harvesters conducting plucking, hulling , polishing , grading and sorting activities, further updating the information of crop variety, temperature used and humidity maintained during the process.

<u>Exporter</u> : Exporters are the organization who exports coffee beans throughout the world. Exporter adds quantity, destination address, ship name, ship number , estimated date and time and exporter id.

<u>Importer</u> : Importers imports the coffee from coffee suppliers and updates quantity, ship name, ship number , transporters information, warehouse name, warehouse address and the importer's address.

<u>Processor</u> : Processors are the organizations who processes raw coffee beans by roasting them on particular temperature and humidity and makes it ready for packaging and to sale into markets. Processor adds the information like quantity, temperature , roasting duration , internal batch number , packaging date time, processor name and processor address.

