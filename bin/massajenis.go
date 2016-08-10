package menu

import(
	"fmt"
	"net/http"
	"strconv"
)

func Menu1Handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, 
	`<html><body>
	<h3>Menghitung Massa Jenis</h3>
	<form	action='/massajenis' method='get'>
	Nilai massa  =<input type="decimal" name="massa">g<br>
	Nilai volume =<input type="decimal" name="volume">cm<br>
	<input type="submit" value="Hitung">
	</form>
	<a href="/">Ke Halaman Menu</a><br>
	</body></html>
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
		<html><body>
		Masa Jenis dari benda tersebut adalah = %.2f g/cm3<br>
		<a href="/menu1">Kembali Menghitung Massa Jenis</a><br>
		</body></html>
		`,massa_jenis)
	}
	
}