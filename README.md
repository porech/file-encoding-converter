# file-encoding-converter

A simple utility with one simple purpose: convert a file from one charset encoding to another.

It can be a one-time conversion, or it can watch an input file and convert it to a new encoding in another file everytime it changes.

## Download
We don't have versioning and well-done releases yet. However, you can find pre-build binaries for various architectures in the `dist` folder of this repository. If you need another architecture, just open an issue and we'll add it.

## Usage
Required parameters:
- `-inputFile`: the path of the input file
- `-inputEncoding`: the encoding of the input file, must be one of the accepted names here below
- `-outputFile`: the path to save the converted file to
- `-outputEncoding`: the encoding of the output file, must be one of the accepted names here below

Optional parameters:
- `-watch`: instead of exiting immediately after conversion, watch for the input file and convert it everytime it changes

## Build
If you want to build the project manually, be sure to have Golang installed.
Then, clone the project and run `make` to build for all the supported architectures in the `dist` folder.

## Supported encodings
|                   Description |                           Accepted names |
| ----------------------------- | ---------------------------------------- |
| UTF-8 |                              utf-8, utf8 |
| UTF-8 with BOM |                       utf-8-bom, utf8bom |
| UTF-16 Little Endian |                       utf-16-le, utf16le |
| UTF-16 Little Endian with BOM |                utf-16-le-bom, utf16lebom |
| UTF-16 Big Endian |                       utf-16-be, utf16be |
| UTF-16 Big Endian with BOM |                utf-16-be-bom, utf16bebom |
| Windows 874 |          windows-874, windows874, win874 |
| Windows 1250 |       windows-1250, windows1250, win1250 |
| Windows 1251 |       windows-1251, windows1251, win1251 |
| Windows 1252 |       windows-1252, windows1252, win1252 |
| Windows 1253 |       windows-1253, windows1253, win1253 |
| Windows 1254 |       windows-1254, windows1254, win1254 |
| Windows 1255 |       windows-1255, windows1255, win1255 |
| Windows 1256 |       windows-1256, windows1256, win1256 |
| Windows 1257 |       windows-1257, windows1257, win1257 |
| Windows 1258 |       windows-1258, windows1258, win1258 |
| IBM Code Page 037 |                       codepage037, cp037 |
| IBM Code Page 437 |                       codepage437, cp437 |
| IBM Code Page 850 |                       codepage850, cp850 |
| IBM Code Page 852 |                       codepage852, cp852 |
| IBM Code Page 855 |                       codepage855, cp855 |
| IBM Code Page 858 |                       codepage858, cp858 |
| IBM Code Page 860 |                       codepage860, cp860 |
| IBM Code Page 862 |                       codepage862, cp862 |
| IBM Code Page 863 |                       codepage863, cp863 |
| IBM Code Page 865 |                       codepage865, cp865 |
| IBM Code Page 866 |                       codepage866, cp866 |
| IBM Code Page 1047 |                     codepage1047, cp1047 |
| IBM Code Page 1140 |                     codepage1140, cp1047 |
| ISO 8859-1 |          iso-8859-1, iso-latin-1, latin1 |
| ISO 8859-2 |          iso-8859-2, iso-latin-2, latin2 |
| ISO 8859-3 |          iso-8859-3, iso-latin-3, latin3 |
| ISO 8859-4 |          iso-8859-4, iso-latin-4, latin4 |
| ISO 8859-5 | iso-8859-5, iso-latin-cyrillic, cyrillic |
| ISO 8859-6 |     iso-8859-6, iso-latin-arabic, arabic |
| ISO 8859-7 |       iso-8859-7, iso-latin-greek, greek |
| ISO 8859-8 |     iso-8859-8, iso-latin-hebrew, hebrew |
| ISO 8859-9 |          iso-8859-9, iso-latin-5, latin5 |
| ISO 8859-10 |         iso-8859-10, iso-latin-6, latin6 |
| ISO 8859-13 |                              iso-8859-13 |
| ISO 8859-14 |                              iso-8859-14 |
| ISO 8859-15 |                              iso-8859-15 |
| ISO 8859-16 |                              iso-8859-16 |
| KOI8-R |                            koi8-r, koi8r |
| KOI8-U |                            koi8-u, koi8u |
| Macintosh |                           macintosh, mac |
| Macintosh Cyrillic |         macintosh-cyrillic, mac-cyrillic |
| X-User-Defined |                      x-user-defined, xud |
| FM radio RDS, G0 table |                              rds-g0, rds |
