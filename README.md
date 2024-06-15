# foundation app

a simple api server written in go

## Table of Contents

1. [Getting Started](#getting-started)
2. [API Documentation](#api-documentation)
   - [Create Nonprofit](#create-nonprofit)
   - [Bulk Send Emails](#bulk-send-emails)
   - [Get All Emails](#get-all-emails)
3. [Implementation Details](#implementation-details)
   - [Data Models](#data-models)
   - [API Handlers](#api-handlers)
   - [In-Memory Stores](#in-memory-stores)
   - [Server Setup](#server-setup)

## Getting Started

To get started with the Foundation App, ensure you have the following pre-requisites installed:

- Go (version 1.16 or newer)
- Git

To install Go, follow the instructions [here](https://golang.org/doc/install).

Now, you can clone the Foundation App repository and navigate to the project directory:

```
$ git clone https://github.com/ShimantaBhuyan/foundation-app
$ cd foundation-app
```

Next, build and run the application:

```
$ go build
$ ./foundation-app
```

After the application starts, you can use tools like ThunderClient in VSCode to test the API endpoints. A sample collection is also included in this repo.

## API Documentation

### Create Nonprofit

`POST /nonprofits`

Create a new nonprofit organization. Returns the created nonprofit with a unique ID.

Sample Request:

```json
{
  "name": "Humane Society",
  "email": "info@humanesociety.org",
  "street": "1234 Shelter Lane",
  "city": "New York",
  "state": "NY",
  "zipcode": "10001",
  "country": "USA"
}
```

### Bulk Send Emails

`POST /emails`

Send a list of emails using a customizable template for each recipient. The request body consists of an array of recipient email ids and the corresponding template string.

Sample Request:

```json
{
  "subject": "New payment for foundations",
  "templateString": "Disbursing to {name} | Address: {address}",
  "recipients": ["info@humanesociety.org", "info@savethechildren.org"]
}
```

### Get All Emails

`GET /emails`

Get the list of all emails sent by the foundations. Returns a list of email data, including the sender's foundation ID, the emails subject and the email body.

Sample Response:

```json
{
  "success": true,
  "data": [
    {
      "id": "cff5fe88-43ee-4d6d-b826-ecf1c13ba4eb",
      "Subject": "New payment for foundations",
      "Body": "Disbursing to Humane Society | Address: 1234 Shelter Lane, New York, NY, "
    }
  ],
  "error": ""
}
```
