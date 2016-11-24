// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/ganggas95/trawanganserver/app"
	controllers "github.com/ganggas95/trawanganserver/app/controllers"
	models "github.com/ganggas95/trawanganserver/app/models"
	tests "github.com/ganggas95/trawanganserver/tests"
	controllers1 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	time "time"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "SetUser",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					65: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					72: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					79: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "AddUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddUserWithSosmed",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "verifyPassword", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AuthApp",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "remember", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetUp",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					194: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "RegisterWithFacebook",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AuthFb",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "code", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "RegisterWithGPlus",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GplusAuth",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "code", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "VerifyAcoount",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "code", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "RegisterWithTwit",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AuthTwit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "oauth_toke", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					339: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UploadFoto",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					358: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Upload",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "avatar", Type: reflect.TypeOf((*[]byte)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					72: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Agent)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					56: []string{ 
						"agent",
						"dahsboard",
						"fotodir",
					},
				},
			},
			&revel.MethodType{
				Name: "ServiceAgent",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					103: []string{ 
						"agent",
						"services",
						"fotodir",
						"service",
					},
				},
			},
			&revel.MethodType{
				Name: "OrderAgent",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					134: []string{ 
						"agent",
						"order",
						"fotodir",
					},
				},
			},
			&revel.MethodType{
				Name: "ChatAgent",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					164: []string{ 
						"agent",
						"chat",
						"fotodir",
					},
				},
			},
			&revel.MethodType{
				Name: "MemberAgent",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					193: []string{ 
						"agent",
						"member",
						"fotodir",
					},
				},
			},
			&revel.MethodType{
				Name: "RegisterAgent",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					208: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "AddAgentFromUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "travelAgent", Type: reflect.TypeOf((*models.AgentTravel)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetService",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "agentService", Type: reflect.TypeOf((*models.AgentService)(nil)) },
					&revel.MethodArg{Name: "foto", Type: reflect.TypeOf((*[]byte)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteService",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "idService", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ActiveService",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "idService", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DisableService",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "idService", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Api)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Auth",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*models.User)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Persons)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					30: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "search", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					58: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Show",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					71: []string{ 
						"persons",
					},
				},
			},
			&revel.MethodType{
				Name: "UnverifyAcc",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					83: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Tambah",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					95: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "AddData",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "person", Type: reflect.TypeOf((**models.Person)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Delete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Ubah",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "nama", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "alamat", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "tempatlahir", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pekerjaan", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "tanggallahir", Type: reflect.TypeOf((*time.Time)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetData",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int64)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					175: []string{ 
						"person",
					},
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/ganggas95/trawanganserver/app/controllers.Agent.SetService": { 
			257: "foto",
			258: "agentService.Service",
			259: "agentService.Kategori",
			260: "agentService.Price",
			262: "foto",
			264: "foto",
			269: "err1",
		},
		"github.com/ganggas95/trawanganserver/app/controllers.Api.AddUser": { 
			44: "password",
		},
		"github.com/ganggas95/trawanganserver/app/controllers.App.AddUser": { 
			108: "password",
		},
		"github.com/ganggas95/trawanganserver/app/controllers.App.AddUserWithSosmed": { 
			135: "verifyPassword",
			136: "verifyPassword",
		},
		"github.com/ganggas95/trawanganserver/app/controllers.App.Upload": { 
			366: "avatar",
			367: "avatar",
			369: "avatar",
			374: "err1",
		},
		"github.com/ganggas95/trawanganserver/app/controllers.Persons.AddData": { 
			99: "person.Nama",
			100: "person.Alamat",
			101: "person.TempatLahir",
			102: "person.TanggalLahir",
			103: "person.Pekerjaan",
		},
		"github.com/ganggas95/trawanganserver/app/models.(*Person).Validation": { 
			26: "person.Nama",
			31: "person.Alamat",
			36: "person.TempatLahir",
		},
		"github.com/ganggas95/trawanganserver/app/models.(*User).Validation": { 
			60: "user.Nama",
			67: "user.Username",
			72: "user.Email",
		},
		"github.com/ganggas95/trawanganserver/app/models.AgentTravel.Validation": { 
			58: "at.NamaAgent",
			60: "at.Alamat",
			65: "at.Notelp",
			71: "at.Email",
			75: "at.Website",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
