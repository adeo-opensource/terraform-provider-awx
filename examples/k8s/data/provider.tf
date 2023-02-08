provider "awx" {
  hostname = "https://awx.172.21.0.2.sslip.io"
  username = "test"
  password = "changeme" # pragma: allowlist secret
}

// Using token auth
#provider "awx" {
#  hostname = "https://awx.172.21.0.2.sslip.io"
#  token = "value"
#}
