package sdk500px

//浏览页面
type IndexPage struct {
	Data []struct {
		UserPictureShareedCount int           `json:"userPictureShareedCount,omitempty"`
		UserLikerState          bool          `json:"userLikerState,omitempty"`
		Rating                  float64       `json:"rating,omitempty"`
		Description             string        `json:"description,omitempty"`
		Title                   string        `json:"title,omitempty"`
		Photos                  []interface{} `json:"photos,omitempty"`
		RatingMax               float64       `json:"ratingMax,omitempty"`
		Operator                struct {
			NickName string `json:"nickName"`
			Avatar   struct {
				BaseURL string `json:"baseUrl"`
			} `json:"avatar"`
			UserFolloweeState      bool   `json:"userFolloweeState"`
			UserName               string `json:"userName"`
			GicEditorialID         string `json:"gicEditorialId"`
			UserfolloweeLookUpload bool   `json:"userfolloweeLookUpload"`
			CoverURL               struct {
				BaseURL string `json:"baseUrl"`
				ID      string `json:"id"`
			} `json:"coverUrl"`
			GicCreativeID         string `json:"gicCreativeId"`
			UserFollowedCount     int    `json:"userFollowedCount"`
			UserfolloweeLookShare bool   `json:"userfolloweeLookShare"`
			UserfolloweeFocus     bool   `json:"userfolloweeFocus"`
			ID                    string `json:"id"`
			UserType              int    `json:"userType"`
			UserFollowedState     bool   `json:"userFollowedState"`
		} `json:"operator,omitempty"`
		PhotoCount     int    `json:"photoCount,omitempty"`
		CommentShareID string `json:"commentShareId,omitempty"`
		UploaderInfo   struct {
			NickName string `json:"nickName"`
			Avatar   struct {
				BaseURL string `json:"baseUrl"`
			} `json:"avatar"`
			UserFolloweeState      bool   `json:"userFolloweeState"`
			UserName               string `json:"userName"`
			GicEditorialID         string `json:"gicEditorialId"`
			UserfolloweeLookUpload bool   `json:"userfolloweeLookUpload"`
			CoverURL               struct {
				BaseURL string `json:"baseUrl"`
				ID      string `json:"id"`
			} `json:"coverUrl"`
			GicCreativeID         string `json:"gicCreativeId"`
			UserFollowedCount     int    `json:"userFollowedCount"`
			UserfolloweeLookShare bool   `json:"userfolloweeLookShare"`
			UserfolloweeFocus     bool   `json:"userfolloweeFocus"`
			ID                    string `json:"id"`
			UserType              int    `json:"userType"`
			UserFollowedState     bool   `json:"userFollowedState"`
		} `json:"uploaderInfo,omitempty"`
		UserPictureShareState bool   `json:"userPictureShareState,omitempty"`
		CreatedTime           int64  `json:"createdTime,omitempty"`
		ID                    string `json:"id,omitempty"`
		State                 int    `json:"state,omitempty"`
		ContestTags           string `json:"contestTags,omitempty"`
		OptTime               int64  `json:"optTime,omitempty"`
		OperatorID            string `json:"operatorId,omitempty"`
		Height                int    `json:"height,omitempty"`
		UploaderID            string `json:"uploaderId,omitempty"`
		Comments              []struct {
			UpdatedTime int64 `json:"updatedTime"`
			UserInfo    struct {
				NickName string `json:"nickName"`
				ID       string `json:"id"`
				Avatar   struct {
					A1      string `json:"a1"`
					BaseURL string `json:"baseUrl"`
				} `json:"avatar"`
				UserName string `json:"userName"`
			} `json:"userInfo"`
			ResourceID    string        `json:"resourceId"`
			Like          bool          `json:"like"`
			IP            string        `json:"ip"`
			CountLike     int           `json:"countLike"`
			FanyiFlag     int           `json:"fanyiFlag"`
			PlatformType  int           `json:"platformType"`
			Sort          int           `json:"sort"`
			Message       string        `json:"message"`
			Type          int           `json:"type"`
			UserID        string        `json:"userId"`
			ParentID      string        `json:"parentId"`
			MessageFanyi  string        `json:"messageFanyi"`
			ID            string        `json:"id"`
			State         int           `json:"state"`
			ChildComments []interface{} `json:"childComments"`
			CreateDate    int64         `json:"createDate"`
		} `json:"comments,omitempty"`
		CommentShareMsg string `json:"commentShareMsg,omitempty"`
		OpenState       string `json:"openState,omitempty"`
		PicturePvCount  int    `json:"picturePvCount,omitempty"`
		SetSetCount     int    `json:"setSetCount,omitempty"`
		URL             struct {
			BaseURL string `json:"baseUrl"`
			ID      string `json:"id"`
		} `json:"url,omitempty"`
		CommentCount       int `json:"commentCount,omitempty"`
		PictureLikeedCount int `json:"pictureLikeedCount,omitempty"`
		ExifInfo           struct {
		} `json:"exifInfo,omitempty"`
		CreatedDate  int64  `json:"createdDate,omitempty"`
		SourceType   string `json:"sourceType"`
		Width        int    `json:"width,omitempty"`
		TribeID      string `json:"tribeId,omitempty"`
		ResourceType int    `json:"resourceType,omitempty"`
		FlowType     string `json:"flowType"`
	} `json:"data"`
	NextStartTime int64  `json:"nextStartTime"`
	Message       string `json:"message"`
	Status        string `json:"status"`
}

//单个照片页面
type PhotoDetail struct {
	UserPictureShareedCount    int     `json:"userPictureShareedCount"`
	DoTsa                      bool    `json:"doTsa"`
	IsSignRecommend            bool    `json:"isSignRecommend"`
	Rating                     float64 `json:"rating"`
	UserLikerState             bool    `json:"userLikerState"`
	RiseUpDate                 int64   `json:"riseUpDate"`
	UploadedDate               int64   `json:"uploadedDate"`
	PictureStrategyCount       int     `json:"pictureStrategyCount"`
	RatingMax                  float64 `json:"ratingMax"`
	HotUpDate                  int64   `json:"hotUpDate"`
	UserPictureContentState    bool    `json:"userPictureContentState"`
	UserPictureCreativityState bool    `json:"userPictureCreativityState"`
	UserPictureTechnicalState  bool    `json:"userPictureTechnicalState"`
	CreatedTime                int64   `json:"createdTime"`
	ContestTags                string  `json:"contestTags"`
	ID                         string  `json:"id"`
	State                      int     `json:"state"`
	Height                     int     `json:"height"`
	DownLoadURL                string  `json:"downLoadUrl"`
	UploaderID                 string  `json:"uploaderId"`
	UserLightState             bool    `json:"userLightState"`
	Sort                       int     `json:"sort"`
	CommentCount               int     `json:"commentCount"`
	PictureLikeedCount         int     `json:"pictureLikeedCount"`
	Tags                       string  `json:"tags"`
	ExtendedField              struct {
		MapShow    bool  `json:"mapShow"`
		RiseUpDate int64 `json:"riseUpDate"`
		HotUpDate  int64 `json:"hotUpDate"`
	} `json:"extendedField"`
	RatingMaxDate int64 `json:"ratingMaxDate"`
	Exif          struct {
	} `json:"exif"`
	Origin                       string `json:"origin"`
	Description                  string `json:"description"`
	Privacy                      int    `json:"privacy"`
	PictureLightedCount          int    `json:"pictureLightedCount"`
	Title                        string `json:"title"`
	UserPictureCreativityedCount int    `json:"userPictureCreativityedCount"`
	PhotoCount                   int    `json:"photoCount"`
	TribeTags                    string `json:"tribeTags"`
	OriginID                     int    `json:"originId"`
	UploaderInfo                 struct {
		UserRoleIds struct {
		} `json:"userRoleIds"`
		NickName string `json:"nickName"`
		Avatar   struct {
			A1      string `json:"a1"`
			A2      string `json:"a2"`
			A3      string `json:"a3"`
			BaseURL string `json:"baseUrl"`
			A4      string `json:"a4"`
		} `json:"avatar"`
		UserFolloweeState      bool   `json:"userFolloweeState"`
		UserName               string `json:"userName"`
		GicEditorialID         string `json:"gicEditorialId"`
		UserfolloweeLookUpload bool   `json:"userfolloweeLookUpload"`
		CoverURL               struct {
			BaseURL string `json:"baseUrl"`
			ID      string `json:"id"`
		} `json:"coverUrl"`
		GicCreativeID         string `json:"gicCreativeId"`
		UserFollowedCount     int    `json:"userFollowedCount"`
		UserfolloweeLookShare bool   `json:"userfolloweeLookShare"`
		UserfolloweeFocus     bool   `json:"userfolloweeFocus"`
		ID                    string `json:"id"`
		State                 int    `json:"state"`
		UserType              int    `json:"userType"`
		UserFollowedState     bool   `json:"userFollowedState"`
	} `json:"uploaderInfo"`
	UserPictureTechnicaledCount   int    `json:"userPictureTechnicaledCount"`
	UserPictureShareState         bool   `json:"userPictureShareState"`
	ProfileSortTime               int64  `json:"profileSortTime"`
	Timestamp                     bool   `json:"timestamp"`
	HasCover                      int    `json:"hasCover"`
	OpenState                     string `json:"openState"`
	UserPicturePutupedCount       int    `json:"userPicturePutupedCount"`
	UserPictureCompositionState   bool   `json:"userPictureCompositionState"`
	PicturePvCount                int    `json:"picturePvCount"`
	UserPictureCompositionedCount int    `json:"userPictureCompositionedCount"`
	URL                           struct {
		P1      string `json:"p1"`
		P2      string `json:"p2"`
		BaseURL string `json:"baseUrl"`
		P5      string `json:"p5"`
		P6      string `json:"p6"`
		ID      string `json:"id"`
	} `json:"url"`
	ExifInfo struct {
		ResourceID string `json:"resourceId"`
		UploadTime int64  `json:"uploadTime"`
	} `json:"exifInfo"`
	CreatedDate               int64  `json:"createdDate"`
	UploaderName              string `json:"uploaderName"`
	Refer                     string `json:"refer"`
	Width                     int    `json:"width"`
	UserPictureContentedCount int    `json:"userPictureContentedCount"`
	ContentLength             int    `json:"contentLength"`
	Category                  struct {
		FivepxID    string `json:"fivepxId"`
		Name        string `json:"name"`
		CreatedTime int64  `json:"createdTime"`
		Description string `json:"description"`
		ID          string `json:"id"`
	} `json:"category"`
	UserPicturePutupState bool   `json:"userPicturePutupState"`
	CategoryID            string `json:"categoryId"`
	ResourceType          int    `json:"resourceType"`
}
