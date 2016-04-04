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

/add
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
Fire each twig/URL from set of kindling. 
Note: Response will be an empty GIF.

Example Usage

`<img src="http://{host}/ignite/234" border="0" />`
