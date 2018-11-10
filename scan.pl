#!/usr/bin/env perl

use warnings;
use strict;

use v5.20;

use IO::Socket::INET;

use experimental qw( switch );

sub scan {
  my ($address, $port, $timeout) = @_;

  my @ips = inet_aton($address) || return "Can't resolve $address";
  my $ip = inet_ntoa(@ips) || return "Can't extract from @ips";
  my $socket = IO::Socket::INET->new(
    PeerHost => $ip,
    PeerPort => $port,
    Timeout  => $timeout,
  );

  if ($@) {
    my $err = '';
    given ($@) {
      when(/connect: timeout$/) { $err = 'closed' }
      when(/connect: connection refused$/i) { $err = 'filtered' }
      default { $err = $@  }
    }
    return $err;
  }
  $socket->close();

  return 'open';
}

say scan('google.com', 443, 3);

