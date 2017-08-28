package gophertile

import (
	"testing"
	"fmt"
)

func TestTile_Ul(t *testing.T) {

	tile := Tile{486,332,10,}
	ll := tile.Ul()
	expected := LngLat{-9.140625, 53.33087298301705}

	assertEq(t,ll.Lng,expected.Lng)
	assertEq(t,ll.Lat,expected.Lat)

}

func TestTile_Bounds(t *testing.T) {

	tile := Tile{486,332,10,}
	expected := LngLatBbox{-9.140625, 53.120405283106564, -8.7890625, 53.33087298301705}
	bbox := tile.Bounds()
	assertEq(t, expected.East, bbox.East)
	assertEq(t, expected.West, bbox.West)
	assertEq(t, expected.North, bbox.North)
	assertEq(t, expected.South, bbox.South)
}

func TestTile_Parent(t *testing.T) {
	expected := Tile{243, 166, 9}
	tile := Tile{486, 332, 10}
	parent := tile.Parent()
	assertEq(t, expected.X,parent.X)
	assertEq(t, expected.Y,parent.Y)
	assertEq(t, expected.Z,parent.Z)
}

func TestTile_Children(t *testing.T) {

	tile := Tile{246,166,9}
	expected := Tile{492, 332, 10}
	children := tile.Children()

	found := false
	for _,t2 := range children {
		if t2.Equals(&expected){
			found = true
		}
	}
	if !found{
		t.Fail()
	}

}

func TestGetTile(t *testing.T) {

	tile := GetTile(20.6852, 40.1222, 9)
	expected := Tile{285,193,9}

	assertEq(t,tile.Z, expected.Z)
	assertEq(t,tile.Y, expected.Y)
	assertEq(t,tile.X, expected.X)

}

func assertEq(t *testing.T, x interface{}, y interface{}){
	if x != y {

		fmt.Printf("%v is not equal to %v", x , y)
		t.Fail()
	}
}