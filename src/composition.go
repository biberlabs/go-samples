/**
 * Composition example based on Go Object Oriented Design
 * see https://nathany.com/good/ for original article
 */
package main

import "fmt"

// define a new type named part
type Part struct {
    Name        string
    Description string
    NeedsSpare  bool
}

// declare String() method on Part type
// implement Stringer interface, which the fmt package makes use of.
func (part Part) String() string {
    return fmt.Sprintf("%s: (%s)", part.Name, part.Description)
}

// parts type: slice of Part values
type Parts []Part

// declare Spares() method on Parts type
func (parts Parts) Spares() (spares Parts) {
    // iterate over parts, ignore the index position (_), filter the parts to return
    for _, part := range parts {
        if part.NeedsSpare {
            spares = append(spares, part)
        }
    }
    return spares
}

// define a new type named Bicycle
type Bicycle struct {
    Size string
    // embed by not specifying a named field for Parts
    Parts
}

// initialize struct values with composite literal syntax
// see http://golang.org/doc/effective_go.html#composite_literals
var (
    RoadBikeParts = Parts{
        {"chain", "10-speed", true},
        {"tire_size", "23", true},
        {"tape_color", "red", true},
    }

    MountainBikeParts = Parts{
        {"chain", "10-speed", true},
        {"tire_size", "2.1", true},
        {"front_shock", "Manitou", false},
        {"rear_shock", "Fox", true},
    }

    RecumbentBikeParts = Parts{
        {"chain", "9-speed", true},
        {"tire_size", "28", true},
        {"flag", "tall and orange", true},
    }
)

// and main function
func main() {
    // short declaration operator (:=) uses type inference to initialize variables
    roadBike := Bicycle{Size: "L", Parts: RoadBikeParts}
    mountainBike := Bicycle{Size: "L", Parts: MountainBikeParts}
    recumbentBike := Bicycle{Size: "L", Parts: RecumbentBikeParts}

    fmt.Println(roadBike.Spares())
    fmt.Println(mountainBike.Spares())
    fmt.Println(recumbentBike.Spares())

    // Parts behaves like a slice.
    // Getting the length, slicing the slice, or combining multiple slices all works as usual.
    comboParts := Parts{}
    comboParts = append(comboParts, mountainBike.Parts...)
    comboParts = append(comboParts, roadBike.Parts...)
    comboParts = append(comboParts, recumbentBike.Parts...)

    fmt.Println(len(comboParts), comboParts[9:])
    fmt.Println(comboParts.Spares())
}