package main

import (
	"bytes"
	"database/sql"
	"time"
)

var (
	db   = createDB()
	row1 *sql.Rows
	res  sql.Result
)

func createPet(pet Pet) error {
	_, err = db.Exec("insert into pets(name, photo, age, feature, type, kind) values (?, ?, ?, ?, ?, ?)", pet.Name, pet.Photo, pet.Age, pet.Feature, pet.Type, pet.Kind)
	return err
}
func updatePet(pet Pet) error {
	//更新时如果没有修改则保留原数据
	_, err = db.Exec("update pets set name = ?,photo = ?,age = ?,feature = ?,type = ? where id = ?", pet.Name, pet.Photo, pet.Age, pet.Feature, pet.Type, pet.Id)
	return err
}
func deletePet(id int) (error, int64) {
	res, err = db.Exec("delete from pets where id = ?", id)
	deletedRows, _ := res.RowsAffected()
	return err, deletedRows
}
func findPets(conditionQuery1 ConditionQuery1) (error, Page[[]Pet]) {
	var buf bytes.Buffer
	buf.WriteString("select id, age, feature, name, photo, type, kind from pets where 1=1")
	pets := make([]Pet, 0)
	if conditionQuery1.Name != "" {
		buf.WriteString(" and name like concat('%',?,'%')")
	}
	if conditionQuery1.Kind != "" {
		buf.WriteString(" and kind like concat('%',?,'%')")
	}
	if conditionQuery1.Feature != "" {
		buf.WriteString(" and feature like concat('%',?,'%')")
	}
	if conditionQuery1.Size == 0 {
		conditionQuery1.Size = 10
	}
	if conditionQuery1.Current == 0 {
		conditionQuery1.Current = 1
	}
	buf.WriteString(" limit ?, ?")
	row1, err = findPetsStrategy(conditionQuery1, buf)
	if err != nil {
		return err, Page[[]Pet]{}
	}
	for row1.Next() {
		var pet2 Pet
		if err = row1.Scan(&pet2.Id, &pet2.Age, &pet2.Feature, &pet2.Name, &pet2.Photo, &pet2.Type, &pet2.Kind); err != nil {
			return err, Page[[]Pet]{}
		}
		pets = append(pets, pet2)
	}
	err1 := make(chan error)
	count1 := make(chan int)
	go func(err1 chan error, count1 chan int) {
		var count2 int
		err1 <- db.QueryRow("select count(*) from pets").Scan(&count2)
		count1 <- count2
	}(err1, count1)

	if <-err1 != nil {
		return err, Page[[]Pet]{}
	}
	return nil, Page[[]Pet]{conditionQuery1.Current, conditionQuery1.Size, <-count1, pets}
}
func findPetsPhoto() (error, []string) {
	row1, err := db.Query("select photo from pets order by rand() limit 4")
	if err != nil {
		return err, nil
	}
	photos := make([]string, 0, 4)
	for row1.Next() {
		var photo string
		err = row1.Scan(&photo)
		if err != nil {
			return err, nil
		}
		photos = append(photos, photo)
	}
	return nil, photos
}
func userRegister(user User) error {
	_, err = db.Exec("insert into user(name, account, password, sex, role, age, photo) values(?, ?, sha(?), ?, ?, ?, ?)", user.Name, user.Account, user.Password, user.Sex, user.Role, user.Age, user.Photo)
	return err
}
func userRegister2(account *string) error {
	return db.QueryRow("select account from user where account = ?", *account).Scan(account)
}
func userLogin(account string, id *int, role *string, password *string) error {
	return db.QueryRow("select id, password, role from user where account = ?", account).Scan(id, password, role)
}
func userEdit(user User) error {
	_, err = db.Exec("update user set name = ?, age = ?, photo = ?, account = ?, password = sha(?), sex = ?, role = ? where id = ?", user.Name, user.Age, user.Photo, user.Account, user.Password, user.Sex, user.Role, user.Id)
	return err
}
func userDelete(id int) (error, int64) {
	res, err = db.Exec("update user set is_deleted = 1 where id = ?", id)
	deletedRows, _ := res.RowsAffected()
	return err, deletedRows
}
func findUserById(id int) (error, Response2) {
	var response2 Response2
	err = db.QueryRow("select account,role from user where id = ?", id).Scan(&response2.Account, &response2.Role)
	return err, response2
}
func findUser(conditionQuery2 ConditionQuery2) (error, Page[[]User]) {
	var count int
	var buf bytes.Buffer
	buf.WriteString("select id, name, age, account, login_time, sex, role, photo from user where is_deleted = 0")
	if conditionQuery2.Name != "" {
		buf.WriteString(" and name like concat('%',?,'%')")
	}
	if conditionQuery2.Size == 0 {
		conditionQuery2.Size = 10
	}
	if conditionQuery2.Current == 0 {
		conditionQuery2.Current = 1
	}
	buf.WriteString(" limit ?, ?")
	row1, err = findUserStrategy(conditionQuery2, buf, db)
	if err != nil {
		return err, Page[[]User]{}
	}
	users := make([]User, 0)
	for row1.Next() {
		var user2 User
		if err = row1.Scan(&user2.Id, &user2.Name, &user2.Age, &user2.Account, &user2.LoginTime, &user2.Sex, &user2.Role, &user2.Photo); err != nil {
			return err, Page[[]User]{}
		}
		users = append(users, user2)
	}
	err = db.QueryRow("select count(*) from user").Scan(&count)
	if err != nil {
		return err, Page[[]User]{}
	}
	return nil, Page[[]User]{conditionQuery2.Current, conditionQuery2.Size, count, users}
}
func findAllUser() (error, []Response7) {
	row1, err := db.Query("select id, name from user where is_deleted = 0")
	if err != nil && err != sql.ErrNoRows {
		return err, nil
	}
	users := make([]Response7, 0)
	for row1.Next() {
		var response7 Response7
		if err = row1.Scan(&response7.Id, &response7.Name); err != nil {
			return err, nil
		}
		users = append(users, response7)
	}
	return nil, users
}
func findPetsStrategy(conditionQuery1 ConditionQuery1, buf bytes.Buffer) (*sql.Rows, error) {
	current := (conditionQuery1.Current - 1) * conditionQuery1.Size
	if conditionQuery1.Name == "" && conditionQuery1.Kind != "" && conditionQuery1.Feature == "" { //010
		return db.Query(buf.String(), conditionQuery1.Kind, current, conditionQuery1.Size)
	} else if conditionQuery1.Name != "" && conditionQuery1.Kind == "" && conditionQuery1.Feature == "" { //100
		return db.Query(buf.String(), conditionQuery1.Name, current, conditionQuery1.Size)
	} else if conditionQuery1.Name != "" && conditionQuery1.Kind != "" && conditionQuery1.Feature == "" { //110
		return db.Query(buf.String(), conditionQuery1.Name, conditionQuery1.Kind, current, conditionQuery1.Size)
	} else if conditionQuery1.Name != "" && conditionQuery1.Kind != "" && conditionQuery1.Feature != "" { //111
		return db.Query(buf.String(), conditionQuery1.Name, conditionQuery1.Kind, conditionQuery1.Feature, current, conditionQuery1.Size)
	} else if conditionQuery1.Name == "" && conditionQuery1.Kind != "" && conditionQuery1.Feature != "" { //011
		return db.Query(buf.String(), conditionQuery1.Kind, conditionQuery1.Feature, current, conditionQuery1.Size)
	} else if conditionQuery1.Name != "" && conditionQuery1.Kind == "" && conditionQuery1.Feature != "" { //101
		return db.Query(buf.String(), conditionQuery1.Name, conditionQuery1.Feature, current, conditionQuery1.Size)
	} else if conditionQuery1.Name == "" && conditionQuery1.Kind == "" && conditionQuery1.Feature != "" { //001
		return db.Query(buf.String(), conditionQuery1.Feature, current, conditionQuery1.Size)
	} else {
		return db.Query(buf.String(), current, conditionQuery1.Size)
	}
}
func findUserStrategy(conditionQuery2 ConditionQuery2, buf bytes.Buffer, db *sql.DB) (*sql.Rows, error) {
	current := (conditionQuery2.Current - 1) * conditionQuery2.Size
	if conditionQuery2.Name == "" {
		return db.Query(buf.String(), current, conditionQuery2.Size)
	} else {
		return db.Query(buf.String(), conditionQuery2.Name, current, conditionQuery2.Size)
	}
}
func setLoginTime(username string) error {
	_, err = db.Exec("update user set login_time = ? where account = ?", time.Now(), username)
	return err
}
func addVaccine(conditionQuery6 ConditionQuery6) error {
	_, err = db.Exec("insert into vaccines(vaccine_type_id,inoculability,pet_id) values(?, ?, ?)", conditionQuery6.VaccineTypeId, conditionQuery6.Inoculability, conditionQuery6.PetId)
	return err
}
func findVaccineByName(conditionQuery3 ConditionQuery3) (error, Page[[]Response4]) {
	if conditionQuery3.Current == 0 {
		conditionQuery3.Current = 1
	}
	if conditionQuery3.Size == 0 {
		conditionQuery3.Size = 10
	}
	var buf bytes.Buffer
	buf.WriteString("select a.id, a.inoculability, a.pet_id, b.id, b.type, c.name from vaccines a join vaccine_type b ")
	buf.WriteString("on a.vaccine_type_id = b.id join pets c on a.pet_id = c.id  where 1=1 ")
	//代码拼接语句太麻烦，使用inner join代替.即使inner join有性能问题
	if conditionQuery3.Id != 0 {
		buf.WriteString("and b.id =? ")
	}
	if conditionQuery3.PetId != 0 {
		buf.WriteString("and a.pet_id = ?")
	}
	buf.WriteString("limit ?, ?")
	row1, err = findVaccineStrategy(conditionQuery3, buf, db)
	if err != nil {
		return err, Page[[]Response4]{}
	}
	response4List := make([]Response4, 0)
	for row1.Next() {
		var response4 Response4
		if err = row1.Scan(&response4.Id, &response4.Inoculability, &response4.Pet.Id,
			&response4.VaccineType.Id, &response4.VaccineType.Type, &response4.Pet.Name); err != nil {
			return err, Page[[]Response4]{}
		}
		response4List = append(response4List, response4)

	}
	err1 := make(chan error)
	count := make(chan int)
	go func(err1 chan error, count chan int) {
		var count0 int
		err1 <- db.QueryRow("select count(*) from vaccines a join vaccine_type b on a.vaccine_type_id = b.id").Scan(&count0)
		count <- count0
	}(err1, count)
	err = <-err1
	if err != nil {
		return err, Page[[]Response4]{}
	}
	return nil, Page[[]Response4]{conditionQuery3.Current, conditionQuery3.Size, <-count, response4List}
}
func findVaccineById(id int) (error, []Vaccine) {
	row1, err = db.Query("select id, vaccine_type_id, inoculability, pet_id from vaccines where pet_id = ?", id)
	if err != nil {
		return err, make([]Vaccine, 0)
	}
	vaccines := make([]Vaccine, 0)
	for row1.Next() {
		var vaccine1 Vaccine
		if err = row1.Scan(&vaccine1.Id, &vaccine1.VaccineType.Id, &vaccine1.Inoculability, &vaccine1.PetId); err != nil {
			return err, make([]Vaccine, 0)
		}
		err = db.QueryRow("select type from vaccine_type where id = ?", vaccine1.VaccineType.Id).Scan(&vaccine1.VaccineType.Type)
		if err != nil {
			return err, make([]Vaccine, 0)
		}
		vaccines = append(vaccines, vaccine1)
	}
	return nil, vaccines
}
func deleteVaccine(id int) (error, int64) {
	res, err := db.Exec("delete from vaccines where id = ?", id)
	deletedRows, _ := res.RowsAffected()
	return err, deletedRows
}
func addVaccineType(vaccineType VaccineType) error {
	_, err = db.Exec("insert into vaccine_type(type) values(?)", vaccineType.Type)
	return err
}
func findAllVaccineType() (error, []VaccineType) {
	row1, err = db.Query("select id, type from vaccine_type")
	vaccineTypes := make([]VaccineType, 0)
	for row1.Next() {
		var vaccineType VaccineType
		err = row1.Scan(&vaccineType.Id, &vaccineType.Type)
		if err != nil {
			return err, make([]VaccineType, 0)
		}
		vaccineTypes = append(vaccineTypes, vaccineType)
	}
	return nil, vaccineTypes
}
func findVaccineType(conditionQuery4 ConditionQuery4) (error, Page[[]VaccineType]) {
	if conditionQuery4.Current == 0 {
		conditionQuery4.Current = 1
	}
	if conditionQuery4.Size == 0 {
		conditionQuery4.Size = 10
	}
	row1, err = db.Query("select id, type from vaccine_type limit ?, ?", (conditionQuery4.Current-1)*conditionQuery4.Size, conditionQuery4.Size)
	vaccineTypes := make([]VaccineType, 0)
	for row1.Next() {
		var vaccineType1 VaccineType
		if err = row1.Scan(&vaccineType1.Id, &vaccineType1.Type); err != nil {
			return err, Page[[]VaccineType]{}
		}
		vaccineTypes = append(vaccineTypes, vaccineType1)
	}
	var count int
	err = db.QueryRow("select count(*) from vaccine_type").Scan(&count)
	if err != nil {
		return err, Page[[]VaccineType]{}
	}
	return nil, Page[[]VaccineType]{conditionQuery4.Current, conditionQuery4.Size, count, vaccineTypes}
}
func deleteVaccineType(id int) (error, int64) {
	res, err = db.Exec("delete from vaccine_type where id = ?", id)
	deletedRows, _ := res.RowsAffected()
	return err, deletedRows
}
func addBadRecords(badRecord BadRecord) error {
	_, err = db.Exec("insert into bad_records(bad_record, photo, user_id) values(?, ?, ?)", badRecord.BadRecord, badRecord.Photo, badRecord.UserId)
	return err
}
func findBadRecords(conditionQuery5 ConditionQuery5) (error, Page[[]Response6]) {
	if conditionQuery5.Current == 0 {
		conditionQuery5.Current = 1
	}
	if conditionQuery5.Size == 0 {
		conditionQuery5.Size = 10
	}
	row1, err = db.Query("select id, bad_record, photo, user_id from bad_records where examined = ? limit ?, ?", conditionQuery5.Examined, (conditionQuery5.Current-1)*conditionQuery5.Size, conditionQuery5.Size)
	if err != nil {
		return err, Page[[]Response6]{}
	}
	badRecordList := make([]Response6, 0)
	for row1.Next() {
		var badRecord Response6
		if err = row1.Scan(&badRecord.Id, &badRecord.BadRecord, &badRecord.Photo, &badRecord.User.Id); err != nil {
			return err, Page[[]Response6]{}
		}
		if err = db.QueryRow("select name from user where id = ?", badRecord.User.Id).Scan(&badRecord.User.Name); err != nil {
			return err, Page[[]Response6]{}
		}
		badRecordList = append(badRecordList, badRecord)
	}
	go func(total *int, err1 *error) {
		*err1 = db.QueryRow("select count(*) from bad_records").Scan(total)
	}(&conditionQuery5.Total, &err)
	if err != nil {
		return err, Page[[]Response6]{}
	}
	return nil, Page[[]Response6]{conditionQuery5.Current, conditionQuery5.Size, conditionQuery5.Total, badRecordList}
}
func updateBadRecords(id int, examined bool) error {
	_, err = db.Exec("update bad_records set examined = ? where id = ?", examined, id)
	return err
}
func deleteBadRecords(id int) (error, int64) {
	res, err = db.Exec("delete from bad_records where id = ?", id)
	deletedRows, _ := res.RowsAffected()
	return err, deletedRows
}
func findVaccineStrategy(conditionQuery3 ConditionQuery3, buf bytes.Buffer, db *sql.DB) (*sql.Rows, error) {
	current := (conditionQuery3.Current - 1) * conditionQuery3.Size
	if conditionQuery3.Id != 0 && conditionQuery3.PetId != 0 {
		return db.Query(buf.String(), conditionQuery3.Id, conditionQuery3.PetId, current, conditionQuery3.Size)
	} else if conditionQuery3.Id != 0 && conditionQuery3.PetId == 0 {
		return db.Query(buf.String(), conditionQuery3.Id, current, conditionQuery3.Size)
	} else if conditionQuery3.Id == 0 && conditionQuery3.PetId != 0 {
		return db.Query(buf.String(), conditionQuery3.PetId, current, conditionQuery3.Size)
	} else {
		return db.Query(buf.String(), current, conditionQuery3.Size)
	}
}
