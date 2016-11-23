package controllers

import (
	"bytes"
	"github.com/ganggas95/trawanganserver/app"
	"github.com/ganggas95/trawanganserver/app/models"
	"github.com/ganggas95/trawanganserver/app/routes"
	"github.com/revel/revel"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
)

type Agent struct {
	App
}

func (a Agent) CheckAgent() *models.AgentTravel {
	user := a.connected()
	if user == nil {
		return nil
	}
	var agent models.AgentTravel
	err := app.GORM.Where("user_id = ?", user.UID).Find(&agent)
	if err.RecordNotFound() {
		return nil
	}
	return &agent

}

func (a Agent) Index() revel.Result {
	agent := a.CheckAgent()
	if agent != nil {
		user := a.GetUserWitId(agent.UserId)
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var foto models.UserFoto
			db := app.GORM.Where("user_id = ?", user.UID).Find(&foto)
			if db.RecordNotFound() {

			}
			dir := "/public/data/" + user.Username + "/"
			fotodir := dir + foto.Foto
			dahsboard := "dashboard"
			a.RenderArgs["agent"] = agent
			a.RenderArgs["dashboard"] = "dashboard"
			a.RenderArgs["fotodir"] = fotodir

			return a.Render(agent, dahsboard, fotodir)
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) ServiceAgent() revel.Result {
	agent := a.CheckAgent()
	user := a.GetUserWitId(agent.UserId)
	if agent != nil {
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var services []models.AgentService
			db := app.GORM.Where("agent_id = ?", agent.IdAgent).Find(&services)
			if db.RecordNotFound() {
				a.FlashParams()
				a.Flash.Error("No Service")
				return a.Redirect(routes.Agent.ServiceAgent())
			}
			var fotoUser models.UserFoto
			db = app.GORM.Where("user_id = ?", user.UID).Find(&fotoUser)
			if db.RecordNotFound() {

			}
			dir := "/public/data/"
			fotodir := dir + fotoUser.Foto
			service := "service"
			for i := 0; i < len(services); i++ {
				var foto models.FotoService
				db1 := app.GORM.Find(&foto, models.FotoService{IdService: services[i].IdService})
				if db1.RecordNotFound() {

				}
				services[i].Foto = foto
				a.RenderArgs["services"] = services
				a.RenderArgs["agent"] = agent
				a.RenderArgs["service"] = service

			}
			a.RenderArgs["fotodir"] = fotodir
			return a.Render(agent, services, fotodir, service)

		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) OrderAgent() revel.Result {
	agent := a.CheckAgent()
	user := a.GetUserWitId(agent.UserId)
	if agent != nil {
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var foto models.UserFoto
			db := app.GORM.Where("user_id = ?", user.UID).Find(&foto)
			if db.RecordNotFound() {

			}
			order := "order"
			dir := "/public/data/"
			fotodir := dir + foto.Foto
			a.RenderArgs["agent"] = agent
			a.RenderArgs["fotodir"] = fotodir
			a.RenderArgs["order"] = order
			return a.Render(agent, order, fotodir)
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) ChatAgent() revel.Result {
	agent := a.CheckAgent()
	user := a.GetUserWitId(agent.UserId)
	if agent != nil {
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var foto models.UserFoto
			db := app.GORM.Where("user_id = ?", user.UID).Find(&foto)
			if db.RecordNotFound() {

			}
			chat := "chat"
			dir := "/public/data/"
			fotodir := dir + foto.Foto
			a.RenderArgs["agent"] = agent
			a.RenderArgs["fotodir"] = fotodir
			a.RenderArgs["chat"] = chat
			return a.Render(agent, chat, fotodir)
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}
func (a Agent) MemberAgent() revel.Result {
	agent := a.CheckAgent()
	user := a.GetUserWitId(agent.UserId)
	if agent != nil {
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var foto models.UserFoto
			db := app.GORM.Where("user_id = ?", user.UID).Find(&foto)
			if db.RecordNotFound() {

			}
			member := "member"
			dir := "/public/data/"
			fotodir := dir + foto.Foto
			a.RenderArgs["agent"] = agent
			a.RenderArgs["fotodir"] = fotodir
			a.RenderArgs["member"] = member
			return a.Render(agent, member, fotodir)
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) RegisterAgent() revel.Result {
	return a.Render()
}

func (a Agent) UniqueHandler(email, website string) bool {
	var travelAgent models.AgentTravel
	db := app.GORM.Debug().Where("website_agent = ? OR email_agent = ?", website, email).Find(&travelAgent)
	if db.RecordNotFound() {
		return false
	} else {
		return true
	}

}

func (a Agent) AddAgentFromUser(travelAgent models.AgentTravel) revel.Result {
	users := a.connected()
	agent := a.UniqueHandler(travelAgent.Email, travelAgent.Website)
	if agent {
		a.Validation.Keep()
		a.FlashParams()
		a.Flash.Error("Your email and website url have been registered!")
		return a.Redirect(routes.Agent.RegisterAgent())
	}
	travelAgent.Validation(a.Validation)
	if a.Validation.HasErrors() {
		a.Validation.Keep()
		a.FlashParams()
		return a.Redirect(routes.Agent.RegisterAgent())
	}

	travelAgent.UserId = users.UID
	err := app.GORM.Create(&travelAgent)
	if err.Error != nil {
		panic(err.Error)
	}
	return a.Redirect(routes.Agent.Index())
}

func (a Agent) GetAgent(userId int64) *models.AgentTravel {
	var agent models.AgentTravel
	err := app.GORM.Where("user_id = ?", userId).Find(&agent)
	if err.Error != nil {
		panic(err.Error)
	}

	return &agent
}

func (a Agent) SetService(agentService models.AgentService, foto []byte) revel.Result {
	a.Validation.Required(foto).Message("foto is required")
	a.Validation.Required(agentService.Service).Message("Service is required")
	a.Validation.Required(agentService.Kategori).Message("Kategori is required")
	a.Validation.Required(agentService.Price).Message("Price is required")

	a.Validation.MinSize(foto, 2*KB).
		Message("Minimum a file size of 2KB expected")
	a.Validation.MaxSize(foto, 2*MB).
		Message("File cannot be larger than 2MB")

	// Check format of the file.
	conf, format, err1 := image.DecodeConfig(bytes.NewReader(foto))
	a.Validation.Required(err1 == nil).Key("foto").
		Message("Incorrect file format")
	a.Validation.Required(format == "jpeg" || format == "png").Key("foto").
		Message("JPEG or PNG file format is expected")

	// Check resolution.
	a.Validation.Required(conf.Height >= 150 && conf.Width >= 150).Key("foto").
		Message("Minimum allowed resolution is 150x150px")
	if a.Validation.HasErrors() {
		a.Validation.Keep()
		a.FlashParams()
		return a.Redirect(routes.Agent.ServiceAgent() + "#modalAddService")
	}
	agent := a.CheckAgent()
	if agent != nil {
		if user := a.GetUserWitId(agent.UserId); !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var img models.FotoService
			m := a.Request.MultipartForm
			for fname, _ := range m.File {
				fheader := m.File[fname]
				for i, _ := range fheader {
					file, err := fheader[i].Open()
					defer file.Close()
					if err != nil {
						panic(err)
					}
					root_path := revel.BasePath
					dst_path := "public/data/service"
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

						img.Foto = fheader[i].Filename
						img.Height = conf.Height
						img.Width = conf.Width
						img.Format = format
						img.Size = len(foto)
						img.Dir = "/" + dst_path + "/"
					}
				}

			}

			agentService.IdAgent = agent.IdAgent
			agentService.Kategori = agentService.Kategori
			agentService.Foto = img
			db := app.GORM.Create(&agentService)

			if db.Error != nil {
				panic(db.Error)
			}
			return a.Redirect(routes.Agent.ServiceAgent())
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) DeleteService(idService int64) revel.Result {
	agent := a.CheckAgent()
	if agent == nil {
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	serviceAgent := models.AgentService{IdService: idService}
	var serviceFoto models.FotoService
	db := app.GORM.Where("service_id = ?", idService).Find(&serviceFoto)
	db = app.GORM.Debug().Delete(&serviceFoto)
	db = app.GORM.Debug().Delete(&serviceAgent)
	if db.Error != nil {
		panic(db.Error)
	}
	a.Flash.Success("Service success deleted!!")
	return a.Redirect(routes.Agent.ServiceAgent())
}

func (a Agent) ActiveService(idService int64) revel.Result {
	agent := a.CheckAgent()
	if agent == nil {
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	var serviceAgent models.AgentService
	db := app.GORM.First(&serviceAgent, idService)
	if db.RecordNotFound() {
		panic(db.RecordNotFound())
	}
	serviceAgent.Status = true
	db = app.GORM.Save(&serviceAgent)
	if db.Error != nil {
		a.Flash.Error("Service failed to active")
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	a.Flash.Success("Service activated")
	return a.Redirect(routes.Agent.ServiceAgent())
}
func (a Agent) DisableService(idService int64) revel.Result {
	agent := a.CheckAgent()
	if agent == nil {
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	var serviceAgent models.AgentService
	db := app.GORM.First(&serviceAgent, idService)
	if db.RecordNotFound() {
		panic(db.RecordNotFound())
	}
	serviceAgent.Status = false
	db = app.GORM.Save(&serviceAgent)
	if db.Error != nil {
		a.Flash.Error("Service failed to active")
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	a.Flash.Success("Service activated")
	return a.Redirect(routes.Agent.ServiceAgent())
}

func (a Agent) Logout() revel.Result {
	if user := a.connected(); user == nil {
		a.Flash.Error("Please log in first")
		return a.Redirect(routes.App.Login())
	}

	for k := range a.Session {
		delete(a.Session, k)
	}
	return a.Redirect(routes.App.Login())
}
