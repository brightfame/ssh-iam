# Brightfame SSH IAM

A light weight client for authenticating SSH users using their IAM public keys.

## Features

 - Developers automatically upload their keys to IAM.
 - SSH authentication audit trail via CloudWatch.

## Usage

```
ssh-iam list-keys rob
ssh-iam install
ssh-iam sync
```

# Development

```
docker build -t brightfame/ssh-iam
```
