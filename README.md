Simple ZIP SFX stub in Go
=========================

This program uses archive/zip from the Go standard library to extract
a ZIP file that has been appended to the binary. The offset of the ZIP
archive is located by searching for 8 magic bytes (`50 4b 03 04 14 00
08 00`).

It has been tested with Linux and Windows (but should work on any
other platform, too).

Building / Usage
----------------

    $ go get github.com/edsrzf/mmap-go
    $ go build
    $ cat go-zipsfx archive.zip > archive.bin && chmod 755 archive.bin

Or, for Windows:

    $ GOOS=windows go build
    $ cat go-zipsfx.exe archive.zip > archive.exe

The resulting program accepts one command line parameter, `-d`, to
specify a destination directory into which the archive contents are
extracted.

License
-------

Copyright (C) 2017  Hilko Bengen <bengen@hilluzination.de>

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
