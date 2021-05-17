# go-websocket-server
Um singelo servidor de Websocket feito em Golang


### Documentation Zebra Weblink

https://techdocs.zebra.com/link-os/2-14/webservices/content/link-os-pg-addendum-en.pdf
https://techdocs.zebra.com/link-os/2-13/webservices/content/Weblink%20WebSocket%20Endpoint%20Configuration.pdf



### Setting up the Zebra printer for Weblink

All commands must be executed on port 9200 vi TCP/IP


##### Enables Weblink on Zebra
```
{}{\"weblink.cloud_connect.enable\":true}
```
##### Checks whether the Weblink is active
```
{}{\"weblink.cloud_connect.enable\":null
```

##### Configure the server address on Zebra
- Defines the number of log lines in Zebra
- Defines the address of the server that Zebra will connect to
- Restart Zebra for settings to take effect

Important: Use WWS or HTTPS for SSL.

```
{}{\"weblink.logging.max_entries\":\"500\",\"weblink.ip.conn1.test.location\":\"wss://192.168.0.19/ws\",\"device.reset\":\"1\"}
```

Enable FTP on Zebra (only if you can't connect ;) )
```
{}{\"ip.ftp.enable\":\"on\"}
```

### Certificate Files 
Zebra certificate files must be at the root of the Zebra file system. Files can be uploaded via FTP.
```
WEBLINK1_CA.NRD
```

## Send File to Zebra With FTP

FTP + IP printer Zebra

Example:
```
# ftp 192.168.0.18
```

How to upload a file via FTP.
```
ftp> put filename
```

To copy multiple files at once, use the mput command.
```
ftp> mput filename [filename ...]
```

##### Check Log at Zebra
```
{}{\"weblink.logging.entries\":\"null\"}
```
