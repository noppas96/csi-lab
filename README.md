# DotGit

FETCH_HEAD (FILE)
```
Example

b19cf10870d984d702b68ed845f37488323d0183                branch 'master' of https://github.com/kubernetes-csi/csi-driver-nfs
67f58120dfcdecd54d98a96c3023442a3f08ed75        not-for-merge   branch 'prow-update-master' of https://github.com/kubernetes-csi/csi-driver-nfs
e4122cc01d3a25d6fdd272759c5e473bd934ab32        not-for-merge   branch 'release-3.0' of https://github.com/kubernetes-csi/csi-driver-nfs
c0bf858eab3b14b3200604488f62091f0b32da3b        not-for-merge   branch 'release-3.1' of https://github.com/kubernetes-csi/csi-driver-nfs
78a6937ae2602143eb1c6cb5fea5ca8cf8d80d8f        not-for-merge   branch 'release-4.0' of https://github.com/kubernetes-csi/csi-driver-nfs
6a119cc144fa394fde03b69da628f555d1f16fb5        not-for-merge   branch 'release_tools2' of https://github.com/kubernetes-csi/csi-driver-nfs
```
HEAD (FILE)
```
Example

ref: refs/heads/master
```

branches (DIR)
config (FILE)
```
Example
[core]
        repositoryformatversion = 0
        filemode = true
        bare = false
[remote "origin"]
        url = https://github.com/kubernetes-csi/csi-driver-nfs.git
        fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
        remote = origin
        merge = refs/heads/master
```
description (FILE)
```
Example
Unnamed repository; edit this file 'description' to name the repository.
```
hooks (DIR)
index (binary)
info (DIR)
```
exlude(file) 
# git ls-files --others --exclude-from=.git/info/exclude
# Lines that start with '#' are comments.
# For a project mostly in C, the following would be a good set of
# exclude patterns (uncomment them if you want to use them):
# *.[oa]
# *~
```
log(DIR)
objects(DIR)
- info(DIR)
- pack(DIR)
packed-refs
```
Example
# pack-refs with: peeled fully-peeled sorted 
0c0f0541fbf78e3f356e1c4ad3e61a00b0c1525e refs/remotes/origin/master
67f58120dfcdecd54d98a96c3023442a3f08ed75 refs/remotes/origin/prow-update-master
e4122cc01d3a25d6fdd272759c5e473bd934ab32 refs/remotes/origin/release-3.0
c0bf858eab3b14b3200604488f62091f0b32da3b refs/remotes/origin/release-3.1
78a6937ae2602143eb1c6cb5fea5ca8cf8d80d8f refs/remotes/origin/release-4.0
6a119cc144fa394fde03b69da628f555d1f16fb5 refs/remotes/origin/release_tools2
8cdbbe13229fb82e9b0f727065d69e4423dddb7c refs/tags/v2.0.0
d1409ab179feb8ee5cffcd63fd236b5bc69c39de refs/tags/v3.0.0
79d4b0eb1cadbd827fbcba689cbc958ee33972e1 refs/tags/v3.1.0
78a6937ae2602143eb1c6cb5fea5ca8cf8d80d8f refs/tags/v4.0.0
935f29fc129f86e0a9c744e40825fb8706eeded3 refs/tags/v4.1.0
```
refs(DIR)

