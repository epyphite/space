package storage

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/epyphite/ulid"
	log "github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

type Client struct {
	Database string
	boltDB   *bolt.DB
}

//OpenBoltDb main structure to open
func (bc *Client) OpenBoltDb(dataDir string, dataDbName string) *Client {

	Client := new(Client)
	var err error

	log.Printf("Opening Database %s, %s \n", dataDir, dataDbName)
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(dataDir, 0770)
	}

	var databaseFileName = ""

	databaseFileName = dataDir + string(os.PathSeparator) + dataDbName

	Client.Database = databaseFileName
	Client.boltDB, err = bolt.Open(databaseFileName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	// reproducible entropy source
	entropy := rand.New(rand.NewSource(time.Unix(1000000, 0).UnixNano()))

	// sub-ms safe ULID generator
	ulidSource = ulid.NewMonotonicULIDsource(entropy)

	return Client
}

//Seed is good for creating basic buckets
func (bc *Client) Seed() {
	bc.initializeBucket()
}

//CloseDB will close the FD to the boltdb file
func (bc *Client) CloseDB() {
	bc.boltDB.Close()
}

//initializeBucket will setup file and buckets in tapestryDB.
//This is a boltdb key/Value Storage
// All collections are created here if they dont exists
func (bc *Client) initializeBucket() {

	err := bc.boltDB.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("EpyphiteSpace"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("Engine"))
		if err != nil {
			return fmt.Errorf("could not create users bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("Rocket"))
		if err != nil {
			return fmt.Errorf("could not create nodes bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("Orbit"))
		if err != nil {
			return fmt.Errorf("could not create nodes bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("SpacePort"))
		if err != nil {
			return fmt.Errorf("could not create nodes bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("FuelType"))
		if err != nil {
			return fmt.Errorf("could not create nodes bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("FuelCycle"))
		if err != nil {
			return fmt.Errorf("could not create objectStore bucket: %v", err)
		}
		return err

	})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	/*
		var tempUser v1.ReUser
		// or error handling
		u2 := uuid.NewV4()

		if err != nil {
			log.Println("UserAdd -> Something went wrong: ", err)
			return
		}
		tempUser.UserName = "root"
		tempUser.Password = []byte("ResilientOne!!") //Default Password CHANGE IN PROD
		tempUser.UserID = u2.String()
		tempUser.Token = ""
		tempUser.NodeID = []string{information.GetUID()}
		tempUser.Approved = "true"
		tempUser.Banned = "no"
		tempUser.Role = "Admin"
		tempUser.IsAdmin = "yes"
		err = bc.UserAdd(tempUser)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
	*/
}
