package main
import (
   "net/http"
   "fmt"
   "log"
   "encoding/json"
   "io/ioutil"
   
)

//modulo kernel de memoria
type MK_mem struct{
	Mtotal int `json:"MemoriaTotal"`
    Mfree int `json:"MemoriaLibre"`
    Percent int `json:"PorcentajeConsumo"`
}


type procs struct{
	PROC_ID  int `json:"PROC_ID"`
    PROC_NAME  string `json:"PROC_NAME"`
    USER_ID  int `json:"USER_ID"`
    PARENT_ID    int `json:"PARENT_ID"`
    ESTATE  string `json:"ESTATE"`
    MEMORY_USED  int `json:"MEMORY_USED"`
}


type MK_proc struct{
    resumen struct{
        TOTAL  int `json:"TOTAL"`
        RUN  int `json:"RUN"`
        SLEEP  int `json:"SLEEP"`
        STOPPED  int `json:"STOPPED"`
        ZOMBIE  int `json:"ZOMBIE"`
    }
}
//historial de modulo kernel de
type H_mem [] MK_mem
type H_proc [] MK_proc

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func getmem() MK_mem{
    fdat, err := ioutil.ReadFile("./mem_grupo3")
    check(err)
    fmt.Println(string(fdat))
    var retdata MK_mem
    json.Unmarshal(fdat,&retdata)
    
    //var retdata = data {Mtotal: 342, Mfree: 100, Percent: 391}
    return retdata
}
func getproc() MK_proc{
    fdat, err := ioutil.ReadFile("./proc_grupo3")
    check(err)
    fmt.Println(string(fdat))
    var retdata MK_proc
    json.Unmarshal([]byte(fdat),&retdata)
    
    //var retdata = data {Mtotal: 342, Mfree: 100, Percent: 391}
    return retdata
}

func getmem_set() []MK_mem{
    var historic H_mem
    var dda MK_mem
    n:=0
    points :=5 //numbero de recolecciones de datos para generar un historico
    for n < points{
        dda =getmem()
        historic = append(historic, dda)
        n +=1
    }
    return historic
}

func getproc_set() []MK_proc{
    var historic H_proc
    var dda MK_proc
    n:=0
    points :=5 //numbero de recolecciones de datos para generar un historico
    for n < points{
        dda =getproc()
        historic = append(historic, dda)
        n +=1
    }
    return historic
}

func mem_set(w http.ResponseWriter, r *http.Request){
    H_mem := getmem_set()
    fmt.Println("Endpoint hit: mem history readings")
    json.NewEncoder(w).Encode(H_mem)
}

func proc_set(w http.ResponseWriter, r *http.Request){
    H_proc := getproc_set()
    fmt.Println("Endpoint hit: proc history readings")
    json.NewEncoder(w).Encode(H_proc)
}

func mem(w http.ResponseWriter, r *http.Request){
    H_mem := getmem()
    fmt.Println("Endpoint hit: mem readings")
    json.NewEncoder(w).Encode(H_mem)
}

func proc(w http.ResponseWriter, r *http.Request){
    H_proc := getproc()
    fmt.Println("Endpoint hit: proc readings")
    json.NewEncoder(w).Encode(H_proc)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/mem", mem)
    http.HandleFunc("/proc", proc)
    http.HandleFunc("/mem_set", mem_set)
    http.HandleFunc("/proc_set", proc_set)
    log.Fatal(http.ListenAndServe(":8081", nil))
}


func main() {
    handleRequests()
}

