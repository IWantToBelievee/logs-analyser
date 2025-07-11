# Logs Analyser

A simple CLI tool built with Go for working with log files. 

This is my personal pet project, created to learn the Go programming language, and I plan to improve it.

## What it Does

- Parses log files.
- Filters entries based on criteria you define.
- Outputs only the fields you need from logs, so you don't read extra.

## How to Use

Run it with your log file, then specify what to filter and which fields to show.

### Example
```bash
./logs-analyser --path logs/access.txt --fields ip,statecode,reqline filter --state-code=200 

+--------------+-----------+---------------------------------+
| IP           | STATECODE | REQLINE                         |
+--------------+-----------+---------------------------------+
| 127.0.0.1    | 200       | "GET /index.html HTTP/1.1"      |
| 192.168.1.15 | 200       | "POST /login HTTP/1.1"          |
| 10.0.0.5     | 200       | "GET /images/logo.png HTTP/1.1" |
| 172.16.0.2   | 200       | "GET /about HTTP/1.1"           |
| 10.0.0.5     | 200       | "GET /favicon.ico HTTP/1.1"     |
| 172.16.0.2   | 200       | "GET /contact HTTP/1.1"         |
+--------------+-----------+---------------------------------+
```

## Future Ideas

- **Aggregation and Counting**: Count occurrences of specific strings, errors, or unique values.
- **Statistical Analysis**: Calculate averages, percentages, or find top N frequent events.
- **Time-Series Analysis**: Identify trends or anomalies over time.
- **Advanced Filtering**: Support for regular expressions, date/time ranges, and logical operators (AND, OR, NOT).
- **Different Log Formats**: Support for JSON, CSV, or other structured log formats.
- **Performance Improvements**: Handle very large log files efficiently.
- **Enhanced Output**: Colored output for better readability.

## Requirements

- Go 1.18+

## Installation

1. Clone the repo:
   ```bash
   git clone https://github.com/yourusername/logs-analyser.git
   ```
2. Go into the folder:
   ```bash
   cd logs-analyser
   ```
3. Build it:
   ```bash
   go build -o logs-analyser
   ```

## License

MIT