package models

//Config is the basic configuration structure for the service
type Config struct {
	CloudNotificationARN     string `json:"CloudNotificationARN"`
	CloudNotificationTopic   string `json:"CloudNotificationTopic"`
	SourceType               string `json:"SourceType"`               // AWS | Local
	CloudSourceRegion        string `json:"CloudSourceRegion"`        // If required, default is us-east-2
	CloudSourceStorage       string `json:"CloudSourceStorage"`       // Source Folder
	CloudSourceProvider      string `json:"CloudSourceProvider"`      // if Enable Cloud you should specify AWS |GCLOUD | AZURE
	CloudSourceKey           string `json:"CloudSourceKey"`           // Source name will hold the file could be null
	CloudDestinationProvider string `json:"CloudDestinationProvider"` // if Enable Cloud you should specify AWS |GCLOUD | AZURE
	CloudDestinationRegion   string `json:"CloudDestinationRegion"`   // If required, default is us-east-2
	CloudDestinationStorage  string `json:"CloudDestinationStorage"`  //Bucket or Root Key folder
	CloudDestinationKey      string `json:"CloudDestinationKey"`      //Folder or Path
	TempDir                  string `json:"TempDir"`                  //TempDir for saving the temporary files locally
	EnableCloud              string `json:"EnableCloud"`              //Enables all Cloud saving
	Debug                    string `json:"Debug"`                    //Enables verbose logging
	ProcessInput             string `json:"ProcessInput"`             //Specify if it needs to read and process more PDF
	WebPort                  string `json:"WebPort"`                  //Specify port to listen to
	WebAddress               string `json:"WebAddress"`               //Address to listen to
	ProcessTable             string `json:"ProcessTable"`             //Job Table to store current jobs
	ReadQueue                string `json:"ReadQueue"`                //Flag to specify if we should read the queue
	ContentDir               string `json:"ContentDir"`               //ContentDir
}
