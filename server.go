package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"strconv"
	"os/exec"
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

type Query struct {
	Command string `json:"command"`
}

var grids []Grid

func callPythonScript(id string, currentPos Position, direction string, query string) string{
	x := strconv.Itoa(currentPos.X)
	y := strconv.Itoa(currentPos.Y)
	out, err := exec.Command("python", "main.py", id, query, x, y, direction).Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	return string(out)
}

func getResponseFromOut(out string) Grid {
	grid := Grid{}
	err := json.Unmarshal([]byte(out), &grid)
	if err != nil {
		 log.Fatal(err)
	}
	return grid
}


func getNewPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var query Query 
	_ = json.NewDecoder(r.Body).Decode(&query) 
	var found bool
	var grid Grid
	for _, item := range grids {
		if item.ID == params["id"] {
			// var out string
			found = true
			out := callPythonScript(item.ID, *item.Position, item.Direction, query.Command)
			grid = getResponseFromOut(out)
			grids = append(grids[:0], grids[0+1:]...)
			grids = append(grids, grid)
		}
	}
	
	if !found {
		var position Position
		position.X = 5
		position.Y = 5
		out := callPythonScript(params["id"], position, "E", query.Command)
		grid = getResponseFromOut(out)
		grids = append(grids, grid)
	}
	json.NewEncoder(w).Encode(grid)
}

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/grids/{id}/feed", getNewPosition).Methods("POST")

	fmt.Println("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", route))
}