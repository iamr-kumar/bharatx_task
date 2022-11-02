class Grid:
  MAX_ROW = 10
  MAX_COL = 10

  # Dictionary used to get the left direction from the current direction
  left = { 'E': 'NE', 'NE': 'N', 'N': 'NW', 'NW': 'W', 'W': 'SW', 'SW': 'S', 'S': 'SE', 'SE': 'E' }
  # Dictionary used to get the right direction from the current direction
  right = { 'E': 'SE', 'SE': 'S', 'S': 'SW', 'SW': 'W', 'W': 'NW', 'NW': 'N', 'N': 'NE', 'NE': 'E' }

  # Dictionary used to get the next row from the current direction
  nextRow = { 'E': 0, 'NE': -1, 'N': -1, 'NW': -1, 'W': 0, 'SW': 1, 'S': 1, 'SE': 1 }
  # Dictionary used to get the next column from the current direction
  nextCol = { 'E': 1, 'NE': 1, 'N': 0, 'NW': -1, 'W': -1, 'SW': -1, 'S': 0, 'SE': 1 }

  def __init__(self, startRow: int, startCol: int, startDirection: str, id: str):

    # Check if the start position is valid
    assert startRow >= 0 and startRow <= self.MAX_ROW 
    assert startCol >= 0 and startCol <= self.MAX_COL

    self._id = id
    self.currentRow = startRow
    self.currentCol = startCol
    self.currentDirection = startDirection
  
  # define a read-only property id
  @property
  def id(self):
    return self._id

  # Calculate the new position based on the command
  # If the player goes out of board, the position is wrapped around the board
  # For example, if the position becomes (-1, 5), it will be wrapped to (9, 5)
  def calculateNewPosition(self, input: str):
    for turn in input:
      if turn == 'L':
        self.currentDirection = self.left[self.currentDirection]
      elif turn == 'R':
        self.currentDirection = self.right[self.currentDirection]
      elif turn == 'F':
        self.currentRow = (self.currentRow + self.nextRow[self.currentDirection]) % self.MAX_ROW
        self.currentCol = (self.currentCol + self.nextCol[self.currentDirection]) % self.MAX_COL

  # Return the current position and direction
  def getGridInfo(self):
    return { 'id': self.id, 'direction': self.currentDirection, 'position': { 'x': self.currentRow, 'y': self.currentCol }}

  