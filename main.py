from grid import Grid
import sys
import json

MAX_ROW = 10
MAX_COL = 10


left = {"E": "NE", "NE": "N", "N": "NW", "NW": "W", "W": "SW", "SW": "S", "S": "SE", "SE": "E"}
right = {"E": "SE", "SE": "S", "S": "SW", "SW": "W", "W": "NW", "NW": "N", "N": "NE", "NE": "E"}

nextRow = {"E": 0, "NE": -1, "N": -1, "NW": -1, "W": 0, "SW": 1, "S": 1, "SE": 1}
nextCol = {"E": 1, "NE": 1, "N": 0, "NW": -1, "W": -1, "SW": -1, "S": 0, "SE": 1}


n = len(sys.argv)

def main():
  # print(sys.args[1])
  id = sys.argv[1]
  command = sys.argv[2]
  currentRow = int(sys.argv[3])
  currentCol = int(sys.argv[4])
  currentDirection = sys.argv[5]
  for turn in command:
      if turn == 'L':
        currentDirection = left[currentDirection]
      elif turn == 'R':
        currentDirection = right[currentDirection]
      elif turn == 'F':
        currentRow = (currentRow + nextRow[currentDirection]) % MAX_ROW
        currentCol = (currentCol + nextCol[currentDirection]) % MAX_COL


  out = { "id": id, "direction": currentDirection, "position": { "x": currentRow, "y": currentCol }}
  print(json.dumps(out))
 
  

if __name__ == '__main__':
  main()
