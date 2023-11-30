# aws-role

This is a simple utility to load aws credentials into a bash (or similar)
environment. What makes this utility different from others is that it will
display a list of all known profiles to choose from. This small change does
not make it better than other utilities, but it simply behaves the way I
want it to.

## Usage

Create an alias to execute: `alias creds='eval $(aws-role)'` and add it to
`~/.bash_profile` or `~/.bashrc`.

```console
$ creds
   1. bar-prod
   2. bar-staging
   3. foo-prod
   4. foo-staging

Select: 4
MFA Code: 123456
Duration in hours [1]:
AWS Creds Loaded!
```
