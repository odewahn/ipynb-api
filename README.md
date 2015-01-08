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

## API Notes

This site seems to summarize the REST API pretty well:

https://github.com/ipython/ipython/wiki/Dev:-URL-mapping-of-IPython-notebook

Most of the functions related to kernels are in these files:

* https://github.com/ipython/ipython/blob/master/IPython/html/services/kernels/handlers.py

* https://github.com/ipython/ipython/blob/master/IPython/html/static/services/kernels/kernel.js

I found it helpful to read the code to figure out what was happening.


### http -v POST 192.168.59.103:8888/api/kernels < kernel.json 

This will start a kernel of a specific type.  You have to post the JSON as part of the payload and specify the name in there.  Here's the contents of `kernel.json`

```
{
   "name": "python2"
}
```

Here's the result you get back:

```
{"name": "python2", "id": "0b4328dd-56ff-4ac8-ad2e-4d149eb04802"}
```


### curl -d '{"name":"python2"}' http://192.168.59.103:8888/api/kernels

This is an alternate way to start a kernel.  The -d flag does this:

```
-d, --data <data>
       (HTTP) Sends the specified data in a POST request  to  the  HTTP
       server,  in  the  same  way  that a browser does when a user has"
```

### http 192.168.59.103:8888/api/kernels

Returns:

```
[{"name": "python3", "id": "18323920-f793-4ee7-846f-23f0d2cbc1f1"}, {"name": "python2", "id": "9497c396-15e5-419a-8450-51a7bb426fc3"}, {"name": "python3", "id": "dc25b724-dc57-4cc7-ae1b-8ff643d0f21a"}]
```


### http http://192.168.59.103:8888/api/contents/

This fetches content, like this:

```
HTTP/1.1 200 OK
Content-Length: 1098
Content-Security-Policy: frame-ancestors 'self'; report-uri /api/security/csp-report;
Content-Type: application/json
Date: Mon, 05 Jan 2015 01:15:23 GMT
Etag: "44a255aeec6952867b6d3495bee8f9cb13b070d9"
Last-Modified: Mon, 05 Jan 2015 00:50:50 GMT
Server: TornadoServer/4.0.2
Set-Cookie: username-192-168-59-103-8888=; expires=Sun, 05 Jan 2014 01:15:23 GMT; Path=/

{
    "content": [
        {
            "content": null, 
            "created": "2014-12-20T23:09:37+00:00", 
            "format": null, 
            "last_modified": "2014-12-20T23:09:37+00:00", 
            "mimetype": null, 
            "name": "examples-md.md", 
            "path": "examples-md.md", 
            "type": "notebook", 
            "writable": true
        }, 
        {
            "content": null, 
            "created": "2015-01-01T13:07:48+00:00", 
            "format": null, 
            "last_modified": "2015-01-01T13:07:46+00:00", 
            "mimetype": null, 
            "name": "simple-line-plots.md", 
            "path": "simple-line-plots.md", 
            "type": "notebook", 
            "writable": true
        }, 
        {
            "content": null, 
            "created": "2014-12-20T23:09:37+00:00", 
            "format": null, 
            "last_modified": "2014-12-20T23:09:37+00:00", 
            "mimetype": null, 
            "name": "toc.md", 
            "path": "toc.md", 
            "type": "notebook", 
            "writable": true
        }, 
        {
            "content": null, 
            "created": "2014-12-20T23:09:37+00:00", 
            "format": null, 
            "last_modified": "2014-12-20T23:09:37+00:00", 
            "mimetype": null, 
            "name": "atlas.json", 
            "path": "atlas.json", 
            "type": "file", 
            "writable": true
        }
    ], 
    "created": "2015-01-05T00:50:50+00:00", 
    "format": "json", 
    "last_modified": "2015-01-05T00:50:50+00:00", 
    "mimetype": null, 
    "name": "", 
    "path": "", 
    "type": "directory", 
    "writable": true
}
```


