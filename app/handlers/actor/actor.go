package actor

import (
	"encoding/json"
	"net/http"

	"github.com/waldyrfelix/mockapi/app/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	// db, err := db.OpenDb()
	// if err != nil {
	// 	log.Fatal("Connection failed: ", err)
	// }
	// defer db.Close()

	actors := []models.Actor{
		{ActorId: 1, FirstName: "Marlon", LastName: "Brando"},
		{ActorId: 2, FirstName: "Morgan", LastName: "Freeman"},
	}
	// db.Raw("select * from actor limit 10").Scan(&actors)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actors)
}

// func Get(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	title := vars["title"]
// 	project := getProjectOr404(db, title, w, r)
// 	if project == nil {
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, project)
// }

// func Create(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	project := model.Project{}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&project); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Save(&project).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusCreated, project)
// }

// func Update(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	title := vars["title"]
// 	project := getProjectOr404(db, title, w, r)
// 	if project == nil {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&project); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Save(&project).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, project)
// }

// func Delete(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	title := vars["title"]
// 	project := getProjectOr404(db, title, w, r)
// 	if project == nil {
// 		return
// 	}
// 	if err := db.Delete(&project).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusNoContent, nil)
// }
