# Nomos
A way to check a variety of arbitrary rules on a go file

Why Nomos?
In Greek mythology, Nomos ... was the daemon of laws, statutes, and ordinances[source](https://en.wikipedia.org/wiki/Nomos_(mythology))

## Features

- CLI tool to run on files to check a variety of rules on Go Files

## 🛠️ Prerequisites

- [Go] To compile the code to a binary

## 📁 Setup

1. Clone this repository:

   ```bash
   git clone https://github.com/jonathon-chew/Nomos.git
   cd Nomos 
   ```

2. Compile the binary:

    `go build .`

## List of Rules to toggle
```json
"functions-have-doc-strings": boolean,
"variable-names": ["camel_case", "snake_case", "kebab-case"],
"function-names": ["camel_case", "snake_case", "kebab-case"],
"readme-file": boolean,
"readme-stats": boolean,
"side-comments": boolean,
"print-f-new-line": boolean,
"print-f-new-line-log-all": boolean,
"ignore-if-in-comments": boolean,
"list-internal-functions": boolean,
"exported-identifiers-have-comments": boolean
```

## 📂 Output

You can define the output behaviour by creating a nomos_rules.json in any folder then running the binary on any file in that folder, if you don't you will be prompted to do so

## 🧠 Notes
This is currently a work in progress with a few inovations planned for the future:
    Readability statstics of files intended for use to explain the code such as README files and documentation.
Issues will be tracked in Github issues.

## 📜 License
This project is licensed under the MIT License. See the LICENSE file for details.
