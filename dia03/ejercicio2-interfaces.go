package main

import "fmt"

const (
	BigProductType    = "BIG_PRODUCT"
	MediumProductType = "MEDIUM_PRODUCT"
	SmallProductType  = "SMALL_PRODUCT"
)

type ProductCommerce struct {
	Products []IProduct
}

type IProduct interface {
	FullPrice() float64
}

func NewProduct(name string, price float64, details ...float64) (product IProduct, ok bool) {
	lenDetails := len(details)
	switch name {
	case BigProductType:
		if lenDetails == 0 || lenDetails > 1 {
			product = BigProduct{}
			return
		}
		product = BigProduct{
			Price:            price,
			MantenerEnTienda: details[0],
		}
		ok = !ok
	case MediumProductType:
		if lenDetails == 0 || lenDetails > 1 {
			product = MediumProduct{}
			return
		}
		product = MediumProduct{
			Price:            price,
			MantenerEnTienda: details[0],
		}
		ok = !ok
	case SmallProductType:
		if len(details) != 0 {
			product = SmallProduct{}
			return
		}
		product = SmallProduct{
			Price: price,
		}
		ok = !ok
	default:
		product = SmallProduct{}
	}
	return
}

type BigProduct struct {
	Price            float64
	MantenerEnTienda float64
}

func (p BigProduct) FullPrice() float64 {
	return p.Price + 2500 + (p.MantenerEnTienda * 0.6)
}

type MediumProduct struct {
	Price            float64
	MantenerEnTienda float64
}

func (p MediumProduct) FullPrice() float64 {
	return p.Price + (p.MantenerEnTienda * 0.3)
}

type SmallProduct struct {
	Price            float64
	MantenerEnTienda float64
}

func (p SmallProduct) FullPrice() float64 {
	return p.Price
}

func main() {
	// productoGrande, ok := NewProduct(BigProductType, 400000.0, 2000.0)
	// if !ok {
	// 	panic("Error en crear el producto")
	// }

	productoMediano, ok := NewProduct(MediumProductType, 100000.0, 1000.0)

	// productoChico, ok := NewProduct(SmallProductType, 100000.0)
	if !ok {
		panic("Error en crear el producto")
	}

	// fmt.Printf("El precio del producto es de %f\n", productoGrande.FullPrice())
	fmt.Printf("El precio del producto es de %f\n", productoMediano.FullPrice())
	// fmt.Printf("El precio del producto es de %f\n", productoChico.FullPrice())
}
