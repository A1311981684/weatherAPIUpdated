package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherImgController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherImgController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherImgController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherImgController"],
        beego.ControllerComments{
            Method: "GetOneIMG",
            Router: `/:id/:size`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherInfoController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherInfoController"],
        beego.ControllerComments{
            Method: "GetAllInfo",
            Router: `/:city`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherInfoController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherInfoController"],
        beego.ControllerComments{
            Method: "GetAllForecast",
            Router: `/weather_info/forecast/:city`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherInfoController"] = append(beego.GlobalControllerRouter["github.com/A1311981684/weatherAPI/controllers:WeatherInfoController"],
        beego.ControllerComments{
            Method: "GetToday",
            Router: `/weather_info/today/:city`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
