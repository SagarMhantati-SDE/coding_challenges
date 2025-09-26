# DNS Resolver

A simple command-line tool to resolve hostnames using a specified DNS server.  
It queries the given DNS server and returns the response in a clean, JSON-formatted output.

---

## 📌 Usage

```bash
./dns_resolver -server <dns-server-ip> -domain <hostname>
```

### Example

**Input:**

```bash
./dns_resolver -server 8.8.8.8 -domain en.wikipedia.org
```

**Output:**

```json
{
   "domain": "en.wikipedia.org",
   "type": "",
   "dns_server": "8.8.8.8",
   "answer": [
      "2001:df2:e500:ed1a::1",
      "103.102.166.224"
   ]
}
```

---

## ⚡ Features
- Query any DNS server (default or custom).
- Supports both IPv4 and IPv6 records.
- Outputs results in **pretty-printed JSON** for readability.
- Lightweight and fast.

---

## 🛠️ Build

```bash
go build -o dns_resolver main.go
```

---

## 🚀 Run

```bash
./dns_resolver -server 8.8.8.8 -domain example.com
```

---

## 📄 License
This project is licensed under the MIT License.
