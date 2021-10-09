package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/zainanizar/appointy-task/models"
)

func main(){

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.createUser)
	r.GET("/Post/:id", pc.GetPost)
	r.POST("/Post", uc.creatPost)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session{
	 
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err !=nil{
		panic(err)
	}
	return s
}