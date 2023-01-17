// Using basic auth
provider "awx" {
  hostname = "https://awx.172.21.0.2.sslip.io"
  #hostname = "http://localhost:8078"
  username = "test"
  password = "changeme"
}

// Using token auth
#provider "awx" {
#  hostname = "https://awx.172.21.0.2.sslip.io"
#  token = "myvalue"
#}
