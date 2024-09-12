# ASCII Art Generator in Go

## Overview

This project is an ASCII art generator implemented in Go. It takes a string as input and outputs the string in a graphical representation using ASCII characters. The program can handle input strings containing numbers, letters, spaces, special characters, and newline characters.

## Features

- Supports ASCII art representation for a wide range of characters
- Uses provided banner files with pre-defined ASCII art templates for each character
- Command-line interface for easy usage
- Written in Go for efficiency and portability
- Simple and easy-to-understand code structure

## Requirements

- Go 1.x or higher (replace x with the minimum version required)

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/ascii-art-generator.git
   ```
2. Navigate to the project directory:
   ```
   cd ascii-art-generator
   ```

## Usage

To use the ASCII art generator, run the program with a string as a command-line argument:

```
go run main.go "Hello, World!"
```

This will output the graphical representation of the input string using ASCII characters.

### Example Output

```
 _    _          _   _                 __          __            _       _   _ 
| |  | |        | | | |                \ \        / /           | |     | | | |
| |__| |   ___  | | | |   ___     ___   \ \  /\  / /    ___   __| |   __| | | |
|  __  |  / _ \ | | | |  / _ \   / _ \   \ \/  \/ /    / _ \ / _` |  / _` | | |
| |  | | |  __/ | | | | | (_) | | (_) |   \  /\  /    |  __/ \__,_| | (_| | |_|
|_|  |_|  \___| |_| |_|  \___/   \___/     \/  \/      \___| \__,_|  \__,_| (_)
```

## Banner Files

The project includes several banner files with pre-defined ASCII art templates for characters:

- `shadow.txt`
- `standard.txt`
- `thinkertoy.txt`

You can customize these banner files or add new ones to extend the supported characters or modify their graphical representation.

## Project Structure

```
ascii-art-generator/
├── main.go
├── banner/
│   ├── shadow.txt
│   ├── standard.txt
│   └── thinkertoy.txt
├── README.md
└── LICENSE
```

## How It Works

1. The program reads the input string from the command-line arguments.
2. It loads the appropriate banner file based on the selected style.
3. For each character in the input string:
   - The program looks up the ASCII art representation in the banner file.
   - It adds the representation to the output.
4. The final ASCII art is printed to the console.

## Contributing

Contributions are welcome! If you have any suggestions, improvements, or bug fixes, please follow these steps:

1. Fork the repository
2. Create a new branch (`git checkout -b feature/AmazingFeature`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
5. Push to the branch (`git push origin feature/AmazingFeature`)
6. Open a Pull Request

Please ensure your code adheres to the existing style and includes appropriate tests.

## Contact
Mehdi/Morningstar/Zangief
Project Link: [https://github.com/yourusername/ascii-art-generator](https://github.com/yourusername/ascii-art-generator)
