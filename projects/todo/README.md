# Todo CLI

A simple and efficient command-line todo application built with Go.

## Features

- ✅ Add, delete, and manage todos
- 🎯 Priority-based sorting
- ✨ Mark todos as done/undone
- 📊 Clean table view
- 🔢 Unique Snowflake ID generation
- 💾 JSON file storage

## Installation

```bash
go mod tidy
go build -o todo
```

## Usage

### Add a todo
```bash
./todo add -c "Buy groceries"
```

### List all todos
```bash
./todo list
```

### Mark as done
```bash
./todo done -i <ID>
```

### Unmark as done
```bash
./todo done -i <ID> -u
```

### Update priority
```bash
./todo update-priority -i <ID> -p <PRIORITY>
```

### Delete a todo
```bash
./todo delete -i <ID>
```

## Example Output

```
┌─────────────────┬──────────┬──────────┬─────────────────────┬─────────────────────┬──────┐
│       ID        │ CONTENT  │ PRIORITY │    CREATION TIME    │     UPDATE TIME     │ DONE │
├─────────────────┼──────────┼──────────┼─────────────────────┼─────────────────────┼──────┤
│ 377902275174400 │ 퇴근하기 │ 1        │ 2025-08-28 10:01:38 │ 2025-08-28 10:01:38 │ ❌   │
│ 377997330685952 │ 밥먹기   │ 2        │ 2025-08-28 10:02:01 │ 2025-08-28 10:02:01 │ ❌   │
│ 378021917696000 │ 운동하기 │ 3        │ 2025-08-28 10:02:07 │ 2025-08-28 10:02:07 │ ❌   │
│ 378033774993408 │ 출근하기 │ 4        │ 2025-08-28 10:02:10 │ 2025-08-28 10:02:32 │ ✅   │
└─────────────────┴──────────┴──────────┴─────────────────────┴─────────────────────┴──────┘
```

## Architecture

- **Snowflake ID Generator**: Distributed unique ID generation
- **Priority System**: Lower numbers = higher priority
- **JSON Storage**: Simple file-based persistence
- **Cobra CLI**: Clean command-line interface

## Limits

- Max todos: 100
- Max content length: 50 characters
- Max priority: 999

## Dependencies

- [cobra](https://github.com/spf13/cobra) - CLI framework
- [tablewriter](https://github.com/olekukonko/tablewriter) - Table formatting

## License

MIT
