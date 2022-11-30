# Build & Install procedure

1. Download [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)
2. Compile tensorflow and its ite
```
$ cd $srctree
$ ./configure
...
$ make 
$ cd $srctree/tensorflow
$ cmake $srctree/tensorflow/lite/c/
$ make 
```
3. copy $srctree/tensorflow/libtensorflowlite_c.dylib /usr/local/lib


