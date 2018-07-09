# TLP Server

This is a transportation layer proxy server that can proxy udp to tcp and tcp to udp.

## HTTP on UDP

The main goal of this server is to receive http request on udp and retransmit it with tcp. Indeed a reliability over udp is vital. So for every packet we send we have the following protocol:

Sequence Number|Payload|
:------------:|:-------------: 
4 Bytes| others
 
This is done in hou package