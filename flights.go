// Package htma provides a Hypertext Markup Abstraction for generating HTML in pure Go.
package htma

// FlightCard creates a new flight-card element.
func FlightCard() Element {
	return newElement("flight-card", false)
}
func SearchCard() Element {
	return newElement("search-card", false)
}
func BliptaFooter() Element {
	return newElement("blipta-footer", false)
}
func BliptaHeader() Element {
	return newElement("blipta-header", false)
}

// FlightCard-Specific Attribute Methods
func (e Element) AirlineLogoTextAttr(value string) Element {
	return e.Attr("airline-logo-text", value)
}

func (e Element) AirlineClassAttr(value string) Element {
	return e.Attr("airline-class", value)
}

func (e Element) FlightNumberAttr(value string) Element {
	return e.Attr("flight-number", value)
}

func (e Element) AirlineNameAttr(value string) Element {
	return e.Attr("airline-name", value)
}

func (e Element) OriginIataAttr(value string) Element {
	return e.Attr("origin-iata", value)
}

func (e Element) OriginCityAttr(value string) Element {
	return e.Attr("origin-city", value)
}

func (e Element) DestIataAttr(value string) Element {
	return e.Attr("dest-iata", value)
}

func (e Element) DestCityAttr(value string) Element {
	return e.Attr("dest-city", value)
}

func (e Element) GateAttr(value string) Element {
	return e.Attr("gate", value)
}

func (e Element) BoardingTimeAttr(value string) Element {
	return e.Attr("boarding-time", value)
}

func (e Element) DepartureTimeAttr(value string) Element {
	return e.Attr("departure-time", value)
}

func (e Element) StatusTextAttr(value string) Element {
	return e.Attr("status-text", value)
}

func (e Element) StatusClassAttr(value string) Element {
	return e.Attr("status-class", value)
}

func (e Element) ArrivalTimeAttr(value string) Element {
	return e.Attr("arrival-time", value)
}
