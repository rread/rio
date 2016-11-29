# rio - IO Benchmark

This is a simple benchmark measures sequential IO along the lines of the "dd"
command. It does a write and read test to a temp file in either the local
subdirectory or the the subdirectory specified by the -dir parameter. The -size
parameter is the size of the file and the -block specifies the block size to use
when doing IOs. The data written are blocks full of zeroes.  The parameters
accept suffixes like "kb" or MiB".
