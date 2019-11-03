package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	sampleJSON := (`{
		"login": "inemtsev",
		"id": 15311236,
		"node_id": "MDQ6VXNlcjE1MzExMjM2",
		"avatar_url": "https://avatars0.githubusercontent.com/u/15311236?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/inemtsev",
		"html_url": "https://github.com/inemtsev",
		"followers_url": "https://api.github.com/users/inemtsev/followers",
		"following_url": "https://api.github.com/users/inemtsev/following{/other_user}",
		"gists_url": "https://api.github.com/users/inemtsev/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/inemtsev/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/inemtsev/subscriptions",
		"organizations_url": "https://api.github.com/users/inemtsev/orgs",
		"repos_url": "https://api.github.com/users/inemtsev/repos",
		"events_url": "https://api.github.com/users/inemtsev/events{/privacy}",
		"received_events_url": "https://api.github.com/users/inemtsev/received_events",
		"type": "User",
		"site_admin": false,
		"name": "Ilya Nemtsev",
		"company": null,
		"blog": "www.eventslooped.com",
		"location": "Bangkok, Thailand",
		"email": null,
		"hireable": null,
		"bio": "I have coded in VB, VB.NET, C, Java, C#, JavaScript, Golang\r\nroughly in that order. ",
		"public_repos": 5,
		"public_gists": 1,
		"followers": 0,
		"following": 0,
		"created_at": "2015-10-25T15:13:08Z",
		"updated_at": "2019-10-27T17:02:04Z",
		"mock": true
	}`)

	http.HandleFunc("/small-get", func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("x-friend-user")
		err := ioutil.WriteFile("output.txt", []byte(h), 0644)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Hello World!`))
	})

	http.HandleFunc("/big-get", func(w http.ResponseWriter, r *http.Request) {
		var rHeaders []string

		for name, headers := range r.Header {
			name = strings.ToLower(name)
			for _, h := range headers {
				rHeaders = append(rHeaders, fmt.Sprintf("%v: %v", name, h))
			}
		}

		f, err := os.Create("test.txt")

		for _, line := range rHeaders {
			_, err = f.WriteString(line + "\n")
			if err != nil {
				panic(err)
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleJSON))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
