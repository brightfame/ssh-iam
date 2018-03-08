package main

type getKeys struct {
	Username string `cli:"arg"`
}

func (r *getKeys) Run() error {
	// open the existing keys file and display all of the keys
	// this file is populated by the sync command

	return nil
}
