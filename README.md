Hercules Command Line
====

## Usage

###Create app

    usage: hercules create <appName>
    Create an application in Hercules.
    Examples:
    	$ hercules create dsf
    	Created dsf with id 1417019037871

###Deploy app

    usage: hercules deploy <appName> [-s <svnURL> -i <imageURI>]
    Options:
    	-s, --svn-url <url>  set the svn url of your code
    	-i, --docker-image <image> set the uri of your docker image
    Deploy an application in Hercules.
    Examples:
    	$ hercules deploy dsf --svn-url http://svnURL
    	Deployed application dsf
	
###Scale app

    usage: hercules scale <appName> [<programName>=<replica> <programName>=<replica>...]
    Options:
    Scale any programs of an application in Hercules.
    Examples:
    	$ hercules scale dsf web=3 db=1
    	Scaled application dsf
