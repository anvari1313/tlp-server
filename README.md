# TLP Server

This is a transportation layer proxy server that can proxy udp to tcp and tcp to udp.

## HTTP on UDP

The main goal of this server is to receive http request on udp and retransmit it with tcp. Indeed a reliability over udp is vital. Also the messages those are sent need to be fragmented to be suited for the server buffer. So if we assume the packet size would be 1024B, for every packet we send we have the following protocol:

Sequence Number|Flages|Fragmentation Offset|Fragmentation ID|Payload|
:------------:|:---:|:---:|:-------:|:-----
4 Bytes|1 Byte|8 Bytes|8 Bytes|1000 Bytes

Flags bits are consist of (From the most significant bits to  the least):

* IsFragmented      (7)
* IsACK             (6)
* IsLastFragment    (5)
* Reserved          (4)
* Reserved          (3)
* Reserved          (2)
* Reserved          (1)
* Reserved          (0)
 
This is implemented in hou package