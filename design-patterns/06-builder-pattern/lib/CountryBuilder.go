package lib

type Country struct {
	name, currency string
}

type countryMod func(*Country)
type CountryBuilder struct {
	actions []countryMod
}

func (cb *CountryBuilder) Name(name string) *CountryBuilder {
	cb.actions = append(cb.actions, func(c *Country) {
		c.name = name
	})
	return cb
}

func (cb *CountryBuilder) Currency(currency string) *CountryBuilder {
	cb.actions = append(cb.actions, func(c *Country) {
		c.currency = currency
	})
	return cb
}

func (cb *CountryBuilder) Build() *Country {
	c := &Country{}
	for _, action := range cb.actions {
		action(c)
	}
	return c
}
