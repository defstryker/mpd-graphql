package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/defstryker/mpd-graphql/graph/generated"
	"github.com/defstryker/mpd-graphql/graph/model"
	"github.com/defstryker/mpd-graphql/internal/library"
)

func (r *mutationResolver) SavePlaylist(ctx context.Context, name string) (*model.Playlist, error) {
	return library.SavePlaylist(name)
}

func (r *mutationResolver) AddToPlaylist(ctx context.Context, name string, uri []string) (*model.Playlist, error) {
	return library.AddToPlaylist(name, uri)
}

func (r *mutationResolver) AddToNowPlaying(ctx context.Context, uri []string) (*model.Playlist, error) {
	return library.AddToNowPlaying(uri)
}

func (r *queryResolver) GetSong(ctx context.Context, title string) (*model.Song, error) {
	return library.GetSong(title)
}

func (r *queryResolver) GetAlbum(ctx context.Context, name string) (*model.Album, error) {
	return library.GetAlbum(name)
}

func (r *queryResolver) GetArtist(ctx context.Context, name string) (*model.Artist, error) {
	// log.Println("get artist: ", name)
	return library.GetArtist(name)
}

func (r *queryResolver) GetPlaylist(ctx context.Context, name string) (*model.Playlist, error) {
	return library.GetPlaylist(name)
}

func (r *queryResolver) ListPlaylists(ctx context.Context) ([]*model.Playlist, error) {
	return library.ListPlaylists()
}

func (r *queryResolver) GetCurrentPlaylist(ctx context.Context) (*model.Playlist, error) {
	return library.GetCurrentPlaylist()
}

func (r *queryResolver) GetAllArtists(ctx context.Context) ([]*model.Artist, error) {
	return library.GetAllArtists()
}

func (r *queryResolver) GetAllAlbums(ctx context.Context, name string) ([]*model.Album, error) {
	return library.GetAllAlbums(name)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
