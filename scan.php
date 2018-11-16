#!/usr/bin/env php
<?php

function scan($address, $port, $timeout) {
  $ips = gethostbynamel($address);
  if ($ips === FALSE) {
    return "Unable to resolve $address";
  }

  $socket = @socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
  if ($socket === FALSE) {
    return socket_strerror(socket_last_error());
  }

  $timeout = array('sec'=>$timeout,'usec'=>$timeout * 1000);
  socket_set_option($socket,SOL_SOCKET,SO_SNDTIMEO,$timeout);
  $result = @socket_connect($socket, $ips[0], $port);
  if ($result === FALSE) {
    $err = socket_strerror(socket_last_error());
    switch ($err) {
    case 'Connection refused':
      return 'filtered';
      break;
    case 'Operation now in progress':
      return 'closed';
      break;
    default :
      print("$err\n");
      break;
    }
  }
  socket_close($socket);
  return 'open';
}

echo scan('google.com', 443, 3);
echo "\n";
?>
