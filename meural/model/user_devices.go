package model

type UserDevices struct {
	Data []struct {
		Alias                string `json:"alias"`
		OriginalAlias        string `json:"originalAlias"`
		AlsEnabled           bool   `json:"alsEnabled"`
		AlsSensitivity       int    `json:"alsSensitivity"`
		AutoPowerMode        bool   `json:"autoPowerMode"`
		AutoPowerModeNetwork bool   `json:"autoPowerModeNetwork"`
		AutoUpgrade          bool   `json:"autoUpgrade"`
		BackgroundColor      string `json:"backgroundColor"`
		Credentials          struct {
			Softap struct {
				Ssid     string `json:"ssid"`
				Password string `json:"password"`
			} `json:"softap"`
		} `json:"credentials"`
		CurrentGallery  int    `json:"currentGallery"`
		CurrentImage    string `json:"currentImage"`
		DisplayLanguage string `json:"displayLanguage"`
		FillMode        string `json:"fillMode"`
		FrameColor      string `json:"frameColor"`
		FrameImageURL   string `json:"frameImageUrl"`
		FrameModel      struct {
			Color                   string `json:"color"`
			HorizontalFrameImageURL string `json:"horizontalFrameImageUrl"`
			ID                      int    `json:"id"`
			Name                    string `json:"name"`
			Resolution              string `json:"resolution"`
			TotalSpace              int    `json:"totalSpace"`
			SizeInches              string `json:"sizeInches"`
			Version                 string `json:"version"`
			VerticalFrameImageURL   string `json:"verticalFrameImageUrl"`
		} `json:"frameModel"`
		FrameStatus struct {
			CurrentGallery   int    `json:"currentGallery"`
			CurrentItem      int    `json:"currentItem"`
			LocalIP          string `json:"localIp"`
			RemoteIP         string `json:"remoteIp"`
			FreeSpace        int    `json:"freeSpace"`
			SyncedAt         string `json:"syncedAt"`
			LastSeen         string `json:"lastSeen"`
			LastSeen2        string `json:"lastSeen2"`
			Brightness       int    `json:"brightness"`
			LightSensor      int    `json:"lightSensor"`
			Sleep            bool   `json:"sleep"`
			CurrentlyPlaying struct {
				Type string `json:"type"`
				ID   int    `json:"id"`
			} `json:"currentlyPlaying"`
			UpNext        interface{} `json:"upNext"`
			ScheduledNext interface{} `json:"scheduledNext"`
		} `json:"frameStatus"`
		FreeSpace           int           `json:"freeSpace"`
		GalleryIds          []int         `json:"galleryIds"`
		AlbumIds            []interface{} `json:"albumIds"`
		GalleryRotation     bool          `json:"galleryRotation"`
		GoesDark            bool          `json:"goesDark"`
		GestureFeedback     bool          `json:"gestureFeedback"`
		GestureFeedbackHelp bool          `json:"gestureFeedbackHelp"`
		GestureFlip         bool          `json:"gestureFlip"`
		ID                  int           `json:"id"`
		ImageDuration       int           `json:"imageDuration"`
		ImageShuffle        bool          `json:"imageShuffle"`
		ItemIds             []int         `json:"itemIds"`
		LocalIP             string        `json:"localIp"`
		Model               int           `json:"model"`
		NewVersionID        interface{}   `json:"newVersionId"`
		Orientation         string        `json:"orientation"`
		OrientationMatch    bool          `json:"orientationMatch"`
		OverlayDuration     int           `json:"overlayDuration"`
		PendingInviteCount  int           `json:"pendingInviteCount"`
		PreviewDuration     int           `json:"previewDuration"`
		ProductKey          string        `json:"productKey"`
		SerialNumber        string        `json:"serialNumber"`
		ShareCount          int           `json:"shareCount"`
		Status              string        `json:"status"`
		SyncedAt            string        `json:"syncedAt"`
		Timezone            string        `json:"timezone"`
		SchedulerEnabled    bool          `json:"schedulerEnabled"`
		UpdatedAt           string        `json:"updatedAt"`
		Used                bool          `json:"used"`
		User                string        `json:"user"`
		UserID              int           `json:"userId"`
		Version             string        `json:"version"`
	} `json:"data"`
	IsPaginated bool   `json:"isPaginated"`
	IsLast      bool   `json:"isLast"`
	Count       int    `json:"count"`
	Message     string `json:"message"`
}
