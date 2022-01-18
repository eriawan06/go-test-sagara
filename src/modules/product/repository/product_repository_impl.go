package repository

import (
	"context"
	"errors"

	"github.com/eriawan06/go-test-sagara/src/modules/product/model/domain"
	"github.com/jmoiron/sqlx"
)

type ProductRepositoryImpl struct {
	DB *sqlx.DB
}

func NewProductRepository(DB *sqlx.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: DB}
}

func (repo *ProductRepositoryImpl) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `
	INSERT INTO product(name, description, price, image)
	VALUES (?, ?, ?, ?)
	`
	result, err := repo.DB.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Image)
	if err != nil {
		return domain.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}

	insertedProduct, _ := repo.FindById(ctx, int(id))
	product = insertedProduct
	return product, nil
}

func (repo *ProductRepositoryImpl) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	query := `
		UPDATE product SET name=?, description=?, price=?, image=?
		WHERE id = ?
	`
	_, err := repo.DB.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Image, product.Id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (repo *ProductRepositoryImpl) Delete(ctx context.Context, product domain.Product) error {
	query := `DELETE FROM product WHERE id = ?`
	_, err := repo.DB.ExecContext(ctx, query, product.Id)
	return err
}

func (repo *ProductRepositoryImpl) FindById(ctx context.Context, productId int) (domain.Product, error) {
	query := `SELECT * FROM product WHERE id = ?`
	row := repo.DB.QueryRowxContext(ctx, query, productId)

	product := domain.Product{}

	if row == nil {
		return product, errors.New("product is not found")
	}

	err := row.StructScan(&product)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (repo *ProductRepositoryImpl) FindAll(ctx context.Context) ([]domain.Product, error) {
	query := `SELECT * FROM product`
	rows, err := repo.DB.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err = rows.StructScan(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
