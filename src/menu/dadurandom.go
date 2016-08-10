package menu
import(
	"fmt"
	"net/http"
	"math/rand"
	"time"
)
func Menu5Handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, 
	`
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
					    padding: 20px;
					    box-sizing: border-box;
					    background: #FFF;
					    border-radius: 5px;
					    color: #2c3e50;
					    text-align: center;
					}

					.content2 {
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
					    float: left;
					    padding: 15px;
					}

					.btn-submit:hover {
					    background-color: #CA665C;
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

				        <div class="content2">
				        <form action='/dadu' method='get'>
				        <h2>Acak Dadu Untuk Permainan Monopoli dengan tambahan Ganjil dan Genap</h2>
				            <div class="">
				                <p>Dadu 1 :</p>

				                <p>Dadu 2 :</p>
				               
				                <p>Status : - </p>
				                <p><input type="radio" name="option" value="normal" checked><label>Normal</label>
				                <input type="radio" name="option" value="ganjil"><label>Ganjil</label>
				                <input type="radio" name="option" value="genap"><label>Genap</label>
				                <br></p>
				               
				                <input type="submit" class="btn btn-submit" value="Roll">
				        </form>
				                 <a href='/'><input  class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
				            </div>
				        </div>

				    </div>
				</body>
				</html>
	`)
}
func DaduRandom(w http.ResponseWriter, r *http.Request){
	var acak1,acak2,stat,hasil1,hasil2,total int
	radio :=  r.FormValue("option")
	
	//untuk fungsi random
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	
	if radio=="genap"{
		var dadu1  = [6]int{1,2,3,4,5,6}
		var dadu2  = [3]int{1,3,5}
		var dadu3  = [3]int{2,4,6}
		acak1 = r1.Intn(6)
		acak2 = r1.Intn(3)
		hasil1=dadu1[acak1]
		if hasil1==1 || hasil1==3 || hasil1==5{
		hasil2=dadu2[acak2]
		}else{
		hasil2=dadu3[acak2]
		}
		
		
	}else if radio=="ganjil"{
		var dadu1  = [3]int{1,3,5}
		var dadu2  = [3]int{2,4,6}
		acak1 = r1.Intn(3)
		acak2 = r1.Intn(3)
		hasil1=dadu1[acak1]
		hasil2=dadu2[acak2]
		
	}else{
		var dadu1  = [6]int{1,2,3,4,5,6}
		var dadu2  = [6]int{1,2,3,4,5,6}
		acak1 = r1.Intn(6)
		acak2 = r1.Intn(6)
		hasil1=dadu1[acak1]
		hasil2=dadu2[acak2]
		
	}
	if hasil1==hasil2{stat=1} 
	total = hasil1+hasil2
	if stat==1{		
		fmt.Fprintf(w, 
		`

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
					    padding: 20px;
					    box-sizing: border-box;
					    background: #FFF;
					    border-radius: 5px;
					    color: #2c3e50;
					    text-align: center;
					}

					.content2 {
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
					    float: left;
					    padding: 15px;
					}

					.btn-submit:hover {
					    background-color: #CA665C;
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
					.btn-quit2{
					    background-color: #2c3e50;
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

				        <div class="content2">
				        <form action='/dadu' method='get'>
				        <h2>Acak Dadu Untuk Permainan Monopoli dengan tambahan Ganjil dan Genap</h2>
				            <div class="">
				                <p>Dadu 1 : %d</p>

				                <p>Dadu 2 : %d </p>
				               
				                <p>Status : <blink><b style="color:red;">Double</b></blink> <br> </p>
				                <p> Total : %d <br> </p>
				                <p><input type="radio" name="option" value="normal" checked><label>Normal</label>
				                <input type="radio" name="option" value="ganjil"><label>Ganjil</label>
				                <input type="radio" name="option" value="genap"><label>Genap</label>
				                <br></p>
				               
				                <input type="submit" class="btn btn-submit" value="ROLL LAGI">
				        </form>
				                 <a href='/'><input class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
				                 
				            </div>
				        </div>

				    </div>
				</body>
				</html>
		`,hasil1,hasil2,total)
	}else{
		fmt.Fprintf(w, 
		`

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
					    padding: 20px;
					    box-sizing: border-box;
					    background: #FFF;
					    border-radius: 5px;
					    color: #2c3e50;
					    text-align: center;
					}

					.content2 {
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
					    float: left;
					    padding: 15px;
					}

					.btn-submit:hover {
					    background-color: #CA665C;
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
					.btn-quit2{
					    background-color: #2c3e50;
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

				        <div class="content2">
				        <form action='/dadu' method='get'>
				        <h2>Acak Dadu Untuk Permainan Monopoli dengan tambahan Ganjil dan Genap</h2>
				            <div class="">
				                <p>Dadu 1 : %d</p>

				                <p>Dadu 2 : %d </p>
				               
				                <p>Status : - </p>
				                <p> Total : %d <br> </p>
				                <p><input type="radio" name="option" value="normal" checked><label>Normal</label>
				                <input type="radio" name="option" value="ganjil"><label>Ganjil</label>
				                <input type="radio" name="option" value="genap"><label>Genap</label>
				                <br></p>
				               
				                <input type="submit" class="btn btn-submit" value="ROLL LAGI">
				        </form>
				                 <a href='/'><input  class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
				                
				            </div>
				        </div>

				    </div>
				</body>
				</html>
		`,hasil1,hasil2,total)
	}
	
	
}