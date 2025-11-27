import socket

target = "scanme.nmap.org"

print(f"Scanning {target}...\n")

for port in range(1, 1025):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.settimeout(0.5)   # timeout to prevent freezing

    result = s.connect_ex((target, port))  # returns 0 if open

    if result == 0:
        print(f"Port {port} is open")

    s.close()
