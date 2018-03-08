# Brightfame SSH IAM

A lightweight client for authenticating SSH users using their IAM public keys.

## Features

 - Developers automatically upload their keys to IAM.
 - SSH authentication audit trail via CloudWatch.

All authorized keys are stored for the `ubuntu` user.

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
