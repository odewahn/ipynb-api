Simple command line client for the IPython notebook.  It assumes you're running a Notebook server under boot2docker, like this:

```
$ docker run -it -p 8888:8888 -v $(pwd):/usr/data odewahn/ipymd
```

This will expose a kernel server running on `192.168.59.103:8888`.  Then, you can use this to little utility, which is basically just a thin wrapper on the Notebook API, to start, stop, restart, and interrupt kernels:

```
$ ./ipynb-api --help
NAME:
   ipynb-api - Client for ipynb api

USAGE:
   ipynb-api [global options] command [command options] [arguments...]

VERSION:
   0.0.0-alpha

COMMANDS:
   show		Show active kernels
   kill		Kills a kernel based on the first few chars of its id
   restart	Restart a kernel based on the first few chars of its id
   interrupt	Interrupt a kernel based on the first few chars of its id
   start	Starts the specified kerlen
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

Here's a session to show how to use it:

```console
admins-air-5:ipynb-api odewahn$ ./ipynb-api show
admins-air-5:ipynb-api odewahn$ ./ipynb-api start python3
Starting python3
Starting 303c81d1-067b-4b27-96a4-2ad6ced30aa8
admins-air-5:ipynb-api odewahn$ ./ipynb-api show
python3 	 303c81d1-067b-4b27-96a4-2ad6ced30aa8 
admins-air-5:ipynb-api odewahn$ ./ipynb-api restart 303
restart on 303c81d1-067b-4b27-96a4-2ad6ced30aa8
admins-air-5:ipynb-api odewahn$ ./ipynb-api kill 303
Killed 303c81d1-067b-4b27-96a4-2ad6ced30aa8
admins-air-5:ipynb-api odewahn$ ./ipynb-api show
admins-air-5:ipynb-api odewahn$
```

## Building

KISS b/c YOLO:

```
go build -o ipynb-api .
```
