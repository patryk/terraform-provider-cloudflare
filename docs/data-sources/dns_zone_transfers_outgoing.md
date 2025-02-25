---
page_title: "cloudflare_dns_zone_transfers_outgoing Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_dns_zone_transfers_outgoing (Data Source)



## Example Usage

```terraform
data "cloudflare_dns_zone_transfers_outgoing" "example_dns_zone_transfers_outgoing" {
  zone_id = "269d8f4853475ca241c4e730be286b20"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String)

### Read-Only

- `checked_time` (String) The time for a specific event.
- `created_time` (String) The time for a specific event.
- `id` (String) The ID of this resource.
- `last_transferred_time` (String) The time for a specific event.
- `name` (String) Zone name.
- `peers` (List of String) A list of peer tags.
- `soa_serial` (Number) The serial number of the SOA for the given zone.


