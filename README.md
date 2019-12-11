1. Install prerequisites:
https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html

2. clone fabric samples from github
git clone https://github.com/hyperledger/fabric-samples.git

3. we use basic-network to run our application
```
./start.sh && docker-compose -f docker-compose.yml up -d cli
```

4. check if all containers are running
```
docker ps
```

5. go to chaincode folder
```
cd ../chaincode
mkdir demo
cd demo
touch demo.go
```
6. write chaincode

7. install chaincode
```
docker exec cli peer chaincode install -n demo -v 1.0 -p "github.com/demo"
docker exec cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n demo -v 1.0 -c '{"Args":[]}' -P "OR ('Org1MSP.member')"
```

8. create application
```
cd ../../
mkdir demo
cd demo
```

9. load the required packageds
```
npm init -y
npm install fabric-ca-client fabric-network
```
10. enroll admin 
```
node enrollAdmin.js
ls wallet
```
11. create file to store jewelry
```
touch createJewelry.js 
node createJewelry.js admin jew1 ringaile ring gold rose
```
11. create file to query jewelry
```
touch queryJewelry.js
node queryJewelry.js admin jew1
```
12. create file to change owner
```
touch changeOwner.js
node changeOwner.js admin jew1 jenny
```
