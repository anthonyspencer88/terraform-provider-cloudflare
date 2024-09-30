---
page_title: "cloudflare_zero_trust_tunnel_cloudflared_virtual_network Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_tunnel_cloudflared_virtual_network (Resource)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) Cloudflare account ID
- `name` (String) A user-friendly name for the virtual network.

### Optional

- `comment` (String) Optional remark describing the virtual network.
- `is_default` (Boolean) If `true`, this virtual network is the default for the account.
- `is_default_network` (Boolean) If `true`, this virtual network is the default for the account.

### Read-Only

- `created_at` (String) Timestamp of when the resource was created.
- `deleted_at` (String) Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.
- `id` (String) UUID of the virtual network.

