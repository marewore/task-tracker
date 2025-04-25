# Task Tracker (Go CLI)

A simple command-line task tracker built with Go.

## 🚀 Getting Started

### Clone the repository

```bash
git clone https://github.com/marewore/task-tracker.git
cd task-tracker
```

### Build and Run

```bash
go build -o task-tracker
./task-tracker --help
```

## 📌 Usage

### Add a ask

```bash
./task-tracker add "Buy groceries"
```

### Update a task

```bash
./task-tracker update 1 "Buy groceries and cook dinner"
```

### Delete a task

```bash
./task-tracker delete 1
```

### Update task status

```bash
./task-tracker mark-in-progress 1
./task-tracker mark-done 1
./task-tracker mark-todo 1
```

### List tasks

```bash
./task-tracker list
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```
