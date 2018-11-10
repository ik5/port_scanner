#!/usr/bin/env ruby

require 'socket'

def tcp_connect(address, port, timeout: 20)
  # making sure we are talking with IP
  connected = false
  addr = Socket.getaddrinfo(address, nil)
  sock_addr = Socket.pack_sockaddr_in(port, addr[0][3])
  Socket.new(Socket.const_get(addr[0][0]), Socket::SOCK_STREAM, 0).tap do |socket|
    begin
      socket.connect_nonblock(sock_addr)
    rescue IO::WaitWritable
      if IO.select(nil, [socket], nil, timeout)
        begin
          # try again, might work
          socket.connect_nonblock(sock_addr)
          connected = :connected # no other exception, then should be yes
        rescue Errno::EISCONN # we have a connection
          connected = :connected
        rescue Errno::ECONNREFUSED # we are filtered
          connected = :filtered
        rescue Errno::ETIMEDOUT # timeout
          connected = :closed
        rescue StandardError # something else :'(
          connected = :error
        end
      else # unable to wait for an answer, but no exception was raised
        connected = :error
      end
    rescue StandardError # ops, something went wrong
      connected = false
    ensure
      socket.close
    end
  end
  connected
end

puts tcp_connect('192.168.97.3', 8823)
