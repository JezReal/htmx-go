package main

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ExpenseItem struct {
	Id            int
	ExpenseName   string
	ExpenseAmount float64
}

func main() {
	expenses := []ExpenseItem{
		{
			Id:            1,
			ExpenseName:   "Transpo",
			ExpenseAmount: 1000,
		},
		{
			Id:            2,
			ExpenseName:   "Food",
			ExpenseAmount: 500},
	}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
		}

		err = templ.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	r.Get("/expenses", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("templates/expense_item.html", "templates/expense_list.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
		}

		err = templ.ExecuteTemplate(w, "expense_list.html", expenses)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	r.Post("/expense", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("error parsing form: %v", err)
		}

		expenseName := r.FormValue("expense_name")
		expenseAmount := r.FormValue("expense_amount")

		parsedAmount, err := strconv.ParseFloat(expenseAmount, 64)
		if err != nil {
			log.Printf("error parsing float: %v", err)
		}

		expense := ExpenseItem{
			Id:            len(expenses) + 1,
			ExpenseName:   expenseName,
			ExpenseAmount: parsedAmount,
		}
		expenses = append(expenses, expense)
		templ, err := template.ParseFiles("templates/expense_item.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
		}

		err = templ.ExecuteTemplate(w, "expense_item.html", expense)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	r.Get("/edit/{id}", func(w http.ResponseWriter, r *http.Request) {
		expenseId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Printf("error parsing id: %v", err)
		}

		var e ExpenseItem
		for i := range expenses {
			if expenses[i].Id == expenseId {
				e = expenses[i]
				break
			}
		}

		templ, err := template.ParseFiles("templates/edit_expense.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
		}

		err = templ.ExecuteTemplate(w, "edit_expense.html", e)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	r.Put("/save/{id}", func(w http.ResponseWriter, r *http.Request) {
		expenseId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Printf("error parsing id: %v", err)
		}

		err = r.ParseForm()
		if err != nil {
			log.Printf("error parsing form: %v", err)
		}

		expenseName := r.FormValue("expense_name")
		expenseAmount := r.FormValue("expense_amount")

		parsedAmount, err := strconv.ParseFloat(expenseAmount, 64)
		if err != nil {
			log.Printf("error parsing float: %v", err)
		}

		var expenseItem *ExpenseItem
		for i := range expenses {
			if expenses[i].Id == expenseId {
				expenseItem = &expenses[i]
				expenseItem.ExpenseName = expenseName
				expenseItem.ExpenseAmount = parsedAmount
			}
		}

		templ, err := template.ParseFiles("templates/expense_item.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
		}

		err = templ.ExecuteTemplate(w, "expense_item.html", expenseItem)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	r.Delete("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		expenseId, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Printf("error parsing id: %v", err)
		}

		err = r.ParseForm()
		if err != nil {
			log.Printf("error parsing form: %v", err)
		}

		for i := range expenses {
			if expenses[i].Id == expenseId {
				expenses = append(expenses[:i], expenses[i+1:]...)
				break
			}
		}

		w.WriteHeader(http.StatusOK)
	})

	r.Get("/message", func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.ParseFiles("message.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
		}

		err = templ.ExecuteTemplate(w, "message.html", nil)
		if err != nil {
			log.Printf("Error executing template: %v", err)
		}
	})

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
