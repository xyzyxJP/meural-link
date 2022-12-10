package model

type GalleryItems struct {
	Count   int    `json:"count"`
	Message string `json:"message"`
	Data    []struct {
		OriginalImage  string      `json:"originalImage"`
		Image          string      `json:"image"`
		Year           string      `json:"year"`
		UserID         int         `json:"userId"`
		ArtistID       int         `json:"artistId"`
		SortedOrder    interface{} `json:"sortedOrder"`
		OriginalWidth  int         `json:"originalWidth"`
		Icon2X         string      `json:"icon2x"`
		ID             int         `json:"id"`
		CroppedY1      int         `json:"croppedY1"`
		Dimensions     interface{} `json:"dimensions"`
		Copyright      interface{} `json:"copyright"`
		Author         string      `json:"author"`
		Hero2X         string      `json:"hero2x"`
		CategoryIds    []int       `json:"categoryIds"`
		CroppedWidth   int         `json:"croppedWidth"`
		Type           string      `json:"type"`
		IsSampler      bool        `json:"isSampler"`
		Medium         string      `json:"medium"`
		Hero           string      `json:"hero"`
		Description    string      `json:"description"`
		Name           string      `json:"name"`
		IsPublic       bool        `json:"isPublic"`
		CroppedX1      int         `json:"croppedX1"`
		Slug           string      `json:"slug"`
		Icon           string      `json:"icon"`
		CategoryNames  []string    `json:"categoryNames"`
		CroppedHeight  int         `json:"croppedHeight"`
		OriginalHeight int         `json:"originalHeight"`
		ArtistName     string      `json:"artistName"`
		IsDetailed     bool        `json:"isDetailed"`
	} `json:"data"`
	IsLast      bool `json:"isLast"`
	IsPaginated bool `json:"isPaginated"`
}
