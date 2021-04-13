package cardobject

import (
	"fmt"
	"os"
	"testing"
)

func writeToFile(fileName string, schema string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(schema)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestKeywordedSchema(t *testing.T) {
	writeToFile("keywordedSchema.json", string(KeywordedSchema()))
}
