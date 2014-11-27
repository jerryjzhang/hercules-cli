Hercules Command Line
====

## Usage

###Create app

Create an application in Hercules

    usage: hercules create <appName>
    Examples:
    	$ hercules create dsf
    	Created application with id 1417019037871

###Deploy app

Deploy an application in Hercules.

    usage: hercules deploy <appName> [-s <svnURL> -i <imageURI> -e <env>=<value> -c <programName>=<cmd> -p <programName>=<entrypoint>]
    Options:
    	-s, --svn-url <url>,  set the svn url of your code
    	-i, --docker-image <image>, set the uri of your docker image
    	-e, --env <env>=<value>, set application-level environment variable
    	-c, --cmd <programName>=<cmd>, set cmd of a program
    	-p, --entrypoint <programName>=<entrypoint>, set entrypoint of a program
    Examples:
    	$ hercules deploy dsf -s http://svnURL -c web=start.sh -e DB_URL=localhost:3006 
    	Created release with id 141701903990
	
###Scale app

Scale any programs of an application in Hercules.

    usage: hercules scale <appName> [<programName>=<replica> <programName>=<replica>...]
    Examples:
    	$ hercules scale dsf web=3 db=1
    	Scaled application dsf

###Set environmet variables

	usage: flynn env [-t <proc>]
flynn env set [-t <proc>] <var>=<val>...
flynn env unset [-t <proc>] <var>...
flynn env get [-t <proc>] <var>
	

	
