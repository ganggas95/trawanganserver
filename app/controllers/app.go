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
	"github.com/ganggas95/trawanganserver/app/models"
	"github.com/ganggas95/trawanganserver/app/routes"
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	*revel.Controller
	Database
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
	CheckError(err.Error)
	return &user
}
func (c App) GetUserWitId(idUser int64) *models.User {
	var users models.User
	err := app.GORM.First(&users, idUser)
	CheckError(err.Error)
	return &users
}

func (c App) AddUser(user models.User, password string, avatar []byte) revel.Result {
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
	CheckError(err3)
	var token models.UserToken
	token.AccessToken = tok
	token.User = user
	err := app.GORM.Create(&token)
	err = app.GORM.initDb().Create(&user)
	CheckError(err.Error)
	c.Flash.Success("We have send to your emails")
	return c.Redirect(routes.App.Login())
}

func (c App) AddUserWithSosmed(email, nama, username, password, verifyPassword, sosmedid, tipe string) revel.Result {
	var user models.User
	user.Nama = nama
	user.Email = email
	user.Username = username
	user.Verify = false
	if tipe == "fb" {
		user.FbId = sosmedid
	} else if tipe == "gplus" {
		user.GplusId = sosmedid
	} else if tipe == "twitter" {
		user.TwitId = sosmedid
	}
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == password).
		Message("Password Not Match")
	user.Validation(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.SetUp(user.Nama, user.Email, user.Username, user.FbId, user.GplusId, user.TwitId))
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
	err = app.GORM.Create(&user)
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
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			if !user.Verify {
				c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
				c.RenderArgs["user"] = username
				return c.Redirect(routes.Persons.UnverifyAcc())

			} else {
				c.Flash.Success("Wellcome here, " + user.Username)
				c.RenderArgs["user"] = username
				return c.Redirect(routes.Persons.List(""))
			}

		}
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login Failed")
	return c.Redirect(routes.App.Login())
}

func (c App) SetUp(nama, email, username, fb, gplus, twit string) revel.Result {
	user := c.connected()
	if user != nil {
		return c.Redirect(routes.Persons.List(""))
	}
	var users models.User
	users.Nama = nama
	users.Username = username
	users.Email = email
	if fb != "" {
		users.FbId = fb
		return c.Render(users)
	} else if gplus != "" {
		users.GplusId = gplus
		return c.Render(users)
	} else if twit != "" {
		users.TwitId = twit
		return c.Render(users)
	}
	return c.Render()

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
	user, _ := jason.NewObjectFromBytes([]byte(str))
	id, _ := user.GetString("id")

	var userfb models.User
	err := app.GORM.Where("fbid = ?", id).Find(&userfb)
	if !err.RecordNotFound() {
		c.Session["user"] = userfb.Username
		c.RenderArgs["user"] = userfb.Username
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
	return c.Redirect(routes.App.SetUp(nama, email, username, id, "", ""))

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
	var user models.User
	err := app.GORM.Where("gplusid = ?", id).Find(&user)
	if !err.RecordNotFound() {
		c.Session["user"] = user.Username
		c.RenderArgs["user"] = user.Username
		return c.Redirect(routes.Persons.List(""))
	}

	var username string
	if people.Nickname == "" {
		username = people.DisplayName
	} else {
		username = people.Nickname
	}

	for a := 0; a < len(people.Emails); a++ {
		return c.Redirect(routes.App.SetUp(nama, people.Emails[0].Value, username, "", id, ""))
	}

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
		c.RenderArgs["user"] = user.Username
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
