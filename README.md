# Placeholder
Simple application for adding watermarks to images

![Clean image](images/test.png)
![Watermarked image](images/output.png)

## Usage
Normal verions:
```console
placeholder [flags] file
  -c    Add watermark in red
  -t    Tilt watermark
```
Simplified:
```console
placeholder file
```

## Building
### Full
To build normal version use branch `master`. 
For now, if you want to use your own watermark, replace `watermark` and `watermarkAngled` variables with output of the following command
```console
cat your_watermark.png | base64 -w 0
```
and compile
```console
go build placeholder.go
```
### Simplified
For simplified, checkout branch `simplified`:
```console
git checkout simplified
```
This version is good for windows users, as it does not take any parameters and chosen image can just be dragged onto the program executable to watermark it:
```console
GOOS=windows GOARCH=amd64 go build -o placeholder.exe placeholder.go
```