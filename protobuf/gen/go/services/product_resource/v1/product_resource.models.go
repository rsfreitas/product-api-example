// Code generated by protoc-gen-fasthttp. DO NOT EDIT.
package product_resourcev1

import (
	"time"
	"github.com/rsfreitas/go-pocket-utils/converters"
	"encoding/json"
)

//
// ///////////////////////// Product ///////////////////////
//

// ProductModel is the entity that a service must manipulate when
// dealing with database operations.
type ProductModel struct {
	Id        string     `json:"id,omitempty" bson:"id,omitempty" gorm:"id,omitempty"`
	Name      string     `json:"name,omitempty" bson:"name,omitempty" gorm:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty" gorm:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty" gorm:"updated_at,omitempty;autoUpdateTime:false"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" gorm:"deleted_at,omitempty"`
}

// ProtoResponse converts a ProductModel into a ProductProto.
func (p *ProductModel) ProtoResponse() *ProductProto {
	if p == nil {
		return nil
	}

	return &ProductProto{
		Id:        p.Id,
		Name:      p.Name,
		CreatedAt: converters.TimeToTimestamp(p.CreatedAt),
		UpdatedAt: converters.TimeToTimestamp(p.UpdatedAt),
		DeletedAt: converters.TimeToTimestamp(p.DeletedAt),
	}
}

//
// ///////////////////////// CreateProductRequest ///////////////////////
//

// CreateProductRequestModel is the entity that a service must manipulate when
// dealing with database operations.
type CreateProductRequestModel struct {
	Name string `json:"name,omitempty" bson:"name,omitempty" gorm:"name,omitempty"`
}

// ProtoResponse converts a CreateProductRequestModel into a CreateProductRequest.
func (c *CreateProductRequestModel) ProtoResponse() *CreateProductRequest {
	if c == nil {
		return nil
	}

	return &CreateProductRequest{
		Name: c.Name,
	}
}

//
// ///////////////////////// GetProductByIDRequest ///////////////////////
//

// GetProductByIDRequestModel is the entity that a service must manipulate when
// dealing with database operations.
type GetProductByIDRequestModel struct {
	Id            string `json:"id,omitempty" bson:"id,omitempty" gorm:"id,omitempty"`
	ReturnDeleted bool   `json:"return_deleted,omitempty" bson:"return_deleted,omitempty" gorm:"return_deleted,omitempty"`
}

// ProtoResponse converts a GetProductByIDRequestModel into a GetProductByIDRequest.
func (g *GetProductByIDRequestModel) ProtoResponse() *GetProductByIDRequest {
	if g == nil {
		return nil
	}

	return &GetProductByIDRequest{
		Id:            g.Id,
		ReturnDeleted: g.ReturnDeleted,
	}
}

//
// ///////////////////////// UpdateProductRequest ///////////////////////
//

// UpdateProductRequestModel is the entity that a service must manipulate when
// dealing with database operations.
type UpdateProductRequestModel struct {
	Id   string `json:"id,omitempty" bson:"id,omitempty" gorm:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty" gorm:"name,omitempty"`
}

// ProtoResponse converts a UpdateProductRequestModel into a UpdateProductRequest.
func (u *UpdateProductRequestModel) ProtoResponse() *UpdateProductRequest {
	if u == nil {
		return nil
	}

	return &UpdateProductRequest{
		Id:   u.Id,
		Name: u.Name,
	}
}

//
// ///////////////////////// DeleteProductRequest ///////////////////////
//

// DeleteProductRequestModel is the entity that a service must manipulate when
// dealing with database operations.
type DeleteProductRequestModel struct {
	Id string `json:"id,omitempty" bson:"id,omitempty" gorm:"id,omitempty"`
}

// ProtoResponse converts a DeleteProductRequestModel into a DeleteProductRequest.
func (d *DeleteProductRequestModel) ProtoResponse() *DeleteProductRequest {
	if d == nil {
		return nil
	}

	return &DeleteProductRequest{
		Id: d.Id,
	}
}

// ///////////////////// CreateProductResponse //////////////////////
//
// HttpResponseOrNil is the struct method responsible for safe casting the
// return of HttpResponse call. It is usually called when a structure has a
// message member, since it will required its specific type when updated.
func (c *CreateProductResponse) HttpResponseOrNil() *CreateProductResponseHttpResponse {
	if r := c.HttpResponse(); r != nil {
		return r.(*CreateProductResponseHttpResponse)
	}

	return nil
}

func (c *CreateProductResponse) HttpResponseBytes() ([]byte, error) {
	response := c.HttpResponse()
	return json.Marshal(response)
}

// HttpResponse is the HttpResponser interface implementation for CreateProductResponse object.
// It converts all members according their desired new types and it is responsible
// for the real output of CreateProductResponse as a HTTP response.
func (c *CreateProductResponse) HttpResponse() interface{} {
	if c == nil {
		return nil
	}

	return &CreateProductResponseHttpResponse{
		Product: c.Product.HttpResponseOrNil(),
	}
}

// CreateProductResponseHttpResponse is the real output of object CreateProductResponse.
// It maps all original members to their new types.
type CreateProductResponseHttpResponse struct {
	Product *ProductHttpResponse `json:"product,omitempty" bson:"product,omitempty" gorm:"product,omitempty"`
}

// ///////////////////// ProductProto //////////////////////
//
// HttpResponseOrNil is the struct method responsible for safe casting the
// return of HttpResponse call. It is usually called when a structure has a
// message member, since it will required its specific type when updated.
func (p *ProductProto) HttpResponseOrNil() *ProductHttpResponse {
	if r := p.HttpResponse(); r != nil {
		return r.(*ProductHttpResponse)
	}

	return nil
}

func (p *ProductProto) HttpResponseBytes() ([]byte, error) {
	response := p.HttpResponse()
	return json.Marshal(response)
}

// HttpResponse is the HttpResponser interface implementation for ProductProto object.
// It converts all members according their desired new types and it is responsible
// for the real output of ProductProto as a HTTP response.
func (p *ProductProto) HttpResponse() interface{} {
	if p == nil {
		return nil
	}

	return &ProductHttpResponse{
		Id:        p.Id,
		Name:      p.Name,
		CreatedAt: converters.ConvertFromTimestampToTimePointer(p.CreatedAt),
		UpdatedAt: converters.ConvertFromTimestampToTimePointer(p.UpdatedAt),
		DeletedAt: converters.ConvertFromTimestampToTimePointer(p.DeletedAt),
	}
}

// ProductProtoHttpResponse is the real output of object ProductProto.
// It maps all original members to their new types.
type ProductHttpResponse struct {
	Id        string     `json:"id,omitempty" bson:"id,omitempty" gorm:"id,omitempty"`
	Name      string     `json:"name,omitempty" bson:"name,omitempty" gorm:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty" gorm:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty" gorm:"updated_at,omitempty;autoUpdateTime:false"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty" gorm:"deleted_at,omitempty"`
}

// ///////////////////// GetProductByIDResponse //////////////////////
//
// HttpResponseOrNil is the struct method responsible for safe casting the
// return of HttpResponse call. It is usually called when a structure has a
// message member, since it will required its specific type when updated.
func (g *GetProductByIDResponse) HttpResponseOrNil() *GetProductByIDResponseHttpResponse {
	if r := g.HttpResponse(); r != nil {
		return r.(*GetProductByIDResponseHttpResponse)
	}

	return nil
}

func (g *GetProductByIDResponse) HttpResponseBytes() ([]byte, error) {
	response := g.HttpResponse()
	return json.Marshal(response)
}

// HttpResponse is the HttpResponser interface implementation for GetProductByIDResponse object.
// It converts all members according their desired new types and it is responsible
// for the real output of GetProductByIDResponse as a HTTP response.
func (g *GetProductByIDResponse) HttpResponse() interface{} {
	if g == nil {
		return nil
	}

	return &GetProductByIDResponseHttpResponse{
		Product: g.Product.HttpResponseOrNil(),
	}
}

// GetProductByIDResponseHttpResponse is the real output of object GetProductByIDResponse.
// It maps all original members to their new types.
type GetProductByIDResponseHttpResponse struct {
	Product *ProductHttpResponse `json:"product,omitempty" bson:"product,omitempty" gorm:"product,omitempty"`
}

// ///////////////////// UpdateProductResponse //////////////////////
//
// HttpResponseOrNil is the struct method responsible for safe casting the
// return of HttpResponse call. It is usually called when a structure has a
// message member, since it will required its specific type when updated.
func (u *UpdateProductResponse) HttpResponseOrNil() *UpdateProductResponseHttpResponse {
	if r := u.HttpResponse(); r != nil {
		return r.(*UpdateProductResponseHttpResponse)
	}

	return nil
}

func (u *UpdateProductResponse) HttpResponseBytes() ([]byte, error) {
	response := u.HttpResponse()
	return json.Marshal(response)
}

// HttpResponse is the HttpResponser interface implementation for UpdateProductResponse object.
// It converts all members according their desired new types and it is responsible
// for the real output of UpdateProductResponse as a HTTP response.
func (u *UpdateProductResponse) HttpResponse() interface{} {
	if u == nil {
		return nil
	}

	return &UpdateProductResponseHttpResponse{
		Product: u.Product.HttpResponseOrNil(),
	}
}

// UpdateProductResponseHttpResponse is the real output of object UpdateProductResponse.
// It maps all original members to their new types.
type UpdateProductResponseHttpResponse struct {
	Product *ProductHttpResponse `json:"product,omitempty" bson:"product,omitempty" gorm:"product,omitempty"`
}

// ///////////////////// DeleteProductResponse //////////////////////
//
// HttpResponseOrNil is the struct method responsible for safe casting the
// return of HttpResponse call. It is usually called when a structure has a
// message member, since it will required its specific type when updated.
func (d *DeleteProductResponse) HttpResponseOrNil() *DeleteProductResponseHttpResponse {
	if r := d.HttpResponse(); r != nil {
		return r.(*DeleteProductResponseHttpResponse)
	}

	return nil
}

func (d *DeleteProductResponse) HttpResponseBytes() ([]byte, error) {
	response := d.HttpResponse()
	return json.Marshal(response)
}

// HttpResponse is the HttpResponser interface implementation for DeleteProductResponse object.
// It converts all members according their desired new types and it is responsible
// for the real output of DeleteProductResponse as a HTTP response.
func (d *DeleteProductResponse) HttpResponse() interface{} {
	if d == nil {
		return nil
	}

	return &DeleteProductResponseHttpResponse{
		Product: d.Product.HttpResponseOrNil(),
	}
}

// DeleteProductResponseHttpResponse is the real output of object DeleteProductResponse.
// It maps all original members to their new types.
type DeleteProductResponseHttpResponse struct {
	Product *ProductHttpResponse `json:"product,omitempty" bson:"product,omitempty" gorm:"product,omitempty"`
}
