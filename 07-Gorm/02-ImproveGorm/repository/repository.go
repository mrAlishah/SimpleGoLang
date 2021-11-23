package repository

import (
	"fmt"
	"simplegorm/entity"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	InitUserData()
	Add(user entity.User) (*entity.User, error)
	Update(user entity.User) error
	Delete(user entity.User) error
	FindFirst(id int) (*entity.User, error)
	FindAll(where string) ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

var models = []interface{}{
	&entity.User{},
}

var users []entity.User = []entity.User{
	entity.User{UserName: "111", FirstName: "ali", LastName: "feyzi"},
	entity.User{UserName: "222", FirstName: "mahtab", LastName: "ebrahimi"},
	entity.User{UserName: "333", FirstName: "ali", LastName: "aslani"},
	entity.User{UserName: "444", FirstName: "mostafa", LastName: "alishah"},
}

// func New(db *gorm.DB) UserRepository{
// 	return &userRepository{
// 		db: db,
// 	}
// }

func CreateUserRepository(db *gorm.DB) (UserRepository, error) {
	repo := &userRepository{
		db,
	}

	fmt.Printf("current db name: %s\n", db.Migrator().CurrentDatabase())
	//err := db.AutoMigrate(models...)
	db.Migrator().DropTable(&entity.User{})
	err := db.Migrator().CreateTable(&entity.User{})
	if err != nil {
		return repo, errors.Wrap(err, "failed to auto migrate models")
	}
	return repo, nil
}

func (repo *userRepository) InitUserData() {
	println("=>InitUserData:: Inserted Users")
	for _, user := range users {
		repo.db.Create(&user)
		fmt.Printf("%d = inserted %v \n", user.ID, user)
	}
	setDivider()
}

func (repo *userRepository) Add(user entity.User) (*entity.User, error) {
	println("=>Add:: Create New User")

	err := repo.db.Create(&user).Error
	if err != nil {
		fmt.Printf("**Err:: Create New User : %s", err.Error())
		return nil, err
	} else {
		fmt.Printf("%v\n", user)
		setDivider()
	}

	return &user, nil
}

func (repo *userRepository) Update(user entity.User) error {
	println("=>Update:: Update User")

	err := repo.db.Save(&user).Error
	if err != nil {
		fmt.Printf("**Err:: Update User : %s", err.Error())
		return err
	} else {
		fmt.Printf("%v\n", user)
		setDivider()
	}

	return nil
}

func (repo *userRepository) Delete(user entity.User) error {
	println("=>Delete:: Delete User")

	err := repo.db.Delete(&user).Error
	if err != nil {
		fmt.Printf("**Err:: Update User : %s", err.Error())
		return err
	} else {
		fmt.Printf("%v\n", user)
		setDivider()
	}

	return nil
}

func (repo *userRepository) FindFirst(id int) (*entity.User, error) {
	println("=>FindFirst:: Find First User ID")

	user := entity.User{}

	err := repo.db.First(&user, id).Error // find user with integer primary key
	if errors.Is(err, gorm.ErrRecordNotFound) {
		println("**Err:: Record Not Found")
		return nil, err
	} else {
		fmt.Printf("%v\n", user)
		setDivider()
	}

	return &user, nil
}

func (repo *userRepository) FindAll(where string) ([]entity.User, error) {
	println("=>FindAll:: Find Query Users")

	users := []entity.User{}
	var result *gorm.DB

	if where == "" {
		result = repo.db.Find(&users)
	} else {
		result = repo.db.Raw(where).Scan(&users)
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		println("**Err:: Record Not Found")
		return nil, result.Error
	} else {
		fmt.Printf("ResultCount =%d\n", result.RowsAffected)
		setDivider()
	}

	return users, nil
}

func setDivider() {
	println("==============================\n")
}

func deepCopy(user *entity.User) (*entity.User, error) {
	other := &entity.User{}

	err := copier.Copy(other, user)
	if err != nil {
		return nil, fmt.Errorf("cannot copy data: %w", err)
	}

	return other, nil
}
