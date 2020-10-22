peer chaincode package -n cc_crcc -p bsn-fabric-sc/chaincode/cmd/ -v 0.1 cc_crcc.0.1.pak

peer chaincode install cc_crcc.0.1.pak

peer chaincode instantiate -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_crcc -v 0.1 -c '{"Args":["init","cc_crcc","cc_cross"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem

peer chaincode upgrade -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_crcc -v 0.2 -c '{"Args":["init","cc_crcc","cc_cross"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem

peer chaincode invoke -o order1.ordernode.bsnbase.com:17051 -C netchannel -n cc_crcc -c '{"Args":["callnft","abc"]}' --tls true --cafile /etc/hyperledger/fabric/certs/ordererOrganizations/ordernode.bsnbase.com/orderers/order1.ordernode.bsnbase.com/tls/tlsintermediatecerts/tls-ca-ordernode-bsnbase-com-15901-2.pem


peer chaincode query -C netchannel -n cc_crcc -c '{"Args":["query","24686dc124d9f3928385ea5d0a9c95ee7328de388f5f066399a6ede534e7f84d"]}'


