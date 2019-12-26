package hostels

import (
	"database/sql"
	"log"
	"errors"
)

type tableHostel struct {
	id int
	name string
}
type tableSpeciality struct {
	id int
	name string
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func fillHostels(s *Store) ([] tableHostel) {
	var hostels []tableHostel
	rowsHostels,err := s.Db.Query(`SELECT * FROM hostels`)
	defer rowsHostels.Close()
	for rowsHostels.Next() {
		var id int
		var name string
		err = rowsHostels.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		hostels = append(hostels, tableHostel{id,name})
	}
	return hostels
}
func getHostelId(s *Store,hostelName string)(int,error){
	var id int
	row := s.Db.QueryRow(`SELECT id FROM hostels WHERE name=$1`,hostelName)
	err:= row.Scan(&id)
	if err != sql.ErrNoRows && err != nil {
		return id,err
	}
	if id==0 {
		return id,errors.New("Invalid Hostel Name")
	}
	return id,nil
}
func getSpecialityId(s *Store,speciality string)(int,error){
	var id int
	row := s.Db.QueryRow(`SELECT id FROM specialities WHERE name=$1`,speciality)
	err:= row.Scan(&id)
	if err != sql.ErrNoRows && err != nil {
		log.Println("Here")
		return id,err
	}
	if id==0 {
		return id,errors.New("Invalid speciality")
	}
	return id,nil
}

func fillSpecialities(s *Store) ([] tableSpeciality) {
	var specialities []tableSpeciality
	rowsSpecialities,err := s.Db.Query(`SELECT * FROM specialities`)
	defer rowsSpecialities.Close()
	for rowsSpecialities.Next() {
		var id int
		var name string
		err = rowsSpecialities.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		specialities = append(specialities, tableSpeciality{id,name})
	}
	return specialities
}
func getNestedValOfSpec(arr interface{}, speciality string)(int){
	return arr.([]interface{})[0].(map[string]interface{})[speciality].(int)
}
func SpecExist(speciality string, arrHostels []interface{})(bool){
	if _,ok := arrHostels[0].([]interface{})[0].(map[string]interface{})["studentsCount"].([]interface{})[0].(map[string]interface{})[speciality];ok{
		return true
	} else {
		return false
	}
}

func findMin(speciality string, arrHostels []interface{}) (map[string]interface{},error){
	if !SpecExist(speciality, arrHostels){
		return nil,errors.New("this speciality not exist")
	}
	min := arrHostels[0].([]interface{})[0].(map[string]interface{})

	for _, arr := range arrHostels {
	obj:=arr.([]interface{})[0].(map[string]interface{})
	if getNestedValOfSpec(obj["studentsCount"], speciality)< getNestedValOfSpec(min["studentsCount"], speciality){
		min= obj
	}
	}
	return min, nil
}

func makeHostelsInterface(s *Store, hostels []tableHostel, specialities []tableSpeciality)([]interface{},error){
	var arrHostels []interface{}

	for _,hostel := range hostels{
		arrHostels = append(arrHostels, []interface{}{
			map[string]interface{}{
				"hostelId": hostel.id,
				"studentsCount": []interface{}{
					map[string]interface{}{
					},
		 		},
			},
		},
	)
		for _, speciality := range specialities {
			rows,err := s.Db.Query("select count(*) as count_st from students where hostelid=$1 and specialityid=$2", hostel.id, speciality.id)
			if(err!= nil){
				log.Println(err)
				return nil,err
			}
			if rows != nil {
				var count_st int
				for rows.Next() {
					rows.Scan(&count_st)
				}
				arrHostels[hostel.id-1].([]interface{})[0].(map[string]interface{})["studentsCount"].([]interface{})[0].(map[string]interface{})[speciality.name]=count_st
			} 
		}
	}
	return arrHostels, nil
}

// GET request
func (s *Store) getBestHostel(hostel *Hostel) (*Response, error) {
	// Init response
	var response Response

	hostels := fillHostels(s)
	specialities := fillSpecialities(s)

	//interface for our hostels
	arrHostels,err := makeHostelsInterface(s,hostels, specialities)
	if err!= nil {
		return nil,err
	}

	hostelWithMinSpec,err := findMin(hostel.Speciality,arrHostels)
	if err!= nil {
		return nil,err
	}
	response.Hostel = hostelWithMinSpec
	return &response, nil
}

// POST request
func (s *Store) setStudent(sStudent *SendStudent) (*Response, error) {
	// Init response
	var response Response

	var result []interface{}
	result = append(result, []interface{}{ 
		map[string]interface{}{
			"result": "",
		},
	})
	// Get Ids from tables hostels, specialities
	 hostelId,err := getHostelId(s,sStudent.Hostel)
	 if err!=nil {
		 return nil,err
	 }
	 specialityId,err := getSpecialityId(s,sStudent.Speciality)
	 if err!=nil {
		return nil,err
	}
		// Insert Data to table students
	_, err = s.Db.Exec(`
  	INSERT INTO students
  	("name", "hostelid", "specialityid")
  	VALUES
		($1, $2, $3)`,
	 sStudent.Name, hostelId, specialityId)

	var resultString = "The student "+ sStudent.Name+" successfully settled in a hostel " + sStudent.Hostel + " for speciality: " + sStudent.Speciality

	result[0].([]interface{})[0].(map[string]interface{})["result"] = resultString

	response.Hostel = result[0].([]interface{})[0].(map[string]interface{})
	return &response, err
}
