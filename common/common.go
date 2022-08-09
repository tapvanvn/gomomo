package common

import "strings"

const (
	DomainDev  = "https://api.mservice.com.vn/openapi"
	DomainProd = "https://openapi.momo.vn"
)

type Attribute string

const (
	AttriAddress          = Attribute("address")
	AttriAddressKYC       = Attribute("addressKyc")
	AttriCity             = Attribute("city")
	AttriCountryCode      = Attribute("countryCode")
	AttriCountryName      = Attribute("countryName")
	AttriCurrentAddress   = Attribute("currentAddress")
	AttriDeviceName       = Attribute("deviceName")
	AttriDeviceOS         = Attribute("deviceOs")
	AttriDOBKYC           = Attribute("dobKyc")
	AttriEmail            = Attribute("email")
	AttriFullNameKyc      = Attribute("fullNameKyc")
	AttriGender           = Attribute("gender")
	AttriGenderKYC        = Attribute("genderKyc")
	AttriLoyaltyPoints    = Attribute("loyaltyPoints")
	AttriMaritalStatus    = Attribute("maritalStatus")
	AttriMicroShop        = Attribute("microShop")
	AttriName             = Attribute("name")
	AttriNationality      = Attribute("nationality")
	AttriNationalityKYC   = Attribute("nationalityKyc")
	AttriNickname         = Attribute("nickname")
	AttriPhone            = Attribute("phone")
	AttriPreferedLanguage = Attribute("preferedLanguage")
	AttriProfressional    = Attribute("profressional")
)

func AttributesToString(atts []Attribute) string {

	tmp := []string{}
	for _, att := range atts {
		tmp = append(tmp, string(att))
	}
	return strings.Join(tmp, ",")
}

func AttributesFromString(attrStr string) []Attribute {

	parts := strings.Split(attrStr, ",")
	atts := []Attribute{}
	for _, el := range parts {
		atts = append(atts, Attribute(el))
	}
	return atts
}
