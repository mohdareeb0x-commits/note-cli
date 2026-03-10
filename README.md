# notes-cli

A simple command-line application for creating, reading, updating, and deleting notes using a JSON file as storage. Written in Go and built with the Cobra CLI library.

---

## Author

**Mohd Areeb**

---

## Features

* Add notes with title and description
* Read all notes
* Read a specific note by ID
* Update a note's title or description
* Delete a note by ID
* Delete all notes
* Uses a local `data.json` file for storage

---

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed (1.18+ recommended).
2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/notes-cli.git
   cd notes-cli
   ```
3. Build the binary:
   ```bash
   go build -o note-cli
   ```
4. (Optional) Move the binary to a directory in your `PATH`:
   ```bash
   mv note-cli /usr/local/bin/
   ```

---

## Usage

Run the `note-cli` command followed by a subcommand and appropriate flags.

### Commands

| Command       | Description                                 | Example |
|---------------|---------------------------------------------|---------|
| `addNote`     | Add a new note                              | `note-cli addNote --title "My Note" --description "This is a note"` |
| `readAll`     | Display all notes                           | `note-cli readAll` |
| `readByID`    | Show note with specified ID                 | `note-cli readByID --id 1` |
| `updateTitle` | Change the title of a note                  | `note-cli updateTitle --id 1 --title "New Title"` |
| `updateDesc`  | Change the description of a note            | `note-cli updateDesc --id 1 --description "New description"` |
| `deleteByID`  | Remove a note with a given ID               | `note-cli deleteByID --id 1` |
| `deleteAll`   | Remove all notes (irreversible)             | `note-cli deleteAll` |

### Flags

Each command accepts specific flags as described in the examples above. Running a command without required flags will result in errors.

---

## Project Structure

```
cmd/            # CLI command definitions
custom/         # Custom error types
logic/          # Core business logic for note operations
model/          # Data model definitions
testing/        # Unit tests for logic functions
main.go         # Application entry point
```

---

## Testing

This project includes unit tests for the logic layer. To run tests:

```bash
go test ./...
```

---

## Notes

* The application writes to `data.json` in the project root by default. Ensure appropriate permissions.
* Error handling is performed with custom error types for invalid IDs, empty notes, and missing titles.

---

## License

This project is open source and available under the [MIT License](LICENSE).

---

Feel free to contribute or open an issue if you encounter any problems!