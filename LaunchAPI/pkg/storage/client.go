package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	c "epyphite/space/v1/LaunchAPI/pkg/crypto"
	"epyphite/space/v1/LaunchAPI/pkg/models"
	modules "epyphite/space/v1/LaunchAPI/pkg/models/modules"

	"github.com/epyphite/ulid"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/crypto/bcrypt"
)

type Client struct {
	Database string
	boltDB   *bolt.DB
	dataDir  string
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
	Client.dataDir = dataDir
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
	bc.PopulateFromDisk()
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
		_, err = root.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return fmt.Errorf("could not create objectStore bucket: %v", err)
		}
		return err

	})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	var tempUser models.User
	// or error handling
	u2 := uuid.NewV4()

	if err != nil {
		log.Println("UserAdd -> Something went wrong: ", err)
		return
	}
	tempUser.Username = "root"
	tempUser.Password = []byte("ResilientOne!!") //Default Password CHANGE IN PROD
	tempUser.UserID = u2.String()
	tempUser.Token = ""
	tempUser.Approved = true
	tempUser.Role = "Admin"
	tempUser.IsAdmin = true
	err = bc.UserAdd(tempUser)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

}

//PopulateFromDisk data from initial pre set
func (bc *Client) PopulateFromDisk() {
	bc.importOrbitDataFromDisk()
	bc.importSpacePortFromDisk()
	bc.importRocketFromDisk()
}

func readJSONFile(filename string) ([]byte, error) {
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	log.Infoln("Successfully Opened ", filename)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	return byteValue, err
}

func (bc *Client) importOrbitDataFromDisk() {
	log.Infoln("Import Orbit information")

	byteValue, err := readJSONFile(bc.dataDir + "OrbitData.json")
	if err != nil {
		log.Errorln("Error importing file ", err)
		return
	}
	var orbits []modules.Orbit
	json.Unmarshal(byteValue, &orbits)

	for _, orbit := range orbits {
		u2 := uuid.NewV4()
		orbit.ID = u2.String()
		err = bc.OrbitDataAdd(orbit)
		if err != nil {
			log.Errorln("Error import space port ", err)
		}
	}
}

func (bc *Client) importSpacePortFromDisk() {
	log.Infoln("Import Space Port information")

	byteValue, err := readJSONFile(bc.dataDir + "SpacePort.json")
	if err != nil {
		log.Errorln("Error importing file ", err)
		return
	}
	var ports []modules.SpacePort

	json.Unmarshal(byteValue, &ports)

	for _, port := range ports {
		u2 := uuid.NewV4()
		port.ID = u2.String()
		err = bc.SpacePortAdd(port)
		if err != nil {
			log.Errorln("Error import space port ", err)
		}
	}
}

func (bc *Client) importRocketFromDisk() {
	log.Infoln("Importing Rocket Information")
	byteValue, err := readJSONFile(bc.dataDir + "RocketData.json")
	if err != nil {
		log.Errorln("Error importing file ", err)
		return
	}

	var RocketData []modules.Rocket
	json.Unmarshal(byteValue, &RocketData)

	for _, rocket := range RocketData {
		u2 := uuid.NewV4()

		if err != nil {
			log.Errorf("Could not Generate UUID for %s \n", rocket.Name)
		}
		rocket.ID = u2.String()

		err = bc.RocketDataAdd(rocket)
		if err != nil {
			log.Errorln("Error inserting Rocket Data", err)
		}
	}
}

func (bc *Client) importEngineDataFromDisk() {
	log.Infoln("Importing Engine Information")
	byteValue, err := readJSONFile(bc.dataDir + "engineSpecs.json")
	if err != nil {
		log.Errorln("Error importing file ", err)
		return
	}

	var EngineData []modules.EngineSpecs
	json.Unmarshal(byteValue, &EngineData)

	for _, engine := range EngineData {
		u2 := uuid.NewV4()

		if err != nil {
			log.Errorf("Could not Generate UUID for %s \n", engine.Name)
		}
		engine.ID = u2.String()

		err = bc.EngineDataAdd(engine)
		if err != nil {
			log.Errorln("Error inserting Rocket Data", err)
		}
	}
}

//EngineDataAdd will add Rocket information
func (bc *Client) EngineDataAdd(engine modules.EngineSpecs) error {

	engineBytes, err := json.Marshal(engine)

	if err != nil {
		return fmt.Errorf("could not marshal config proto: %v", err)
	}
	err = bc.boltDB.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("Engine")).Put([]byte(engine.ID), engineBytes)
		if err != nil {
			log.Errorf("%s \n", err.Error())
			return fmt.Errorf("could not set Engine: %v", err)
		}
		return nil
	})
	return err
}

//RocketDataAdd will add Rocket information
func (bc *Client) RocketDataAdd(rocket modules.Rocket) error {

	rocketBytes, err := json.Marshal(rocket)

	if err != nil {
		return fmt.Errorf("could not marshal config proto: %v", err)
	}
	err = bc.boltDB.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("Rocket")).Put([]byte(rocket.ID), rocketBytes)
		if err != nil {
			log.Errorf("%s \n", err.Error())
			return fmt.Errorf("could not set Rocket: %v", err)
		}
		return nil
	})
	return err
}

//SpacePortAdd will add Rocket information
func (bc *Client) SpacePortAdd(port modules.SpacePort) error {

	portBytes, err := json.Marshal(port)

	if err != nil {
		return fmt.Errorf("could not marshal config proto: %v", err)
	}
	err = bc.boltDB.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("SpacePort")).Put([]byte(port.ID), portBytes)
		if err != nil {
			log.Errorf("%s \n", err.Error())
			return fmt.Errorf("could not set Port: %v", err)
		}
		return nil
	})
	return err
}

//OrbitDataAdd will add Rocket information
func (bc *Client) OrbitDataAdd(orbit modules.Orbit) error {

	orbitBytes, err := json.Marshal(orbit)

	if err != nil {
		return fmt.Errorf("could not marshal config proto: %v", err)
	}
	err = bc.boltDB.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("Orbit")).Put([]byte(orbit.ID), orbitBytes)
		if err != nil {
			log.Errorf("%s \n", err.Error())
			return fmt.Errorf("could not set Orbit: %v", err)
		}
		return nil
	})
	return err
}

//RocketGetAll Get all database registered Rocket
func (bc *Client) RocketGetAll() ([]*modules.Rocket, error) {
	var rockets []*modules.Rocket

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys

		b := tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("Rocket"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var rocket modules.Rocket
			json.Unmarshal(v, &rocket)
			rockets = append(rockets, &rocket)
		}
		return nil
	})
	return rockets, err
}

//SpacePortGetAll Get all database registered Rocket
func (bc *Client) SpacePortGetAll() ([]*modules.SpacePort, error) {
	var ports []*modules.SpacePort

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys

		b := tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("SpacePort"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var port modules.SpacePort
			json.Unmarshal(v, &port)
			ports = append(ports, &port)
		}
		return nil
	})
	return ports, err
}

//OrbitGetAll Get all database registered Rocket
func (bc *Client) OrbitGetAll() ([]*modules.Orbit, error) {
	var orbits []*modules.Orbit

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys

		b := tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("Orbit"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var orbit modules.Orbit
			json.Unmarshal(v, &orbit)
			orbits = append(orbits, &orbit)
		}
		return nil
	})
	return orbits, err
}

//EngineGetAll Get all database registered Rocket
func (bc *Client) EngineGetAll() ([]*modules.EngineSpecs, error) {
	var engines []*modules.EngineSpecs

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys

		b := tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("Engine"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var engine modules.EngineSpecs
			json.Unmarshal(v, &engine)
			engines = append(engines, &engine)
		}
		return nil
	})
	return engines, err
}

//CheckToken for checkig validity of admin token
func (bc *Client) CheckToken(tokenstring string) (models.Token, error) {
	var _token models.Token
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys

		b := tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("users"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var token models.Token
			json.Unmarshal(v, &token)
			if token.TokenID == tokenstring {

				_token = token
			}

		}
		return nil
	})
	return _token, err
}

//TokenAdd will add a token to the tokens bucket
func (bc *Client) TokenAdd(token models.Token) error {
	//log.Println("UserAdd --> password ", user.Password)
	tokenBytes, err := json.Marshal(&token)

	if err != nil {
		return fmt.Errorf("could not marshal config proto: %v", err)
	}
	err = bc.boltDB.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("users")).Put([]byte(token.TokenID), tokenBytes)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return fmt.Errorf("could not set Token: %v", err)
		}
		return nil
	})
	return nil
}

//UserAdd will add a user using models.users
func (bc *Client) UserAdd(user models.User) error {
	_, err := bc.CheckUserCExists(&user)
	if err != nil {
		return fmt.Errorf("could not marshal config proto: %v", err)
	}
	// or error handling
	u2 := uuid.NewV4()
	if err != nil {
		log.Println("UserAdd -> Something went wrong: ", err)
		return err
	}

	user.UserID = u2.String()
	user.Token = c.CreateHash(user.Email)

	user.Password, _ = bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	//log.Println("UserAdd --> password ", user.Password)
	userBytes, err := json.Marshal(&user)

	if err != nil {
		return fmt.Errorf("could not marshal config proto: %v", err)
	}
	err = bc.boltDB.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("users")).Put([]byte(user.Email), userBytes)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return fmt.Errorf("could not set USER: %v", err)
		}
		return nil
	})
	return nil
}

func decodeUser(data []byte) (models.User, error) {
	var p models.User
	err := json.Unmarshal(data, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

//CheckUserCExists usage is to check the hash of the IOC and search in the database.
func (bc *Client) CheckUserCExists(user *models.User) (*models.User, error) {

	var iuser models.User

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		iochhash := tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("users")).Get([]byte(user.Email))
		if len(iochhash) > 0 {
			var err error
			iuser, err = decodeUser(iochhash)
			if err != nil {
				return fmt.Errorf("Bucket exists 1")
			}
			return nil
		}
		return nil
	})

	return &iuser, err
}

//CheckUser usage for authentication, returns true or false.
func (bc *Client) CheckUser(user models.User) (models.User, bool, error) {

	var iuser models.User
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		iochhash := tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("users")).Get([]byte(user.Email))
		if len(iochhash) > 0 {
			var err error
			iuser, err = decodeUser(iochhash)
			if err != nil {
				return fmt.Errorf("Bucket exists 1")
			}
			return nil
		}
		return fmt.Errorf("User Not Found")
	})
	if err != nil {
		return iuser, false, err
	}

	err = bcrypt.CompareHashAndPassword(iuser.Password, user.Password)

	if err == nil {
		return iuser, true, err
	}
	//log.Println("Error in compare password ", err)
	return iuser, false, err
}

//CheckUserByID will check a user by their ID
func (bc *Client) CheckUserByID(userID string) (models.User, error) {
	var _user models.User
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys

		b := tx.Bucket([]byte("EpyphiteSpace")).Bucket([]byte("users"))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var user models.User
			json.Unmarshal(v, &user)
			if user.UserID == userID {

				_user = user
			}

		}
		return nil
	})
	return _user, err
}
