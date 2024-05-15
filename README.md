# query_craft

query_craft is a CLI tool designed to convert a list of user IDs stored in a CSV file into a SQL query, facilitating quick and error-free generation of SQL statements for database operations.

## Features

- **Read User IDs from CSV**: Takes a CSV file as input and reads list of user IDs.
- **Generate SQL Statements**: Based on user inputs, generates SQL queries (`SELECT`, `UPDATE`, or `DELETE`).
- **Interactive CLI**: Interactive prompts to customize the SQL query based on user needs.
- **Error Handling**: Gracefully handles file read errors and user input errors.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (Golang) - Make sure you have Go installed on your system

### Installing

First, clone the repository to your local machine:

```bash
git clone https://github.com/o8n/query_craft.git
cd query_craft
```

Next, build the project using:

```bash
go build -o query_craft
```

Now, you can run the tool using:

```bash
./query_craft path/to/your/file.csv
```

### Usage

To use query_craft, you need a CSV file with user IDs listed, one per line. Here’s how you can run the tool:

```bash
./query_craft sample.csv
```

You will be prompted with a series of questions:

1. **Do you want to generate a SQL statement? [yes/no]** - Answer `yes` to proceed.
2. **Please choose: SELECT, UPDATE, or DELETE** - Choose the type of SQL operation you need.
3. **Enter the table name:** - Specify the table to query against.
4. If you choose `UPDATE`, you will also be asked to **Enter the set clause (e.g., `set column = value`):**.

### Example

Given a `sample.csv` file with the following content:

```
111
222
333
```

Running `./query_craft sample.csv` and choosing to generate a `SELECT` statement for the `users` table, you might see:

```
Do you want to generate a SQL statement? [yes/no]
yes
Please choose: SELECT, UPDATE, or DELETE
SELECT
Enter the table name:
users
Generated SQL: SELECT * FROM users WHERE user_id IN (111, 222, 333);
```

<!-- ## Contributing

We welcome contributions to query_craft! You can contribute in several ways:

1. **Report Issues**: Report a bug or a feature request [here](https://github.com/o8n/query_craft/issues).
2. **Send Pull Request**: Feel free to fork the repo and open pull requests. -->