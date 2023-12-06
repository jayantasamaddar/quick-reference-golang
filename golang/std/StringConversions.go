package std

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*************************************************************************************************/
// String Conversions
/*************************************************************************************************/
func StringConversionsDemo() {
	// Variables to work with
	numberString := "132"
	number := 132
	// str := "Mr. Jock, TV quiz PhD, bags few lynx."

	/********************************************************************************************************************/
	// Atoi: Ascii to Integer. Equivalent to ParseInt(s, 10, 0) converted to type int
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strconv.Atoi, strconv.ParseInt" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	// Demo with `strconv.Atoi`
	i, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf(Green+"STRING TO INTEGER(using `strconv.Atoi`): "+Reset+"%v%d%v of type: %v%T%v\n", Yellow, i, Reset, Yellow, i, Reset)

	// Demo with `strconv.ParseInt`
	i64, err := strconv.ParseInt(numberString, 10, 0)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf(Green+"STRING TO INTEGER(using `strconv.ParseInt`): "+Reset+"%v%d%v of type: %v%T%v\n", Yellow, i64, Reset, Yellow, i64, Reset)

	// NaN: Unlike ParseFloat, ParseInt cannot handle "NaN"
	i64, err = strconv.ParseInt("NaN", 10, 0)
	if err != nil {
		fmt.Printf(Green+"\"NaN\" TO INTEGER(using `strconv.ParseInt`) (error): "+Reset+"%v%s%v\n", Red, err.Error(), Reset)
	}
	fmt.Printf(Green+"\"NaN\" TO INTEGER(using `strconv.ParseInt`): "+Reset+"%v%d%v of type: %v%T%v\n", Yellow, i64, Reset, Yellow, i64, Reset)

	/********************************************************************************************************************/
	// Itoa: Integer to ASCII. Equivalent to FormatInt(int64(i), 10)
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strconv.Itoa, strconv.FormatInt" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	convStr := strconv.Itoa(number)

	// Demo with `strconv.Itoa`
	fmt.Printf(Green+"INTEGER TO STRING (using `strconv.Itoa`): "+Reset+"%v%q%v of type: %v%T%v\n", Yellow, convStr, Reset, Yellow, convStr, Reset)

	// Demo with `FormatInt(int64(i), 10)`
	convStr = strconv.FormatInt(int64(number), 10)
	fmt.Printf(Green+"INTEGER TO STRING (using `strconv.FormatInt`): "+Reset+"%v%q%v of type: %v%T%v\n", Yellow, convStr, Reset, Yellow, convStr, Reset)

	/********************************************************************************************************************/
	// ParseBool: Returns the boolean value represented by the string.
	// Truthy values: "1", "t", "T", "TRUE", "true", "True"
	// Falsy values: "0", "f", "F", "FALSE", "false", "False"
	// Any other value returns an error
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strconv.ParseBool" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	truthy := []string{
		"1",
		"t",
		"T",
		"TRUE",
		"true",
		"True",
	}
	falsy := []string{
		"0",
		"f",
		"F",
		"FALSE",
		"false",
		"False",
	}

	// Evaluate each element in a slice of strings and run `ParseBool` over them
	boolEval := func(boolStr []string, sep string) string {
		result := []string{}
		for _, v := range boolStr {
			isTruthy, err := strconv.ParseBool(v)
			if err != nil {
				log.Fatalln(err.Error())
			} else {
				result = append(result, fmt.Sprintf("%q is %t", v, isTruthy))
			}
		}
		return strings.Join(result, sep)
	}

	fmt.Printf(Green+"PARSE BOOLEAN (truthy): "+Reset+"%v%s%v\n", Yellow, boolEval(truthy, ", "), Reset)
	fmt.Printf(Green+"PARSE BOOLEAN (falsy): "+Reset+"%v%s%v\n", Yellow, boolEval(falsy, ", "), Reset)

	// Error Case
	isTruthy, err := strconv.ParseBool("abc")
	if err != nil {
		fmt.Printf(Green+"PARSE BOOLEAN (error): "+Reset+"%v%s%v\n", Red, err.Error(), Reset)
	} else {
		fmt.Printf(Green+"PARSE BOOLEAN (error): "+Reset+"%v%v%v\n", Yellow, isTruthy, Reset)
	}

	/********************************************************************************************************************/
	// ParseFloat: Converts the string s to a float with the precision specified by bitSize: 32 or 64
	/********************************************************************************************************************/
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)
	fmt.Println(Purple + "\tDemo: strconv.ParseFloat" + Reset)
	fmt.Println(Purple + "-------------------------------------------------------------------------------------------------" + Reset)

	v := "3.1415926535"

	if s, err := strconv.ParseFloat(v, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("NaN", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	// ParseFloat is case insensitive
	if s, err := strconv.ParseFloat("nan", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("inf", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("+Inf", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("-Inf", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("+Infinity", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("-Infinity", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("-0", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("+0", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}
