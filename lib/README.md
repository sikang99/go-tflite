# Build & Install procedure

1. Download [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)
2. Compile tensorflow lite
```
$ cd $srctree/tensorflow/tensorflow
$ ./configure
$ make 
$ cd github.com/tensorflow/tensorflow/tensorflow/lite
$ cmake $srctree/tensorflow/tensorflow/lite/c/
$ make 
```
3. copy libtensorflowlite_c.dylib /usr/local/lib


