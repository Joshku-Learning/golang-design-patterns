package abstract_factory

type nike struct{}

func (n *nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 15,
		},
	}
}

func (n *nike) MakeShirt() IShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 29,
		},
	}
}


