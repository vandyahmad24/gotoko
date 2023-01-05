package seeder

import (
	"fmt"
	cash "vandyahmad/newgotoko/cashier"
	category "vandyahmad/newgotoko/category"
	payment "vandyahmad/newgotoko/payment"

	"gorm.io/gorm"
)

type Seeder interface {
	CashierSeeder() error
	CategorySeeder() error
	PaymentSeeder() error
}

type seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *seeder {
	return &seeder{
		db: db,
	}
}

func (s *seeder) CashierSeeder() error {
	fmt.Println("seeder cashier")
	var count int64
	var oneCashier cash.Cashier
	var allCashier []cash.Cashier
	err := s.db.Model(&oneCashier).Select("count(id)").Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		for i := 1; i <= 11; i++ {
			oneCashier.Name = fmt.Sprintf("Kasir %d", i)
			oneCashier.Passcode = "123456"
			allCashier = append(allCashier, oneCashier)
		}
		err = s.db.Create(&allCashier).Error
		if err != nil {
			return err
		}

	}

	fmt.Println("selesai seeder cashier")

	return nil
}

func (s *seeder) CategorySeeder() error {
	fmt.Println("seeder category")
	var count int64
	var oneCategory category.Category
	var allCategory []category.Category
	err := s.db.Model(&oneCategory).Select("count(id)").Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		cat := category.Category{
			Name: "Body Treatment",
		}
		allCategory = append(allCategory, cat)
		cat = category.Category{
			Name: "Snack",
		}
		allCategory = append(allCategory, cat)
		cat = category.Category{
			Name: "Utilities",
		}
		allCategory = append(allCategory, cat)
		err = s.db.Create(&allCategory).Error
		if err != nil {
			return err
		}

	}

	fmt.Println("selesai seeder category")

	return nil
}

func (s *seeder) PaymentSeeder() error {
	fmt.Println("seeder payment")
	var count int64
	var onePayment payment.Payment
	var allPayment []payment.Payment
	err := s.db.Model(&onePayment).Select("count(id)").Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		cat := payment.Payment{
			Name: "Cash",
			Type: "Cash",
			Logo: "",
		}
		allPayment = append(allPayment, cat)
		logo := "https://rm.id/images/berita/750x390/genjot-layanan-ovo-buka-peluang-kerja-sama-dengan-berbagai-pihak_22246.jpg"
		cat = payment.Payment{
			Name: "OVO",
			Type: "E-WALLET",
			Logo: logo,
		}
		allPayment = append(allPayment, cat)
		logo = "https://asset.kompas.com/crops/ZGHCMpTmxXAigLjYIA42Isbap6Y=/0x0:0x0/780x390/data/photo/2017/01/11/1631493logo-black780x390.jpg"
		cat = payment.Payment{
			Name: "OVO",
			Type: "E-WALLET",
			Logo: logo,
		}
		allPayment = append(allPayment, cat)
		err = s.db.Create(&allPayment).Error
		if err != nil {
			return err
		}

	}

	fmt.Println("selesai seeder payment")

	return nil
}
