package main
 
import (
    "net/url"
    "net/http"
    "service"
)
 
type IndexController struct {
}
 
func (IndexController *IndexController) Index(request map[string]url.Values, headers http.Header) (statusCode int, response interface{}) {
    return 200, map[string]string{"first": "webApp"}
}
 
type BarController struct {
}

 
func main() {
    var serv = service.Prepare()
    serv.Get("index", "IndexController@Index")
    controllers := []interface{}{&IndexController{}}
    serv.RegisterControllers(controllers)
    serv.RegisterController(&BarController{})
    serv.Serve(3000)
}

