
📝 MyCLI – A Simple Todo CLI in Go

mycli is a command-line tool to manage your todos with priorities, completion status, and persistent storage in JSON/YAML.
It’s built in Go using Cobra and Viper.

🚀 Features

Add, list, and mark todos as done ✅

Prioritize tasks 🔼

Stores data persistently in JSON file

Configurable datafile location (via --datafile flag or environment variable)

Cross-platform support (Linux, Mac, Windows)

📦 Installation

Clone the repo and build:
```bash
git clone https://github.com/yourusername/mycli.git
cd mycli
go build -o mycli
```
⚙️ Configuration

You need to tell mycli where to store the todos.

Option 1: Using flag
```bash
mycli list --datafile /home/vivek/golang/tridos.json
```
Option 2: Using environment variable
```bash
export TRI_DATAFILE=$HOME/golang/tridos.json
```
Option 3: Using config file (tri.yaml)
```bash
datafile: /home/vivek/golang/tridos.json
```
🛠️ Usage
Add a todo
```bash
mycli add "Buy groceries" --priority 2
```
List todos
```bash
mycli list --datafile /home/vivek/golang/tridos.json
```
Mark as done
```bash
mycli done 1
```
📂 Example JSON (tridos.json)
```bash
[
  {
    "text": "Buy groceries",
    "priority": 2,
    "done": false,
    "created_at": "2025-08-28T11:30:48+05:30"
  }
]
```

🧑‍💻 Development

Run locally without building:
```bash
go run main.go list --datafile ./tridos.json
```
📜 License

MIT License. Feel free to use and modify.
