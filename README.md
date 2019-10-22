# Kyma Drone

This project contains the project data to connect a Parrot Mambo drone with kyma using an Linux host, SAP CP API Management and SAP Cloud Connector.

## Requirements

- Drone: Parrot Mambo
- Linux Host (or VM)
- BLE Dongle (on OSX it is not possible to use the host Bluetooth. A dedicated hardware dongle is required)

## Build Project

```bash
go build -o server cmd/kyma-drone-server-v2/main.go
go build -o connector cmd/kyma-connector/kyma-connector.go
```

## Setup

*Detailes Setup Instructions in the [Hybris Wiki](https://wiki.hybris.com/display/ps/Drone+Setup)*

The Project contains of two commands. The server and the Registration cli. During the setup, the server can run in TEST mode.

Run server in test mode:

```bash
sudo TEST_API=true KYMA_CONFIG=./config ./server --port 8080 --host 0.0.0.0
```

The server is running on port `8080`. The server is exposed via SAP CP API Management. To setup the connection, the server needs to be available `From Cloud to OnPremise` via SAP Cloud Connector.

1. Setup Cloud Connector (use `drone` as virtual host name)
2. Attach service to API Management (use drone as system name)
3. Import API's (see `api-management`folder)
4. Bundle API's as a product and subscribe via developer portal

Int he next step, the api and events needs to be connected to the Kyma Application connector.

Establish Connection (exchange certificates). This is generating new certificates and stores the connection metadata as well as the certificates in the `config`folder:

```bash
./connector connect -u <APPLICATION CONNECT URL>
```

The Open API Defenition and Event definition is part of the project. It needs to be registered in the Application Connector:

```bash
./connector register -a api-docs.json -e event-docs.json
```

This is registering a new service on kyma side. The compiled service definition is stored in the config folder using the service UUID (\<UUID\>.json). To inject the connection details (API Gateway URL, Credentials etc.) adjust the \<UUID\>.json file in the `config` folder by adding the required fields and update the service.

```bash
./connector update -i <UUID>
```

To run the server in production mode connect the BLE Dongle with the VM (in VMWare it is required to disable the "Share Bluetooth connection with VM option"). Start the server in none test mode referencing the drone  bluethooth name and config directory.

```bash
sudo KYMA_DRONE="Mambo_12345" KYMA_CONFIG=./config ./server --port 8080 --host 0.0.0.0
```

## Kyma Setup

`fly.js`: Simple Lambda functions which is starting the drone and landing it after a few seconds if called.

`takeoff.js`: Simple Lambda function to notify slack on takeoff. Subscribe to `drone.takeOff` event.

`shippackage.js`: Is shipping a package including the OrderID. Subscribe to `order.created` event from commerce.

`setcompletedstate.js`: Is setting the order completed date in commerce as soon as the drone is reporting the package as shipped. Subscribe to `drone.shipped` event.

## Troubleshooting

1. Drone is not connecting:

The server is using the gobot library to connect the drone. A good startingpoint for debuging is the example program documented on the gobot webside: <https://gobot.io/documentation/platforms/minidrone/>
