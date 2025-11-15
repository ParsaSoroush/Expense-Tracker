# Expense Tracker

A simple **Expense Tracker** application written in **Go** to manage your finances. This command-line tool allows users to add, update, delete, view expenses, and see a summary of all expenses stored in a JSON file.

- **Repository URL:** [https://github.com/ParsaSoroush/Expense-Tracker.git](https://github.com/ParsaSoroush/Expense-Tracker.git)  
- **Project Page:** [https://roadmap.sh/projects/expense-tracker](https://roadmap.sh/projects/expense-tracker)

---

## Features

- Add a new expense with description and amount
- List all expenses in a table format
- View a summary of all expenses (total amount)
- Update an existing expense by ID
- Delete an expense by ID
- Stores data in a JSON file (`expenses.json`)  

---

## Installation

1. **Clone the repository**

```bash
git clone https://github.com/ParsaSoroush/Expense-Tracker.git
cd Expense-Tracker
```

2. **Build the application**
```bash
go build -o expense-tracker
```

## Usage
- Run the application from the command line using the following commands:

### Add an Expense
```bash
./expense-tracker add --description "Lunch" --amount 20.5
```

- Example output
    - `Added expense: Lunch - $20.50`

### List All Expenses
```bash
./expense-tracker list
```

- Example Output
    ```bash
    ID  Date       Description     Amount
    1   2025-11-15 Lunch           $20.50
    ```
### View Expense Summary
```bash
./expense-tracker summary
```
- Example Output
    `All Amount: $55.50`

### Update an Expense
```bash
./expense-tracker update --id 1 --description "Brunch" --amount 22.0
```
- --id specifies which expense to update.
- --description and --amount are optional; only provide the fields you want to update.

- Examle Output
    - `Expense 1 updated successfully`

### Delete an Expense
```bash
./expense-tracker delete --id 2
```
- --id specifies which expense to delete.

- Example Output
    - `Expense 2 deleted successfully`