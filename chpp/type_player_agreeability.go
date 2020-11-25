package chpp

// PlayerAgreeability ...
type PlayerAgreeability uint

// List of PlayerAgreeability constants.
const (
	PlayerAgreeabilityNastyFellow         PlayerAgreeability = 0
	PlayerAgreeabilityControversialPerson PlayerAgreeability = 1
	PlayerAgreeabilityPleasantGuy         PlayerAgreeability = 2
	PlayerAgreeabilitySympatheticGuy      PlayerAgreeability = 3
	PlayerAgreeabilityPopularGuy          PlayerAgreeability = 4
	PlayerAgreeabilityBelovedTeamMember   PlayerAgreeability = 5
)

// String returns a string representation of the type.
func (a PlayerAgreeability) String() string {
	switch a {
	case PlayerAgreeabilityNastyFellow:
		return "nasty fellow"
	case PlayerAgreeabilityControversialPerson:
		return "controversial person"
	case PlayerAgreeabilityPleasantGuy:
		return "pleasant guy"
	case PlayerAgreeabilitySympatheticGuy:
		return "sympathetic guy"
	case PlayerAgreeabilityPopularGuy:
		return "popular guy"
	case PlayerAgreeabilityBelovedTeamMember:
		return "beloved team member"
	}

	return ""
}
