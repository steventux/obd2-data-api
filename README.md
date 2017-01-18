# OBD2 Data API

A Torque Pro centric OBD2 data logging API written in Go and backed by MongoDB.

## Installation

This app expects to read your secrets from the following environment variables:

 - `MONGODB_HOSTS` - Hostnames for your MongoDB instance or cluster
 - `MONGODB_DB`    - The database to use
 - `MONGODB_COLLECTION` - The collection name to use
 - `TORQUE_ID`     - The Torque Pro application ID from the device sending the data

## Endpoints

 - `GET /obd2-data/create` The data upload endpoint, it's a GET because Torque Pro.
 - `GET /obd2-data/show/:session_id`   (NOT IMPLEMENTED) Endpoint to view data.

## Tests

You need to specify the environment variables above a la:

```
MONGODB_HOST=localhost MONGODB_DATABASE=test MONGODB_COLLECTION=obd2_data make
```
