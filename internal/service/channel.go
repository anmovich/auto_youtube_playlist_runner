package service
/*
import (
	"io"
	"net/http"
	"fmt"
)
*/
/*
func (s *Service) GetLink(channel_link string) error {
	data, err := s.getAllVideos("6767")
	if err != nil{
		return err
	}
	fmt.Println(string(data))
	return nil
}

/*
func (s *Service) getAllVideos(channelLink string) ([]byte, error){
	resp, err := http.Get("https://api.example.com/data")
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	return data, nil
} */
