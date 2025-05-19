package main

//Not good
var testTiles = [][]*Tile{
	{{0, 0, 0, Hidden}, {1, 0, 0, Hidden}, {2, 0, 0, Hidden}, {3, 0, 0, Hidden}, {4, 0, 0, Hidden}, {5, 0, 1, Hidden}, {6, 0, -1, Hidden}, {7, 0, 1, Hidden}, {8, 0, 0, Hidden}, {9, 0, 0, Hidden}},
	{{0, 1, 0, Hidden}, {1, 1, 0, Hidden}, {2, 1, 0, Hidden}, {3, 1, 1, Hidden}, {4, 1, 1, Hidden}, {5, 1, 2, Hidden}, {6, 1, 2, Hidden}, {7, 1, 2, Hidden}, {8, 1, 1, Hidden}, {9, 1, 0, Hidden}},
	{{0, 2, 0, Hidden}, {1, 2, 0, Hidden}, {2, 2, 0, Hidden}, {3, 2, 1, Hidden}, {4, 2, -1, Hidden}, {5, 2, 3, Hidden}, {6, 2, -1, Hidden}, {7, 2, 3, Hidden}, {8, 2, 2, Hidden}, {9, 2, 0, Hidden}},
	{{0, 3, 0, Hidden}, {1, 3, 0, Hidden}, {2, 3, 1, Hidden}, {3, 3, 2, Hidden}, {4, 3, 3, Hidden}, {5, 3, 4, Hidden}, {6, 3, 4, Hidden}, {7, 3, -1, Hidden}, {8, 3, 2, Hidden}, {9, 3, 0, Hidden}},
	{{0, 4, 0, Hidden}, {1, 4, 0, Hidden}, {2, 4, 1, Hidden}, {3, 4, -1, Hidden}, {4, 4, 3, Hidden}, {5, 4, -1, Hidden}, {6, 4, 4, Hidden}, {7, 4, 3, Hidden}, {8, 4, 2, Hidden}, {9, 4, 0, Hidden}},
	{{0, 5, 0, Hidden}, {1, 5, 0, Hidden}, {2, 5, 1, Hidden}, {3, 5, 2, Hidden}, {4, 5, 3, Hidden}, {5, 5, 4, Hidden}, {6, 5, -1, Hidden}, {7, 5, 2, Hidden}, {8, 5, 1, Hidden}, {9, 5, 0, Hidden}},
	{{0, 6, 0, Hidden}, {1, 6, 0, Hidden}, {2, 6, 0, Hidden}, {3, 6, 0, Hidden}, {4, 6, 1, Hidden}, {5, 6, 2, Hidden}, {6, 6, 2, Hidden}, {7, 6, 2, Hidden}, {8, 6, 1, Hidden}, {9, 6, 0, Hidden}},
	{{0, 7, 0, Hidden}, {1, 7, 0, Hidden}, {2, 7, 0, Hidden}, {3, 7, 0, Hidden}, {4, 7, 1, Hidden}, {5, 7, -1, Hidden}, {6, 7, 2, Hidden}, {7, 7, -1, Hidden}, {8, 7, 1, Hidden}, {9, 7, 0, Hidden}},
	{{0, 8, 0, Hidden}, {1, 8, 0, Hidden}, {2, 8, 0, Hidden}, {3, 8, 0, Hidden}, {4, 8, 1, Hidden}, {5, 8, 2, Hidden}, {6, 8, 3, Hidden}, {7, 8, 2, Hidden}, {8, 8, 1, Hidden}, {9, 8, 0, Hidden}},
	{{0, 9, 0, Hidden}, {1, 9, 0, Hidden}, {2, 9, 0, Hidden}, {3, 9, 0, Hidden}, {4, 9, 0, Hidden}, {5, 9, 1, Hidden}, {6, 9, -1, Hidden}, {7, 9, 1, Hidden}, {8, 9, 0, Hidden}, {9, 9, 0, Hidden}},
}
