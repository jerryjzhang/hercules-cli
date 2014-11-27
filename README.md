Hercules Command Line
====

To list available commands, run "hercules help".

###create

Create an application in Hercules.

    Usage: hercules create <appName>
    Examples:
    	$ hercules create dsf
    	Created application with id 1417019037871

###delete

Delete an application in Hercules.

	Usage: hercules delete <appName>
	Examples:
		$ hercules delete dsf
		Deleted application dsf

###deploy

Deploy an application in Hercules.

    Usage: hercules deploy <appName> [-s <svnURL> -i <imageURI> -e <env>=<value> -c <programName>=<cmd> -p <programName>=<entrypoint>]
    Options:
    	-s, --svn-url <url>,  set the svn url of your code
    	-i, --docker-image <image>, set the uri of your docker image
    	-e, --env <env>=<value>, set application-level environment variable
    	-c, --cmd <programName>=<cmd>, set cmd of a program
    	-p, --entrypoint <programName>=<entrypoint>, set entrypoint of a program
    Examples:
    	$ hercules deploy dsf -s http://svnURL -c web=start.sh -e DB_URL=localhost:3006 
    	Created release with id 141701903990

###env

Set environment variables of an application in Hercules.

	Usage: hercules env [-t <proc>]
		   hercules env set [-t <proc>] <var>=<val>...
		   hercules env unset [-t <proc>] <var>...
		   hercules env get [-t <proc>] <var>
	Examples: 
	
###cmd

Set program commands of an application in Hercules.

	Usage: hercules cmd <appName> set -p <programName>=<cmd>...
	   	   hercules cmd <appName> unset -p <programName>=<cmd>...
	   	   hercules cmd <appName> get -p <programName>=<cmd>
	Examples: 
	
	
###entrypoint

Set program entrypoints of an application in Hercules

	Usage: hercules entrypoint <appName> set -p <programName>=<entrypoint>...
		   hercules entrypoint <appName> unset -p <programName>=<entrypoint>...
		   hercules entrypoint <appName> get -p <programName>=<entrypoint>...
	Examples: 

###scale

Scale programs of an application in Hercules.

    Usage: hercules scale <appName> [<programName>=<replica> <programName>=<replica>...]
    Examples:
    	$ hercules scale dsf web=3 db=1
    	Scaled application dsf

###ps

List all running processes of an application.

	Usage: hercules ps <appName>
	Examples:
	    $ hercules ps dsf
	    ID            TYPE
		141701903990  web
		141701903991  web
		141701903992  db
