# b64

### ```b64``` is a command for encoding given string/file into base64 string and vice-versa

## Build
### ```go build -ldflags="-s -w"```

## Command Usage
### Print usage
### ```b64 -h```
### Output:
![Usage](https://github.com/thakurabhiv/b64/blob/main/screenshots/usage.png)
##

### Encode string
### ```b64 "some string"```
### Output:
![Encoding](https://github.com/thakurabhiv/b64/blob/main/screenshots/encoding_normal.png)
##

### Encode file
### ```b64 -f clipboard.go```
### File content
![File Content](https://github.com/thakurabhiv/b64/blob/main/screenshots/file_content.png)
### Output:
![File Encoding](https://github.com/thakurabhiv/b64/blob/main/screenshots/file_enoding.png)

# Adding More