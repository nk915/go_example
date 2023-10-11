package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/copier"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const TableNameTblSaasAdmin = "saas.tbl_saas_admin"
const TableNameTblSaasAdminHis = "saas.tbl_saas_admin_his"

// TblSaasAdmin mapped from table <tbl_saas_admin>
type TblSaasAdmin struct {
	Seq              int64  `gorm:"column:seq;type:bigint;primaryKey;comment:서비스 포탈 관리자 SEQ" json:"seq"`                                 // 서비스 포탈 관리자 SEQ
	Userid           string `gorm:"column:userid;type:character varying(32);not null;comment:관리자 계정" json:"userid"`                      // 관리자 계정
	Password         string `gorm:"column:password;type:character varying(128);not null;comment:비밀번호" json:"password"`                   // 비밀번호
	Grade            string `gorm:"column:grade;type:character varying(128);not null;comment:등급(:ADMIN_GRADE)" json:"grade"`             // 등급(:ADMIN_GRADE)
	Name             string `gorm:"column:name;type:character varying(32);not null;comment:성명" json:"name"`                              // 성명
	Department       string `gorm:"column:department;type:character varying(64);comment:부서" json:"department"`                           // 부서
	Position         string `gorm:"column:position;type:character varying(32);comment:직급" json:"position"`                               // 직급
	Rank             string `gorm:"column:rank;type:character varying(32);comment:직책" json:"rank"`                                       // 직책
	Status           string `gorm:"column:status;type:character varying(128);not null;comment:상태(:ADMIN_STATUS)" json:"status"`          // 상태(:ADMIN_STATUS)
	Email            string `gorm:"column:email;type:character varying(64);not null;comment:메일주소" json:"email"`                          // 메일주소
	Phone            string `gorm:"column:phone;type:character varying(32);not null;comment:전화번호" json:"phone"`                          // 전화번호
	IPCheck          bool   `gorm:"column:ip_check;type:boolean;not null;default:true;comment:IP 검사 여부(true: 검사)" json:"ip_check"`       // IP 검사 여부(true: 검사)
	MacCheck         bool   `gorm:"column:mac_check;type:boolean;not null;default:true;comment:MAC 검사 여부(true: 검사)" json:"mac_check"`    // MAC 검사 여부(true: 검사)
	UpdateTime       string `gorm:"column:update_time;type:character(14);not null;comment:수정일시(YYYYMMDDhhmmss)" json:"update_time"`      // 수정일시(YYYYMMDDhhmmss)
	RegistrationDate string `gorm:"column:registration_date;type:character(8);not null;comment:등록일자(YYYYMMDD)" json:"registration_date"` // 등록일자(YYYYMMDD)
}

// TableName TblSaasAdmin's table name
func (*TblSaasAdmin) TableName() string {
	return TableNameTblSaasAdmin
}

// TblSaasAdminHis mapped from table <tbl_saas_admin_his>
type TblSaasAdminHis struct {
	Seq              int64  `gorm:"column:seq;type:bigint;primaryKey;comment:서비스 포탈 관리자 이력 SEQ" json:"seq"`                              // 서비스 포탈 관리자 이력 SEQ
	AdminSeq         int64  `gorm:"column:admin_seq;type:bigint;not null;comment:서비스 포탈 관리자 SEQ" json:"admin_seq"`                       // 서비스 포탈 관리자 SEQ
	Userid           string `gorm:"column:userid;type:character varying(32);not null;comment:관리자 계정" json:"userid"`                      // 관리자 계정
	Password         string `gorm:"column:password;type:character varying(128);not null;comment:비밀번호" json:"password"`                   // 비밀번호
	Grade            string `gorm:"column:grade;type:character varying(128);not null;comment:등급(:ADMIN_GRADE)" json:"grade"`             // 등급(:ADMIN_GRADE)
	Name             string `gorm:"column:name;type:character varying(32);not null;comment:성명" json:"name"`                              // 성명
	Department       string `gorm:"column:department;type:character varying(64);comment:부서" json:"department"`                           // 부서
	Position         string `gorm:"column:position;type:character varying(32);comment:직급" json:"position"`                               // 직급
	Rank             string `gorm:"column:rank;type:character varying(32);comment:직책" json:"rank"`                                       // 직책
	Status           string `gorm:"column:status;type:character varying(128);not null;comment:상태(:ADMIN_STATUS)" json:"status"`          // 상태(:ADMIN_STATUS)
	Email            string `gorm:"column:email;type:character varying(64);not null;comment:메일주소" json:"email"`                          // 메일주소
	Phone            string `gorm:"column:phone;type:character varying(32);not null;comment:전화번호" json:"phone"`                          // 전화번호
	IPCheck          bool   `gorm:"column:ip_check;type:boolean;not null;comment:IP 검사 여부(true: 검사)" json:"ip_check"`                    // IP 검사 여부(true: 검사)
	MacCheck         bool   `gorm:"column:mac_check;type:boolean;not null;comment:MAC 검사 여부(true: 검사)" json:"mac_check"`                 // MAC 검사 여부(true: 검사)
	UpdateTime       string `gorm:"column:update_time;type:character(14);not null;comment:수정일시(YYYYMMDDhhmmss)" json:"update_time"`      // 수정일시(YYYYMMDDhhmmss)
	RegistrationDate string `gorm:"column:registration_date;type:character(8);not null;comment:등록일자(YYYYMMDD)" json:"registration_date"` // 등록일자(YYYYMMDD)
}

// TableName TblSaasAdminHis's table name
func (*TblSaasAdminHis) TableName() string {
	return TableNameTblSaasAdminHis
}

func main() {
	dsn := "host=localhost user=hsck password=hsck@2301 dbname=test_tenant port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Set logger mode as needed
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	// AutoMigrate will create necessary tables if they don't exist
	// db.AutoMigrate(&TblSaasAdmin{}, &AuditLog{})

	// Set up hooks
	db.Callback().Create().After("gorm:create").Register("track_create", Track)
	db.Callback().Update().After("gorm:update").Register("TrackUpdate", Track)
	db.Callback().Delete().After("gorm:delete").Register("TrackUpdate", Track)

	// Example usage
	admin := TblSaasAdmin{Seq: 1, Userid: "kng", Password: "password", Grade: "AA"}
	fmt.Println("Insert admin start")
	db.Create(&admin)
	db.Create(&TblSaasAdmin{Seq: 2, Userid: "kng1"})
	fmt.Printf("Insert admin end \n\n\n")

	//	admins := []TblSaasAdmin{{Seq: 3, Userid: "kng3", Password: "password2"}, {Seq: 5}}
	//	db.Create(admins)

	////	// Update
	//	admin.Name = "NamGyu"
	//	fmt.Println("Update admin")
	//	//db.Save(&admin)
	//	db.Save(&TblSaasAdmin{})
	//	fmt.Printf("Update admin end\n\n\n")
	//
	//	// Delete
	//	fmt.Println("Delete admin")
	//	db.Delete(&admin)
}

func Track(db *gorm.DB) {
	if db.Error != nil {
		fmt.Println(db.Error)
		return
	}

	switch dest := db.Statement.Dest.(type) {
	case *TblSaasAdmin:
		data := TblSaasAdminHis{}
		copier.Copy(&data, dest)

		fmt.Printf("data:%#v \n", data)
		if err := db.Create(&data).Error; err != nil {
			fmt.Println("err: ", err)
		} else {
			fmt.Println("succ ")
		}
	}
}

func TrackCreate(tx *gorm.DB) {
	fmt.Println("--> TrackCreate")
	if tx.Error == nil && tx.Statement.Schema != nil {
		fmt.Println(tx.Statement.Model)
	}
}

// 만일 쿼리가 에러난다면
func TrackUpdate(tx *gorm.DB) {
	fmt.Println("--> TrackUpdate")

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	if tx.Statement.Schema != nil {
		table := tx.Statement.Schema.Table
		model := tx.Statement.Model
		fmt.Println(table)
		fmt.Println(model)
		//fmt.Println(reflect.New(reflect.TypeOf(model).Elem()).Interface())
		//tx.Statement.DB.Session(&gorm.Session{LogMode: true}).Create(&log)
	}
}

func TrackDelete(tx *gorm.DB) {
	fmt.Println("--> TrackDelete")

	if tx.Error == nil && tx.RowsAffected > 0 {
		table := tx.Statement.ReflectValue.MethodByName("TableName").Call([]reflect.Value{})[0].String()
		fmt.Println(table)
		//		log := AuditLog{
		//			Action:   "DELETE",
		//			Table:    table,
		//			RecordID: uint(tx.Statement.DB.Model(tx.Statement.Model).Select("id")),
		//		}
		//		tx.Statement.DB.Session(&gorm.Session{LogMode: true}).Create(&log)
	}
}

func GenSeq(PDB *gorm.DB, table string) (int64, error) {
	fmt.Printf("GenSeq(%s)", table)

	var tblSeq int64
	seqName := strings.ToUpper(table) + "_SEQ"
	sqlString := `SELECT nextval('` + seqName + `')`

	result := PDB.Raw(sqlString).Scan(&tblSeq)
	if result.Error != nil {
		fmt.Printf("Fail: GenSeq (%s): %v", seqName, result.Error)
		return 0, fmt.Errorf("Fail: GenSeq (%s)", seqName)
	}

	fmt.Printf("Succ: GenSeq(%s) %d", table, tblSeq)
	return tblSeq, nil
}
