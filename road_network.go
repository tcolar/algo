// HisTory: Nov 22 13 tcolar Creation

package algo

// Test data for graph stuff
// Network of few cities with connectng roads
var roadNetwork = RoadNetwork{
  Cities: []City{
    City{1, "Tacoma"},
    City{2, "Tukwilla"},
    City{3, "Lynnwood"},
    City{4, "Everett"},
    City{5, "Spokane"},
    City{6, "Portland"},
    City{7, "HermisTon"},
    City{8, "Ellensburg"},
    City{9, "Seattle"},
    City{10, "Bellevue"},
  },
  Roads: []Road{
    // tacoma portland
    Road{From: 1, To: 6, Dist: 144, Name: "I 5"},
    // tacoma tukwilla
    Road{From: 1, To: 2, Dist: 23, Name: "I 5"},
    // tukwilla seattle
    Road{From: 2, To: 9, Dist: 10, Name: "I 5"},
    // tukwilla-bellevue
    Road{From: 2, To: 10, Dist: 12, Name: "I 405"},
    // lynnwood-belevue
    Road{From: 3, To: 10, Dist: 19, Name: "I 405"},
    // lynnwood everett
    Road{From: 3, To: 4, Dist: 11, Name: "I 5"},
    // lynnwood seattle
    Road{From: 3, To: 9, Dist: 21, Name: "I 5"},
    // everett spokane
    Road{From: 5, To: 4, Dist: 279, Name: "US 2"},
    // spokane hermisTon
    Road{From: 5, To: 7, Dist: 170, Name: "US 395"},
    // spokane ellensburg
    Road{From: 5, To: 8, Dist: 173, Name: "I 90"},
    // portland hermisTon
    Road{From: 6, To: 7, Dist: 187, Name: "I 84"},
    // hermisTon ellensburg
    Road{From: 7, To: 8, Dist: 142, Name: "I 82"},
    // ellensburg bellevue
    Road{From: 8, To: 10, Dist: 98, Name: "I 90"},
    // seattle bellevue
    Road{From: 9, To: 10, Dist: 7, Name: "I 90"},
  },
}

// a road edge
type Road struct {
  From int
  To   int
  Dist int // distance (~weight)
  Name string
}

func (r Road) NdFrom() int {
  return r.From
}

func (r Road) NdTo() int {
  return r.To
}

// a City node
type City struct {
  Id   int
  Name string
}

func (c City) NdId() int {
  return c.Id
}

type RoadNetwork struct {
  Cities []City
  Roads  []Road
}

func (r RoadNetwork) Nodes() (nodes []Node) {
  for _, city := range r.Cities {
    nodes = append(nodes, Node(city))
  }
  return nodes
}

func (r RoadNetwork) Edges() (edges []Edge) {
  for _, road := range r.Roads {
    edges = append(edges, Edge(road))
  }
  return edges
}