<h3 align="center">Golang Server with Python Script</h3>
<h5 align="center">Coding task by BharatX</h5>

---

<p align="center"> 
  The project is built as part of BharatX recruitment process. A Golang REST API that calls a Python script.  
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Implementation Choice](#implementation_choice)
- [Screenshots](#screenshots)
- [Built Using](#built_using)
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

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

What things you need to install the software and how to install them.

```
Python 3.8
Golang 
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
  
## 💡 Implementation Choice <a name = "implementation_choice"></a>
  
  <p>
    The main challenge was how to make the Go server communicate with the Python script. The way I have implemented it is that the request handler runs the Python script as a shell command using exec.Command and passes the necessary arguments as command line arguments. The output is the read from the stdout in the form of bytes which is then converted into string and then the corresponsing Grid struct type. 
    <br>
    This perhaps is not the best way to implement it in the long run. I did find another way which could have been better that used C API to run a python program using C header files. But I just could not get the neccessary setup required to make that work. Also, it needed quite a few specific configurations which would have been hard to replicate to run this system locally and therefore test. 
    <br>
    One other possibility I thought of was to expose the neccessary functions as a Flash API but that would've invalidated the purpose of having a Go server. 
    <br>
    Therefore, I made the decision to keep things simple and instead run the python program as a shell script.
  </p>

## 🎈 Screenshots <a name="screenshots"></a>

  ![Screenshot 2022-11-02 163511](https://user-images.githubusercontent.com/58480195/199474876-4c0d1258-2c59-43e7-8236-f726b7f9783f.png)
   ![Screenshot 2022-11-02 163554](https://user-images.githubusercontent.com/58480195/199474861-eed5827d-3bc9-4be1-b1e6-5db6cdc9109b.png)
    ![Screenshot 2022-11-02 163632](https://user-images.githubusercontent.com/58480195/199474832-f0ac0b14-8171-4b6d-818e-fa4dec52dc90.png)
  ![Screenshot 2022-11-02 163653](https://user-images.githubusercontent.com/58480195/199474821-fd034dd3-31ff-4012-b73e-1944b54b1bd5.png)
![Screenshot 2022-11-02 163719](https://user-images.githubusercontent.com/58480195/199474806-e6c26f73-c1ac-43d7-930f-c242e7afa85f.png)







## ⛏️ Built Using <a name = "built_using"></a>

- [Python3](https://www.python.org//) - Script
- [Golang](https://go.dev/) - Server Environment
- [Gorilla/Mux](https://pkg.go.dev/github.com/gorilla/mux) - Request Response handler

## ✍️ Authors <a name = "authors"></a>

- [@iamr-kumar](https://github.com/iamr-kumar) - Initial Work

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Shyam Murugan
- Mehul Jindal
