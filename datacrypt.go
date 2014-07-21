package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Plugin struct {
	name string
}

func (p *Plugin) RecieveRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Plugin found.</h1>")
}

func main() {
	plugins := registerPlugins()

	rtr := mux.NewRouter()
	rtr.Handle("/{plugin}", pluginHandler(plugins))

	http.Handle("/", rtr)
	http.ListenAndServe(":3000", nil)
}

func pluginHandler(plugins map[string]*Plugin) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		pname := vars["plugin"]
		plugin, ok := plugins[pname]

		if ok {
			plugin.RecieveRequest(w, r)
		} else {
			fmt.Fprintf(w, "<h1>Plugin not found.</h1>")
		}
	})
}

func registerPlugins() map[string]*Plugin {
	plugin := &Plugin{name: "test"}
	plugins := map[string]*Plugin{
		"test": plugin,
	}
	return plugins
}
