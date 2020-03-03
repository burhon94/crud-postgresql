package app

import (
	"crud/pkg/crud/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func showBurgers(receiver *server, writer http.ResponseWriter, tpl *template.Template) (err error) {
	list, err := receiver.burgersSvc.BurgersList()
	if err != nil {
		log.Printf("can't execute Burgers list sevice: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}

	data := struct {
		Title   string
		Burgers []models.Burger
	}{
		Title:   "KFS, so good",
		Burgers: list,
	}

	err = tpl.Execute(writer, data)
	if err != nil {
		log.Printf("can't execute print burgers data: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}

	return nil
}

func (receiver *server) handleBurgersList() func(http.ResponseWriter, *http.Request) {
	tpl, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
	if err != nil {
		log.Printf("can't parse index page: %v", err)
	}

	return func(writer http.ResponseWriter, request *http.Request) {

		err := showBurgers(receiver, writer, tpl)
		if err != nil {
			log.Printf("Error while print burgers: %v", err)
		}

	}
}

func (receiver *server) handleBurgersSave() func(responseWriter http.ResponseWriter, request *http.Request) {
	tpl, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
	if err != nil {
		log.Printf("can't parse index page: %v", err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {

		name := request.FormValue("name")
		price := request.FormValue("price")

		parsedPrice, err := strconv.Atoi(price)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = receiver.burgersSvc.Save(models.Burger{
			Name:  name,
			Price: parsedPrice,
		})

		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// TODO: посмотреть, можно ли переделать на GET
		err = showBurgers(receiver, writer, tpl)
		if err != nil {
			log.Printf("Error while print burgers: %v", err)
		}

		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
		return
	}
}

func (receiver *server) handleBurgersRemove() func(responseWriter http.ResponseWriter, request *http.Request) {
	tpl, err := template.ParseFiles(filepath.Join(receiver.templatesPath, "index.gohtml"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {

		idBurger := request.FormValue("id")

		idNumBurger, err := strconv.Atoi(idBurger)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = receiver.burgersSvc.RemoveById(int64(idNumBurger))

		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = showBurgers(receiver, writer, tpl)
		if err != nil {
			log.Printf("Error while print burgers: %v", err)
		}

		// TODO: посмотреть, можно ли переделать на GET
		http.Redirect(writer, request, "/", http.StatusPermanentRedirect)
		return
	}
}

func (receiver *server) handleFavicon() func(http.ResponseWriter, *http.Request) {
	// TODO: handle concurrency
	file, err := ioutil.ReadFile(filepath.Join(receiver.assetsPath, "favicon.ico"))
	if err != nil {
		log.Printf("can't read favicon file: %v", err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write(file)
		if err != nil {
			log.Print(err)
		}
	}
}
