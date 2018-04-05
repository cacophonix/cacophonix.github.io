---
layout: post
title: "In case if i forget how to setup octopress in another machine"
date: 2018-02-12 02:51:51 +0600
comments: true
categories: ["blogging"]
---

Recently I had not been in touch with my blog and  forgot the whole process of using my blog in a new machine. 

Initialization
    git clone -b source git@github.com:cacophonix/cacophonix.github.io octopress
    cd octopress
    git clone git@github.com:cacophonix/cacophonix.github.io _deploy
    gem install bundler
    rbenv rehash    # If you use rbenv, rehash to be able to run the bundle command
    bundle install
    rake setup_github_pages

Pushing changes from one machine
    rake generate
    git add .
    git commit -am "Some comment here."
    git push origin source  # update the remote source branch
    rake deploy             # update the remote master branch
    
From another machine
    cd octopress
    git pull origin source  # update the local source branch
    cd ./_deploy
    git pull origin master  # update the local master branch


```
http://blog.zerosharp.com/clone-your-octopress-to-blog-from-two-places/
```
or this

```
http://stackoverflow.com/questions/20765692/how-to-setup-octopress-if-you-already-have-one-in-github
```
 I am just writing it down in case i forgot again.
