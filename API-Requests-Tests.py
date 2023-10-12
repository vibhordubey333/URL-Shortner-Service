import requests

def URLShortnerPOST():
  data = '{"long_url": "https://www.google.com","user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"}'
  url = 'http://0.0.0.0:7777/api/shorturls'
  headers = {"ACCEPT":"application/json"}
  response = requests.get(url,data = data)
  print ("Code:"+ str(response.status_code))
  print("Response:"+ str(response.text))

URLShortnerPOST()
