# Get Started

- Remove all existing docker images, containers, or volumes to make sure there are no conflicting images. Use the ```-f``` tag if any of the following commands fail.
```bash
docker rm $(docker ps -aq)
docker rmi $(docker images -aq)
docker volumes rm $(docker volume ls)
```
- Execute the following command to download samples, binaries, and docker images.
```bash
curl -sSL https://bit.ly/2ysbOFE | bash -s -- 1.4.7 1.4.7 0.4.20
```
- To generate crypto certs and channel artifacts execute the generate script in the terminal. 
```bash
./generate.sh
```

- Once the relevant files have been generated bring up the network using docker compose 
```bash
docker-compose up
```

- With the network up and running, open  a new terminal window and execute the cli-script
```bash
./cli-script.sh
```

# Network Design Details

<u>Channel</u> : scchannel

<u>Orderer Channel</u> : syschannel

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