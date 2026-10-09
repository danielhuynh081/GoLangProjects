# BudgetMe

BudgetMe is a personal GoLang project I built to learn the fundamentals of GoLang. This app helps people visualize monthly xpenses and have more control of their finances.

## Current Features

- A couple common frameworks. (50/30/20, pay yourself first, etc)
- currently focused on budgeting money after monthly expenses
- **NEW**: Automatic expense calculation and categorization from bank CSV statements (last 1, 2, and 3 months).

## Built With

- GoLang
- ftm library

## Planned Features

This project is still a work in progress. Some features I'd like to add include:

- budgeting both monthly expenses and leftovers, monthly expenses are usually the reason people dont have leftovers
- importing bank statements as csv and automatically calculating expenses
  - tracking subscriptions
  - tracking eating out
  - unnecessary online purchases
- budget friendly recipes
- LLM API implementation for custom questions or recipes
- GUI

## Running the Project

```bash
go build -o main *.go
./main
```
