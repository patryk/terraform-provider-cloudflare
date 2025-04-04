
resource "cloudflare_spectrum_application" "%[3]s" {
  zone_id  = "%[1]s"
  protocol = "tcp/22"

  dns = {
  type = "ADDRESS"
    name = "%[3]s.%[2]s"
}

  origin_direct = ["tcp://128.66.0.4:23"]
  origin_port   = 22
  edge_ips = {
  type = "static"
	ips = ["172.65.64.13"]
}
}