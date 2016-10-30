// Runs a Git clone in the current folder for all Gitlab repositories,
// optionally only for projects in a given Gitlab group.
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"

	"github.com/tomnomnom/linkheader"
)

// Project contains minimal information of a Gitlab project.
type Project struct {
	FullPath  string `json:"path_with_namespace"`
	Namespace struct {
		Name string
	}
	URL string `json:"http_url_to_repo"`
}

func main() {
	token := os.Getenv("GITLAB_TOKEN")
	if len(token) < 1 {
		log.Fatal("GITLAB_TOKEN not set in environment")
	}
	url := os.Getenv("GITLAB_URL")
	if len(url) < 1 {
		log.Fatal("GITLAB_URL not set in environment")
	}
	jobstr := os.Getenv("GLGET_JOBS")
	if len(jobstr) < 1 {
		log.Fatal("GLGET_JOBS not set in environment")
	}
	jobs, err := strconv.Atoi(jobstr)
	if err != nil {
		log.Fatal(err)
	}
	if jobs < 1 {
		log.Fatal("GLGET_JOBS must be > 0")
	}
	var groupName string
	if len(os.Args) > 1 {
		groupName = os.Args[1]
	}
	client := &http.Client{}
	var projects []Project
	for link := url + "/projects?per_page=100"; len(link) > 0; {
		req, err := http.NewRequest("GET", link, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("PRIVATE-TOKEN", token)
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var page []Project
		err = json.Unmarshal(body, &page)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, page...)
		log.Printf("Got %v projects", len(projects))
		link = ""
		for _, l := range linkheader.Parse(res.Header.Get("Link")) {
			if l.Rel == "next" {
				link = l.URL
			}
		}
	}
	cmds := make(chan *exec.Cmd, 1)
	workers := sync.WaitGroup{}
	for i := 0; i < jobs; i++ {
		go func() {
			for cmd := range cmds {
				err := cmd.Run()
				dir := cmd.Dir
				if len(dir) == 0 {
					dir = "."
				}
				if err == nil {
					log.Printf("%v in %v", cmd.Args, dir)
				} else {
					log.Printf("Error: %v: %v in %v", err, cmd.Args, dir)
				}
			}
			workers.Done()
		}()
		workers.Add(1)
	}
	for _, project := range projects {
		if (len(groupName) == 0) || (len(groupName) > 0 && groupName == project.Namespace.Name) {
			var cmd *exec.Cmd
			if _, err := os.Stat(project.FullPath + "/.git"); os.IsNotExist(err) {
				cmd = exec.Command("git", "clone", project.URL, project.FullPath)
			} else {
				cmd = exec.Command("git", "pull")
				cmd.Dir = project.FullPath
			}
			cmds <- cmd
		}
	}
	close(cmds)
	workers.Wait()
}
