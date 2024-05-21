package main

type Nike struct {
}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe{
			logo: "nike",
			size: 40,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt{
			logo: "nike",
			size: 40,
		},
	}
}
