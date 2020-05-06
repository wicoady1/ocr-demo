# OCR Demo
A script for small demo on how we utilize Tesseract OCR in Golang.
This script was used for small demo on how OCR works in Golang, during [April 2020 Golang Singapore Meetup](https://engineers.sg/video/optical-character-recognition-ocr-implementation-with-golang-gosg--4031)

------

## Prerequistes
- [Tesseract OCR](https://tesseract-ocr.github.io/tessdoc/Home.html) (4.1.1)
**Recommended** to use the exact version, since different version may causing different result and accuracy.
- [Golang](https://golang.org/) (1.13.4 or Above)

## Supporting Libraries
- [Gosseract](https://github.com/otiai10/gosseract)

## How to Build
1. Make sure you already successfully installed Tesseract OCR and functioning properly via Command Prompt
You can test it with `tesseract -v`
And it will output something like this
```
tesseract 4.1.1
 leptonica-1.79.0
  libgif 5.2.1 : libjpeg 9d : libpng 1.6.37 : libtiff 4.1.0 : zlib 1.2.11 : libwebp 1.1.0 : libopenjp2 2.3.1
 Found AVX2
 Found AVX
 Found FMA
 Found SSE
```
2. Git clone this repository into your Go working path
3. `go get && go build && ./ocr-demo`

## Documentation
More documentation, you can check [Gosseract's godoc](https://godoc.org/github.com/otiai10/gosseract)

## Author
Kennard Wicoady (wicoady1)