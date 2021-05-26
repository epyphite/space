package models

//Config is the basic configuration structure for the service
type Config struct {
	Version      string `json:"version"`      // Version for this agent to run
	TempDir      string `json:"TempDir"`      //TempDir for saving the temporary files locally
	EnableCloud  string `json:"EnableCloud"`  //Enables all Cloud saving
	Debug        string `json:"Debug"`        //Enables verbose logging
	ProcessInput string `json:"ProcessInput"` //Specify if it needs to read and process more PDF
	WebPort      string `json:"WebPort"`      //Specify port to listen to
	WebAddress   string `json:"WebAddress"`   //Address to listen to
	ReadQueue    string `json:"ReadQueue"`    //Flag to specify if we should read the queue
	ContentDir   string `json:"ContentDir"`   //ContentDir
}
