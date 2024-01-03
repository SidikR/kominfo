package seeder

import (
	"fmt"
	"main/model"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) {
	// // Buat pengguna dengan peran "admin"
	// adminUser := model.User{
	// 	Email:    "admin@gmail.com",
	// 	Username: "admin123",
	// 	Role:     "admin",
	// }

	// // Hash password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123@"), bcrypt.DefaultCost)
	// if err != nil {
	// 	panic("Failed to hash password")
	// }
	// adminUser.Password = string(hashedPassword)

	// // Tambahkan pengguna admin ke database
	// if err := db.Create(&adminUser).Error; err != nil {
	// 	panic("Failed to seed admin user")
	// }

	// // Buat pengguna dengan peran "user"
	// user := model.User{
	// 	Email:    "user@gmail.com",
	// 	Username: "user123",
	// 	Role:     "user",
	// }

	// // Hash password untuk pengguna "user"
	// hashedUserPassword, err := bcrypt.GenerateFromPassword([]byte("user123@"), bcrypt.DefaultCost)
	// if err != nil {
	// 	panic("Failed to hash user password")
	// }
	// user.Password = string(hashedUserPassword)

	// // Tambahkan pengguna "user" ke database
	// if err := db.Create(&user).Error; err != nil {
	// 	panic("Failed to seed user")
	// }

	// Seed untuk mendapatkan data acak yang konsisten
	rand.Seed(time.Now().UnixNano())

	// Membuat 200 data dummy
	for i := 4; i <= 24; i++ {
		kecamatan := fmt.Sprintf("18.01.%02d", i)

		for j := 1; j <= 2; j++ {
			desa := fmt.Sprintf("%s.%04d", kecamatan, 2000+j)

			stunting := model.Stunting{
				NIK:       generateRandomNIK(),
				Name:      generateRandomName(),
				Koodinat:  generateRandomCoordinate(),
				Kecamatan: kecamatan,
				Desa:      desa,
				Status:    "Belum ditangani",
			}

			db.Create(&stunting)
		}
	}

	program:=model.Program{
		Name: "Stunting",
		Deskripsi: "Penyakit anak",
	}
	db.Create(&program);
	
}

// Fungsi untuk menghasilkan NIK acak (hanya contoh, sesuaikan sesuai kebutuhan)
func generateRandomNIK() string {
	return strconv.Itoa(rand.Intn(99999999-10000000+1) + 10000000)
}

// Fungsi untuk menghasilkan nama acak (hanya contoh, sesuaikan sesuai kebutuhan)
func generateRandomName() string {
	names := []string{"Abdullah", "Budi", "Citra", "Dewi", "Eko", "Faisal", "Gita", "Hadi", "Ina", "Joko"}
	return names[rand.Intn(len(names))]
}

// Fungsi untuk menghasilkan koordinat acak (hanya contoh, sesuaikan sesuai kebutuhan)
func generateRandomCoordinate() string {
	latitude := rand.Float64()*10 - 5
	longitude := rand.Float64()*10 + 100
	return fmt.Sprintf("%.6f, %.6f", latitude, longitude)
}