### URL-Shortner-Service
URL shortner service using redis as database. Storing shortened URL's are stored in the file `pkg/repository/filestorage/url-data.txt`

#### Run
1. make build
2. make test
#### APIs
 
0. Generate short URL.<br/>
   `Method:` POST <br/>
   `URL:` http://localhost:7777/api/shorturls <br/>
   `Body:` ```{
   "long_url": "https://www.google.com",
   "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
   }``` <br/>
<br/>

1. `Method:` GET <br/>
   `URL:` http://localhost:7777/api/shorturls/:bPlNFG <br/>

<details>
<summary>References </summary>
    A. https://github.com/bxcodec/go-clean-arch <br/>
    B. https://www.eddywm.com/
</details>
