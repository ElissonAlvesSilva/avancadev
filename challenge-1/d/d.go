package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Coupon struct {
	Code string
}

type RemovedCoupons struct {
	Coupon []Coupon
}

func (c RemovedCoupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return "REMOVED"
		}
	}
	return "OK"
}

type Result struct {
	Status string
}

var coupons RemovedCoupons

func main() {
	// Removed cupons
	coupon := Coupon{
		Code: "abc",
	}

	coupons.Coupon = append(coupons.Coupon, coupon)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	valid := coupons.Check(coupon)
	result := Result{Status: valid}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}
