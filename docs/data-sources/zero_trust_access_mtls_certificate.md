---
page_title: "cloudflare_zero_trust_access_mtls_certificate Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_access_mtls_certificate (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `certificate_id` (String) UUID
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `associated_hostnames` (List of String) The hostnames of the applications that will use this certificate.
- `created_at` (String)
- `expires_on` (String)
- `fingerprint` (String) The MD5 fingerprint of the certificate.
- `id` (String) The ID of the application that will use this certificate.
- `name` (String) The name of the certificate.
- `updated_at` (String)

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Optional:

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

