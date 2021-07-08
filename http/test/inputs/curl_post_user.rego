package http.authz

input_curl_post_user_pass = {
  "attributes": {
      "request": {
          "http": {
              "method": "POST"
          }
      }
  },
  "parsed_path": [
      "user",
      "1"
  ],
  "parsed_query": {}
}

input_curl_post_user_err_1 = {
  "attributes": {
      "request": {
          "http": {
              "method": "POST"
          }
      }
  },
  "parsed_path": [
      "user",
      "-1"
  ],
  "parsed_query": {}
}

input_curl_post_user_err_2 = {
  "attributes": {
      "request": {
          "http": {
              "method": "POST"
          }
      }
  },
  "parsed_path": [
      "user",
      "a"
  ],
  "parsed_query": {}
}