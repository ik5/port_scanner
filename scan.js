#!/usr/bin/env node

/*jshint esversion: 6 */

function scan(address, port, timeout, status) {
  const net = require('net');
  const dns = require('dns');

  dns.lookup(address, (err, host, family) => {
    if (err) {
      status('Resolved error');
      return;
    }

    const client = net.createConnection({host, port}, () => {
      status('open');
    });
    client.on('error', (err) => {
      switch (err.code) {
        case 'ECONNREFUSED':
          status('filtered');
          break;
        default:
          status(err);
          break;
      }
    });
    client.setTimeout(timeout * 1000, () => {
      if (client) {
        client.destroy();
      } else {
        status('closed');
      }
    });
  });

}

scan('google.com', 443, 3, (status) => console.log(status));
