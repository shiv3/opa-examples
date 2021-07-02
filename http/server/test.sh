curl --location --request POST 'http://localhost:1323/user/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "name": "test user 1",
    "email": "test1@test.com"
}'


curl --location --request POST 'http://localhost:1323/user/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 2,
    "name": "test user 2",
    "email": "test2@test.com"
}'

