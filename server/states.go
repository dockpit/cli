package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dockpit/pit/model"

	"github.com/zenazn/goji/web"
)

var Builds = map[string]*model.Build{}
var Runs = map[string]*model.Run{}

func (s *Server) OneBuild(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["id"]
	if b, ok := Builds[id]; !ok {
		http.NotFound(w, r)
		return
	} else {
		enc := json.NewEncoder(w)
		err := enc.Encode(b)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode: %s", err), http.StatusInternalServerError)
		}
	}
}

func (s *Server) BuildDepState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	sid := c.URLParams["state_id"]
	state := dep.GetState(sid)
	if state == nil {
		http.NotFound(w, r)
		return
	}

	b, err := model.NewBuild(*dep, *state)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create build from state '%s': %s", state.Name, err), http.StatusInternalServerError)
	}

	//builds are run in the background
	Builds[b.ID] = b
	go func() {
		iname, err := s.client.Build(b)
		if err != nil {
			b.Error = err
			return
		}
		b.ImageName = iname
	}()

	enc := json.NewEncoder(w)
	err = enc.Encode(b)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: %s", err), http.StatusInternalServerError)
	}
}

func (s *Server) OneRun(c web.C, w http.ResponseWriter, r *http.Request) {
	id := c.URLParams["id"]
	if run, ok := Runs[id]; !ok {
		http.NotFound(w, r)
		return
	} else {
		enc := json.NewEncoder(w)
		err := enc.Encode(run)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode: %s", err), http.StatusInternalServerError)
		}
	}
}

func (s *Server) RunDepState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	sid := c.URLParams["state_id"]
	state := dep.GetState(sid)
	if state == nil {
		http.NotFound(w, r)
		return
	}

	run, err := model.NewRun(*state)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create build from state '%s': %s", state.Name, err), http.StatusInternalServerError)
	}

	//remove any state containers that are still running
	err = s.client.Remove(state)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to remove running container for state '%s': %s", state.Name, err), http.StatusInternalServerError)
	}

	//runs are started in the background
	Runs[run.ID] = run
	go func() {
		_, err = s.client.Start(run)
		if err != nil {
			run.Error = err
			return
		}
	}()

	enc := json.NewEncoder(w)
	err = enc.Encode(run)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: %s", err), http.StatusInternalServerError)
	}

}

func (s *Server) CreateState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newstate *model.State
	err = dec.Decode(&newstate)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed decode new state: %s", err), http.StatusBadRequest)
		return
	}

	state, err := model.NewStateFromTemplate(newstate.Name, dep.Template)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create state: %s", err), http.StatusBadRequest)
		return
	}

	dep.AddState(state)
	err = s.model.UpdateDep(dep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed update dep with name '%s': %s", name, err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) OneDepState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	sid := c.URLParams["state_id"]
	state := dep.GetState(sid)
	if state == nil {
		http.NotFound(w, r)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(state)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode: %s", err), http.StatusInternalServerError)
	}
}

func (s *Server) UpdateDepState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	sid := c.URLParams["state_id"]
	state := dep.GetState(sid)
	if state == nil {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newstate *model.State
	err = dec.Decode(&newstate)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed decode state for updating '%s': %s", sid, err), http.StatusBadRequest)
		return
	}

	dep.UpdateState(state, newstate)

	err = s.model.UpdateDep(dep)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update dep: %s", err), http.StatusInternalServerError)
		return
	}
}

func (s *Server) RemoveDepState(c web.C, w http.ResponseWriter, r *http.Request) {
	name := c.URLParams["name"]
	dep, err := s.model.FindDepByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to find dep with name '%s': %s", name, err), http.StatusBadRequest)
		return
	}

	if dep == nil {
		http.NotFound(w, r)
		return
	}

	sid := c.URLParams["state_id"]
	err = s.model.RemoveDepStateByID(dep, sid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to remove state from dep: %s", err), http.StatusInternalServerError)
	}
}
