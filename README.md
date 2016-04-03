Multi-Fire
==========
Allows for multiple concurrent URLs to be fired off from a single client connection in order to decrease the processing responsibilities of the client.

Setup
=====
1. Run `go get` to retrieve packages
2. Run `go install` to compile
3. Run `go build && ./Multi-Fire` to build and run
4. Access at `http://localhost:8080`


Nomenclature
============
* `twig`: A URL
* `kindling`: A set of twigs
* `kindlings`: Multiple sets of kindling
* `ignite`: To trigger a firing of a set of kindling

Usage
=====

/add/{tbd}
----------
Add new set of URLs (kindling). Returns kindlingId

`http://localhost:8080/add`

Payload
```
[
    {
        "url":"http://foo.com"
    },
    {
        "url":"http://foo2.com"
    },
    {
        "url":"http://foo3.com"
    }
]
```

/get/{kindlingId}
-----------------
Get existing set of URLs (kindling)

`http://localhost:8080/get/{kindlingId}`

Response
```
{
  "kindling_id": 1,
  "twigs": [
    {
      "url": "http://foo.com"
    },
    {
      "url": "http://foo2.com"
    },
    {
      "url": "http://foo3.com"
    }
  ]
}
```

/ignite/{kindlingId}
--------------------
NOTE: This hasn't been implemented yet
Fire URLS (kindling) by kindlingId

`http://localhost:8080/ignite/{kindlingId}`
