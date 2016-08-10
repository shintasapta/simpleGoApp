package menu

import(
	"fmt"
	"net/http"
	"strconv"
	
)

func Menu2Handler(w http.ResponseWriter,r *http.Request){
	//http.Handle("/menu/", http.StripPrefix("/menu/", http.FileServer(http.Dir("menu"))))
	//<img src="/menu/grafiklinier.png" alt="Grafik Linier" style="width:450px;height:228px;">
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
			        <h2>Menghitung Fuzzyfikasi Grafik Linear</h2>
			        <img src="http://2.bp.blogspot.com/-vLz00B1EwdU/Vf5w4g-loBI/AAAAAAAAAM8/CwWCUrqwD6w/s1600/prt_4.png" alt="Grafik Linier Turun" style="width:300px;height:200px;">
			        <img src="http://2.bp.blogspot.com/-GrvI4dzEOhY/Tf604040o2I/AAAAAAAAACM/PZiSN_bbaxI/s1600/naik.png" alt="Grafik Linier Naik" style="width:300px;height:200px;">
			            <div class="">
			            	<form	action='/linier' method='get'>
			                	<p>Batas Bawah</p>
				                <p><input class="input" type="decimal" name="bb"></p>

				                <p>Batas Atas</p>
				                <p><input class="input" type="decimal" name="ba"></p>

				                <p>Nilai X</p>
				                <p><input class="input" type="decimal" name="x"></p>

				                <p>Grafik Liner :</p>
				                <input type="radio" name="grafik" value="naik" checked> Naik
				                <input type="radio" name="grafik" value="menurun"> Turun
				                <br>
				                <br>
				               
				                <input type="submit" class="btn btn-submit" value="Hitung"><br><br>
				            </form>

			                <p>Keterangan : Batas bawah adalah nilai a<br>
			                Batas atas adalah nilai b pada semua grafik seperti gambar</p>
			                  <a href="/"><input type="submit" class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
			            </div>
			        </div>

			    </div>
			</body>
			</html>

	`)
}
func FuzzyLinier(w http.ResponseWriter, r *http.Request){
	var hasil float64
	batasbawah :=  r.FormValue("bb")
	batasatas :=  r.FormValue("ba")
	nilaix :=  r.FormValue("x")
	grafik :=  r.FormValue("grafik") 
	bb, err := strconv.ParseFloat(batasbawah, 32)
	ba, err := strconv.ParseFloat(batasatas, 32)
	x, err := strconv.ParseFloat(nilaix, 32)
	
	if err == nil{
		if grafik=="naik"{
			if x<=bb{
				hasil=0.0
			}else if x>=ba{
				hasil=1.0
			}else{
				hasil=(x-bb)/(ba-bb)
			}
		}else if grafik=="menurun"{
			if x<=bb{
				hasil=1.0
			}else if x>=ba{
				hasil=0.0
			}else{
				hasil=(ba-x)/(ba-bb)
			}
		}else{
			fmt.Fprintf(w,"Salah Grafik")
		}
		fmt.Fprintf(w,
		`

				<html>
				<body>
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

					.content3 {
					    min-height: 400px;
					    padding-top: 50px;
					    padding-bottom: 50px;
					    padding-left: 150px;
					    padding-right: 150px;
					    box-sizing: border-box;
					    background: #FFF;
					    border-radius: 5px;
					    color: #2c3e50;
					    text-align:center;
					}



					body{
					    margin: 0;
					    background : #798E82;
					}

					.btn {
					    width: auto;
					    color: #fff;
					    border: none;
					    cursor: pointer;
					    padding: 10px;
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
			    <div class="container">
			    <h1>APLIKASI HITUNG SERBA GUNA</h1>
			    	<div class="content3">
				        <h3>Nilai X[%.2f] = %.2f <br><br></h3>
				        <a href="/"><input type="submit" class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
				        <a href="/menu2"><input type="submit" class="btn btn-quit2" name="" value="KEMBALI HITUNG FUZZYFIKASI LINIER"></a>
				     </div>
			    </div>
			</body>
			</html>
		`,x,hasil)
	}
		
		
}