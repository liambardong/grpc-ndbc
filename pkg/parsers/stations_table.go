package parsers

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)


func parseResponse(response *http.Response) ([]){
	
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line,"#"){
			continue
		}
		
		fmt.Println(line)
		
		columns := strings.Split(line, "|")

		
	}
}
