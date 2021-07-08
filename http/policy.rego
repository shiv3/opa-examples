package http.authz

import input.attributes.request.http as http_request

default allow = false

allow {
	input.attributes.request.http.method == "GET"
	input.parsed_path == ["ping"]
}

# user_idが正数であることをチェック
user_id_check := user_id {
    some user_id_string
    input.parsed_path = ["user", user_id_string] # リクエストパスのuser_idを取得
    user_id := to_number(user_id_string) # user_idをnumberに
    is_number(user_id) # user_idがnumberかを確認
    user_id > 0 # user_idが正数かを確認
}

allow {
	input.attributes.request.http.method == "POST"
	user_id_check
}

# 認証が必要になる処理

rt = opa.runtime()
jwt_cert_string = rt.env.JWT_CERT

decoded_payload := payload {
    io.jwt.verify_hs256(bearer_token, jwt_cert_string)
    [_, payload, _] := io.jwt.decode(bearer_token)
}

bearer_token = t {
    v := input.attributes.request.http.headers.authorization
    startswith(v, "Bearer ")
    t := substring(v, count("Bearer "), -1)
}

allow {
	input.attributes.request.http.method == "GET"
	decoded_payload.user_id == user_id_check
}

allow {
	input.attributes.request.http.method == "GET"
	input.parsed_path == ["users"]
	decoded_payload.admin
}