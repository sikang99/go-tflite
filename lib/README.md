## Build & install procedure (Apple M2 MacAir)

1. Download [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)
2. Compile tensorflow and its lite version on your platform
```
$ cd $srctree
$ ./configure
...
$ make 
$ cd $srctree/tensorflow
$ cmake $srctree/tensorflow/lite/c/
$ make 
```
3. copy tensorflow lite library to the library search folder
```
$ cp $srctree/tensorflow/libtensorflowlite_c.dylib /usr/local/lib
```

