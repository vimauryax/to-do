package helpers

// import (
// 	"errors"
// 	"strconv"

// 	"github.com/Samratakgec/to-do-go-api/config"
// 	"github.com/Samratakgec/to-do-go-api/models"
// )

// func GetNextTaskID() (string, error) {
// 	var counter models.Counter

// 	// Start a transaction for atomicity
// 	err := config.DB.Transaction(func(tx *gorm.DB) error {
// 		// Lock the row for update
// 		if err := tx.Clauses(gorm.Locking{Strength: "UPDATE"}).First(&counter, "name = ?", "task_id").Error; err != nil {
// 			// If not found, insert initial value
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				counter = models.Counter{Name: "task_id", Seq: 1}
// 				if err := tx.Create(&counter).Error; err != nil {
// 					return err
// 				}
// 			} else {
// 				return err
// 			}
// 		} else {
// 			// Increment sequence
// 			counter.Seq++
// 			if err := tx.Save(&counter).Error; err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		return "", err
// 	}

// 	return strconv.Itoa(counter.Seq), nil
// }
