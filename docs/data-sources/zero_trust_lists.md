---
page_title: "cloudflare_zero_trust_lists Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_lists (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)

### Optional

- `max_items` (Number) Max items to fetch, default: 1000
- `type` (String) The type of list.

### Read-Only

- `result` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `created_at` (String)
- `description` (String) The description of the list.
- `id` (String) API Resource UUID tag.
- `list_count` (Number) The number of items in the list.
- `name` (String) The name of the list.
- `type` (String) The type of list.
- `updated_at` (String)

