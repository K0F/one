# /dev/one

## Description
One writes the byte value 0x01 (which represents the character '1' in ASCII) to a device file `/dev/one`, creating a stream of ones.
This mimics the functionality of `/dev/zero`, but instead of null bytes, it produces a stream of ones.

go build -o one

sudo mknod /dev/one c 1 3
sudo ./one
