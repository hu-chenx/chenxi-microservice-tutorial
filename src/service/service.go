package service

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "reflect"
    "strings"
)

const (
    GETMethod     = "GET"
    POSTMethod    = "POST"
    PUTMethod     = "PUT"
    DELETEMethod  = "DELETE"
)

const (
    HTTPMethodNotAllowed = 405
)

type APIService struct {
    controllerRegistry map[string]interface{}
    registeredPathAndController map[string]map[string]map[string]string
    requestForm                 map[string]url.Values
}

func (api *APIService) Get(path, controllerWithActionString string) {
    mapping := api.mappingRequestMethodWithControllerAndActions(GETMethod, path, controllerWithActionString)
    api.registeredPathAndController[path] = mapping
}

func (api *APIService) Post(path, controllerWithActionString string) {
    mapping := api.mappingRequestMethodWithControllerAndActions(POSTMethod, path, controllerWithActionString)
    api.registeredPathAndController[path] = mapping
}

func (api *APIService) Put(path, controllerWithActionString string) {
    mapping := api.mappingRequestMethodWithControllerAndActions(PUTMethod, path, controllerWithActionString)
    api.registeredPathAndController[path] = mapping
}

func (api *APIService) Delete(path, controllerWithActionString string) {
    mapping := api.mappingRequestMethodWithControllerAndActions(DELETEMethod, path, controllerWithActionString)
    api.registeredPathAndController[path] = mapping
}

func (api *APIService) mappingRequestMethodWithControllerAndActions(requestMethod, path, controllerWithActionString string) map[string]map[string]string {
    mappingResult := make(map[string]map[string]string)
    if length := len(api.registeredPathAndController[path]); length > 0 {
        mappingResult = api.registeredPathAndController[path]
    }
    controllerAndActionSlice := strings.Split(controllerWithActionString, "@")
    controller := controllerAndActionSlice[0]
    action := controllerAndActionSlice[1]
    controllerAndActionMap := map[string]string{controller: action}
    mappingResult[requestMethod] = controllerAndActionMap
    return mappingResult
}

func (api *APIService) HandleRequest(controllers map[string]map[string]string) http.HandlerFunc {
    return func(rw http.ResponseWriter, request *http.Request) {
        request.ParseForm()
        method := request.Method
        api.requestForm["query"] = request.Form
        api.requestForm["form"] = request.PostForm
        macthedControllers, ok := controllers[method]
        if !ok {
            rw.WriteHeader(HTTPMethodNotAllowed)
        }
        for k, v := range macthedControllers {
            controllerKey := "*" + k
            controller := api.controllerRegistry[controllerKey]
            in := make([]reflect.Value, 2)
            in[0] = reflect.ValueOf(api.requestForm)
            in[1] = reflect.ValueOf(request.Header)
            returnValues := reflect.ValueOf(controller).MethodByName(v).Call(in)
            statusCode := returnValues[0].Interface()
            intStatusCode := statusCode.(int)
            response := returnValues[1].Interface()
            responseHeaders := http.Header{}
            if len(returnValues) == 3 {
                responseHeaders = returnValues[2].Interface().(http.Header)
            }
            api.JSONResponse(rw, intStatusCode, response, responseHeaders)
        }
    }
}


func (api *APIService) RegisterHandleFunc() {
    for k, v := range api.registeredPathAndController {
        path := k
        if !strings.HasPrefix(k, "/") {
            path = fmt.Sprintf("/%v", k)
        }
        http.HandleFunc(path, api.HandleRequest(v))
    }
}

func (api *APIService) RegisterControllers(controllers []interface{}) {
    for _, v := range controllers {
        api.RegisterController(v)
    }
}


func (api *APIService) RegisterController(controller interface{}) {
    controllerType := getType(controller)
    api.controllerRegistry[controllerType] = controller
}


func getType(value interface{}) string {
    if t := reflect.TypeOf(value); t.Kind() == reflect.Ptr {
        return "*" + t.Elem().Name()
    } else {
        return t.Name()
    }
}


func (api *APIService) Serve(port int) {
    api.RegisterHandleFunc()
    fullPort := fmt.Sprintf(":%d", port)
    http.ListenAndServe(fullPort, nil)
}


func (api *APIService) JSONResponse(rw http.ResponseWriter, statusCode int, response interface{}, headers http.Header) {
    for k, v := range headers {
        for _, header := range v {
            rw.Header().Add(k, header)
        }
    }
    rw.Header().Add("Content-Type", "application/json")
    rw.WriteHeader(statusCode)
    rsp, err := json.Marshal(response)
    if err != nil {
        fmt.Println("error:", err)
    }
    rw.Write(rsp)
}


func Prepare() *APIService {
    var apiService = new(APIService)
    apiService.controllerRegistry = make(map[string]interface{})
    apiService.registeredPathAndController = make(map[string]map[string]map[string]string)
    apiService.requestForm = make(map[string]url.Values)
    return apiService
}
