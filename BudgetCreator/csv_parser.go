package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	Date           time.Time
	Description    string
	Amount         float64
	Classification string
}

type RRReader struct {
	r io.Reader
}

func (r *RRReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		if p[i] == '\r' {
			p[i] = '\n'
		}
	}
	return n, err
}

func analyzeCSV(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(&RRReader{file})
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1

	// Read header
	_, err = reader.Read()
	if err != nil {
		fmt.Printf("Error reading header: %v\n", err)
		return
	}

	// Use a fixed reference date since the data seems to be from 2026
	// In a real app, this would be time.Now()
	now := time.Date(2026, 10, 9, 0, 0, 0, 0, time.UTC)
	oneMonthAgo := now.AddDate(0, -1, 0)
	twoMonthsAgo := now.AddDate(0, -2, 0)
	threeMonthsAgo := now.AddDate(0, -3, 0)

	type Totals struct {
		Month1 float64
		Month2 float64
		Month3 float64
	}

	categoryTotals := make(map[string]*Totals)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		// CSV structure: Account Number,Post Date,Check,Description,Debit,Credit,Status,Balance,Classification
		// Index: 0, 1, 2, 3, 4, 5, 6, 7, 8
		if len(record) < 9 {
			continue
		}

		dateStr := record[1]
		debitStr := record[4]
		classification := record[8]

		if debitStr == "" || classification == "" {
			continue
		}

		// Parse Date
		date, err := time.Parse("1/2/2006", dateStr)
		if err != nil {
			continue
		}

		// Parse Amount
		amount, err := strconv.ParseFloat(debitStr, 64)
		if err != nil {
			continue
		}

		if _, ok := categoryTotals[classification]; !ok {
			categoryTotals[classification] = &Totals{}
		}

		// Check if date is within 1 month, 2 months, or 3 months from 'now'
		if date.After(oneMonthAgo) {
			categoryTotals[classification].Month1 += amount
		}
		if date.After(twoMonthsAgo) {
			categoryTotals[classification].Month2 += amount
		}
		if date.After(threeMonthsAgo) {
			categoryTotals[classification].Month3 += amount
		}
	}

	fmt.Println("\n--- CSV Expense Analysis (Debits Only) ---")
	fmt.Printf("%-20s | %-15s | %-15s | %-15s\n", "Category", "Last 1 Month", "Last 2 Months", "Last 3 Months")
	fmt.Println(strings.Repeat("-", 75))

	for cat, totals := range categoryTotals {
		if totals.Month3 > 0 {
			fmt.Printf("%-20s | $%14.2f | $%14.2f | $%14.2f\n", cat, totals.Month1, totals.Month2, totals.Month3)
		}
	}

	fmt.Print("\nPress Enter to continue")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
