package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"mini-project/database"
	"mini-project/repository"
	"mini-project/structs"
	"regexp"
	"strconv"
)

func Segitiga(c *gin.Context) {

	alasStr := c.Query("alas")
	tinggiStr := c.Query("tinggi")

	alas, err1 := strconv.ParseFloat(alasStr, 64)
	tinggi, err2 := strconv.ParseFloat(tinggiStr, 64)

	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{
			"error": "Paramater harus numerik!",
		})
		return
	}

	hitung := c.Query("hitung")

	if hitung == "luas" {
		luas := 0.5 * alas * tinggi

		c.JSON(200, gin.H{
			"bangun_datar": "segitiga sama sisi",
			"alas":         alas,
			"tinggi":       tinggi,
			"luas":         luas,
		})
	} else if hitung == "keliling" {
		keliling := 3 * alas

		c.JSON(200, gin.H{
			"bangun_datar": "segitiga sama sisi",
			"alas":         alas,
			"tinggi":       tinggi,
			"keliling":     keliling,
		})
	} else {
		c.JSON(400, gin.H{
			"error": "Parameter tidak ditemukan !",
		})
	}
}

func Persegi(c *gin.Context) {
	sisiStr := c.Query("sisi")

	sisi, err := strconv.ParseFloat(sisiStr, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Parameter harus numerik",
		})
		return
	}

	hitung := c.Query("hitung")

	if hitung == "luas" {
		luas := math.Pow(sisi, 2)

		c.JSON(200, gin.H{
			"bangun_datar": "persegi",
			"sisi":         sisi,
			"luas":         luas,
		})
	} else if hitung == "keliling" {
		keliling := 4 * sisi

		c.JSON(200, gin.H{
			"bangun_datar": "persegi",
			"sisi":         sisi,
			"keliling":     keliling,
		})
	} else {
		c.JSON(400, gin.H{
			"error": "Parameter tidak ditemukan!",
		})
	}
}

func PersegiPanjang(c *gin.Context) {
	panjangStr := c.Query("panjang")
	lebarStr := c.Query("lebar")

	panjang, err1 := strconv.ParseFloat(panjangStr, 64)
	lebar, err2 := strconv.ParseFloat(lebarStr, 64)

	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{
			"error": "Parameter harus numerik!",
		})
		return
	}

	hitung := c.Query("hitung")

	if hitung == "luas" {
		luas := panjang * lebar

		c.JSON(200, gin.H{
			"bangun_datar": "persegi_panjang",
			"panjang":      panjang,
			"lebar":        lebar,
			"luas":         luas,
		})
	} else if hitung == "keliling" {
		keliling := 2 * (panjang + lebar)

		c.JSON(200, gin.H{
			"bangun_datar": "persegi_panjang",
			"panjang":      panjang,
			"lebar":        lebar,
			"keliling":     keliling,
		})
	} else {
		c.JSON(400, gin.H{
			"error": "Parameter tidak ditemukan !",
		})
	}
}

func Lingkaran(c *gin.Context) {

	jariJariStr := c.Query("jari-jari")

	jariJari, err := strconv.ParseFloat(jariJariStr, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Parameter harus numerik !",
		})
		return
	}

	hitung := c.Query("hitung")

	if hitung == "luas" {
		luas := math.Pi * math.Pow(jariJari, 2)

		c.JSON(200, gin.H{
			"bangun_datar": "lingkaran",
			"jari_jari":    jariJari,
			"luas":         luas,
		})
	} else if hitung == "keliling" {
		keliling := 2 * math.Pi * jariJari

		c.JSON(200, gin.H{
			"bangun_datar": "lingkaran",
			"jari_jari":    jariJari,
			"keliling":     keliling,
		})
	}
}

func GetCat(c *gin.Context) {
	var (
		result gin.H
	)

	category, err := repository.GetCat(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}
	c.JSON(200, result)
}

func PostCat(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.PostCat(database.DbConnection, category)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, gin.H{
		"result": "succes add category",
	})
}

func UpdateCat(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.Id = id

	err = repository.UpdateCat(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"result": "Success Update Category",
	})
}

func DeleteCat(c *gin.Context) {
	var category structs.Category

	id, err := strconv.Atoi(c.Param("id"))

	category.Id = id

	err = repository.DeleteCat(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"result": "Success Delete Category",
	})
}

func GetCatId(c *gin.Context) {
	var (
		result gin.H
	)
	var books structs.Books
	books.CategoryId, _ = strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	book, err1 := repository.GetCatId(database.DbConnection, books)

	if err1 != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}
	c.JSON(200, result)
}

func GetBooks(c *gin.Context) {
	var (
		result gin.H
	)

	book, err := repository.GetBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}
	c.JSON(200, result)
}

func PostBooks(c *gin.Context) {
	var books structs.Books

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	pattern := `^(https?|ftp)://[^\s/$.?#].[^\s]*$`

	match, err1 := regexp.MatchString(pattern, books.ImageUrl)
	if err1 != nil {
		panic(err1)
	}
	if match {
		c.Next()
	} else {
		c.AbortWithStatusJSON(400, gin.H{"error": "url tidak valid"})
	}

	if books.ReleaseYear < 1980 || books.ReleaseYear > 2021 {
		c.AbortWithStatusJSON(400, gin.H{"error": "release year tidak ditemukan"})
	} else {
		c.Next()
	}

	var thickness string

	if books.TotalPage <= 100 {
		thickness = "tipis"
	} else if books.TotalPage > 100 && books.TotalPage < 200 {
		thickness = "sedang"
	} else if books.TotalPage >= 201 {
		thickness = "tebal"
	}

	books.Thickness = thickness

	err = repository.PostBooks(database.DbConnection, books)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"result": "Succes add books",
	})
}

func UpdateBooks(c *gin.Context) {
	var books structs.Books
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	books.Id = id

	err = repository.UpdateBooks(database.DbConnection, books)

	c.JSON(200, gin.H{
		"result": "Succes updated books",
	})
}
func DeleteBooks(c *gin.Context) {
	var books structs.Books
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&books)
	if err != nil {
		panic(err)
	}

	books.Id = id

	err = repository.DeleteBooks(database.DbConnection, books)

	c.JSON(200, gin.H{
		"result": "Succes deleted books",
	})
}
