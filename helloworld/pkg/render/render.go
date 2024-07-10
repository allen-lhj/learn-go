package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template from cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// for finer grained error checking
	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	// render the tempalte
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	// parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.html")
	// err := parseTemplate.Execute(w, nil)
	// if err != nil {
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we alerday have the template in our catche
// 	_, inMap := tc[t]
// 	log.Println(inMap)
// 	if !inMap {
// 		log.Println("create the template")
// 		// need to create the template
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// have template in the cache
// 		log.Println("using cache template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	log.Println("create")
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.html",
// 	}
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }
