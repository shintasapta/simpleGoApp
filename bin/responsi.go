package main
import(
	"fmt"
	"net/http"
	//"strconv"
	//"math/rand"
	//"time"
	"menu"
)

func RootHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, 
	`
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Aplikasi Serba Guna</title>
    <style>
    *{
	font-family: sans-serif;
	font-weight: 100;
}
.container {
    width: 960px;
	max-width: 100%;
    margin: 0 auto;
    background:#2980b9;
    border-radius: 5px;
    padding: 10px;
    box-shadow: 0px 0px 5px black;
    color: #FFF;
}
h1{
    text-align: center;
}
a {
    text-decoration: none;
    color: #FFF;
}

.content {
    width: 100%;
    padding: 20px;
    box-sizing: border-box;
    background: #FFF;
    border-radius: 5px;
    color: #2c3e50;
    text-align: center;
}

.content2 {
    width: 100%;
    min-height: 400px;
    padding-top: 50px;
    padding-bottom: 50px;
    padding-left: 150px;
    padding-right: 150px;
    box-sizing: border-box;
    background: #FFF;
    border-radius: 5px;
    color: #2c3e50;
}



body{
    margin: 0;
    background : #798E82;
}

.input, textarea, select {
    width: 100%;
    padding: 10px 5px;
    box-sizing: border-box;
    font-size: large;
}



textarea {
    min-height: 80px;
}

.btn {
    width: auto;
    color: #fff;
    border: none;
    cursor: pointer;
    padding: 10px;
}

.btn-submit {
    background-color: #C35045;
    float: right;
    padding: 15px;
}

.btn-submit:hover {
    background-color: #CA665C;
}

.btn-blue{
    padding: 40px;
    margin: 20px;
    font-family: verdana;
    font-size:20px;
    background-color:#5DA7D8;
}
.btn-red{
   padding: 40px;
    margin: 20px;
    font-family: verdana;
    font-size:20px;
    background-color: #e74c3c;
}
.btn-yellow{
   padding: 40px;
    margin: 20px;
    font-family: verdana;
    font-size:20px;
    background-color: #f1c40f;
}

.btn-green{
   padding: 40px;
    margin: 20px;
    font-family: verdana;
    font-size:20px;
    background-color: #2ecc71;
}
.btn-purple{
   padding: 40px;
    margin: 20px;
    font-family: verdana;
    font-size:20px;
    background-color: #9b59b6;
}

.btn-quit {
    background-color: #f1c40f;
    float:center;
    margin-top:80px;
    padding-top: 20px;
    padding-bottom: 20px;
    padding-right: 30px;
    padding-left: 30px;

}

    </style>
</head>
<body>
    <div class="container clearfix">
        <h1>APLIKASI HITUNG SERBA GUNA</h1>

        <div class="content">
            <div class="">
                <a href="/menu1"><button class="btn btn-blue">MASSA JENIS</button></a>
                <a href="/menu2"><button class="btn btn-red">FUZZYFIKASI GRAFIK LINIER</button></a>
                <a href="/menu3"><button class="btn btn-yellow">FUZZYFIKASI GRAFIK SEGITIGA</button></a>
                <a href="/menu4"><button class="btn btn-green">INDEX MASSA TUBUH</button></a>
                <a href="/menu5"><button class="btn btn-purple">ACAK DADU</button></a>
            </div>

        <p style="text-align:left;">Responsi Workshop dengan Go Language Genap 2015/2016 :<br><br>
         1. (135410299)    Faishal Abrari<br>
         2. (135410307)    Sefty Nindyastuti<br>
         3. (135410310)    Fitri Atun<br>
         4. (135410296)    Shinta Saptarini<br>
         5. (125410010)    Arif Nur Rohman<br></p>
        </div>

    </div>
</body>
</html>

	`)
}

func main(){
	//http.Handle("/menu3", http.FileServer(http.Dir("/menu")))
	//http.Handle("/menu3/", http.StripPrefix("/menu3/", http.FileServer(http.Dir("menu"))))
	http.HandleFunc("/",RootHandler)
	http.HandleFunc("/menu1",menu.Menu1Handler)
	http.HandleFunc("/massajenis",menu.MassaJenis)
	http.HandleFunc("/menu2",menu.Menu2Handler)
	http.HandleFunc("/linier",menu.FuzzyLinier)
	http.HandleFunc("/menu3",menu.Menu3Handler)
	http.HandleFunc("/segitiga",menu.FuzzySegitiga)
	http.HandleFunc("/menu4",menu.Menu4Handler)
	http.HandleFunc("/massatubuh",menu.Indexmassatubuh)
	http.HandleFunc("/menu5",menu.Menu5Handler)
	http.HandleFunc("/dadu",menu.DaduRandom)
	http.ListenAndServe(":8000",nil)
}