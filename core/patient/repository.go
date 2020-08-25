package patient

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tooth-fairy/infrastructure/database"
)

// Repository - represents a abstraction to interact with database
type Repository interface {
	FindAllPatients() ([]Patient, error)
	GetPatient(uid uint32) (*Patient, error)
	CreatePatient(p *Patient) (*Patient, error)
	UpdatePatient(p *Patient, uid uint32) (*Patient, error)
	DeletePatient(uid uint32) (int64, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) CreatePatient(p *Patient) (*Patient, error) {
	err := r.db.Create(p).Error
	if err != nil {
		panic(err)
	}
	return p, nil
}

func (r *repository) FindAllPatients() ([]Patient, error) {
	var patients []Patient
	err := r.db.Model(Patient{}).Limit(100).Find(&patients).Error
	if err != nil {
		panic(err)
	}

	return patients, nil
}

func (r *repository) GetPatient(uid uint32) (*Patient, error) {
	p := Patient{}
	err := r.db.Model(Patient{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		panic(err)
	}

	return &p, nil
}

func (r *repository) UpdatePatient(p *Patient, uid uint32) (*Patient, error) {
	err := r.db.Model(&Patient{}).Where("id = ?", uid).Updates(
		map[string]interface{}{
			"name":       p.Name,
			"phone":      p.Phone,
			"email":      p.Email,
			"age":        p.Age,
			"gender":     p.Gender,
			"updated_at": time.Now(),
		},
	).Error

	if err != nil {
		return &Patient{}, err
	}

	err = r.db.Debug().Model(&Patient{}).Where("id = ?", uid).Take(&p).Error
	if err != nil {
		return &Patient{}, err
	}
	return p, nil
}

func (r *repository) DeletePatient(uid uint32) (int64, error) {
	db := r.db.Model(Patient{}).Where("id = ?", uid).Take(&Patient{}).Delete(&Patient{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// NewRepository - responsible to create a new repository
func newRepository(db *database.Database) Repository {
	_db := db.GetGormClient()

	return &repository{
		db: _db,
	}
}
