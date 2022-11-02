<h3 align="center">Golang Server with Python Script</h3>
<h5 align="center">Coding task by BharatX</h5>

---

<p align="center"> 
  The project is built as part of BharatX recruitment process. A Golang REST API that calls a Python script.  
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Deployment](#deployment)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

<p>
  A python program is built that contains a grid with a player in one of its cell (starting position is [5,5]), and facing a direction (starting direction is East). The program takes in a command string containing three letters - L, R and F. L stands for 'turn left by 45 degree', R stands for 'turn right by 45 degree' and F stands for 'move one cell forward in the currently facing direction'. The task of the program is to process the input command and return the new position of the player along with the direction.
  <br> 
  The program is not called directly. A HTTP server is build using Golang which serves a single POST route - <italic>'/grids/[id]/feed'<italic> and takes in a body with <italic>command</italic> containing the input command. The route handler calles in the python script as a shell script and passes the id and command as command line arguments. The output is then read in the form of bytes, which is then converted to string and json unmarshalled to a Grid struct type in Golang. The final output is return in the response body.
  <br>
  The server maintains a Grid array containing current position and direction of all active grids. If a request comes with an existing ID, changes are made to the existing grid. If a request with a new ID comes in, a new Grid is created with the default position and direction.
</p>

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them.

```
Python 3.8 installed
Golang installed
```

### Installing

A step by step series of examples that tell you how to get a development env running.

Say what the step will be

```bash
git clone [repository-link]

go mod init example.com/bharatx_task

go get github.com/gorilla/mux

go build server.go

go run server.go
```

## 🎈 Screenshots <a name="usage"></a>

Add notes about how to use the system.

## ⛏️ Built Using <a name = "built_using"></a>

- [Python3](https://www.python.org//) - Script
- [Golang](https://go.dev/) - Server Environment
- [Gorilla/Mux](https://pkg.go.dev/github.com/gorilla/mux) - Request Response handler

## ✍️ Authors <a name = "authors"></a>

- [@iamr-kumar](https://github.com/iamr-kumar) - Initial Work

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
- Inspiration
- References
