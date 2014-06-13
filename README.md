ssdbAdmin
=========

![image](https://github.com/jiorry/ssdbAdmin/raw/master/release/ssdbAdmin.png)

## Run
find [realese] folder and unzip application package.
modify app/app.conf to connect ssdb server

```
[ssdb]
host=127.0.0.1
port=8888
```

## Installation
Make sure you have the a working Go environment. See the [install instructions](http://golang.org/doc/install.html).

First, you must install dependence:

	go get github.com/jiorry/gos

And install ssdbAdmin, simply run:

	go get github.com/jiorry/ssdbAdmin

Or compile it from source:
	
	git clone git://github.com/jiorry/ssdbAdmin
    cd ssdbAdmin && go build

if windows, you only need double click build.bat

## Good Luck
