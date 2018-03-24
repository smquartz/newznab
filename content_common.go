package entry

import "net/url"

// TVMovieMusicBook describes information common to TV, Movie, Music,
// and Book Content implementations
type TVMovieMusicBook struct {
	// embed common information
	TVMovieMusic
	// url for a cover image
	CoverImage url.URL
}

// TVMovieMusic describes information common to TV, Movie, and Music Content
// implementations
type TVMovieMusic struct {
	// url for a backdrop cover image
	BackdropCoverImage url.URL
	// audio codec used
	AudioCodec string
	// audio languages available
	Languages []string
}

// MovieMusicBook describes information common to Movie, Music, and Book
// Content implementations
type MovieMusicBook struct {
	// critics review score
	ReviewScore float32
}

// MusicBook describes information common to Music and Book Content
// implementations
type MusicBook struct {
	// embed common information
	MovieMusicBook
	// embed common information
	TVMovieMusicBook
	// name of the publisher
	Publisher string
}

// TVMovie describes information common to TV and Movie Content implementations
type TVMovie struct {
	// embed common information
	TVMovieMusicBook
	// embed common information
	TVMovieMusic
	// video codec used
	VideoCodec string
	// video resolution
	Resolution string
	// framerate
	Framerate float32
	// subtitles available
	Subtitles []string
	// genre of the content
	Genre string
}
