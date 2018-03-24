package entry

// Category defines a category defined by the newznab spec
type Category struct {
	Code int
	Text string
}

// categories defined by the newznab spec
var (
	CategoryReserved          = Category{Code: 0000, Text: "Reserved"}
	CategoryConsole           = Category{Code: 1000, Text: "Console"}
	CategoryConsoleNDS        = Category{Code: 1010, Text: "Console/NDS"}
	CategoryConsolePSP        = Category{Code: 1020, Text: "Console/PSP"}
	CategoryConsoleWii        = Category{Code: 1030, Text: "Console/Wii"}
	CategoryConsoleXbox       = Category{Code: 1040, Text: "Console/XBox"}
	CategoryConsoleXbox360    = Category{Code: 1050, Text: "Console/XBox 360"}
	CategoryConsoleWiiware    = Category{Code: 1060, Text: "Console/Wiiware"}
	CategoryConsoleXbox360DLC = Category{Code: 1070, Text: "Console/XBox 360 DLC"}
	CategoryMovies            = Category{Code: 2000, Text: "Movies"}
	CategoryMoviesForeign     = Category{Code: 2010, Text: "Movies/Foreign"}
	CategoryMoviesOther       = Category{Code: 2020, Text: "Movies/Other"}
	CategoryMoviesSD          = Category{Code: 2030, Text: "Movies/SD"}
	CategoryMoviesHD          = Category{Code: 2040, Text: "Movies/HD"}
	CategoryMoviesUHD         = Category{Code: 2045, Text: "Movies/UHD"}
	CategoryMoviesBluRay      = Category{Code: 2050, Text: "Movies/BluRay"}
	CategoryMovies3D          = Category{Code: 2060, Text: "Movies/3D"}
	CategoryAudio             = Category{Code: 3000, Text: "Audio"}
	CategoryAudioMP3          = Category{Code: 3010, Text: "Audio/MP3"}
	CategoryAudioVideo        = Category{Code: 3020, Text: "Audio/Video"}
	CategoryAudioAudiobook    = Category{Code: 3030, Text: "Audio/Audiobook"}
	CategoryAudioLossless     = Category{Code: 3040, Text: "Audio/Lossless"}
	CategoryPC                = Category{Code: 4000, Text: "PC"}
	CategoryPCZeroDay         = Category{Code: 4010, Text: "PC/0day"}
	CategoryPCISO             = Category{Code: 4020, Text: "PC/ISO"}
	CategoryPCMac             = Category{Code: 4030, Text: "PC/Mac"}
	CategoryPCMobileOther     = Category{Code: 4040, Text: "PC/Mobile-Other"}
	CategoryPCGames           = Category{Code: 4050, Text: "PC/Games"}
	CategoryPCMobileiOS       = Category{Code: 4060, Text: "PC/Mobile-iOS"}
	CategoryPCMobileAndroid   = Category{Code: 4070, Text: "PC/Mobile-Android"}
	CategoryTV                = Category{Code: 5000, Text: "TV"}
	CategoryTVForeign         = Category{Code: 5020, Text: "TV/Foreign"}
	CategoryTVSD              = Category{Code: 5030, Text: "TV/SD"}
	CategoryTVHD              = Category{Code: 5040, Text: "TV/HD"}
	CategoryTVUHD             = Category{Code: 5045, Text: "TV/UHD"}
	CategoryTVOther           = Category{Code: 5050, Text: "TV/Other"}
	CategoryTVSport           = Category{Code: 5060, Text: "TV/Sport"}
	CategoryTVAnime           = Category{Code: 5070, Text: "TV/Anime"}
	CategoryTVDocumentary     = Category{Code: 5080, Text: "TV/Documentary"}
	CategoryXXX               = Category{Code: 6000, Text: "XXX"}
	CategoryXXXDVD            = Category{Code: 6010, Text: "XXX/DVD"}
	CategoryXXXWMV            = Category{Code: 6020, Text: "XXX/WMV"}
	CategoryXXXXvid           = Category{Code: 6030, Text: "XXX/XviD"}
	CategoryXXXx264           = Category{Code: 6040, Text: "XXX/x264"}
	CategoryXXXPack           = Category{Code: 6050, Text: "XXX/Pack"}
	CategoryXXXImageSet       = Category{Code: 6060, Text: "XXX/ImgSet"}
	CategoryXXXOther          = Category{Code: 6070, Text: "XXX/Other"}
	CategoryBooks             = Category{Code: 7000, Text: "Books"}
	CategoryBooksMags         = Category{Code: 7010, Text: "Books/Mags"}
	CategoryBooksEbook        = Category{Code: 7020, Text: "Books/EBook"}
	CategoryBooksComics       = Category{Code: 7030, Text: "Books/Comics"}
	CategoryOther             = Category{Code: 8000, Text: "Other"}
	CategoryOtherMisc         = Category{Code: 8010, Text: "Other/Misc"}
)

// CategoryFromCode takes an integer category code, and returns the
// corresponding Category struct
func CategoryFromCode(code int) Category {

	var categories = map[int]Category{
		0:    CategoryReserved,
		1000: CategoryConsole,
		1010: CategoryConsoleNDS,
		1020: CategoryConsolePSP,
		1030: CategoryConsoleWii,
		1040: CategoryConsoleXbox,
		1050: CategoryConsoleXbox360,
		1060: CategoryConsoleWiiware,
		1070: CategoryConsoleXbox360DLC,
		2000: CategoryMovies,
		2010: CategoryMoviesForeign,
		2020: CategoryMoviesOther,
		2030: CategoryMoviesSD,
		2040: CategoryMoviesHD,
		2045: CategoryMoviesUHD,
		2050: CategoryMoviesBluRay,
		2060: CategoryMovies3D,
		3000: CategoryAudio,
		3010: CategoryAudioMP3,
		3020: CategoryAudioVideo,
		3030: CategoryAudioAudiobook,
		3040: CategoryAudioLossless,
		4000: CategoryPC,
		4010: CategoryPCZeroDay,
		4020: CategoryPCISO,
		4030: CategoryPCMac,
		4040: CategoryPCMobileOther,
		4050: CategoryPCGames,
		4060: CategoryPCMobileiOS,
		4070: CategoryPCMobileAndroid,
		5000: CategoryTV,
		5020: CategoryTVForeign,
		5030: CategoryTVSD,
		5040: CategoryTVHD,
		5045: CategoryTVUHD,
		5050: CategoryTVOther,
		5060: CategoryTVSport,
		5070: CategoryTVAnime,
		5080: CategoryTVDocumentary,
		6000: CategoryXXX,
		6010: CategoryXXXDVD,
		6020: CategoryXXXWMV,
		6030: CategoryXXXXvid,
		6040: CategoryXXXx264,
		6050: CategoryXXXPack,
		6060: CategoryXXXImageSet,
		6070: CategoryXXXOther,
		7000: CategoryBooks,
		7010: CategoryBooksMags,
		7020: CategoryBooksEbook,
		7030: CategoryBooksComics,
		8000: CategoryOther,
		8010: CategoryOtherMisc,
	}
	if category, ok := categories[code]; ok {
		return category
	}
	return Category{Code: code}
}
