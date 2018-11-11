# Text Reminder

This program sends a amazing cat fact in version 2. at the same time each day.

## Getting Started

create the file 'config.go'
```
package main
 
const sid = ""           // twilio account sid
const token = ""         // twilio auth token
const from = ""          // sender's number, format: "+15558889999"
const reminderTime = ""  // time of day to send reminder (UTC), format: 21:00
const to = ""            // receiver's number
const message = ""       // message to send
```

## Build Docker

```
docker build -t registry.spofcloud.com/text-reminder:2.0.0 .

```

```
docker run -d --restart=unless-stopped /
--name=text-reminder registry.spofcloud.com/text-reminder:2.0.0
```
