# gofetch
This could have been a bash one-liner but guess what. It's a small Go tool that lists the trending CVEs from cvetrends.com

By default, it will output just the CVE IDs so it's easier to pipe into your workflows.

Optionally, you can use the `-info` flag to print the CVE descriptions and severities (if assigned).

# Install

```bash
go install -v github.com/tr3ss/gofetch@latest
```

# Usage

```zsh
❯ gofetch
Usage: gofetch day|week
```

## Get top 10 CVEs trending in the last 24 hours:

```
❯ gofetch day
CVE-2022-2143
CVE-2022-32788
CVE-2022-29455
CVE-2022-21571
CVE-2021-30737
CVE-2022-34907
CVE-2022-34906
CVE-2022-34576
CVE-2022-36375
CVE-2022-34575
```

## Get top 10 CVEs trending in the last week:

```
❯ gofetch week
CVE-2022-26138
CVE-2022-34918
CVE-2022-22047
CVE-2022-33891
CVE-2022-2294
CVE-2022-20857
CVE-2022-29455
CVE-2022-34169
CVE-2022-20861
CVE-2022-22280
```

## Get CVE info

```
❯ gofetch day -info
CVE-2022-27924
 Description:  Zimbra Collaboration (aka ZCS) 8.8.15 and 9.0 allows an unauthenticated attacker to inject arbitrary memcache commands into a targeted instance. These memcache commands becomes unescaped, causing an overwrite of arbitrary cached entries.
 Severity : HIGH
-----------------------------------
CVE-2022-29582
 Description:  In the Linux kernel before 5.17.3, fs/io_uring.c has a use-after-free due to a race condition in io_uring timeouts. This can be triggered by a local user who has no access to any user namespace; however, the race condition perhaps can only be exploited infrequently.
 Severity : HIGH
 -----------------------------------
 
 [...]
 ```
