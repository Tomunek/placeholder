# Placeholder
Simple application for adding watermarks to images

![Clean image](images/test.png)
![Watermarked image](images/output.png)

## Usage
```console
# coming soon
```

## Examples
```console
# coming soon
```

## Building
For now, if you want to use your own watermark, replace `watermark` and `watermarkAngled` variables with output of the following command
```console
cat your_watermark.png | base64 -w 0
```
and compile
```console
go build placeholder.go
```