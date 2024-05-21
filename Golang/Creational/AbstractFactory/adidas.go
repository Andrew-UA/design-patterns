package main

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe{
			logo: "adidas",
			size: 40,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt{
			logo: "adidas",
			size: 40,
		},
	}
}
