package library

import (
	"log"

	"github.com/defstryker/mpd-graphql/conn"
	"github.com/defstryker/mpd-graphql/graph/model"
	"github.com/fhs/gompd/mpd"
)

func GetAllArtists() ([]*model.Artist, error) {
	var artists []*model.Artist
	out, err := conn.Conn.List(`artist`)
	if err != nil {
		return nil, err
	}
	for _, a := range out {
		if a == "" {
			continue
		}
		art, err := GetArtist(a)
		if err != nil {
			log.Println("GetAllArtists")
			return nil, err
		}
		artists = append(artists, art)
	}
	return artists, nil
}

func GetAllAlbums(name string) ([]*model.Album, error) {
	var albums []*model.Album
	var out []string
	var err error

	if name == "" {
		return nil, nil
	}

	out, err = conn.Conn.List("album", "artist", name)
	if err != nil {
		return nil, err
	}

	for _, name := range out {
		album, err := GetAlbum(name)
		if err != nil {
			log.Println("GetAllAlbums")
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func GetAllSongs(name string) ([]*model.Song, error) {
	var songs []*model.Song

	out, err := conn.Conn.List("title", "album", name)
	if err != nil {
		return nil, err
	}
	for _, song := range out {
		songs = append(songs, &model.Song{Title: song})
	}
	return songs, nil
}

func ListPlaylists() ([]*model.Playlist, error) {
	var playlists []*model.Playlist
	out, err := conn.Conn.ListPlaylists()
	if err != nil {
		return nil, err
	}
	for _, pl := range out {
		name := pl["playlist"]

		data, err := conn.Conn.PlaylistContents(name)
		if err != nil {
			return nil, err
		}

		songs := toSong(data)

		playlists = append(
			playlists,
			&model.Playlist{
				Name:  name,
				Songs: songs,
			},
		)
	}
	// conn.Conn.PlaylistContents
	return playlists, nil
}

func GetPlaylist(name string) (*model.Playlist, error) {
	data, err := conn.Conn.PlaylistContents(name)
	if err != nil {
		return nil, err
	}

	playlist := &model.Playlist{
		Name:  name,
		Songs: toSong(data),
	}

	return playlist, nil
}

func GetCurrentPlaylist() (*model.Playlist, error) {
	out, err := conn.Conn.PlaylistInfo(-1, -1)
	if err != nil {
		return nil, err
	}

	playlist := &model.Playlist{
		Name:  "Now Playing",
		Songs: toSong(out),
	}

	return playlist, nil
}

func GetArtist(name string) (*model.Artist, error) {
	albums, err := GetAllAlbums(name)
	if err != nil {
		return nil, err
	}
	return &model.Artist{
		Name:   name,
		Albums: albums,
	}, nil
}

func GetAlbum(name string) (*model.Album, error) {
	songs, err := GetAllSongs(name)
	if err != nil {
		log.Println("GetAlbum")
		return nil, err
	}
	return &model.Album{
		Name:  name,
		Songs: songs,
	}, nil
}

func GetSong(title string) (*model.Song, error) {
	data, err := conn.Conn.Find("title", title)
	if err != nil {
		return nil, err
	}
	return toSong(data)[0], nil
}

func AddToPlaylist(name string, songs []string) (*model.Playlist, error) {
	for _, song := range songs {
		if err := conn.Conn.PlaylistAdd(name, song); err != nil {
			log.Println("AddToPlaylist")
			return nil, err
		}
	}
	pl, err := GetPlaylist(name)
	if err != nil {
		log.Println("AddToPlaylist")
		return nil, err
	}
	return pl, nil
}

func AddToNowPlaying(songs []string) (*model.Playlist, error) {
	for _, song := range songs {
		if err := conn.Conn.Add(song); err != nil {
			log.Println("AddToNowPlaying")
			return nil, err
		}
	}
	pl, err := GetCurrentPlaylist()
	if err != nil {
		log.Println("AddToNowPlaying")
		return nil, err
	}
	return pl, nil
}

func SavePlaylist(name string) (*model.Playlist, error) {
	if err := conn.Conn.PlaylistSave(name); err != nil {
		return nil, err
	}
	pl, err := GetPlaylist(name)
	if err != nil {
		log.Println("SavePlaylist")
		return nil, err
	}
	return pl, nil
}

func toSong(data []mpd.Attrs) []*model.Song {
	var songs []*model.Song
	for _, d := range data {
		track := d["Track"]
		uri := d["file"]
		songs = append(songs, &model.Song{
			Title:    d["Title"],
			Track:    &track,
			Album:    d["Album"],
			Artist:   d["Artist"],
			Duration: d["duration"],
			URI:      &uri,
		})
	}
	return songs
}
