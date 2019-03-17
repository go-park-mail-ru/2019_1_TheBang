package main

//
//import (
//	_ "crypto/md5"
//	_ "github.com/gorilla/mux"
//	_ "io"
//	_ "os"
//	_ "strconv"
//)
//
//func ChangeProfileAvatarHMTLHandler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte(HTML))
//}
//
//func ChangeProfileAvatarHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	nickname, err := NicknameFromCookie(w, r)
//	if err != nil {
//		info := InfoText{Data: err.Error()}
//		err = json.NewEncoder(w).Encode(info)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			log.Printf("MyProfileInfoHandler: %v\n", err.Error())
//
//			return
//		}
//
//		return
//	}
//
//	file, header, err := r.FormFile("photo")
//	if err != nil {
//		w.WriteHeader(http.StatusUnprocessableEntity)
//		info := InfoText{Data: "image was failed in form!"}
//		err := json.NewEncoder(w).Encode(info)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			log.Println(err.Error())
//
//			return
//		}
//
//		return
//	}
//	defer file.Close()
//
//	hasher := md5.New()
//	_, err = io.Copy(hasher, file)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Println(err.Error())
//
//		return
//	}
//	filename := string(hasher.Sum(nil))
//
//	filein, err := header.Open()
//	if err != nil {
//		w.WriteHeader(http.StatusUnprocessableEntity)
//		info := InfoText{Data: "image was failed in form!"}
//		err := json.NewEncoder(w).Encode(info)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			log.Println(err.Error())
//
//			return
//		}
//
//		return
//	}
//	defer filein.Close()
//
//	fileout, err := os.OpenFile("tmp/"+filename, os.O_WRONLY|os.O_CREATE, 0644)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Println("ChangeProfileAvatarHandler: ", "file for img was not created!")
//
//		return
//	}
//	defer fileout.Close()
//
//	_, err = io.Copy(fileout, filein)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Println("ChangeProfileAvatarHandler: ", "img was not saved on disk!")
//
//		return
//	}
//
//	storageProf.mu.Lock()
//	defer storageProf.mu.Unlock()
//
//	updatedProf := storageProf.data[nickname]
//	deletePhoto(updatedProf.Photo)
//
//	updatedProf.Photo = filename
//	storageProf.data[nickname] = updatedProf
//
//	w.WriteHeader(http.StatusAccepted)
//	//toDo возвращать профиль
//	info := InfoText{Data: filename}
//	err = json.NewEncoder(w).Encode(info)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Println(err.Error())
//
//		return
//	}
//}

