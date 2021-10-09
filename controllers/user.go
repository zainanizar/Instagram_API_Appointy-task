package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"labix.org/v2/mgo/bson"
    "strconv"
    "time"
	"github.com/zainanizar/appointy-task/models"
	"github.com/zainanizar/appointy-task/timestamp"
)

type UserController struct {
	session *mgo.Session
}
type PostController struct {
	session *mgo.Session
}

func NewUserController(s * mgo.Session) *UserController{
return &UserController{s}
}

func NewPostController(s * mgo.Session) *PostController{
	return &PostController{s}
	}
	

func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo.golang").C("users").FindId(oid).One(&u); err != null{
		w.WriteHeader(404)
		return
	}

	uj, err :=json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}


func (pc PostController) GetPost (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.Post{}

	if err := pc.session.DB("mongo.golang").C("post").FindId(oid).One(&u); err != null{
		w.WriteHeader(404)
		return
	}

	uj, err :=json.Marshal(u)
	if err!= nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (pc PostController) CreatePost (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := models.Post{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	pc.session.DB("mongo-golang").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

type Timestamp timestamp.Timestamp

func (t *Timestamp) MarshalJSON() ([]byte, error) {
    ts := time.Time(*t).Unix()
    stamp := fmt.Sprint(ts)

    return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
    ts, err := strconv.Atoi(string(b))
    if err != nil {
        return err
    }

    *t = Timestamp(time.Unix(int64(ts), 0))

    return nil
}

func (t Timestamp) GetBSON() (interface{}, error) {
    if time.Time(*t).IsZero() {
        return nil, nil
    }

    return time.Time(*t), nil
}

func (t *Timestamp) SetBSON(raw bson.Raw) error {
    var tm time.Time

    if err := raw.Unmarshal(&tm); err != nil {
        return err
    }

    *t = Timestamp(tm)

    return nil
}

func (t *Timestamp) String() string {
    return time.Time(*t).String()
}
