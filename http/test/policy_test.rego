package http.authz

test_ping {
    allow with input as curl_ping_input
}

test_ping_err {
    not allow with input as curl_ping_input_post
}

test_post_user {
  allow with input as input_curl_post_user_pass
  not allow with input as input_curl_post_user_err_1
  not allow with input as input_curl_post_user_err_2
}

test_get_user {
  allow with input as input_jwt_header_request
}

test_get_users {
  allow with input as input_jwt_admin_header_request
  not allow with input as input_jwt_admin_header_request_err
}