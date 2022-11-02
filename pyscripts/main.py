from grid import Grid
import sys
import json


n = len(sys.argv)

# Gets the input from the command line arguments
# Input validation is left to the server
def main():
  id = sys.argv[1]
  command = sys.argv[2]
  currentRow = int(sys.argv[3])
  currentCol = int(sys.argv[4])
  currentDirection = sys.argv[5]
  grid = Grid(currentRow, currentCol, currentDirection, id)
  grid.calculateNewPosition(command)
  out = grid.getGridInfo()
  print(json.dumps(out))
 
  

if __name__ == '__main__':
  main()
