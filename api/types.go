package api

import (
	"encoding/json"
	"fmt"
)

//StringBoolean type
type StringBoolean string

//Socket struct for item sockets
type Socket struct {
	Group  int64         `json:"group"`
	Attr   StringBoolean `json:"attr"`
	Colour string        `json:"sColour"`
}

//PropertyValues is values of properties
type PropertyValues struct {
	StrVal string
	IntVal int64
}

//Property of item
type Property struct {
	Name        string           `json:"name"`
	Values      []PropertyValues `json:"values"`
	DisplayMode int64            `json:"displayMode"`
	Type        int64            `json:"type"`
	Progress    float64          `json:"progress"`
}

//ItemCategory struct
type ItemCategory struct {
	Info map[string][]string
}

//Item struct
type Item struct {
	AbyssJewel            bool         `json:"abyssJewel,omitempty"`
	AdditionalPropreties  []Property   `json:"additionalProperties,omitempty"`
	ArtFileName           string       `json:"artFileName,omitempty"`
	Category              ItemCategory `json:"category"`
	Corrupted             bool         `json:"corrupted,omitempty"`
	CosmeticMods          []string     `json:"cosmeticMods,omitempty"`
	CraftedMods           []string     `json:"craftedMods,omitempty"`
	DescrText             string       `json:"descrText,omitempty"`
	Duplicated            bool         `json:"duplicated,omitempty"`
	Elder                 bool         `json:"elder,omitempty"`
	EnchantMods           []string     `json:"enchantMods,omitempty"`
	ExplicitMods          []string     `json:"explicitMods,omitempty"`
	FlavourText           []string     `json:"flavourText,omitempty"`
	FrameType             int64        `json:"frameType"`
	H                     int64        `json:"h"`
	ID                    string       `json:"id"`
	Icon                  string       `json:"icon"`
	Identified            bool         `json:"identified"`
	Ilvl                  int64        `json:"ilvl"`
	ImplicitMods          []string     `json:"implicitMods,omitempty"`
	InventoryID           string       `json:"inventoryId,omitempty"`
	IsRelic               bool         `json:"isRelic,omitempty"`
	League                string       `json:"league"`
	LockedToCharacter     bool         `json:"lockedToCharacter,omitempty"`
	MaxStackSize          int64        `json:"maxStackSize,omitempty"`
	Name                  string       `json:"name"`
	NextLevelRequirements []Property   `json:"nextLevelRequirements,omitempty"`
	Note                  string       `json:"note,omitempty"`
	Properties            []Property   `json:"properties,omitempty"`
	ProphecyDiffText      string       `json:"prophecyDiffText,omitempty"`
	ProphecyText          string       `json:"prophecyText,omitempty"`
	Requirements          []Property   `json:"requirements,omitempty"`
	SecDescrText          string       `json:"secDescrText,omitempty"`
	Shaper                bool         `json:"shaper,omitempty"`
	SocketedItems         []Item       `json:"socketedItems,omitempty"`
	Sockets               []Socket     `json:"sockets"`
	StackSize             int64        `json:"stackSize,omitempty"`
	Support               bool         `json:"support,omitempty"`
	TalismanTier          int64        `json:"talismanTier,omitempty"`
	TypeLine              string       `json:"typeLine"`
	UtilityMods           []string     `json:"utilityMods,omitempty"`
	Verified              bool         `json:"verified"`
	W                     int64        `json:"w"`
	X                     int64        `json:"x"`
	Y                     int64        `json:"y"`
}

//Stash tab struct
type Stash struct {
	AccountName       string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	ID                string `json:"id"`
	StashName         string `json:"stash"`
	StashType         string `json:"stashType"`
	Items             []Item `json:"items"`
	Public            bool   `json:"public"`
}

// PublicStashesResponse is pathofexile respone struct
type PublicStashesResponse struct {
	NextChangeID string  `json:"next_change_id"`
	Stashes      []Stash `json:"stashes"`
}

//UnmarshalJSON for StringBoolan
//That sucks, but Attr can be false for abyss jewels
func (s *StringBoolean) UnmarshalJSON(buf []byte) error {
	*s = StringBoolean(buf)
	return nil
}

//UnmarshalJSON for PropertyValues
func (p *PropertyValues) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&p.StrVal, &p.IntVal}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
	}
	return nil
}
