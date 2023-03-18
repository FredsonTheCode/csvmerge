package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	custom := [][]string{}
	product := [][]string{}
	outpt := [][]string{}
	outString := "Version*;Currency ID*;Customer ID*;Product ID*;Key Figure*;Price\n"
	customers, err := os.Open("csv/lista1.csv")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(customers)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !strings.Contains(line, "Customer ID;Customer") {
			cus := strings.Split(line, ";")
			if len(cus) >= 2 {
				custom = append(custom, []string{cus[0], cus[1]})
			}
		}
	}
	products, err := os.Open("csv/lista2.csv")
	if err != nil {
		panic(err)
	}
	fileScanner = bufio.NewScanner(products)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if !strings.Contains(line, "Customer;Product ID*;Price;Currency ID*") {
			prd := strings.Split(line, ";")
			if len(prd) >= 4 {
				product = append(product, []string{prd[0], prd[1], prd[2], prd[3]})
			}
		}
	}
	for _, cu := range custom {
		for _, pr := range product {
			if cu[1] == pr[0] {
				outpt = append(outpt, []string{"", pr[3], cu[0], pr[1], "", pr[2]})
			}
		}
	}
	for _, ou := range outpt {
		outString += strings.Join(ou, ";") + "\n"
	}
	err = os.WriteFile("csv/output.csv", []byte(outString), 0644)
	if err != nil {
		panic(err)
	}
}
