package sitemapcrawl

const (
	urlSetXMLNameLocal = "urlset"
	indexXMLNameLocal  = "sitemapindex"
)

type URL struct {
	Loc        string  `xml:"loc" json:"loc"`
	LastMod    string  `xml:"lastmod,omitempty" json:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty" json:"changefreq,omitempty"`
	Priority   float32 `xml:"priority,omitempty" json:"priority,omitempty"`
	Image      []Image `xml:"image,omitempty" json:"image,omitempty"`
	Video      []Video `xml:"video,omitempty" json:"video,omitempty"`
}

type Image struct {
	Loc         string `xml:"loc" json:"loc"`
	Title       string `xml:"title,omitempty" json:"title,omitempty"`
	Caption     string `xml:"caption,omitempty" json:"caption,omitempty"`
	GeoLocation string `xml:"geo_location,omitempty" json:"geoLocation,omitempty"`
	License     string `xml:"license,omitempty" json:"license,omitempty"`
}

type Video struct {
	ThumbnailLocation    string   `xml:"thumbnail_location" json:"thumbnailLocation"`
	Title                string   `xml:"title,omitempty" json:"title,omitempty"`
	Description          string   `xml:"description,omitempty" json:"description,omitempty"`
	ContentLoc           string   `xml:"content_loc,omitempty" json:"contentLoc,omitempty"`
	PlayerLoc            string   `xml:"player_loc,omitempty" json:"playerLoc,omitempty"`
	Duration             int      `xml:"duration,omitempty" json:"duration,omitempty"`
	ExpirationDate       string   `xml:"expiration_date,omitempty" json:"expirationDate,omitempty"`
	Rating               float32  `xml:"rating,omitempty" json:"rating,omitempty"`
	ViewCount            int64    `xml:"view_count,omitempty" json:"viewCount,omitempty"`
	PublicationDate      string   `xml:"publication_date,omitempty" json:"publication_date,omitempty"`
	FamilyFriendly       string   `xml:"family_friendly,omitempty" json:"familyFriendly,omitempty"`
	Restriction          string   `xml:"restriction,omitempty" json:"restriction,omitempty"`
	Platform             string   `xml:"platform,omitempty" json:"platform,omitempty"`
	Price                float32  `xml:"price,omitempty" json:"price,omitempty"`
	RequiresSubscription string   `xml:"requires_subscription,omitempty" json:"requiresSubscription,omitempty"`
	Uploader             string   `xml:"uploader,omitempty" json:"uploader,omitempty"`
	Live                 string   `xml:"live,omitempty" json:"live,omitempty"`
	Tag                  []string `xml:"tag,omitempty" json:"tag,omitempty"`
	Category             string   `xml:"category,omitempty" json:"category,omitempty"`
}
