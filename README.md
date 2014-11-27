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
    	-s, --svn-url <url>,  the svn url of your code
    	-i, --docker-image <image>, the uri of your docker image
    	-e, --env <env>=<value>, the application-level environment variable
    	-c, --cmd <programName>=<cmd>, the cmd of a program
    	-p, --entrypoint <programName>=<entrypoint>, the entrypoint of a program
    Examples:
    	$ hercules deploy dsf -s http://svnURL -c web=start.sh -e DB_URL=localhost:3006 
    	Created release with id 141701903990

###env

Manage environment variables of an application in Hercules.

	Usage: hercules env <appName> set [-p <programName>] -e <env>=<val>...
		   hercules env <appName> unset [-p <programName>] -e <env>...
		   hercules env <appName> get [-p <programName] -e <env>
	Options:
		-p, --program <programName>, the name of the program to be set/unset/get
		-e, --env <env>, the environment variable
	Examples: 
	
###cmd

Manage program commands of an application in Hercules.

	Usage: hercules cmd <appName> set -p <programName>=<cmd>...
	   	   hercules cmd <appName> get -p <programName>
	Options:
		-p, --program <programName>, the name of the program to be set/get
	Examples: 
	
###entrypoint

Manage program entrypoints of an application in Hercules

	Usage: hercules entrypoint <appName> set -p <programName>=<entrypoint>...
		   hercules entrypoint <appName> get -p <programName>
	Options:
		-p, --program <programName>, the name of the program to be set/get
	Examples: 

###release

List/Rollback application releases in Hercules

	Usage: hercules release <appName> list
	       hercules release <appName> rollback -v <releaseVersion>
	Options:
	       -v, --version <releaseVersion>, the version of the release to be rolled back
	Examples: 

###scale

Scale programs of an application in Hercules.

    Usage: hercules scale <appName> [<programName>=<replica> <programName>=<replica>...]
    Examples:
    	$ hercules scale dsf web=3 db=1
    	Scaled application dsf

###process

Manage processes of an application.

	Usage: hercules process <appName> ps
		   hercules process <appName> kill -p <processID>
	Examples:
	    $ hercules ps dsf
	    ID            Program
		141701903990  web
		141701903991  web
		141701903992  db


