package model

type VideoData struct {
	ID                     uint    `gorm:"primarykey" csv:"id"`
	Date                   string  `csv:"date"`
	Title                  string  `csv:"title"`
	Status                 string  `csv:"status"`
	PubTime                string  `gorm:"pub_time" csv:"pub_time"`
	PlaybackVolume         int     `gorm:"playback_volume" csv:"playback_volume"`
	CompleteRatio          float32 `gorm:"complete_ratio" csv:"complete_ratio"`
	AverageDuration        float32 `gorm:"average_duration" csv:"average_duration"`
	Like                   int     `csv:"like"`
	Comment                int     `csv:"comment"`
	Share                  int     `csv:"share"`
	Access                 int     `csv:"access"`
	Fans                   int     `csv:"fans"`
	CompleteLike           float32 `gorm:"complete_like" csv:"complete_like"`
	CompleteShare          float32 `gorm:"complete_share" csv:"complete_share"`
	DayPlaybackVolume      int     `gorm:"day_playback_volume" csv:"day_playback_volume"`
	CompleteLikeIncrement  float32 `gorm:"complete_like_increment" csv:"complete_like_increment"`
	CompleteShareIncrement float32 `gorm:"complete_share_increment" csv:"complete_share_increment"`
}

func (c *VideoData) TableName() string {
	return database + ".video_data"
}
