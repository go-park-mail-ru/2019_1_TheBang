package main

type Login struct {
	Nickname string `json:"nickname"`
	Passwd   string `json:"passwd"`
}
//
//
//func CreateAccount(w http.ResponseWriter, r *http.Request) (prof Profile, err error) {
//	signup := Signup{}
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Println(err.Error())
//
//		return prof, err
//	}
//
//	err = json.Unmarshal(body, &signup)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Println(err.Error())
//
//		return prof, err
//	}
//
//	prof = Profile{
//		Nickname: signup.Nickname,
//		Name:     signup.Name,
//		Surname:  signup.Surname,
//		DOB:      signup.DOB,
//	}
//	passwd := signup.Passwd
//
//	storageAcc.mu.Lock()
//	defer storageAcc.mu.Unlock()
//
//	storageProf.mu.Lock()
//	defer storageProf.mu.Unlock()
//
//	if _, ok := storageAcc.data[prof.Nickname]; ok {
//		w.WriteHeader(http.StatusConflict)
//		err := errors.New("This user already exists!")
//
//		return prof, err
//	}
//
//	prof.Photo = DefaultImg
//
//	storageAcc.data[prof.Nickname] = passwd
//	storageProf.data[prof.Nickname] = prof
//
//	return prof, nil
//}
////
//func LoginAcount(username, passwd string) (string, error) {
//	storageAcc.mu.Lock()
//	defer storageAcc.mu.Unlock()
//
//	if pw, ok := storageAcc.data[username]; !ok || pw != passwd {
//		err := errors.New("Wrong answer or password!")
//		return "", err
//	}
//
//	claims := customClaims{
//		username,
//		jwt.StandardClaims{
//			Issuer: ServerName,
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	ss, err := token.SignedString(SECRET)
//	if err != nil {
//		log.Printf("Error with JWT tocken generation: %v\n", err.Error())
//	}
//
//	return ss, nil
//}
//
//func deletePhoto(filename string) {
//	if filename == DefaultImg {
//		return
//	}
//
//	err := os.Remove("tmp/" + filename)
//	if err != nil {
//		log.Printf("Can not remove file tmp/%v\n", filename)
//	}
//}
