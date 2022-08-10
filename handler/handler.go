package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	/*message := "Hello World"
	w.Write([]byte(message))
	*/

	tmpl, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title": "Golang Web Pertamaku",
		"content": "GOLANG WEB DTS PROA 3",
		"name": "JOKO SUSANTO - 152308829101-227",
		"link": "/product",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	message := "About Page"
	w.Write([]byte(message))
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	message := "Profile Page"
	w.Write([]byte(message))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"title": "eCommerce",
		"ID": "1",
		"Product": "Microtik RB 450g x4",
		"Price": 2000000,
		"Stock":10,
	}

	tmpl, err := template.ParseFiles(path.Join("views","product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}

}