/*
 * MVCGo
 * Copyright (c) 2021.
 * author Muhammad Farid H
 */

package models

import (
	"mvcgo/config"
	"mvcgo/systems/handler"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Product struct {
	Id    int    `json:"id" `
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

func (*Product) Validate(request Product) {
	if err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Stock, validation.Required, validation.Min(1)),
	); err != nil {
		panic(handler.ValidationError{
			Message: err.Error(),
		})
	}
}

func (*Product) Create(request Product) error {
	_, err := config.Db().Model(&Product{
		Name:  request.Name,
		Stock: request.Stock,
	}).Insert()

	if err != nil {
		handler.PanicIfNeeded(err)
	}

	return nil
}
func (*Product) Update(id int, request Product) error {
	data := &Product{
		Id:    id,
		Name:  request.Name,
		Stock: request.Stock}
	_, err := config.Db().Model(data).WherePK().Update()
	handler.PanicIfNeeded(err)

	return nil

}
func (*Product) Delete(id int) error {
	var data Product
	_, err := config.Db().Model(&data).Where("id = ?", id).Delete()
	if err != nil {
		handler.PanicIfNeeded(err)
	}
	return nil
}

func (*Product) GetOne(id int) (Product, error) {
	var data Product
	if err := config.Db().Model(&data).Where("id=?", id).Select(); err != nil {
		handler.PanicIfNeeded(err)
	}
	return data, nil

}

func (*Product) List() ([]Product, error) {
	var data []Product
	if err := config.Db().Model(&data).Select(); err != nil {
		handler.PanicIfNeeded(err)
	}
	return data, nil

}
