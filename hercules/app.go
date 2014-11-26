package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/flynn/flynn/Godeps/_workspace/src/github.com/flynn/go-docopt"
	"github.com/jmcvetta/napping"
	"github.com/tragicjun/hercules-cli/types"
)

var (
	//apiserverURL = os.Getenv("HERCULES_URL")
	apiserverURL = "http://localhost:58080"
)

func init() {
	register("create", runCreate, `
usage: hercules create <appName>

Create an application in Hercules.

Examples:

	$ hercules create dsf
	Created dsf with id 1417019037871
`)
	register("deploy", runDeploy, `
usage: hercules deploy <appName> [-s <svnURL> -i <imageURI>]

Options:
	-s, --svn-url <url>  set the svn url of your code
	-i, --docker-image <image> set the uri of your docker image

Deploy an application in Hercules.

Examples:

	$ hercules deploy dsf --svn-url http://svnURL
	Deployed application dsf
`)
	register("scale", runScale, `
usage: hercules scale <appName> [<programName>=<replica> <programName>=<replica>...]

Options:

Scale any programs of an application in Hercules.

Examples:

	$ hercules scale dsf web=3 db=1
	Scaled application dsf
`)
}

func runCreate(args *docopt.Args) error {
	var appName = args.String["<appName>"]
	//exec.Command("git", "remote", "remove", "flynn").Run()
	//exec.Command("git", "remote", "add", "flynn", gitURLPre(clusterConf.GitHost)+app.Name+gitURLSuf).Run()
	resp, _ := napping.Post(apiserverURL+"/apps/create/"+appName, nil, nil, nil)
	log.Printf("Created app %s with id %s", appName, resp.RawText())
	return nil
}

func runDeploy(args *docopt.Args) error {
	var appName = args.String["<appName>"]
	req := types.DeployRequest{}

	svn := args.String["--svn-url"]
	if svn != "" {
		req.SvnURL = svn
	}

	img := args.String["--docker-image"]
	if img != "" {
		req.DockerImage = img
	}

	resp, _ := napping.Post(apiserverURL+"/apps/deploy/"+appName, &req, nil, nil)
	if resp.Status() != 200 {
		log.Printf("Deploy request failed")
		log.Fatal(resp.RawText())
	}
	log.Printf("Deployed app %s", appName)
	return nil
}

func runScale(args *docopt.Args) error {
	var appName = args.String["<appName>"]
	req := types.ScaleRequest{}
	req.ProgramReplica = make(map[string]int)
	for _, arg := range args.All["<programName>=<replica>"].([]string) {
		i := strings.IndexRune(arg, '=')
		if i < 0 {
			fmt.Println(commands["scale"].usage)
		}
		val, err := strconv.Atoi(arg[i+1:])
		if err != nil {
			fmt.Println(commands["scale"].usage)
		}
		//replica, err := strconv.Atoi(val)
		req.ProgramReplica[arg[:i]] = val
	}
	resp, _ := napping.Post(apiserverURL+"/apps/scale/"+appName, &req, nil, nil)
	if resp.Status() != 200 {
		log.Printf("Scale request failed")
		log.Fatal(resp.RawText())
	}
	log.Printf("Scaled app %s", appName)
	return nil
}

func execCmd(cmd string) ([]byte, error) {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]
	return exec.Command(head, parts...).Output()
}
