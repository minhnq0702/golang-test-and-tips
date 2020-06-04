package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Partner basic info
var Partner = Type("Partner", func() {
	Description("Partner definition")
	Attribute("ID", Integer, func() {
		Description("Partner Identity")
		Example(1)
	})
	Attribute("Name", String, func() {
		Description("Partner Name")
	})
	Attribute("Code", String, func() {
		Description("Unique partner code")
	})
	Attribute("PartnerType", String, func() {
		Description("Type of partner: shipping address/invoice address/employee")
	})
	Attribute("Title", String)
	Attribute("JobPosition", String)
	Attribute("Street", String)
	Attribute("Street2", String)
	Attribute("District", String)
	Attribute("City", String)
	Attribute("Country", String)
	Attribute("Zip", String)
	Attribute("IsCompany", Boolean, func() {
		Description("Check Partner is company or invidual")
	})
	Attribute("Active", Boolean)
})

// PartnerMedia is using for list partner data
var PartnerMedia = MediaType("application/vnd.goa.partner+json", func() {
	Description("Partner is address/customer/supplier/employee info")
	Reference(Partner)
	Attributes(func() {
		Attribute("ID", Integer)
		Attribute("Name", String)
		Attribute("Code", String)
		Attribute("PartnerType", String)
		Attribute("Title", String)
		Attribute("JobPosition", String)
		Attribute("Street", String)
		Attribute("Street2", String)
		Attribute("District", String)
		Attribute("City", String)
		Attribute("Country", String)
		Attribute("Zip", String)
		Attribute("IsCompany", Boolean)
		Attribute("Active", Boolean)
	})
	View("default", func() {
		Attribute("ID", Integer)
		Attribute("Name", String)
		Attribute("Code", String)
	})
})

var _ = Resource("Partner", func() {
	BasePath("/partners")
	DefaultMedia(PartnerMedia)

	Action("List", func() {
		Description("Get all partner")
		Routing(GET("/"))
		Response(OK, PartnerMedia)
		Response(InternalServerError)
	})
})
