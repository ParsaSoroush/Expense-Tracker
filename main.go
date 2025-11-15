package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

const fileName = "expenses.json"

func readExpenses() ([]Expense, error) {
	var expenses []Expense
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&expenses)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func writeExpenses(expenses []Expense) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(expenses)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'list', 'summary', 'update', or 'delete' subcommands")
		os.Exit(1)
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addDescription := addCmd.String("description", "", "Description of the expense")
	addAmount := addCmd.Float64("amount", 0.0, "Amount of the expense")

	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	summaryMonth := summaryCmd.Int("month", 0, "Month to summarize (1-12)")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateId := updateCmd.Int("id", 0, "Id of the Expense")
	updateDescription := updateCmd.String("description", "", "Description of the Expense")
	updateAmount := updateCmd.Float64("amount", 0.0, "Amount of the Expense")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.Int("id", 0, "Id of the Expense")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addDescription == "" || *addAmount == 0 {
			fmt.Println("Please provide --description and --amount")
			return
		}

		expenses, err := readExpenses()
		if err != nil {
			fmt.Println("Error reading expenses:", err)
			return
		}

		newID := 1
		if len(expenses) > 0 {
			newID = expenses[len(expenses)-1].ID + 1
		}
		newExpense := Expense{
			ID:          newID,
			Date:        time.Now().Format("2006-01-02"),
			Description: *addDescription,
			Amount:      *addAmount,
		}

		expenses = append(expenses, newExpense)

		err = writeExpenses(expenses)
		if err != nil {
			fmt.Println("Error writing expenses:", err)
			return
		}

		fmt.Printf("Added expense: %s - $%.2f\n", newExpense.Description, newExpense.Amount)

	case "list":
		listCmd.Parse(os.Args[2:])
		expenses, err := readExpenses()
		if err != nil {
			fmt.Println("Error reading expenses:", err)
			return
		}

		fmt.Printf("# %-3s %-10s %-15s %s\n", "ID", "Date", "Description", "Amount")
		for _, e := range expenses {
			fmt.Printf("# %-3d %-10s %-15s $%.2f\n", e.ID, e.Date, e.Description, e.Amount)
		}

	case "summary":
		summaryCmd.Parse(os.Args[2:])
		expenses, err := readExpenses()
		if err != nil {
			fmt.Println("Error reading expenses:", err)
			return
		}

		total := 0.0
		for _, exp := range expenses {
			if *summaryMonth == 0 {
				total += exp.Amount
			} else {
				expDate, err := time.Parse("2006-01-02", exp.Date)
				if err != nil {
					fmt.Println("Error parsing date:", err)
					return
				}
				if int(expDate.Month()) == *summaryMonth {
					total += exp.Amount
				}
			}
		}

		if *summaryMonth == 0 {
			fmt.Printf("Total expenses: $%.2f\n", total)
		} else {
			fmt.Printf("Total expenses for %s: $%.2f\n", time.Month(*summaryMonth), total)
		}

	case "update":
		updateCmd.Parse(os.Args[2:])
		if *updateId == 0 {
			fmt.Println("Please provide --id to update")
			return
		}

		expenses, err := readExpenses()
		if err != nil {
			fmt.Println("Error reading expenses:", err)
			return
		}

		updated := false
		for i, exp := range expenses {
			if exp.ID == *updateId {
				if *updateDescription != "" {
					expenses[i].Description = *updateDescription
				}
				if *updateAmount != 0 {
					expenses[i].Amount = *updateAmount
				}
				updated = true
				break
			}
		}

		if !updated {
			fmt.Printf("Expense with ID %d not found\n", *updateId)
			return
		}

		err = writeExpenses(expenses)
		if err != nil {
			fmt.Println("Error saving updated expenses:", err)
			return
		}

		fmt.Printf("Expense %d updated successfully\n", *updateId)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteId == 0 {
			fmt.Println("Please provide --id to delete")
			return
		}

		expenses, err := readExpenses()
		if err != nil {
			fmt.Println("Error reading expenses:", err)
			return
		}

		deleted := false
		var updatedExpenses []Expense

		for _, exp := range expenses {
			if exp.ID == 0 {
				fmt.Println("Please add a valid Id")
				return
			}

			if exp.ID == *deleteId {
				deleted = true
				continue
			}

			updatedExpenses = append(updatedExpenses, exp)
		}

		if !deleted {
			fmt.Printf("Expense with ID %d not found\n", *deleteId)
			return
		}

		err = writeExpenses(updatedExpenses)
		if err != nil {
			fmt.Println("Error saving expenses:", err)
			return
		}

		fmt.Printf("Expense %d deleted successfully\n", *deleteId)

	default:
		fmt.Println("Unknown command. Use 'add', 'list', 'summary', 'update', or 'delete'.")
	}
}
