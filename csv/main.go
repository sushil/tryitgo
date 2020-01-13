package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

/*
	Problem taken from: https://www.fluentcpp.com/2017/10/23/results-expressive-cpp17-coding-challenge/

	The task proposed in the challenge was to write a command line tool that takes in a CSV file, overwrites all the
	data of a given column by a given value, and outputs the results into a new CSV file.

	More specifically, this command line tool should accept the following arguments:

		the filename of a CSV file, the name of the column to overwrite in that file, the string that will be used as a
		replacement for that column, the filename where the output will be written.

	For instance, if the CSV file had a column “City” with various values for the entries in the file, calling the tool
	with the name of the input file, City, London and the name of output file would result in a copy of the initial
	file, but with all cities set equal to “London”:

	Input file:

	First Name,Last Name,Age,City,Eyes color,Species
	John,Doe,32,Tokyo,Blue,Human
	Flip,Helm,12,Canberra,Red,Unknown
	Terdos,Bendarian,165,Cracow,Blue,Magic tree
	Dominik,Elpos,33,Paris,Blue,Human
	Ewan,Grath,51,New Delhi,Green,Human

	will be converted to (City -> London)

	Output file:

	First Name,Last Name,Age,City,Eyes color,Species
	John,Doe,32,London,Blue,Human
	Flip,Helm,12,London,Red,Unknown
	Terdos,Bendarian,165,London,Blue,Magic tree
	Dominik,Elpos,33,London,Blue,Human
	Ewan,Grath,51,London,Green,Human

	Here was how to deal with edge cases:

		if the input file is empty, the program should write “input file missing” to the console. if the input file does
		not contain the specified column, the program should write “column name doesn’t exists in the input file” to the
		console.

	In both cases, there shouldn’t be any output file generated.

	And if the program succeeds but there is already a file having the name specified for output, the program should
	overwrite this file.
	The goal of the challenge was double: using as many C++17 features as possible (as long as they were useful to solve
	the case), and write the clearest code possible with them.
*/

var ErrColumnNotFound = errors.New("column not found")

func main() {
	// inputFile, column, replacement, outputFile
	if len(os.Args[1:]) == 0 {
		printUsage()
		return
	}

	var (
		inputFile   = os.Args[1]
		column      = os.Args[2]
		replacement = os.Args[3]
		outputFile  = os.Args[4]
	)

	run(inputFile, column, replacement, outputFile)

}

func run(inputFile, column, replacement, outputFile string) {

	inFile, err := os.Open(inputFile)
	if err != nil {
		reportError("cannot open file", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		reportError("cannot open file", err)
		return
	}
	defer outFile.Close()

	csvReader := csv.NewReader(inFile)
	csvWriter := csv.NewWriter(outFile)
	defer csvWriter.Flush()

	colIndex := -1
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			return
		}

		if err != nil {
			reportError("error reading row", err)
			return
		}

		if colIndex < 0 {
			colIndex, err = columnIndex(row, column)
			if err != nil {
				reportError("column not found", err)
				return
			}

			if err := csvWriter.Write(row); err != nil {
				reportError("cannot write row %s", err)
				return
			}

		} else {
			row[colIndex] = replacement
			if err := csvWriter.Write(row); err != nil {
				reportError("cannot write row %s", err)
				return
			}
		}
	}
}

func columnIndex(row []string, columnName string) (int, error) {
	for i := 0; i < len(row); i++ {
		if columnName == row[i] {
			return i, nil
		}
	}

	return 0, ErrColumnNotFound
}

func printUsage() {
	fmt.Println("<this program> <input file> <column> <replacement> <output file>")
}

func reportError(message string, err error) {
	fmt.Printf("%s\ndetails:\n%s\n", message, err)
}
