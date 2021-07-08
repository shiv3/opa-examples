package grpc.authz


test_get_user {
  allow with input as input_jwt_header_request
}

test_get_users {
  allow with input as input_jwt_admin_header_request
}