resource "cloudflare_dns_firewall" "example_dns_firewall" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "My Awesome DNS Firewall cluster"
  upstream_ips = ["192.0.2.1", "198.51.100.1", "2001:DB8:100::CF"]
  attack_mitigation = {
    enabled = true
    only_when_upstream_unhealthy = false
  }
  deprecate_any_requests = true
  ecs_fallback = false
  maximum_cache_ttl = 900
  minimum_cache_ttl = 60
  negative_cache_ttl = 900
  ratelimit = 600
  retries = 2
}
