package material

import (
	"encoding/json"

	"../utils"
)

//MediaType 媒体文件类型
type MediaType string

const (
	//MediaTypeImage 媒体文件:图片
	MediaTypeImage MediaType = "image"
	//MediaTypeVoice 媒体文件:声音
	MediaTypeVoice = "voice"
	//MediaTypeVideo 媒体文件:视频
	MediaTypeVideo = "video"
	//MediaTypeThumb 媒体文件:缩略图
	MediaTypeThumb = "thumb"
)

const (
	mediaUploadURL      = "https://api.weixin.qq.com/cgi-bin/media/upload"
	mediaUploadImageURL = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	mediaGetURL         = "https://api.weixin.qq.com/cgi-bin/media/get"
)

//Media 临时素材上传返回信息
type Media struct {
	Type      MediaType `json:"type"`
	MediaID   string    `json:"media_id"`
	CreatedAt int64     `json:"created_at"`
}

//MediaUpload 临时素材上传
func MediaUpload(filename, uri string) (media Media, err error) {

	var response []byte
	response, err = util.PostFile("media", filename, uri)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &media)
	if err != nil {
		return
	}
	return media, err
}
