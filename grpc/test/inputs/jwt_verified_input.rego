package grpc.authz

user_id_1_token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcl9pZCI6MSwiaWF0IjoxNTE2MjM5MDIyfQ.DuIbFMj_3Hoak3s7jb7B1-KYjpZlQjxvwhAc6nE-cYA"

input_jwt_header_request = {
 "attributes": {
     "request": {
         "http": {
             ":path": "/server.GPRCServiceAPI/ShowUser",
             "headers": {
                 "authorization": user_id_1_token
             }
         }
     }
 },
 "parsed_body": {
     "id": "1"
 },
 "parsed_path": [
     "server.GPRCServiceAPI",
     "ShowUser"
 ],
}

user_id_0_admin_token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwidXNlcl9pZCI6MCwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.auwl26C0A3b2PeN1pjHow2GvsIwANibLSWv6r6pn6K4"

input_jwt_admin_header_request = {
 "attributes": {
     "request": {
         "http": {
             ":path": "/server.GPRCServiceAPI/ShowUsers",
             "headers": {
                 "authorization": user_id_0_admin_token
             }
         }
     }
 },
 "parsed_body": {},
  "parsed_path": [
      "server.GPRCServiceAPI",
      "ShowUsers"
  ],
}

