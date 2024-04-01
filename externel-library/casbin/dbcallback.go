package main

import (
	"context"
	"fmt"
	"time"

	psqlwatcher "github.com/IguteChung/casbin-psql-watcher"
)

func dbCallback() {

	// prepare the watcher with local psql server.
	// for demo purpose, enable NotifySelf to receive a local change callback.
	conn := "host=192.168.1.205 port=5432 user=hsck password=hsck@2301 database=test_tenant sslmode=disable"
	w, _ := psqlwatcher.NewWatcherWithConnString(context.Background(), conn,
		psqlwatcher.Option{NotifySelf: true, Verbose: true})

	// prepare the enforcer.
	e, err := getEnforcerByDB(conn)
	if err != nil {
		fmt.Println("db: ", err)
		return
	}

	//e, _ := casbin.NewEnforcer("../testdata/rbac_model.conf", "../testdata/rbac_policy.csv")

	// set the watcher for enforcer.
	_ = e.SetWatcher(w)

	// set the default callback to handle policy changes.
	_ = w.SetUpdateCallback(psqlwatcher.DefaultCallback(e))

	// update the policy and notify other enforcers.
	_ = e.SavePolicy()

	// wait for callback.
	for {
		fmt.Println("")
		policies := e.GetPolicy()
		for _, policy := range policies {
			fmt.Println("--> ", policy)
		}
		fmt.Println("")

		time.Sleep(time.Second)
	}
	//fmt.Scanln()
}
