# Build & Install procedure

1. Download [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)
2. Compile tensorflow lite
```
$ cd $srctree
$ ./configure
...
$ make 
$ cd $srctree/tensorflow/lite
$ cmake $srctree/tensorflow/lite/c/
$ make 
```
3. copy $srctree/tensorflow/tensorflow/libtensorflowlite_c.dylib /usr/local/lib


