### URL-Shortner-Service

#### Run

0. Build docker image: `docker build --rm -t url-shortner-service -f build/Dockerfile .`
1. `docker-compose -f ./build/docker-compose.yml up -d`

#### APIs
 
0. Generate short URL.<br/>
   `Method:` POST <br/>
   `URL:` http://localhost:7777/create-short-url\ <br/>
   `Body:` ```{
   "long_url": "https://www.google.com",
   "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
   }``` <br/>
<br/>

1. `Method:` GET <br/>
   `URL:` http://localhost:7777/shortUrl/:bPlNFG <br/>

<details>
<summary>References </summary>
    A. https://github.com/bxcodec/go-clean-arch <br/>
    B. https://www.eddywm.com/
</details>