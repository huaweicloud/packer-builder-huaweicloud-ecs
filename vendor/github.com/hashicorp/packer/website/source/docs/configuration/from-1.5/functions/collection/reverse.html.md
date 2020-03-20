---
layout: "docs"
page_title: "reverse - Functions - Configuration Language"
sidebar_current: "configuration-functions-collection-reverse"
description: |-
  The reverse function reverses a sequence.
---

# `reverse` Function


`reverse` takes a sequence and produces a new sequence of the same length
with all of the same elements as the given sequence but in reverse order.

## Examples

```
> reverse([1, 2, 3])
[
  3,
  2,
  1,
]
```

## Related Functions

* [`strrev`](../string/strrev.html) reverses a string.
