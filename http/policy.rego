package envoy.authz

import input.attributes.request.http as http_request

default allow = false

allow {
	input.attributes.request.http.method == "GET"
	input.attributes.request.http.path == "/ping"
}

