package api

import (
	"strconv"

	"github.com/lucianoq/hattrick/chpp"
	"github.com/lucianoq/hattrick/chpp/id"
)

// TransferSearchSkillFilter is an additional (secondary/tertiary/etc.)
// skill filter for GetTransferSearch. The difference between Min and Max
// can be up to 4 levels.
type TransferSearchSkillFilter struct {
	SkillType chpp.SkillID
	Min       chpp.SkillLevel
	Max       chpp.SkillLevel
}

// TransferSearchFilters holds the optional filters for GetTransferSearch.
// Zero-valued fields are omitted (left unspecified) except PageSize, which
// defaults to 25 if left 0, matching the CHPP's own default.
type TransferSearchFilters struct {
	AgeDaysMin uint
	AgeDaysMax uint

	// Up to 3 additional skill filters, beyond the required primary one.
	AdditionalSkills []TransferSearchSkillFilter

	Specialty       chpp.SpecialtyID
	NativeCountryID id.Country

	TSIMin uint
	TSIMax uint

	PriceMin chpp.Money
	PriceMax chpp.Money

	PageSize  uint
	PageIndex uint
}

// GetTransferSearch searches the transfer market for players matching the
// given age range and primary skill filter, plus any additional optional
// filters. Per the CHPP doc, minSkillValue1/maxSkillValue1 must be within 4
// levels of each other; set pieces, leadership and experience require a
// minimum level of 5, and keeper a minimum of 3.
func (a *API) GetTransferSearch(ageMin, ageMax uint, skillType1 chpp.SkillID, minSkillValue1, maxSkillValue1 chpp.SkillLevel, filters TransferSearchFilters) (*chpp.TransferSearchResults, error) {
	values := map[string]string{
		"ageMin":         strconv.FormatUint(uint64(ageMin), 10),
		"ageMax":         strconv.FormatUint(uint64(ageMax), 10),
		"skillType1":     strconv.FormatUint(uint64(skillType1), 10),
		"minSkillValue1": strconv.FormatUint(uint64(minSkillValue1), 10),
		"maxSkillValue1": strconv.FormatUint(uint64(maxSkillValue1), 10),
	}

	if filters.AgeDaysMin != 0 {
		values["ageDaysMin"] = strconv.FormatUint(uint64(filters.AgeDaysMin), 10)
	}
	if filters.AgeDaysMax != 0 {
		values["ageDaysMax"] = strconv.FormatUint(uint64(filters.AgeDaysMax), 10)
	}

	for i, skill := range filters.AdditionalSkills {
		if i >= 3 {
			break
		}
		n := strconv.Itoa(i + 2)
		values["skillType"+n] = strconv.FormatUint(uint64(skill.SkillType), 10)
		values["minSkillValue"+n] = strconv.FormatUint(uint64(skill.Min), 10)
		values["maxSkillValue"+n] = strconv.FormatUint(uint64(skill.Max), 10)
	}

	if filters.Specialty != 0 {
		values["specialty"] = strconv.FormatUint(uint64(filters.Specialty), 10)
	}
	if filters.NativeCountryID != 0 {
		values["nativeCountryId"] = filters.NativeCountryID.String()
	}
	if filters.TSIMin != 0 {
		values["tsiMin"] = strconv.FormatUint(uint64(filters.TSIMin), 10)
	}
	if filters.TSIMax != 0 {
		values["tsiMax"] = strconv.FormatUint(uint64(filters.TSIMax), 10)
	}
	if filters.PriceMin != 0 {
		values["priceMin"] = strconv.Itoa(int(filters.PriceMin))
	}
	if filters.PriceMax != 0 {
		values["priceMax"] = strconv.Itoa(int(filters.PriceMax))
	}
	if filters.PageSize != 0 {
		values["pageSize"] = strconv.FormatUint(uint64(filters.PageSize), 10)
	}
	values["pageIndex"] = strconv.FormatUint(uint64(filters.PageIndex), 10)

	res, err := a.parsed.GetTransferSearchXML(values)
	if err != nil {
		return nil, err
	}

	return &res.TransferSearch, nil
}
