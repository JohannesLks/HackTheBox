# AES Decryption Tool for Napper Challenge

This repository contains a Go script developed to assist in decrypting data for the "Napper" challenge on Hack The Box. It's designed to decrypt data that was encrypted using AES-128 in CFB mode. The decryption process requires a seed to regenerate the encryption key and the encrypted data in base64 format.

## Overview

The script leverages the AES cryptographic standard with a Cipher Feedback (CFB) mode of operation, providing secure decryption capabilities. It accepts a numerical seed and base64-encoded encrypted data as input, both of which are necessary to accurately reconstruct the original plaintext.

## Prerequisites

To use this decryption tool, you will need:

- Go (Golang) installed on your system (version 1.14 or newer recommended).

## Usage

The script is executed from the command line, where you will pass the seed and encrypted data as arguments. Here's the basic syntax:

```bash
go run decrypt.go -seed=<seed> -data="<base64-encoded-data>"
```
## Example

To decrypt data with a seed of 46385390 and encrypted data tbjZvSCUhZtSmOqEYO1TFmX-ibTWLnMJc6CQJHZ_aM6alBTptvEaiMEvjv_Jfx33T7spOEMKOXg=, run the following command:

```bash
go run decrypt.go -seed=46385390 -data="tbjZvSCUhZtSmOqEYO1TFmX-ibTWLnMJc6CQJHZ_aM6alBTptvEaiMEvjv_Jfx33T7spOEMKOXg="
```