#!/bin/bash
. env.sh

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $5)
    local CP=$(one_line_pem $6)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${P1PORT}/$3/" \
        -e "s/\${CAPORT}/$4/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
	-e "s/\${ORGPK}/$7/" \
        $CONN_CONF_PATH/networkConfig-template.json 
}


ORG=1
P0PORT=7051
P1PORT=8051
CAPORT=7054
PEERPEM=crypto-config/peerOrganizations/org1.fabric-iot.edu/tlsca/tlsca.org1.fabric-iot.edu-cert.pem
CAPEM=crypto-config/peerOrganizations/org1.fabric-iot.edu/ca/ca.org1.fabric-iot.edu-cert.pem
ORGPK=`ls crypto-config/peerOrganizations/org1.fabric-iot.edu/users/Admin@org1.fabric-iot.edu/msp/keystore`

echo "$(json_ccp $ORG $P0PORT $P1PORT $CAPORT $PEERPEM $CAPEM $ORGPK)" > $CONN_CONF_PATH/networkConfig.json 

cp conn-conf/networkConfig.json ../../caliper-workspace/networks/
