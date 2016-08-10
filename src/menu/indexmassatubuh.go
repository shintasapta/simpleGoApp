package menu

import(
	"fmt"
	"net/http"
	"strconv"
)
func Menu4Handler(w http.ResponseWriter,r *http.Request){
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
			        <h2>Menghitung Index Massa Tubuh</h2>
			            <div class="">
			            	<form	action='/massatubuh' method='get'>
				                <p>Berat Badan</p>
				                <p><input class="input" type="text" name="bb"></p>

				                <p>Tinggi Badan</p>
				                <p><input class="input" type="text" name="tb"></p>
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
func Indexmassatubuh(w http.ResponseWriter, r *http.Request){
	var imt,temp float64
	var cetak string
	bb :=  r.FormValue("bb")
	tb :=  r.FormValue("tb")
	bbf, err := strconv.ParseFloat(bb, 32)
	tbf, err := strconv.ParseFloat(tb, 32)
	if err == nil{
		temp = tbf/100.0
		imt = bbf/(temp*temp)
		switch {
			case imt<=17.0:
				cetak ="Gizi Kurang, Sangat Kurus"
			case (imt>17.0) && (imt<=18.5):
				cetak =" Gizi Kurang, Kurus"
			case (imt>18.5) && (imt<=25.0):
				cetak =" Gizi Baik, Normal"
			case (imt>25.0) && (imt<=27.0):
				cetak =" Gizi Lebih, Gemuk"
			case imt>27.0:
				cetak =" Gizi Lebih, Sangat Gemuk"
				default:
				{
					cetak ="ada kesalahan input"
				}
			}
		//fmt.Fprintf(w,"%f",imt)
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
				        <h3>Hasil : %f %s </h3><br>
				        <a href="/"><input type="submit" class="btn btn-quit" name="" value="KEMBALI KE MENU UTAMA"></a>
				        <a href="/menu4"><input type="submit" class="btn btn-quit2" name="" value="KEMBALI HITUNG MASSA TUBUH"></a>
				     </div>
			    </div>
			</body>
			</html>`,imt,cetak)
		
	}
}