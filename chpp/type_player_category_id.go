package chpp

// PlayerCategoryID ...
type PlayerCategoryID uint

// List of PlayerCategoryID constants.
const (
	PlayerCategoryNoCategorySet   PlayerCategoryID = 0
	PlayerCategoryKeeper          PlayerCategoryID = 1
	PlayerCategoryWingBack        PlayerCategoryID = 2
	PlayerCategoryCentralDefender PlayerCategoryID = 3
	PlayerCategoryWinger          PlayerCategoryID = 4
	PlayerCategoryInnerMidfielder PlayerCategoryID = 5
	PlayerCategoryForward         PlayerCategoryID = 6
	PlayerCategorySubstitute      PlayerCategoryID = 7
	PlayerCategoryReserve         PlayerCategoryID = 8
	PlayerCategoryExtra1          PlayerCategoryID = 9
	PlayerCategoryExtra2          PlayerCategoryID = 10
)
