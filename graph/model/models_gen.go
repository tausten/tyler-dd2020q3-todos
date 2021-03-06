// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// An object with an ID
type Node interface {
	IsNode()
}

type Constellation struct {
	ID        string            `json:"id"`
	CreatedAt string            `json:"createdAt"`
	UpdatedAt *string           `json:"updatedAt"`
	Name      *string           `json:"name"`
	Galaxies  *GalaxyConnection `json:"galaxies"`
}

func (Constellation) IsNode() {}

type ConstellationConnection struct {
	Nodes      []*Constellation     `json:"nodes"`
	Edges      []*ConstellationEdge `json:"edges"`
	PageInfo   *PageInfo            `json:"pageInfo"`
	TotalCount int                  `json:"totalCount"`
}

type ConstellationEdge struct {
	Cursor string         `json:"cursor"`
	Node   *Constellation `json:"node"`
}

type CreateConstellationInput struct {
	Name string `json:"name"`
}

type CreateConstellationPayload struct {
	Constellation *Constellation `json:"constellation"`
}

type CreateConstellationsPayload struct {
	Constellations []*Constellation `json:"constellations"`
}

type Galaxy struct {
	ID        string            `json:"id"`
	CreatedAt string            `json:"createdAt"`
	UpdatedAt *string           `json:"updatedAt"`
	Name      *string           `json:"name"`
	Planets   *PlanetConnection `json:"planets"`
	Stars     *StarConnection   `json:"stars"`
}

func (Galaxy) IsNode() {}

type GalaxyConnection struct {
	Nodes      []*Galaxy     `json:"nodes"`
	Edges      []*GalaxyEdge `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

type GalaxyEdge struct {
	Cursor string  `json:"cursor"`
	Node   *Galaxy `json:"node"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
}

type Planet struct {
	ID          string          `json:"id"`
	CreatedAt   string          `json:"createdAt"`
	UpdatedAt   *string         `json:"updatedAt"`
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	PlanetType  *PlanetTypeEnum `json:"planetType"`
	Galaxy      *Galaxy         `json:"galaxy"`
}

func (Planet) IsNode() {}

type PlanetConnection struct {
	Nodes      []*Planet     `json:"nodes"`
	Edges      []*PlanetEdge `json:"edges"`
	PageInfo   *PageInfo     `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

type PlanetEdge struct {
	Cursor string  `json:"cursor"`
	Node   *Planet `json:"node"`
}

type Star struct {
	ID        string            `json:"id"`
	CreatedAt string            `json:"createdAt"`
	UpdatedAt *string           `json:"updatedAt"`
	Name      *string           `json:"name"`
	Class     *StellarClassEnum `json:"class"`
}

func (Star) IsNode() {}

type StarConnection struct {
	Nodes      []*Star     `json:"nodes"`
	Edges      []*StarEdge `json:"edges"`
	PageInfo   *PageInfo   `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

type StarEdge struct {
	Cursor string `json:"cursor"`
	Node   *Star  `json:"node"`
}

// A user in the system
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PlanetTypeEnum string

const (
	PlanetTypeEnumDoublePlanet  PlanetTypeEnum = "DOUBLE_PLANET"
	PlanetTypeEnumDwarfPlanet   PlanetTypeEnum = "DWARF_PLANET"
	PlanetTypeEnumExoplanet     PlanetTypeEnum = "EXOPLANET"
	PlanetTypeEnumExtragalactic PlanetTypeEnum = "EXTRAGALACTIC"
	PlanetTypeEnumMajorPlanet   PlanetTypeEnum = "MAJOR_PLANET"
	PlanetTypeEnumOuterPlanet   PlanetTypeEnum = "OUTER_PLANET"
	PlanetTypeEnumPulsarPlanet  PlanetTypeEnum = "PULSAR_PLANET"
	PlanetTypeEnumRoguePlanet   PlanetTypeEnum = "ROGUE_PLANET"
)

var AllPlanetTypeEnum = []PlanetTypeEnum{
	PlanetTypeEnumDoublePlanet,
	PlanetTypeEnumDwarfPlanet,
	PlanetTypeEnumExoplanet,
	PlanetTypeEnumExtragalactic,
	PlanetTypeEnumMajorPlanet,
	PlanetTypeEnumOuterPlanet,
	PlanetTypeEnumPulsarPlanet,
	PlanetTypeEnumRoguePlanet,
}

func (e PlanetTypeEnum) IsValid() bool {
	switch e {
	case PlanetTypeEnumDoublePlanet, PlanetTypeEnumDwarfPlanet, PlanetTypeEnumExoplanet, PlanetTypeEnumExtragalactic, PlanetTypeEnumMajorPlanet, PlanetTypeEnumOuterPlanet, PlanetTypeEnumPulsarPlanet, PlanetTypeEnumRoguePlanet:
		return true
	}
	return false
}

func (e PlanetTypeEnum) String() string {
	return string(e)
}

func (e *PlanetTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlanetTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlanetTypeEnum", str)
	}
	return nil
}

func (e PlanetTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StellarClassEnum string

const (
	StellarClassEnumClass0 StellarClassEnum = "CLASS_0"
	StellarClassEnumClassA StellarClassEnum = "CLASS_A"
	StellarClassEnumClassB StellarClassEnum = "CLASS_B"
)

var AllStellarClassEnum = []StellarClassEnum{
	StellarClassEnumClass0,
	StellarClassEnumClassA,
	StellarClassEnumClassB,
}

func (e StellarClassEnum) IsValid() bool {
	switch e {
	case StellarClassEnumClass0, StellarClassEnumClassA, StellarClassEnumClassB:
		return true
	}
	return false
}

func (e StellarClassEnum) String() string {
	return string(e)
}

func (e *StellarClassEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StellarClassEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StellarClassEnum", str)
	}
	return nil
}

func (e StellarClassEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
