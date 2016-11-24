package controllers

import (
	"bytes"
	"github.com/antonholmquist/jason"
	fbook "github.com/huandu/facebook"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	//"log"
	"github.com/ganggas95/trawanganserver/app"
	"github.com/ganggas95/trawanganserver/app/job"
	"github.com/ganggas95/trawanganserver/app/models"
	"github.com/ganggas95/trawanganserver/app/routes"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	*revel.Controller
	job.GoogleHandler
	job.FbHandler
	job.GplusHandler
	job.TwitHandler
}

type AccessToken struct {
	Token  string
	Expiry int64
}

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func (c App) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.GetUser(username)
	}
	return nil
}

func (c App) SetUser() revel.Result {
	user := c.connected()
	if user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}
func (c App) Index() revel.Result {
	user := c.connected()
	if user != nil {
		return c.Redirect(routes.Persons.List(""))
	}
	return c.Render()
}
func (c App) Login() revel.Result {
	user := c.connected()
	if user != nil {
		return c.Redirect(routes.Persons.List(""))
	}
	return c.Render()
}
func (c App) Register() revel.Result {
	user := c.connected()
	if user != nil {
		return c.Redirect(routes.Persons.List(""))
	}
	return c.Render()
}

func (c App) GetUser(username string) *models.User {
	var user models.User
	err := app.GORM.Where("username = ? OR email = ?", username, username).Find(&user)
	if err != nil {
		panic(err.Error)
	}
	return &user
}
func (c App) GetUserWitId(idUser int64) *models.User {
	var users models.User
	err := app.GORM.First(&users, idUser)
	if err != nil {
		panic(err.Error)
	}
	return &users
}

func (c App) AddUser(user models.User, password string) revel.Result {
	var usr models.User
	db := app.GORM.Where("email = ? OR username = ?", user.Email, user.Username).Find(&usr)
	if !db.RecordNotFound() {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Error("Your email have been registered!")
		return c.Redirect(routes.App.Register())
	}
	c.Validation.Required(password)
	user.Validation(c.Validation)
	user.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	tok := job.RandomToken(32)
	err3 := job.SendToken(user.Email, tok)
	if err3 != nil {
		panic(err3)
	}
	var token models.UserToken
	token.AccessToken = tok
	token.User = user
	db = app.GORM.Create(&token)
	if db.Error != nil {
		panic(db.Error)
	}
	c.Flash.Success("We have send to your emails")
	return c.Redirect(routes.App.Login())
}

func (c App) AddUserWithSosmed(user models.User, password, verifyPassword string) revel.Result {
	var usr models.User
	db := app.GORM.Where("email = ?", user.Email).Find(&usr)
	if !db.RecordNotFound() {
		c.Flash.Error("You was registered in system")
		return c.Redirect(routes.App.Login())
	}
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == password).
		Message("Password Not Match")
	user.Validation(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.SetUp(user))
	}
	user.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	tok := job.RandomToken(32)
	err3 := job.SendToken(user.Email, tok)
	if err3 != nil {
		panic(err3)
	}
	var token models.UserToken
	token.AccessToken = tok
	token.Expiry = job.TokenExpiry()
	token.User = user
	err := app.GORM.Create(&token)
	if err.Error != nil {
		panic(err.Error)
	}
	c.Flash.Success("We have send to your emails")
	return c.Redirect(routes.App.Login())
}
func (c App) AuthApp(username, password string, remember bool) revel.Result {
	user := c.GetUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = user.Username
			c.RenderArgs["user"] = user
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			if !user.Verify {
				c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
				return c.Redirect(routes.Persons.UnverifyAcc())

			} else {
				c.Flash.Success("Wellcome here, " + user.Username)
				return c.Redirect(routes.Persons.List(""))
			}

		}
	}
	c.Flash.Error("Login Failed")
	return c.Redirect(routes.App.Login())
}

func (c App) SetUp(user models.User) revel.Result {
	usr := c.connected()
	if usr != nil {
		return c.Redirect(routes.Persons.List(""))
	}
	c.RenderArgs["user"] = user
	return c.Render(user)

}

func (c App) RegisterWithFacebook() revel.Result {
	url := c.FbHandler.GetUrlFb()
	return c.Redirect(url)
}
func (c App) AuthFb(code string) revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	token := c.GetToken(code)
	response := c.GetResponse(token)
	str := job.ReadHttpBody(response)
	usr, _ := jason.NewObjectFromBytes([]byte(str))
	id, _ := usr.GetString("id")

	var userfb models.User
	err := app.GORM.Where("fbid = ?", id).Find(&userfb)
	if !err.RecordNotFound() {
		c.Session["user"] = userfb.Username
		c.RenderArgs["user"] = userfb
		return c.Redirect(routes.Persons.List(""))
	}

	res1, _ := fbook.Get("/"+id, fbook.Params{
		"fields":       "name",
		"access_token": token.AccessToken,
	})
	res2, _ := fbook.Get("/"+id, fbook.Params{
		"fields":       "email",
		"access_token": token.AccessToken,
	})

	var email string
	if res2["email"] != nil {
		email = res2["email"].(string)

	} else {
		email = ""
	}
	username := strings.Split(email, "@")[0]
	nama := res1["name"].(string)
	user := new(models.User)
	user.Nama = nama
	user.Email = email
	user.Username = username
	user.FbId = id
	return c.Redirect(routes.App.SetUp(&user))

}

/*
func (c App) RegisterWithGoogle() revel.Result {
	url := c.GetUrl()
	fmt.Println(url)
	return c.Redirect(url)
}

func (c App) AuthGmail(code string) revel.Result {
	token := c.GetTokenGmail(code)
	fmt.Println(token)
	gmailService := c.GetClient(token)
	_, err1 := gmailService.Users.GetProfile("me").Do()
	if err1 != nil {
		panic(err1)
	}
	return c.Redirect(routes.App.Login())
}
*/
func (c App) RegisterWithGPlus() revel.Result {
	url := c.GetUrlPlus()
	return c.Redirect(url)
}

func (c App) GplusAuth(code string) revel.Result {
	token := c.GetTokenPlus(code)
	client := c.GetClientPlus(token)
	plusService := c.GetServicePlus(client)
	people := c.GetPeoplePlus(plusService)
	nama := people.Name.FamilyName
	id := people.Id
	var usr models.User
	err := app.GORM.Where("gplusid = ?", id).Find(&usr)
	if !err.RecordNotFound() {
		c.Session["user"] = usr.Username
		c.RenderArgs["user"] = usr
		return c.Redirect(routes.Persons.List(""))
	}

	var username string
	if people.Nickname == "" {
		username = people.Name.FamilyName
	} else {
		username = people.Nickname
	}

	user := new(models.User)
	user.Nama = nama
	user.Username = username
	user.GplusId = id
	if len(people.Emails) > 0 {
		user.Email = people.Emails[0].Value
		return c.Redirect(routes.App.SetUp(&user))
	}
	return c.Redirect(routes.App.Login())
}

func (c App) LoginWithGPlus() revel.Result {
	url := c.GetUrlPlus()
	return c.Redirect(url)
}

func (c App) GplusLogin(code string) revel.Result {
	token := c.GetTokenPlus(code)
	client := c.GetClientPlus(token)
	plusService := c.GetServicePlus(client)
	people := c.GetPeoplePlus(plusService)
	nama := people.Name.FamilyName
	id := people.Id
	var usr models.User
	err := app.GORM.Where("gplusid = ?", id).Find(&usr)
	if !err.RecordNotFound() {
		c.Session["user"] = usr.Username
		c.RenderArgs["user"] = usr
		c.Flash.Success("Wellcome " + nama)
		return c.Redirect(routes.Persons.List(""))
	}
	c.Flash.Error("You not registered. Pleas sign up!")
	return c.Redirect(routes.App.Login())
}

func (c App) LoginWithFacebook() revel.Result {
	url := c.FbHandler.GetUrlFb()
	return c.Redirect(url)
}
func (c App) LoginFb(code string) revel.Result {
	u := c.connected()
	if u != nil {
		c.Session["user"] = u.Username
		c.RenderArgs["user"] = u
		return c.Redirect(routes.Persons.List(""))
	}
	token := c.GetToken(code)
	response := c.GetResponse(token)
	str := job.ReadHttpBody(response)
	usr, _ := jason.NewObjectFromBytes([]byte(str))
	id, _ := usr.GetString("id")
	var userfb models.User
	err := app.GORM.Where("fbid = ?", id).Find(&userfb)
	if !err.RecordNotFound() {
		c.Session["user"] = userfb.Username
		c.RenderArgs["user"] = userfb
		return c.Redirect(routes.Persons.List(""))
	}
	c.Flash.Error("You not registered. Pleas sign up!")
	return c.Redirect(routes.App.Login())
}

func (c App) VerifyAcoount(code string) revel.Result {
	var token models.UserToken
	err := app.GORM.Where(models.UserToken{AccessToken: code}).Find(&token)
	if err.Error != nil {
		panic(err.Error)
	}
	var user models.User
	err = app.GORM.Find(&user, &token.User)
	if err.Error != nil {
		panic(err.Error)
	}
	if !token.Status {
		token.Status = true
		token.User = user
		user.Verify = true
		err = app.GORM.Save(&token)
		err = app.GORM.Save(&user)
		if err.Error != nil {
			panic(err.Error)
		}
		c.Session["user"] = user.Username
		c.RenderArgs["user"] = user
		return c.Redirect(routes.Persons.List(""))
	} else {
		c.Flash.Error("Your Account was active. ")
		return c.Redirect(routes.App.Login())
	}
}

func (c App) RegisterWithTwit() revel.Result {
	url, _, _ := c.GetUrlTwit()
	return c.Redirect(url)
}
func (c App) AuthTwit(oauth_toke string) revel.Result {
	user := c.VerifyTwit(oauth_toke)

	return c.Render(user)
}
func (a App) Logout() revel.Result {
	if user := a.connected(); &user == nil {
		a.Flash.Error("Please log in first")
		return a.Redirect(routes.App.Login())
	}

	for k := range a.Session {
		delete(a.Session, k)
	}
	return a.Redirect(routes.App.Login())
}

func (c App) UploadFoto() revel.Result {
	user := c.connected()
	if &user == nil {
		return c.Redirect(routes.App.Index())
	}
	return c.Render()
}

func (c App) Upload(avatar []byte) revel.Result {
	user := c.connected()
	if user == nil {
		return c.Redirect(routes.App.Index())
	}
	c.Validation.Required(avatar)
	c.Validation.MinSize(avatar, 2*KB).
		Message("Minimum a file size of 2KB expected")
	c.Validation.MaxSize(avatar, 2*MB).
		Message("File cannot be larger than 2MB")

	// Check format of the file.
	conf, format, err1 := image.DecodeConfig(bytes.NewReader(avatar))
	c.Validation.Required(err1 == nil).Key("avatar").
		Message("Incorrect file format")
	c.Validation.Required(format == "jpeg" || format == "png").Key("avatar").
		Message("JPEG or PNG file format is expected")

	// Check resolution.
	c.Validation.Required(conf.Height >= 150 && conf.Width >= 150).Key("avatar").
		Message("Minimum allowed resolution is 150x150px")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.UploadFoto())
	}

	m := c.Request.MultipartForm
	for fname, _ := range m.File {
		fheader := m.File[fname]
		for i, _ := range fheader {
			file, err := fheader[i].Open()
			defer file.Close()
			if err != nil {
				panic(err)
			}
			root_path := revel.BasePath
			dst_path := "public/data/" + user.Username
			path := filepath.Join(root_path, dst_path)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				err2 := os.Mkdir(path, os.ModePerm)
				if err2 != nil {
					panic(err2)
				}
			} else {
				dst, err := os.Create(path + "/" + fheader[i].Filename)
				if err != nil {
					panic(err)
				}
				if _, err := io.Copy(dst, file); err != nil {
					panic(err)
				}
				defer dst.Close()
				defer file.Close()
				namaFile, ok := job.RenameFile(dst, "profile")
				if ok != nil {
					panic(ok)
				}
				var foto models.UserFoto
				foto.Foto = namaFile
				foto.Height = conf.Height
				foto.Width = conf.Width
				foto.Format = format
				foto.Size = len(avatar)
				foto.UserId = user.UID
				db := app.GORM.Create(&foto)
				if db.Error != nil {
					panic(db.Error)
				}

			}

		}
	}
	c.Flash.Success("Upload Success")
	return c.Redirect(routes.Agent.Index())

}
