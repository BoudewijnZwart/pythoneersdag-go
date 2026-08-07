package main

use (
	"rand"
	"assignment2/shared/utils"
)

const modalSalary = 48000

type Child struct {
	name string
	wish string
	salaryParentsInEuros int
	hasBeenGood bool
	present string
}

func main(){}

func generator() {

}

func investigateChild(c *Child) *Child {
	c.hasBeenGood := rand.Intn(2) == 1
	return c
}

func selectPresentGoodKid(c *Child) *Child {

	if c.salaryParentsInEuros > modalSalary {
		c.present = c.wish
		utils.DoWork("")		
		return c
	}

	c.present = nil
	return c
}

