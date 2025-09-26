# Port Scanner

A lightweight command-line **port scanner** that probes specific TCP/UDP ports on a target host and reports their status in a clean JSON output.

---

## ❓ What is a port scanner?

A port scanner sends a network request to connect to a specific TCP or UDP port on a remote machine and records the response.  
In short: it sends a packet to a port to check whether the service listening there is reachable.

Example use case: to check whether a web server is running, probe port **80** on that server — if it’s open and listening, the server is likely up.

---

## 📚 Quick TCP port overview

- **Well-known ports (0–1023)** — reserved for standard services such as:
  - FTP (21), SSH (22), HTTP (80)
  - These ports are standardized and assigned by IANA.
- **Registered ports (1024–49151)** — available for services and applications (can be registered with IANA).
- **Dynamic/private ports (49152–65535)** — generally free to use for ephemeral connections.

---

## ⚙️ Port scanning basics

A port scanner sends either UDP or TCP packets and interprets the target’s replies. Typical responses:

- **Open / Accepted** — the host responded and a service is listening (e.g., “anything I can do for you?”).
- **Closed / Not listening** — the host replied but the port is not available for connections.
- **Filtered / Dropped / Blocked** — no response (likely filtered by a firewall or dropped).

---

## 📌 Usage

```bash
./portscan -port=80,81,82,83,84 64.29.17.131
```

---

## 🔍 Example

**Input:**

```bash
./portscan -port=80,81,82,83,84 64.29.17.131
```

**Output:**

```json
data: {
   "string": "64.29.17.131",
   "ports_scanned": [
      "80",
      "81",
      "82",
      "83",
      "84"
   ],
   "open_ports": [
      "80"
   ],
   "closed_ports": [
      "81",
      "83",
      "84",
      "82"
   ],
   "elasped_time": "10s",
   "timestamp": "2025-09-26T19:50:38.467338+05:30"
}
```

> Note: The input and output above are preserved exactly as requested.

---

## ✨ Features
- Scan arbitrary lists of ports (comma-separated).
- Report open and closed ports in a simple JSON structure.
- Show elapsed time and timestamp for traceability.
- Minimal and fast — suitable for quick checks and automation.

---

## 🛠️ Build

```bash
go build -o portscan main.go
```

---

## 🚀 Run

```bash
./portscan -port=80,81,82,83,84 64.29.17.131
```

---

## 📄 License
This project is licensed under the MIT License.

---