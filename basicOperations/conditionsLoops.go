package basicoperations

import "fmt"

func ifState() {
	fmt.Println("This is the if statement: ")
	var getCountry bool = true
	var countryName string = ""
	if getCountry {
		countryName = "Kenya"
	}
	println(countryName)
}

func traditionalFor() {
	k := 1
	for ; k < 10; k++ {
		fmt.Println(k)
	}

	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}
}

func rangeFor() {
	strDict := map[string]string{"Japan": "Tokyo", "Kenya": "Nairobi", "China": "Beijing", "France": "Paris"}
	fmt.Println(strDict)
	for index, element := range strDict {
		fmt.Println("Country:", index, " City:", element)
	}
}
