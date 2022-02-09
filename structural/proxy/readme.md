# Structural patterns in Go: Proxy

Sometimes we need to add more control over our objects usage. We can easily do this using Proxy: some kind of layer-on-top of the real object which acts like it, but provides additional functionality. 

It should implement and provide the same interface as our original object, meaning we can use them interchangeably. The calling code should not tell the difference. The proxy contains the original inside, it calls the original when it finishes doing its extra stuff.

Proxy is useful when we need to add something *before* and/or *after* we call our original object methods, but we don't want to change the original. So we wrap it up and add new stuff to the proxy. We can add a lot of things, from security-related (check permissions, check transaction), logging, profiling to caching and lazy-init.