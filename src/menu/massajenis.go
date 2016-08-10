package menu

import(
	"fmt"
	"net/http"
	"strconv"
)

func Menu1Handler(w http.ResponseWriter,r *http.Request){
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
			        <h2>Menghitung Massa Jenis</h2>
			            <div class="">
			            	<form	action='/massajenis' method='get'>
				                <p>Nilai Massa</p>
				                <p><input class="input" type="decimal" name="massa"></p>

				                <p>Nilai Volume</p>
				                <p><input class="input" type="decimal" name="volume"></p>
				                <input type="submit" class="btn btn-submit" value="Hitung">
				            </form>
			                <a href="/"><input type="submit" class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
			            </div>
			        </div>

			    </div>
			</body>
			</html>
	`)
}

func MassaJenis(w http.ResponseWriter, r *http.Request){
	var massa_jenis float64
	m :=  r.FormValue("massa")
	v :=  r.FormValue("volume")
	massa, err := strconv.ParseFloat(m, 32)
	volume, err := strconv.ParseFloat(v, 32)
	if err == nil{
		massa_jenis = massa/volume
		fmt.Fprintf(w,`


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
				        <h3>Masa Jenis dari benda tersebut adalah = %.2f g/cm3<br><br>
				        <a href="/"><input type="submit" class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
				        <a href="/menu1"><input type="submit" class="btn btn-quit2" name="" value="KEMBALI HITUNG MASSA JENIS"></a>
				     </div>
			    </div>
			</body>
			</html>
		`,massa_jenis)
	}
	
}