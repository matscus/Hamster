package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"
	"github.com/matscus/Hamster/Package/Services/service"
)

//GetAllServiceBins -  handle function, for return ALL servicebins info
func GetAllServiceBins(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	projects := jwttoken.GetUserProjects(strings.TrimSpace(splitToken[1]))
	pgclient := client.PGClient{}.New()
	projectsID, err := pgclient.GetProjectsIDtoString(projects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"Get all ServiceBins error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Upload file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	bins, err := pgclient.GetAllServiceBinsByOwner(projectsID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"Get all ServiceBins error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Upload file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	err = json.NewEncoder(w).Encode(bins)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}

//GetAllServiceBinsType -  handle function, for return ALL servicebins info
func GetAllServiceBinsType(w http.ResponseWriter, r *http.Request) {
	bins, err := client.PGClient{}.New().GetAllServiceBinsType()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"Get all ServiceBins types error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Upload file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	err = json.NewEncoder(w).Encode(bins)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}

//ServiceBins -  handle function, for new,update and delete host
func ServiceBins(w http.ResponseWriter, r *http.Request) {
	var s service.Service
	id, err := strconv.Atoi(r.FormValue("serviceID"))
	if err != nil {
		s.ID = int64(id)
	}
	s.Type = r.FormValue("serviceType")
	s.RunSTR = r.FormValue("runSTR")
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	own := jwttoken.GetUser(strings.TrimSpace(splitToken[1]))
	switch r.Method {
	case "POST":
		file, header, err := r.FormFile("uploadFile")
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Upload file error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Upload file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		s.Name = header.Filename
		err = s.CreateBin(file, own)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Host created \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Host created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "PUT":
		if own == s.Own || own == "admin" {
			err := s.UpdateBin(own)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\":\"Host updated \"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Host updated, but Not Writing to ResponseWriter due: %s", errWrite.Error())
			}
		} else {
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\": You are not a owner for this service}"))
			if errWrite != nil {
				log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}

	case "DELETE":
		if own == s.Own || own == "admin" {
			err := s.DeleteBin()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\":\"Host deleted \"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Host deleted, but Not Writing to ResponseWriter due: %s", errWrite.Error())
			}
		} else {
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\": You are not a owner for this service}"))
			if errWrite != nil {
				log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
	}
}
