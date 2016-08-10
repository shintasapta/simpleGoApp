package menu

import(
	"fmt"
	"net/http"
	"strconv"
	
)

func Menu3Handler(w http.ResponseWriter,r *http.Request){
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
				        <h2>Menghitung Fuzzyfikasi Grafik Segitiga</h2>
				        <img src="http://1.bp.blogspot.com/-6ITo2V0l7o4/Tf62E8VdfgI/AAAAAAAAACU/D75wK3hLsfk/s1600/kurva.png" alt="Grafik Segitiga" style="width:400px;height:300px;>
				            <div class="">
				            	<form	action='/segitiga' method='get'>
				                <p>Batas Bawah</p>
				                <p><input class="input" type="text" name="bb"></p>

				                <p>Batas Atas</p>
				                <p><input class="input" type="text" name="ba"></p>

				                <p>Nilai Tengah</p>
				                <p><input class="input" type="text" name="tengah"></p>

				                 <p>Nilai X</p>
				                <p><input class="input" type="text" name="x"></p>
				               
				                <input type="submit" class="btn btn-submit" value="Hitung"><br><br>
				                </form>

				                <p>Keterangan : Batas bawah adalah nilai a<br>
				                Batas atas adalah nilai c pada semua grafik seperti gambar</p>
				                <a href="/"><input type="submit" class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
				            </div>
				        </div>

				    </div>
				</body>
				</html>
	`)
}
func FuzzySegitiga(w http.ResponseWriter, r *http.Request){
	var hasil float64
	var grafik string
	batasbawah :=  r.FormValue("bb")
	batasatas :=  r.FormValue("ba")
	t :=  r.FormValue("tengah")
	nilaix :=  r.FormValue("x")
	bb, err := strconv.ParseFloat(batasbawah, 32)
	ba, err := strconv.ParseFloat(batasatas, 32)
	tengah, err := strconv.ParseFloat(t, 32)
	x, err := strconv.ParseFloat(nilaix, 32)
	
	if x<tengah{
		grafik="naik"
	}else{
		grafik="menurun"
	}
	
	if err == nil{
		if x==tengah{
				hasil=1
		}else if grafik=="menurun"{
			if x==tengah{
				hasil=1.0
			}else if x>=ba{
				hasil=0.0
			}else{
				hasil=(ba-x)/(ba-tengah)
			}
		}else if grafik=="naik"{
			if x<=bb{
				hasil=0.0
			}else if x==tengah{
				hasil=1.0
			}else{
				hasil=(x-bb)/(tengah-bb)
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
				        <a href="/menu3"><input type="submit" class="btn btn-quit2" name="" value="KEMBALI HITUNG FUZZYFIKASI SEGITIGA"></a>
				     </div>
			    </div>
			</body>
			</html>

		`,x,hasil)
	}
		
		
}