package models

import (
	"api_get_products/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"time"
)

type Product struct {
	Id    int64
	Skus  []string
	Names []string
	Price float64	
	CreatedAt  time.Time

}

type Products struct {
	Products []Product `json:"product"`
}

var con *gorm.DB

func GetProd1(productID string) Products {
	con := db.CreateCon()
	
	//Command SQL
	productQuery := `
		SELECT 
			products.id, 
			array(
				SELECT pv.sku 
				FROM product_variants AS pv 
				WHERE pv.product_id = products.id
			) AS skus,
			products.price, 
			array(
				SELECT pd.name 
				FROM product_descriptions AS pd 
				WHERE pd.product_id = products.id
			) AS names,
			products.created_at
		FROM products
		WHERE products.id = ?
	`

	rows, err := con.Raw(productQuery, productID).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := Products{}

	for rows.Next() {
		product := Product{}

		err := rows.Scan(
			&product.Id,
			pq.Array(&product.Skus),
			&product.Price,
			pq.Array(&product.Names),
			&product.CreatedAt,
		)
		if err != nil {
			panic(err)
		}
		result.Products = append(result.Products, product)
	}
	return result
}
