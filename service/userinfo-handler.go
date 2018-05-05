package service

import (
    "net/http"
    "strconv"
    //"fmt"
    "github.com/MBControlGroup/login/entities"
    "github.com/MBControlGroup/login/token"
    "github.com/unrolled/render"
    //"github.com/dgrijalva/jwt-go/request"
    //"github.com/dgrijalva/jwt-go"

)

func signinHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        if len(req.Form["account"]) == 0 {
            formatter.JSON(w, http.StatusBadRequest, struct{ success bool }{false})
            return
        }
        u := entities.LoginService.AdminFindByAccount(req.Form["account"][0])
        //u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
        
        if u.Admin_password != req.Form["password"][0] {
            formatter.JSON(w, http.StatusBadRequest, struct{ success bool}{false})
        } else {
            //fmt.Println(u.Admin_id)
            tokenString, err := token.Generate(u.Admin_id)
            checkErr(err)
            formatter.JSON(w, http.StatusOK, struct{Token string}{tokenString})
        }
    }
}

func addAdminHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        var admin entities.Admin
        admin.Admin_password = req.Form["admin_password"][0]
        admin.Admin_account = req.Form["admin_account"][0]
        admin.Admin_type = req.Form["admin_type"][0]
        entities.LoginService.AdminSave(&admin)
    }
}

func testToken(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {

        cookie, err := req.Cookie("token")
	    if err != nil || cookie.Value == ""{
	        formatter.JSON(w, 403, struct{Error string}{"token not found."})
	        return;
	    }

	    user_id, err := token.Valid(cookie.Value)

	    if err != nil {
	        formatter.JSON(w, 403, struct{Error string}{"bad token"})
	        return;
        }

        id, err := strconv.Atoi(user_id)
        admin := entities.LoginService.AdminFindById(id)
        //fmt.Println(user_id)

        formatter.JSON(w, http.StatusOK, 
            struct{ Success bool;
                    Content string;
                    AdminInfo entities.Admin}{
                    true, 
                    "The token is valid.",
                    *admin})
    }
}

func tokenValid(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()

        if len(req.Form["token"]) == 0 {
            formatter.JSON(w, http.StatusBadRequest, 
                struct{ Success bool;
                        Detail  string;
                        id      int}{
                        false, "Notfound token", -1,
                        })
            return
        }

        tokenString := req.Form["token"][0]

	    user_id, err := token.Valid(tokenString)

	    if err != nil {
            formatter.JSON(w, 403, 
                struct{ Success bool;
                        Detail  string;
                        id      int}{
                        false, "Invalid token", -1,
                        })
            return
        }

        id, _ := strconv.Atoi(user_id)
        formatter.JSON(w, http.StatusOK, 
            struct{ Success bool;
                    Content string;
                    Id      int}{
                    true, 
                    "The token is valid.",
                    id})
    }
}