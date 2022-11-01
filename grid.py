class Grid:
  MAX_ROW = 10
  MAX_COL = 10

  left = { 'E': 'NE', 'NE': 'N', 'N': 'NW', 'NW': 'W', 'W': 'SW', 'SW': 'S', 'S': 'SE', 'SE': 'E' }
  right = { 'E': 'SE', 'SE': 'S', 'S': 'SW', 'SW': 'W', 'W': 'NW', 'NW': 'N', 'N': 'NE', 'NE': 'E' }

  nextRow = { 'E': 0, 'NE': -1, 'N': -1, 'NW': -1, 'W': 0, 'SW': 1, 'S': 1, 'SE': 1 }
  nextCol = { 'E': 1, 'NE': 1, 'N': 0, 'NW': -1, 'W': -1, 'SW': -1, 'S': 0, 'SE': 1 }

  def __init__(self, startRow: int, startCol: int, startDirection: str, id: str):
    assert startRow >= 0 and startRow <= self.MAX_ROW 
    assert startCol >= 0 and startCol <= self.MAX_COL

    self._id = id
    self.currentRow = startRow
    self.currentCol = startCol
    self.currentDirection = startDirection
  
  @property
  def id(self):
    return self._id

  def calculateNewPosition(self, input: str):
    for turn in input:
      if turn == 'L':
        self.currentDirection = self.left[self.currentDirection]
      elif turn == 'R':
        self.currentDirection = self.right[self.currentDirection]
      elif turn == 'F':
        self.currentRow = (self.currentRow + self.nextRow[self.currentDirection]) % self.MAX_ROW
        self.currentCol = (self.currentCol + self.nextCol[self.currentDirection]) % self.MAX_COL

  def getGridInfo(self):
    return {
      'direction': self.currentDirection,
      'position': {
        'x': self.currentRow,
        'y': self.currentCol
      }

    }

  