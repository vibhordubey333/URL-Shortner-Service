import requests
import json
def URLShortnerPOST():
  data = '{"long_url": "https://www.google.com","user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"}'
  url = 'http://localhost:7777/api/shorturls'
  response = requests.post(url,data=data,headers={"Content-Type":"application/json"})
  print ("Code:"+ str(response.status_code))
  print("Response:"+ str(response.text))

def URLShortnerGET():
  url = 'http://localhost:7777/api/shorturls/:bPlNFG'
  response = requests.get(url,headers={"Content-Type":"application/json"})
  print ("GET Code:"+ str(response.status_code))
  print("Response:"+ str(response.text))

URLShortnerPOST()
URLShortnerGET()
