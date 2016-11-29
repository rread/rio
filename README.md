# rio - IO Benchmark

This is a simple benchmark measures sequential IO along the lines of the "dd"
command. It does a write and read test to a temp file in either the local
subdirectory or the the subdirectory specified by the -dir parameter. The -size
parameter is the size of the file and the -block specifies the block size to use
when doing IOs. The data written are blocks full of zeroes.  The parameters
accept suffixes like "kb" or MiB".

Currently this tool attempts to purge memory between the Write and Read steps, to attempt to eliminate caching from influencing the benchmark, though this has only been implemented on Linux and Mac. On both platforms it uses "sudo", so this will prompt for password when it runs. This is annoying and needs to be fixed somehow.

## Example

    $ ./rio -size 16GiB -block 128Kib
    WRITE 16 GiB 128 KiB/io 8.81s 14873.69 IO/s 1859.21 MiB/s
    Password:
    READ 16 GiB 128 KiB/io 8.36s 15684.30 IO/s 1960.54 MiB/s
