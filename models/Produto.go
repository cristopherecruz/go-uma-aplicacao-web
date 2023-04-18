package models

import (
	"database/sql"
)
import "github.com/goumaaplicacaoweb/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarProdutos() []Produto {

	database := db.ConectarBancoDados()

	dbProdutos, err := database.Query("SELECT * FROM produtos ORDER BY id")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	var produtos []Produto

	for dbProdutos.Next() {
		var id int
		var nome string
		var descricao string
		var quantidade int
		var preco float64

		err = dbProdutos.Scan(&nome, &descricao, &preco, &quantidade, &id)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}

	defer func(database *sql.DB) {
		var err = database.Close()
		if err != nil {
			panic(err.Error())
		}
	}(database)

	return produtos
}

func CadastrarProduto(nome string, descricao string, preco float64, quantidade int) {

	database := db.ConectarBancoDados()

	inserirDadosBanco, err := database.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = inserirDadosBanco.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}

	defer func(database *sql.DB) {
		var err = database.Close()
		if err != nil {
			panic(err.Error())
		}
	}(database)
}

func DeletarProduto(idProduto string) {

	database := db.ConectarBancoDados()

	deletarProduto, err := database.Prepare("DELETE FROM produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	_, err = deletarProduto.Exec(idProduto)
	if err != nil {
		panic(err.Error())
	}

	defer func(database *sql.DB) {
		var err = database.Close()
		if err != nil {
			panic(err.Error())
		}
	}(database)
}

func BuscarProdutoPorId(idBuscar string) Produto {

	database := db.ConectarBancoDados()

	buscarProduto, err := database.Query("SELECT * FROM produtos WHERE id = $1", idBuscar)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for buscarProduto.Next() {

		var id int
		var nome string
		var descricao string
		var quantidade int
		var preco float64

		err = buscarProduto.Scan(&nome, &descricao, &preco, &quantidade, &id)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Quantidade = quantidade
		produto.Preco = preco

	}

	return produto
}

func AtualizarProduto(id int, nome string, descricao string, preco float64, quantidade int) {

	database := db.ConectarBancoDados()

	atualizarDadosBanco, err := database.Prepare("UPDATE produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}

	_, err = atualizarDadosBanco.Exec(nome, descricao, preco, quantidade, id)
	if err != nil {
		panic(err.Error())
	}

	defer func(database *sql.DB) {
		var err = database.Close()
		if err != nil {
			panic(err.Error())
		}
	}(database)

}
