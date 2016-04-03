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
* `kindling`: A set of URLs
* `kindlings`: Multiple sets of kindling
* `ignite`: To trigger a firing of a set of kindling

Usage
=====

/ignite/{kindlingId}
--------------------
Fire URLS (kindling) by kindlingId

```http://localhost:8080/ignite/{kindlingId}```

/get/{kindlingId}
-----------------
Get existing set of URLs (kindling)

```http://localhost:8080/get/{kindlingId}```

/add/{tbd}
----------
Add new set of URLs (kindling). Returns kindlingId

```http://localhost:8080/add/{tbd}```