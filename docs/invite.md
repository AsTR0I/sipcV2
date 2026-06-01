# SIP INVITE

Send SIP INVITE request

---

## Description

INVITE is used to initiate a SIP call between two users.

---

| Flag          | Required | Description     |
| ------------- | -------- | --------------- |
| --user-port   | no       | Local SIP port  |
| --server-host | yes      | SIP server host |
| --proxy       | no       | SIP proxy       |
| --from        | no       | Caller SIP user |
| --to          | no       | Called SIP user |
| --realm       | no       | Auth realm      |
| --username    | yes      | Auth username   |
| --password    | yes      | Auth password   |
| --user-agent  | no       | SIP User-Agent  |

## Usage

```bash
sipc invite \
  --server-host sip.astro.local:5060 \
  --from 1001 \
  --to 1002 \
  --username 1001 \
  --password secret \
  --realm astro.local \
  --log-level debug
```
