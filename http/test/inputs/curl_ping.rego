package http.authz

curl_ping_input = {
  "attributes": {
      "request": {
          "http": {
              "method": "GET"
          }
      }
  },
  "parsed_path": [
      "ping"
  ],
  "parsed_query": {}
}

curl_ping_input_post = {
  "attributes": {
      "request": {
          "http": {
              "method": "POST" # /ping はPOSTを受け付けない
          }
      }
  },
  "parsed_path": [
      "ping"
  ],
  "parsed_query": {}
}