package main

import (
   "fmt"
    "log"
    "os/exec"
   //  "reflect"
   "encoding/json"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Grid struct {
	ID string `json:"id"`
	Direction string `json:"direction"`
	Position *Position `json:"position"`
}

func getResponseFromOut(out []byte) {
   grid := Grid{}
	var response string = string(out)
   err := json.Unmarshal([]byte(response), &grid)
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println(grid)
   
}

func main() {
   out, err := exec.Command("python", "main.py", "1", "LLFFFRRRLRFFLRLFFFLF", "5", "5", "E").Output()
      if err != nil {
         log.Fatal(err)
      }
      getResponseFromOut(out)
  
}