##### Block Multi_hosts
is two array.
check each host in the Hosts array by all ports in the array Ports.

###### Example
Multi_hosts:
    hosts: [ya.ru,google.ru,localhost]
    ports: [22,443,500]


##### Block Single_host
if you need to check only one host, or to multiple hosts, on different ports,can be filled in this way:

##### Example
Single_host:
  - 
    host: 'yandex.ru'
    port: [444,333]
  - 
    host: 'google.com'
    ports: [444,333,443]

###Use of both one block and both is allowed