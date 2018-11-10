#!/usr/bin/env python3
'''
Example on how to do TCP port scanner in Python
'''

import socket


def scan(address, port, timeout=3):
    '''TCP Port Scanner'''
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    sock.settimeout(timeout)
    try:
        ip = socket.gethostbyname(address)
        sock.connect((ip, port))
        return 'connected'
    except socket.gaierror:
        return 'Unresolved address'
    except ConnectionRefusedError:
        return 'filtered'
    except socket.timeout:
        return 'closed'
    except socket.error as e:
        return 'General socket error: %s' % e
    except Exception as e:
        return 'General error: %s' % e


print(scan('google.com', 443))
