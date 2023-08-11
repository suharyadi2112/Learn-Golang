package main

import (
  "fmt"
)

// type name interface {
//
// }
//
// type name struct {
//
// }
var (
  // UserDraft page 1 completed (need to fill identity form next)
	UserDraft = "new_draft"
	// UserIdentityDraft page 2 completed (need to fill financial form next)
	UserIdentityDraft = "identity_draft"
	// UserFinancialDraft page 3 completed (need to fill risk profile next)
	UserFinancialDraft = "financial_draft"
	// UserAdminVerification after write data to OAO2 success need verification by admin
)
func main() {

  var interAny interface{}//interface biasa
  interAny = "Suharyadi"
  cekvar := []string{
    UserDraft,
    UserIdentityDraft,
    UserFinancialDraft,
  }

  sub := `users.status IN (`
	for i, r := range cekvar {
		if i == len(cekvar)-1 {
			sub += `'`+ r +`'`
		} else {
			sub += `'`+ r +`'`+`, `
		}
	}
	sub += `)`

	query := 	`SELECT DISTINCT
					users.phone,
					users.name,
					users.updated_at,
					users.status,
					NOW( ) AS Sekarang,
					EXTRACT ( epoch FROM ( NOW( ) - users.updated_at ) ) AS KeDetik,
				CASE
						WHEN EXTRACT ( epoch FROM ( NOW( ) - users.updated_at ) ) > 3600 THEN
						'> 1 hour' ELSE'<1 Hours'
					END AS HasilJam
				FROM
					"users"
				WHERE ( SELECT EXTRACT ( epoch FROM ( NOW( ) - users.updated_at ) ) > 3600 ) AND ` + sub


  fmt.Println(query)

  fmt.Printf("var1 = %T\n", cekvar)

  fmt.Println(interAny)

  interAny = []string{"tes1","tes2"}
  fmt.Println(interAny)

//disiapkan variabel data dengan tipe map[string]interface{}, yaitu sebuah koleksi dengan key bertipe string dan nilai bertipe interface kosong interface{}.
  var data map[string]interface{}
  data = map[string]interface{}{
    "nama" : "suharyadi",
    "no_hp" : 98765678765,//invalid digit '9' in octal literal jika diawali dengan 0
    "k" : interAny,
  }
  fmt.Println(data)

//interface kosong bisa dituliskan dengan menggunakan any
  var wp map[string]any
  wp = map[string]any{
    "nama" : 12,
    "j" : interAny,
  }
  fmt.Println(wp)

}
