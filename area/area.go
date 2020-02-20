package area

import (
	"errors"
	"github.com/gregmus2/village-of-life/character"
	"github.com/lafriks/go-tiled"
)

type Obstacle struct{}
type ObstaclesArea map[int]map[int]*Obstacle
type CharactersArea map[int]map[int]*character.Character
type SpritesArea [][]*Sprite
type Sprite struct {
	ID          uint32
	TileSetName string
}

type Area struct {
	W          int
	H          int
	tileW      int
	tileH      int
	Obstacles  ObstaclesArea
	Characters CharactersArea
	Sprites    SpritesArea
	TileSets   map[string]*tiled.Tileset
}

const SpritesLayer string = "sprites"
const ObstaclesLayer string = "obstacles"
const CharactersLayer string = "characters"

func (s *Service) ParseArea(file string) (*Area, error) {
	a := &Area{
		Obstacles:  make(ObstaclesArea),
		Characters: make(CharactersArea),
		Sprites:    nil,
		TileSets:   make(map[string]*tiled.Tileset),
	}

	gameMap, err := tiled.LoadFromFile(file)
	if err != nil {
		return nil, err
	}

	a.tileW = gameMap.TileWidth
	a.tileH = gameMap.TileHeight
	a.W = gameMap.Width
	a.H = gameMap.Height

	for _, tileSet := range gameMap.Tilesets {
		a.TileSets[tileSet.Name] = tileSet
	}

	for _, layer := range gameMap.Layers {
		switch layer.Name {
		case SpritesLayer:
			s.ParseSprites(a, layer.Tiles)
		case ObstaclesLayer:
			s.ParseObstacles(a, layer.Tiles)
		case CharactersLayer:
			s.ParseCharacters(a, layer.Tiles)
		default:
			return nil, errors.New("unknown layer: " + layer.Name)
		}
	}

	return a, nil
}

func (s *Service) ParseSprites(a *Area, tiles []*tiled.LayerTile) {
	a.Sprites = make(SpritesArea, 0, tiles[0].Tileset.Columns)
	for x := 0; x < a.W; x++ {
		a.Sprites = append(a.Sprites, make([]*Sprite, 0, a.H))
		for y := 0; y < a.H; y++ {
			tile := tiles[y+a.W*x]
			a.Sprites[x] = append(a.Sprites[x], &Sprite{tile.ID, tile.Tileset.Name})
		}
	}
}

func (s *Service) ParseObstacles(a *Area, tiles []*tiled.LayerTile) {
	for x := 0; x < a.W; x++ {
		a.Obstacles[x] = make(map[int]*Obstacle)
		for y := 0; y < a.H; y++ {
			tile := tiles[y+a.W*x]
			if tile.IsNil() {
				continue
			}

			a.Sprites[x][y] = &Sprite{tile.ID, tile.Tileset.Name}
			a.Obstacles[x][y] = &Obstacle{}
		}
	}
}

func (s *Service) ParseCharacters(a *Area, tiles []*tiled.LayerTile) {
	for x := 0; x < a.W; x++ {
		a.Characters[x] = make(map[int]*character.Character)
		for y := 0; y < a.H; y++ {
			tile := tiles[y+a.W*x]
			if tile.IsNil() {
				continue
			}

			a.Sprites[x][y] = &Sprite{tile.ID, tile.Tileset.Name}
			a.Characters[x][y] = s.c.NewCharacter(s.u.GenerateName(), 0)
		}
	}
}
