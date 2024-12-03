export FABRIC_CA_CLIENT_HOME=/vagrant/veracity_node/identity_manager/

Enroll the admin:

fabric-ca-client enroll -u http://admin:adminpw@<fabric-ca-server-url> -M <admin-msp-path>

eg:
fabric-ca-client enroll -u http://admin:adminpw@localhost:7054 -M adminMSP

key pairs and certificates are in adminMSP




Register New Identity - By Admin

fabric-ca-client register --id.name <identity-name> --id.secret <identity-password> --id.type <type> --id.affiliation <affiliation> -u http://<fabric-ca-server-url> -M <admin-msp-path>

eg:

fabric-ca-client register --id.name tharindu1 --id.secret user1pw --id.type client --id.affiliation 
org1 -u http://localhost:7054 -M ./adminMSP


Enroll the new identity - By new user.

fabric-ca-client enroll -u http://<identity-name>:<identity-password>@<fabric-ca-server-url> -M <identity-msp-path>

eg:

fabric-ca-client enroll -u http://tharindu1:user1pw@localhost:7054 -M userMSP


# Artifacts Created for the Client Identity

## Client Certificate
- **Location:** `/vagrant/veracity_node/identity_manager/userMSP/signcerts/cert.pem`
- **Description:** This is the signed certificate for the client identity (`tharindu1`).

## Root CA Certificate
- **Location:** `/vagrant/veracity_node/identity_manager/userMSP/cacerts/localhost-7054.pem`
- **Description:** This is the CA certificate used to verify certificates signed by the CA.

## Issuer Public Key
- **Location:** `/vagrant/veracity_node/identity_manager/userMSP/IssuerPublicKey`
- **Description:** Used for certificate verification and issuance checks.

## Issuer Revocation Public Key
- **Location:** `/vagrant/veracity_node/identity_manager/userMSP/IssuerRevocationPublicKey`
- **Description:** Used for revocation operations.

## Private Key (Important)
- **Location:** `/vagrant/veracity_node/identity_manager/userMSP/keystore/`
- **File:** `<random-string>_sk`


In Fabric CA identities the user's public key is embedded in the client certificate generated during the enrollment process

How to get the public key of the user.

openssl x509 -in userMSP/signcerts/cert.pem -pubkey -noout


The private-public key pair generation algorithm used in your case is ECDSA (Elliptic Curve Digital Signature Algorithm) with a 256-bit key size.