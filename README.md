Hercules Command Line
====

## Usage

###Create app

    usage: hercules create <appName>
    Descption: Create an application in Hercules.
    Examples:
    	$ hercules create dsf
    	Created application with id 1417019037871

###Deploy app

    usage: hercules deploy <appName> [-s <svnURL> -i <imageURI> -e <env>=<value> -c <programName>=<cmd> -p <programName>=<entrypoint>]
    Options:
    	-s, --svn-url <url>,  set the svn url of your code
    	-i, --docker-image <image>, set the uri of your docker image
    	-e, --env <env>=<value>, set application-level environment variable
    	-c, --cmd <programName>=<cmd>, set cmd of a program
    	-p, --entrypoint <programName>=<entrypoint>, set entrypoint of a program
    Descption: Deploy an application in Hercules.
    Examples:
    	$ hercules deploy dsf -s http://svnURL -c web="start.sh" -e DB_URL=localhost:3006 
    	Created release with id 141701903990
	
###Scale app

    usage: hercules scale <appName> [<programName>=<replica> <programName>=<replica>...]
    Descption: Scale any programs of an application in Hercules.
    Examples:
    	$ hercules scale dsf web=3 db=1
    	Scaled application dsf
