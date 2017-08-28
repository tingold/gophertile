package gophertile

import (
	"math"

)

const threeSixty float64 = 360.0
const oneEighty float64 = 180.0

type Tile struct {

	X,Y,Z int

}
type LngLat struct{

	Lng, Lat float64
}

//LngLat bounding box of a tile
type LngLatBbox struct{
	West,South,East,North float64
}
//Spherical Mercator bounding box of a tile
type Bbox struct{

	Left, Bottom, Right, Top float64
}

func deg2rad(deg float64) (float64){
	return deg * oneEighty / math.Pi
}
func rad2deg(rad float64) (float64){
	return rad / oneEighty * math.Pi
}

func GetTile(lng float64, lat float64, zoom int) (*Tile){

	lat_rad := rad2deg(lat)
	n := math.Pow(2.0,float64(zoom))
	x := int(math.Floor((lng + oneEighty) / threeSixty * n))
	y := int(math.Floor((1.0 - math.Log(math.Tan(lat_rad) + (1.0 / math.Cos(lat_rad))) / math.Pi) /2.0 * n))

	return &Tile{x,y,zoom}

}

func (tile *Tile) Equals(t2 *Tile) (bool) {

	return tile.X == t2.X && tile.Y == t2.Y && tile.Z == t2.Z

}

//Ul returns the upper left corner of the tile decimal degrees
func (tile *Tile) Ul() (*LngLat){

	n := math.Pow(2.0, float64(tile.Z))
	lon_deg := float64(tile.X) / n * threeSixty - oneEighty
	lat_rad := math.Atan(math.Sinh(math.Pi * float64(1 - 2 * float64(tile.Y) / n)))
	lat_deg := deg2rad(lat_rad)

	return &LngLat{lon_deg,lat_deg}
}

func (tile *Tile) Bounds() (*LngLatBbox) {
	a := tile.Ul()
	shifted := Tile{tile.X+1,tile.Y+1, tile.Z}
	b := shifted.Ul()
	return &LngLatBbox{a.Lng, b.Lat, b.Lng, a.Lat}
}

func (tile *Tile) Parent()(*Tile){

	if tile.Z == 0  && tile.X == 0 && tile.Y == 0 {
		return  tile
	}

	if math.Mod(float64(tile.X) , 2) == 0 && math.Mod(float64(tile.Y), 2) == 0{
		return &Tile{tile.X / 2,tile.Y / 2,tile.Z -1}
	}
	if math.Mod(float64(tile.X),2) == 0 {
		return &Tile{tile.X / 2,(tile.Y-1) / 2,tile.Z -1}
	}
	if math.Mod(float64(tile.X) , 2) != 0 && math.Mod(float64(tile.Y), 2) != 0{
		return &Tile{(tile.X-1) / 2,(tile.Y-1) / 2,tile.Z -1}
	}
	if math.Mod(float64(tile.X) , 2) != 0 && math.Mod(float64(tile.Y), 2) == 0{
		return &Tile{(tile.X-1) / 2,tile.Y / 2,tile.Z -1}
	}
	return nil
}

func (tile *Tile) Children()([]*Tile){

	kids := []*Tile{
		&Tile{tile.X * 2, tile.Y * 2, tile.Z + 1},
		&Tile{tile.X * 2 + 1, tile.Y * 2, tile.Z + 1},
		&Tile{tile.X * 2 + 1, tile.Y * 2 + 1, tile.Z + 1},
		&Tile{tile.X * 2, tile.Y * 2 + 1, tile.Z + 1},
	}
	return kids
}