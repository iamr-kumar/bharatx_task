package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"strconv"
	"os/exec"
	"errors"
	"strings"
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

type Error struct {
	Message string `json:"message"`
}

var grids []Grid

// Function to validate the input. Check if ID is negative
// and if the input command is within range and contains only L, F and R 
func validateInput(command string, id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		return errors.New("Invalid ID")
	}
	if len(command) > 250 {
		return errors.New("Command length should be less than 250")
	}
	command = strings.ToUpper(command)
	for _, char := range command {
		if char != 'L' && char != 'R' && char != 'F' {
			return errors.New("Command should only contain L/l, R/r, F/f")
		}
	}
	return nil
}

// Function to run the python script and get the output from stdout.
// The output is read in the form of bytes, which is converted into a string.
// The string is then unmarshalled into a Grid struct and returned.
func callPythonScript(id string, currentPos Position, direction string, query string) (Grid, error){
	x := strconv.Itoa(currentPos.X)
	y := strconv.Itoa(currentPos.Y)
	grid := Grid{}
	out, err := exec.Command("python", "pyscripts/main.py", id, query, x, y, direction).Output()
	if err != nil {
		return grid, errors.New("Some error occurred...")
	}
	response := string(out)
	err = json.Unmarshal([]byte(response), &grid)
	if err != nil {
		return grid, errors.New("Some error occurred...")
	}
	return grid, nil
	
}


// Function to handle the request to get the new position of the grid
// 1. Checks if the input are valid
// 2. Looks for the grid with the given ID in the existing Grid array.
// 3. If found, runs the operation on the existing grid and returns the new position
// 4. If not found, creates a new grid and returns the new position
// Note - the position in the grid is 0 indexed
func getNewPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var query Query 
	_ = json.NewDecoder(r.Body).Decode(&query) 
	err := validateInput(query.Command, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error()})
		return
	}
	var found bool
	var grid Grid
	for _, item := range grids {
		if item.ID == params["id"] {
			found = true
			grid, err = callPythonScript(item.ID, *item.Position, item.Direction, query.Command)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Error{Message: err.Error()})
				return
			}
			grids = append(grids[:0], grids[0+1:]...)
			grids = append(grids, grid)
		}
	}
	
	if !found {
		var position Position
		position.X = 5
		position.Y = 5
		grid, err = callPythonScript(params["id"], position, "E", query.Command)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{Message: err.Error()})
			return
		}
		grids = append(grids, grid)
	}
	json.NewEncoder(w).Encode(grid)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("API running...")
}


// Main function to setup the server
func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", handleHome).Methods("GET")
	route.HandleFunc("/grids/{id}/feed", getNewPosition).Methods("POST")

	fmt.Println("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", route))
}