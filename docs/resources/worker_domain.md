---
page_title: "cloudflare_worker_domain Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_worker_domain (Resource)



## Example Usage

```terraform
resource "cloudflare_worker_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "subdomain.example.com"
  service    = "my-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)
- `environment` (String) Worker environment associated with the zone and hostname.
- `hostname` (String) Hostname of the Worker Domain.
- `id` (String) Identifer of the Worker Domain.
- `service` (String) Worker service associated with the zone and hostname.
- `zone_id` (String) Identifier of the zone.

### Read-Only

- `zone_name` (String) Name of the zone.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_worker_domain.example <account_id>/<worker_domain_id>
```