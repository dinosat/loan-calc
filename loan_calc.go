package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strconv"
)

type LoanCalculator struct {
	TotalCapital       float64
	LoanCapital        float64
	PaybackYears       float64
	MonthlyRent        float64
	TotalProfit        float64
	totalProfitLabel   *walk.Label // New field to display total profit
	roiLabel           *walk.Label // New field to display ROI
	totalCapitalEditor *walk.Label
}

func (lc *LoanCalculator) CalculateProfit() {
	v := lc.LoanCapital / lc.PaybackYears
	lc.TotalProfit = 0.0

	for year := 0.0; year <= lc.PaybackYears; year++ {
		profitFunction := ((lc.LoanCapital - v*year) / lc.TotalCapital) * lc.MonthlyRent * 12
		lc.TotalProfit += profitFunction
		fmt.Printf("Profit for Year %.0f: %.2f\n", year, profitFunction)
	}
}

func (lc *LoanCalculator) CalculateROI() float64 {
	initialInvestment := lc.LoanCapital
	netProfit := lc.TotalProfit
	return (netProfit / initialInvestment) * 100
}

func (lc *LoanCalculator) displayTotalProfit() {
	lc.totalProfitLabel.SetText(fmt.Sprintf("Total Profit: %.2f", lc.TotalProfit))
}

func (lc *LoanCalculator) displayROI() {
	roi := lc.CalculateROI()
	lc.roiLabel.SetText(fmt.Sprintf("ROI: %.2f%%", roi))
}

func main() {
	var mainWindow *walk.MainWindow
	calc := &LoanCalculator{}
	var totalCapitalEditor, loanCapitalEditor, payBackYearsEditor, rentEditor *walk.LineEdit

	if _, err := (MainWindow{
		AssignTo: &mainWindow,
		Title:    "Loan Calculator",
		Size:     Size{Width: 400, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			Label{Text: "Loan Calculator", Font: Font{PointSize: 14, Bold: true}},
			Label{Text: "Enter totalCapital value:"},
			LineEdit{AssignTo: &totalCapitalEditor, Text: "0"},
			Label{Text: "Enter loanCapital value:"},
			LineEdit{AssignTo: &loanCapitalEditor, Text: "0"},
			Label{Text: "Enter how many years you want to pay back:"},
			LineEdit{AssignTo: &payBackYearsEditor, Text: "0"},
			Label{Text: "Enter rent value:"},
			LineEdit{AssignTo: &rentEditor, Text: "0"},
			PushButton{
				Text: "Calculate",
				OnClicked: func() {
					// Extract input values from LineEdit widgets and update the fields in the LoanCalculator struct
					calc.TotalCapital, _ = strconv.ParseFloat(totalCapitalEditor.Text(), 64)
					calc.LoanCapital, _ = strconv.ParseFloat(loanCapitalEditor.Text(), 64)
					calc.PaybackYears, _ = strconv.ParseFloat(payBackYearsEditor.Text(), 64)
					calc.MonthlyRent, _ = strconv.ParseFloat(rentEditor.Text(), 64)

					calc.CalculateProfit()
				},
			},
			PushButton{
				Text:      "Display Total Profit",
				OnClicked: calc.displayTotalProfit,
			},
			PushButton{
				Text:      "Display ROI",
				OnClicked: calc.displayROI,
			},
			Label{AssignTo: &calc.totalProfitLabel},
			Label{AssignTo: &calc.roiLabel},
		},
	}.Run()); err != nil {
		fmt.Println("Error:", err)
	}
}
