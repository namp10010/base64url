# base64url
a no-brainer base64url

### install

```golang
  go install https://github.com/namp10010/base64url
```

### usage

encrypt

```bash
  echo hello | base64url
```

decrypt

```bash
  echo hello | base64url -d
```

url encrypt

```bash
  echo hello | base64url -u
```

url decrypt

```bash
  echo hello | base64url -u -d
```
