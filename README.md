Hosts
===

A modest application for storing details about a bunch of random servers or virtual servers that are distributed around the world.  It is a *pita* to keep track of where and who each is hosted.

This was put together while on a 1h 32m flight from SEA to SJC; so don't have high expections.

## API
Currently there is only an API mapped to a mock database served from `db.go` that contains simplified host information for initial development.

* GET `/` - bah
* GET `/hosts` - list hosts 
* GET `/hosts/:id` - show host details

