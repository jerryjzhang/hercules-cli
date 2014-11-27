Hercules Command Line
====

## Usage

###Create app

```bash
    usage: hercules create <appName>
    Descption: Create an application in Hercules.
    Examples:
    	$ hercules create dsf
    	Created application with id 1417019037871
```

###Deploy app

    usage: hercules deploy <appName> [-s <svnURL> -i <imageURI>]
    Options:
    	-s, --svn-url <url>  set the svn url of your code
    	-i, --docker-image <image> set the uri of your docker image
    Descption: Deploy an application in Hercules.
    Examples:
    	$ hercules deploy dsf --svn-url http://svnURL
    	Created release with id 141701903990
	
###Scale app

    usage: hercules scale <appName> [<programName>=<replica> <programName>=<replica>...]
    Descption: Scale any programs of an application in Hercules.
    Examples:
    	$ hercules scale dsf web=3 db=1
    	Scaled application dsf
