# Hyperledger Profile Management
Hyperledger Profile Management 

## Setup Hyperledger Network
### Prerequisites
#### Docker and Docker Compose
MacOSX, *nix, or Windows 10: Docker Docker version 17.06.2-ce or greater is required.
Older versions of Windows: Docker Toolbox - again, Docker version Docker 17.06.2-ce or greater is required.

#### Go Programming Language
Hyperledger Fabric uses the Go Programming Language for many of its components.
`Go version 1.12.x is required.`

First, you must set the environment variable GOPATH to point at the Go workspace containing the downloaded Fabric code base, with something like:
`export GOPATH=$HOME/go`

Second, you should (again, in the appropriate startup file) extend your command search path to include the Go bin directory, such as the following example for bash under Linux:
`export PATH=$PATH:$GOPATH/bin`

#### Node.js Runtime and NPM

#### Python
Install
`sudo apt-get install python`

Check your version
`python --version`

### Install Samples, Binaries, and Docker Images
If you want a specific release, pass a version identifier for Fabric and Fabric-CA docker images. The command below demonstrates how to download the latest production releases - Fabric v2.2.1 and Fabric CA v1.4.9

`curl -sSL https://bit.ly/2ysbOFE | bash -s -- <fabric_version> <fabric-ca_version>`

`curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.2.1 1.4.9`

You may want to add that to your PATH environment variable so that these can be picked up without fully qualifying the path to each binary. e.g.:
`export PATH=<path to download location>/bin:$PATH`

## Setup Your Project
### Repare resource
Copy `chaincode/fabcar` to `fabric-samples/chaincode`

Copy `apiserver` to `fabric-samples/fabcar`

### Run Network

In `fabric-samples/fabcar` run `startFabric.sh`

Wait a few minutes... take a cup of coffee!

### Run API server

#### Install dependency 
In `fabric-samples/fabcar/apiserver` run `npm install`

#### Run
In `fabric-samples/fabcar/apiserver` run `npx nodemon apiserver.js`

#### User Enrollment into Fabric Network

The two codes `enrollAdmin.js` and `registerUser.js` are responsible for admin and user enrollment on the Fabric network before we can interact the chaincode. The result is key pair and certificates for admin and for user.

`node enrollAdmin.js`

`node registerUser.js`

Now, you can try: `localhost:8081/api/login`

## Setup Blockchain Explorer (using codebase)

### Prerequisites

* Nodejs 10 and 12 (10.19 and 12.16 tested)
* PostgreSQL 9.5 or greater
* [jq](https://stedolan.github.io/jq)
* Linux-based operating system, such as Ubuntu or MacOS
* golang (optional)
  * For e2e testing

### Database Setup

```
$ cd blockchain-explorer/app
```

* Modify `app/explorerconfig.json` to update PostgreSQL database settings.

    ```json
    "postgreSQL": {
        "host": "127.0.0.1",
        "port": "5432",
        "database": "fabricexplorer",
        "username": "minhnd",
        "passwd": "minhndpassword"
    }
    ```

  * Another alternative to configure database settings is to use environment variables, example of settings:

    ```shell
    export DATABASE_HOST=127.0.0.1
    export DATABASE_PORT=5432
    export DATABASE_DATABASE=fabricexplorer
    export DATABASE_USERNAME=minhnd
    export DATABASE_PASSWD=minhndpassword
    ```
  **Important** repeat after every git pull (in some case you may need to apply permission to db/ directory, from blockchain-explorer/app/persistence/fabric/postgreSQL run: `chmod -R 775 db/`

### Update configuration

* Modify `app/platform/fabric/config.json` to define your fabric network connection profile:

    ```json
    {
        "network-configs": {
            "first-network": {
                "name": "firstnetwork",
                "profile": "./connection-profile/first-network.json",
                "enableAuthentication": false
            }
        },
        "license": "Apache-2.0"
    }
    ```

  * `first-network` is the name of your connection profile, and can be changed to any name
  * `name` is a name you want to give to your fabric network, you can change only value of the key `name`
  * `profile` is the location of your connection profile, you can change only value of the key `profile`

* Modify connection profile in the JSON file `app/platform/fabric/connection-profile/first-network.json`:
  * Change `fabric-path` to your fabric network disk path in the first-network.json file:
  * Provide the full disk path to the adminPrivateKey config option, it ussually ends with `_sk`, for example:
    `/fabric-path/fabric-samples/first-network/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore/aaacd899a6362a5c8cc1e6f86d13bfccc777375365bbda9c710bb7119993d71c_sk`
  * `adminUser` and `adminPassword` is the credential for user of Explorer to login the dashboard
  * `enableAuthentication` is a flag to enable authentication using a login page, setting to false will skip authentication.

### Run create database script:

* **Ubuntu**

    ```
    $ cd blockchain-explorer/app/persistence/fabric/postgreSQL/db
    $ sudo -u postgres ./createdb.sh
    ```

* **MacOS**

    ```
    $ cd blockchain-explorer/app/persistence/fabric/postgreSQL/db
    $ ./createdb.sh
    ```

Connect to the PostgreSQL database and run DB status commands:

```shell
$ sudo -u postgres psql -c '\l'
                                List of databases
      Name      |  Owner   | Encoding | Collate |  Ctype  |   Access privileges
----------------+----------+----------+---------+---------+-----------------------
 fabricexplorer | hppoc    | UTF8     | C.UTF-8 | C.UTF-8 |
 postgres       | postgres | UTF8     | C.UTF-8 | C.UTF-8 |
 template0      | postgres | UTF8     | C.UTF-8 | C.UTF-8 | =c/postgres          +
                |          |          |         |         | postgres=CTc/postgres
 template1      | postgres | UTF8     | C.UTF-8 | C.UTF-8 | =c/postgres          +
                |          |          |         |         | postgres=CTc/postgres
(4 rows)

$ sudo -u postgres psql fabricexplorer -c '\d'
                   List of relations
 Schema |           Name            |   Type   | Owner
--------+---------------------------+----------+-------
 public | blocks                    | table    | minhnd
 public | blocks_id_seq             | sequence | minhnd
 public | chaincodes                | table    | minhnd
 public | chaincodes_id_seq         | sequence | minhnd
 public | channel                   | table    | minhnd
 public | channel_id_seq            | sequence | minhnd
 public | orderer                   | table    | minhnd
 public | orderer_id_seq            | sequence | minhnd
 public | peer                      | table    | minhnd
 public | peer_id_seq               | sequence | minhnd
 public | peer_ref_chaincode        | table    | minhnd
 public | peer_ref_chaincode_id_seq | sequence | minhnd
 public | peer_ref_channel          | table    | minhnd
 public | peer_ref_channel_id_seq   | sequence | minhnd
 public | transactions              | table    | minhnd
 public | transactions_id_seq       | sequence | minhnd
 public | write_lock                | table    | minhnd
 public | write_lock_write_lock_seq | sequence | minhnd
(18 rows)

```

### Build Hyperledger Explorer

**Important:** repeat the below steps after every git pull

* `./main.sh install`
  * To install, run tests, and build project
- `./main.sh clean`
  * To clean the /node_modules, client/node_modules client/build, client/coverage, app/test/node_modules
   directories

Or

```
$ cd blockchain-explorer
$ npm install
$ cd client/
$ npm install
$ npm run build
```

### Run Hyperledger Explorer

#### Run Locally in Same Location

* Modify `app/explorerconfig.json` to update sync settings.

    ```json
    "sync": {
      "type": "local"
    }
    ```

* `npm start`
  * It will have the backend and GUI service up

* `npm run app-stop`
  * It will stop the node server

**Note:** If Hyperledger Fabric network is deployed on other machine, please define the following environment variable

```
$ DISCOVERY_AS_LOCALHOST=false npm start
```

#### Run Standalone in Different Location

* Modify `app/explorerconfig.json` to update sync settings.

    ```json
    "sync": {
      "type": "host"
    }
    ```

* If the Hyperledger Explorer was used previously in your browser be sure to clear the cache before relaunching

* `./syncstart.sh`
  * It will have the sync node up

* `./syncstop.sh`
  * It will stop the sync node

**Note:** If Hyperledger Fabric network is deployed on other machine, please define the following environment variable

```
$ DISCOVERY_AS_LOCALHOST=false ./syncstart.sh
```

### Logs

* Please visit the `./logs/console` folder to view the logs relating to console and `./logs/app` to view the application logs and visit the `./logs/db` to view the database logs.
