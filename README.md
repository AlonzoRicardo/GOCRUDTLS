# go_auth
go CRUD TLS server + SQLITE3 db

## Helper Content
[Production Ready Go Service in 30 Minutes](https://www.youtube.com/watch?v=wxkEQxvxs3w)         
[gopherconuk](https://github.com/dlsniper/gopherconuk)         
[Exposing Go to the Internet](https://blog.cloudflare.com/exposing-go-on-the-internet/)      
[Style guideline for Go packages](https://rakyll.org/style-packages/)       
[Creating and managing CA certs with mkcert](https://github.com/FiloSottile/mkcert)       
[GO TLS server examples](https://gist.github.com/denji/12b3a568f092ab951456)        
[ORM sqlite](https://www.youtube.com/watch?v=VAGodyl84OY) - [code](https://gist.github.com/elliotforbes/e241eaa8cc9d7bf3ec75b333e891d422)       
[sqlite3 in a docker image](https://devopsheaven.com/sqlite/backup/restore/dump/databases/docker/2017/10/10/sqlite-backup-restore-docker.html)

# cURL
Create a self-signed certificate and tell cURL where to look for it
```
mkcert localhost
curl --capath $(mkcert -CAROOT) https://localhost:8080/
```

# Tips
- As a rule of thumb, keep types closer to where they are used. This makes it easy for any maintainer (not just the original author) to find a type. 

- A common practise from other languages is to organize types together in a package called models or types. In Go, we organize code by their functional responsibilities.

- Main packages are not importable, so exporting identifiers from main packages is unnecessary. Don't export identifiers from a main package if you are building the package to a binary.

- Package names should be lowercase. Don't use snake_case or camelCase in package names. The Go blog has a comprehensive [guide](https://blog.golang.org/package-names) about naming packages with a good variety of examples.

- In go, package names are not plural. This is surprising to programmers who came from other languages and are retaining an old habit of pluralizing names. Don't name a package httputils, but httputil!

- Always document the package. Package documentation is a top-level comment immediately preceding the package clause.



