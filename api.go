package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}


func WriteJson(w http.ResponseWriter, status int, v any) error{
	w.WriteHeader(status)
	w.Header().Set("Content-Type","application/json")
	return json.NewEncoder(w).Encode(v)
}


func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
   if err  := f(w,r); err != nil{
		WriteJson(w,400,ApiError{Error:err.Error()})
		}
	}
}

type APIServer struct{
	listenAddr string
}

func NewApiServer(listenAddr string) *APIServer{
	return  &APIServer{
     listenAddr: listenAddr,
	}
}

func (s *APIServer) Run(){
	router := mux.NewRouter();

	router.HandleFunc("/account",makeHTTPHandlerFunc(s.handleAccount))
	
	log.Println("json api server is running")

	http.ListenAndServe(s.listenAddr,router)
}


func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error{

	if r.Method == "GET"{
		return s.handleGetAccount(w,r)
	}

	if r.Method == "POST"{
		return s.handleCreateAccount(w,r)
	}

	if r.Method == "DELETE"{
		return s.handleDeleteAccount(w,r)
	}


	return fmt.Errorf("Method not allowed %s",r.Method);
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error{
	account:= NewAccount("Midhun","Binoy")
	return WriteJson(w,http.StatusOK,account)
}


func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}


func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error{
	return nil
}


