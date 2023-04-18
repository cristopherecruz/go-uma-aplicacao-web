package controllers

import (
	"github.com/goumaaplicacaoweb/models"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, _ *http.Request) {

	/*produto := Produto{"Fone", "Muito bom", 59, 2}

	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		produto,
		{"Produto novo", "Muito legal", 1.99, 1},
	}*/

	produtos := models.BuscarProdutos()

	err := templates.ExecuteTemplate(w, "Index", produtos)
	if err != nil {
		panic(err.Error())
	}

}

func NovoProduto(w http.ResponseWriter, _ *http.Request) {

	err := templates.ExecuteTemplate(w, "NovoProduto", nil)
	if err != nil {
		panic(err.Error())
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		parsePreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}

		parseQuantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.CadastrarProduto(nome, descricao, parsePreco, parseQuantidade)
	}

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")
	models.DeletarProduto(idProduto)

	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")

	produto := models.BuscarProdutoPorId(idProduto)

	err := templates.ExecuteTemplate(w, "Update", produto)
	if err != nil {
		panic(err.Error())
	}

}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		parseId, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}

		parsePreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}

		parseQuantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.AtualizarProduto(parseId, nome, descricao, parsePreco, parseQuantidade)
	}

	http.Redirect(w, r, "/", 301)

}
