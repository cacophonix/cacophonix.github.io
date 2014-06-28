---
layout: post
title: "testing domain names in localhost"
date: 2014-06-22 19:39:52 +0600
comments: true
categories: [Linux]
---
I was testing my local Apache configuration. In those tutorials the examples were using domain names like
www.example.com or www.test.com. I was wondering how can i test these examples. then I just googled and found
that I can easily test them in my PC without purchasing the domain names and configuring DNS.

the magic thing is there is a file in /etc/hosts in my ubuntu that contains
``` bash

127.0.0.1  localhost
```

The format here is
``` bash
<ip-address> <space> <site-domain>
```

If I add
``` bash
127.1.1.1 test.com
```

then if I write in the address bar of my browser "test.com" then the browser will automatically load the 127.1.1.1 ip address.

Isn't it awesome ?


